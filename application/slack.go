package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/juntaki/firestarter/domain"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

const (
	seletType  = "select"
	buttonType = "button"
)

type SlackBot struct {
	VerificationToken string
	API               *slack.Client
	ConfigRepository  domain.ConfigRepository
	Log               *zap.SugaredLogger
	Session           *Session
	channelCache      map[string]string
	sqsMode           bool
}

func NewSlackBot(
	VerificationToken string,
	API *slack.Client,
	ConfigRepository domain.ConfigRepository,
	Log *zap.SugaredLogger,
	sqsMode bool,
) *SlackBot {
	return &SlackBot{
		VerificationToken: VerificationToken,
		API:               API,
		ConfigRepository:  ConfigRepository,
		Log:               Log,
		Session:           NewSession(),
		channelCache:      make(map[string]string),
		sqsMode:           sqsMode,
	}
}

func (s *SlackBot) InteractiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Error check
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.Log.Errorf("Failed to read request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		s.Log.Errorf("Failed to unespace request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		s.Log.Errorf("Failed to decode json message from slack: %s", jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if message.Token != s.VerificationToken {
		s.Log.Errorf("Invalid token: %s", message.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Load config
	config, err := s.ConfigRepository.GetConfigList()
	if err != nil {
		s.Log.Error("Get config map failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	q := config.FindByCallbackID(strings.Split(message.CallbackID, "@")[0])
	if q == nil {
		s.Log.Error("Config not found")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sess, ok := s.Session.Get(message.CallbackID)
	if !ok {
		s.Log.Errorw("Session expired", zap.String("Session ID", strings.Split(message.CallbackID, "@")[1]))
		s.responseMessage(w, message.OriginalMessage, ":x: Session is expired", "", message.Channel)
		return
	}

	action := message.Actions[0]
	s.Log.Infow("Request verified", zap.String("action", action.Name))

	switch action.Name {
	case actionSelect:
		value := action.SelectedOptions[0].Value
		s.Log.Infow("Update Session", zap.String("callbackID", message.CallbackID), zap.String("value", value))
		sess.value = value
		s.Session.Set(message.CallbackID, sess)

		if q.Confirm {
			// Overwrite original drop down message.
			originalMessage := message.OriginalMessage
			originalMessage.Attachments[0].Text =
				fmt.Sprintf("OK to select %s ?", strings.Title(sess.value))
			originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
				{
					Name:  actionStart,
					Text:  "Yes",
					Type:  "button",
					Value: "start",
					Style: "primary",
				},
				{
					Name:  actionCancel,
					Text:  "No",
					Type:  "button",
					Style: "danger",
				},
			}

			w.Header().Add("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&originalMessage)

			if s.sqsMode {
				_, _, _, err := s.API.SendMessage(
					message.Channel.ID,
					slack.MsgOptionUpdate(originalMessage.Timestamp),
					slack.MsgOptionAttachments(originalMessage.Attachments[0]),
					slack.MsgOptionText(originalMessage.Text, false),
				)
				if err != nil {
					s.Log.Error(err)
				}
			}
			return
		} else {
			err := s.SendRequest(q, sess)
			if err != nil {
				s.Log.Errorw("Send request failed", zap.Error(err))
				s.responseMessage(w, message.OriginalMessage, ":x: "+err.Error(), "", message.Channel)
			} else {
				title := fmt.Sprintf(":ok: @%s start this, %s", message.User.Name, sess.value)
				s.responseMessage(w, message.OriginalMessage, title, "", message.Channel)
			}
			return
		}
	case actionStart: // 3. OK button
		err = s.SendRequest(q, sess)
		if err != nil {
			s.Log.Errorw("Send request failed", zap.Error(err))
			s.responseMessage(w, message.OriginalMessage, ":x: "+err.Error(), "", message.Channel)
		} else {
			title := fmt.Sprintf(":ok: @%s confirmed, %s", message.User.Name, (sess).value)
			s.responseMessage(w, message.OriginalMessage, title, "", message.Channel)
		}
		return
	case actionCancel: // 3. Cancel button
		s.Log.Infow("Request canceled", zap.String("Session ID", sess.id))
		title := fmt.Sprintf(":x: @%s canceled the request", message.User.Name)
		s.responseMessage(w, message.OriginalMessage, title, "", message.Channel)
		return
	default:
		s.Log.Errorf("Invalid action was submitted: %s", action.Name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *SlackBot) responseMessage(w http.ResponseWriter, original slack.Message, title, value string, channel slack.Channel) {
	original.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
	original.Attachments[0].Fields = []slack.AttachmentField{
		{
			Title: title,
			Value: value,
			Short: false,
		},
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&original)

	if s.sqsMode {
		_, _, _, err := s.API.SendMessage(
			channel.ID,
			slack.MsgOptionUpdate(original.Timestamp),
			slack.MsgOptionAttachments(original.Attachments[0]),
			slack.MsgOptionText(original.Text, false),
		)
		if err != nil {
			s.Log.Error(err)
		}
	}
}

func (s *SlackBot) ProcessNonInteractiveRequest(c *domain.Config, sess *SessionValue, channel string) error {
	err := s.SendRequest(c, sess)
	if err != nil {
		_, _, cause := s.API.PostMessage(channel,
			":x: "+err.Error(),
			slack.PostMessageParameters{})
		if cause != nil {
			return errors.Wrap(cause, "post message failed")
		}
	} else {
		text, err := c.TextCompile(sess.matched)
		if err != nil {
			return err
		}

		_, _, cause := s.API.PostMessage(channel, text, slack.PostMessageParameters{})
		if cause != nil {
			return errors.Wrap(cause, "post message failed")
		}
	}
	return nil
}

func (s *SlackBot) ProcessInteractiveRequest(c *domain.Config, sess *SessionValue, channel string) error {
	opt := make([]slack.AttachmentActionOption, 0)
	for _, a := range c.Actions {
		opt = append(opt, slack.AttachmentActionOption{
			Text:  a,
			Value: a,
		})
	}
	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			slack.Attachment{
				Text:       "Select your choice",
				Color:      "#f9a41b",
				CallbackID: c.CallbackID + "@" + sess.id,
				Actions: []slack.AttachmentAction{
					{
						Name:    actionSelect,
						Type:    seletType,
						Options: opt,
					},
					{
						Name:  actionCancel,
						Text:  "Cancel",
						Type:  "button",
						Style: "danger",
					},
				},
			},
		},
	}

	text, err := c.TextCompile(sess.matched)
	if err != nil {
		return err
	}

	_, _, err = s.API.PostMessage(channel, text, params)
	if err != nil {
		return errors.Wrap(err, "post message failed")
	}
	s.Log.Info("Response posted")
	return nil
}

func (s *SlackBot) SendRequest(c *domain.Config, sess *SessionValue) error {
	url, err := c.URLCompile(sess.value, sess.matched, c.Secrets)
	if err != nil {
		return err
	}

	body, err := c.BodyCompile(sess.value, sess.matched, c.Secrets)
	if err != nil {
		return err
	}

	s.Log.Infow("Send Request",
		zap.String("url", url),
		zap.String("body", body),
	)

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(body)),
	)
	if err != nil {
		return errors.Wrap(err, "Cannot make request")
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("POST request failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		s.Log.Infof("Send request failed status: %d", resp.StatusCode)
		return errors.Errorf("Send request failed status: %d", resp.StatusCode)
	}
	s.Log.Info("Send request success")
	return nil
}

func (s *SlackBot) getChannelName(channelID string) (string, error) {
	if id, ok := s.channelCache[channelID]; ok {
		return id, nil
	}

	ch, err := s.API.GetConversationInfo(channelID, false)
	if err != nil {
		return "", err
	}
	return ch.Name, nil
}

func (s *SlackBot) Run() error {
	rtm := s.API.NewRTM()
	go rtm.ManageConnection()

	auth, err := s.API.AuthTest()
	if err != nil {
		return err
	}
	bot, err := s.API.GetUserInfo(auth.UserID)
	if err != nil {
		return err
	}
	s.Log.Debugw("Firestarter bot ID", zap.String("ID", bot.Profile.BotID))

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			s.Log.Info("Hello Event")
		case *slack.MessageEvent:
			s.Log.Debugw("Message bot ID", zap.String("ID", ev.Msg.BotID))
			if ev.Msg.BotID == bot.Profile.BotID {
				break
			}
			// Get config on each event, it may be updated.
			config, err := s.ConfigRepository.GetConfigList()
			if err != nil {
				return err
			}

			name, err := s.getChannelName(ev.Msg.Channel)
			if err != nil {
				return err
			}

			message := ev.Msg.Text
			if message == "" { // IFTTT message with title
				message = ev.Msg.Attachments[0].Text
			}
			if message == "" { // IFTTT message only
				message = ev.Msg.Attachments[0].Pretext
			}

			s.Log.Debugw("Message to be parsed", zap.String("message", message))
			c := config.FindMatched(name, message)
			if c == nil {
				break
			}
			s.Log.Infow("RTM Match", zap.String("id", c.CallbackID),
				zap.String("regexp", c.Regexp.String()),
				zap.String("message", ev.Msg.Text),
			)

			// Create Session for matched request
			sess := s.Session.Create(c.Regexp.FindStringSubmatch(ev.Msg.Text))
			s.Log.Infow("Create Session", zap.String("SessionID", sess.id))

			// No Action means non interactive request
			if len(c.Actions) == 0 {
				err := s.ProcessNonInteractiveRequest(c, sess, ev.Channel)
				if err != nil {
					return errors.Wrap(err, "process non interactive")
				}
			} else {
				err := s.ProcessInteractiveRequest(c, sess, ev.Channel)
				if err != nil {
					return errors.Wrap(err, "process interactive")
				}
			}
		case *slack.InvalidAuthEvent:
			return errors.New("Invalid credentials")
		}
	}
	return errors.New("Never happen")
}
