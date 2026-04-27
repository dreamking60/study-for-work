package store

import (
	"edge-gateway/internal/session"
)

type SessionStore struct {
	sessions map[string]*session.Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: map[string]*session.Session{},
	}
}

func (s *SessionStore) Save(sess *session.Session) {
	s.sessions[sess.ConnID] = sess
}

func (s *SessionStore) Get(connID string) (*session.Session, bool) {
	sess, ok := s.sessions[connID]
	return sess, ok
}

func (s *SessionStore) Delete(connID string) {
	delete(s.sessions, connID)
}

func (s *SessionStore) All() map[string]*session.Session {
	return s.sessions
}
