package main

import (
	"log"
	"os"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/nlopes/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				// Only response mention to bot. Ignore else.
				if strings.HasPrefix(ev.Msg.Text, "Test") {
					break
				}

				pp.Println(ev)
				attachment := slack.Attachment{
					Text:       "Which beer do you want? :beer:",
					Color:      "#f9a41b",
					CallbackID: "beer",
					Actions: []slack.AttachmentAction{
						{
							Name: actionSelect,
							Type: "select",
							Options: []slack.AttachmentActionOption{
								{
									Text:  "Asahi Super Dry",
									Value: "Asahi Super Dry",
								},
								{
									Text:  "Kirin Lager Beer",
									Value: "Kirin Lager Beer",
								},
								{
									Text:  "Sapporo Black Label",
									Value: "Sapporo Black Label",
								},
								{
									Text:  "Suntory Malts",
									Value: "Suntory Malts",
								},
								{
									Text:  "Yona Yona Ale",
									Value: "Yona Yona Ale",
								},
							},
						},

						{
							Name:  actionCancel,
							Text:  "Cancel",
							Type:  "button",
							Style: "danger",
						},
					},
				}

				params := slack.PostMessageParameters{
					Attachments: []slack.Attachment{
						attachment,
					},
				}

				api.PostMessage(ev.Channel, "Test", params)
				//rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

type BothubConfig interface {
	Match(message, channelID string) bool
	Run() error
}

type InteractiveConfig struct {
}

func main() {
	token := os.Getenv("BOTHUB_TOKEN")
	api := slack.New(token)
	os.Exit(run(api))
}
