package room

import (
	"edge-gateway/internal/apperr"
	"edge-gateway/internal/session"
	"edge-gateway/internal/store"
	"fmt"
)

type Service struct {
	store        *store.RoomStore[Room]
	sessionStore *store.SessionStore
}

func NewService() *Service {
	return &Service{
		store:        store.NewRoomStore[Room](),
		sessionStore: store.NewSessionStore(),
	}
}

func (s *Service) Join(roomID string, sess *session.Session) *Room {
	room, ok := s.store.Get(roomID)
	if !ok {
		room = &Room{
			ID:      roomID,
			Members: map[string]*session.Session{},
		}
		s.store.Save(room.ID, room)
	}

	sess.RoomID = roomID
	if sess.Lifecycle == "" {
		sess.Lifecycle = session.LifecycleConnected
	}
	if sess.Activity == "" {
		sess.Activity = session.ActivityInactive
	}
	s.sessionStore.Save(sess)
	room.Members[sess.ConnID] = sess
	return room
}

func (s *Service) BroadcastTargets(roomID string) []*session.Session {
	room, ok := s.store.Get(roomID)
	if !ok {
		return nil
	}

	out := make([]*session.Session, 0, len(room.Members))
	for _, sess := range room.Members {
		out = append(out, sess)
	}

	return out
}

func (s *Service) Disconnect(connID string) {
	if sess, ok := s.sessionStore.Get(connID); ok {
		sess.RoomID = ""
		sess.Lifecycle = session.LifecycleDisconnected
		sess.Activity = session.ActivityInactive
	}
	s.sessionStore.Delete(connID)

	for roomID, room := range s.store.All() {
		delete(room.Members, connID)
		if len(room.Members) == 0 {
			s.store.Delete(roomID)
		}
	}
}

func (s *Service) Leave(roomID, connID string) error {
	room, ok := s.store.Get(roomID)
	if !ok {
		return fmt.Errorf("%w: %q", apperr.ErrRoomNotFound, roomID)
	}
	delete(room.Members, connID)
	if len(room.Members) == 0 {
		s.store.Delete(roomID)
	}
	if sess, ok := s.sessionStore.Get(connID); ok {
		sess.RoomID = ""
		sess.Lifecycle = session.LifecycleAuthenticated
	}

	return nil
}

func (s *Service) MemberCount(roomID string) (int, error) {
	room, ok := s.store.Get(roomID)
	if !ok {
		return 0, fmt.Errorf("%w: %q", apperr.ErrRoomNotFound, roomID)
	}

	return len(room.Members), nil
}
