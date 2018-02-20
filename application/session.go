package application

import (
	"strings"
	"time"

	"github.com/juntaki/expiresync"
	"github.com/rs/xid"
)

type Session struct {
	store  *expiresync.Map
	expire time.Duration
}

const expire = 1 * time.Hour

func NewSession() *Session {
	return &Session{
		store:  expiresync.NewMap(),
		expire: expire,
	}
}

type SessionValue struct {
	matched []string
	value   string
	id      string
}

func (s *Session) Get(callbackID string) (*SessionValue, bool) {
	sessionID := s.getSessionID(callbackID)
	sess, ok := s.store.Get(sessionID)
	if !ok {
		return nil, false
	}
	return sess.(*SessionValue), true
}

func (s *Session) Set(callbackID string, sess *SessionValue) {
	sessionID := s.getSessionID(callbackID)
	s.store.Set(sessionID, sess, s.expire)
}

func (s *Session) Create(matched []string) *SessionValue {
	sessionID := xid.New().String()
	sess := &SessionValue{
		matched: matched,
		value:   "",
		id:      sessionID,
	}
	s.store.Set(sessionID, sess, s.expire)
	return sess
}

func (s *Session) getSessionID(callbackID string) string {
	return strings.Split(callbackID, "@")[1]
}
