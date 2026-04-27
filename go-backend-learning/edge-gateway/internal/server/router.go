package server

import (
	"edge-gateway/internal/apperr"
	"edge-gateway/internal/auth"
	"edge-gateway/internal/protocol"
	"edge-gateway/internal/room"
	"edge-gateway/internal/session"
	"fmt"
	"time"
)

func Route(msg protocol.Message, sess *session.Session, roomSvc *room.Service) error {
	if err := validateAuth(msg, sess); err != nil {
		return err
	}

	switch msg.Type {
	case protocol.MessageTypeJoin:
		// handle join
		return handleJoin(msg, sess, roomSvc)
	case protocol.MessageTypeChat:
		// handle chat
		return handleChat(msg, sess, roomSvc)
	case protocol.MessageTypeHeartbeat:
		return handleHeartbeat(msg, sess)
	case protocol.MessageTypeLeave:
		return handleLeave(msg, sess, roomSvc)
	default:
		return fmt.Errorf("%w: type=%q", apperr.ErrInvalidMessage, msg.Type)
	}
}

func handleJoin(msg protocol.Message, sess *session.Session, roomSvc *room.Service) error {
	roomID := msg.RoomID
	if roomID == "" {
		return fmt.Errorf("%w: room_id is empty", apperr.ErrInvalidMessage)
	}

	roomSvc.Join(roomID, sess)
	sess.Lifecycle = session.LifecycleJoined
	return nil
}

func handleChat(msg protocol.Message, sess *session.Session, roomSvc *room.Service) error {
	roomID := msg.RoomID
	if roomID == "" {
		roomID = sess.RoomID
	}
	if roomID == "" {
		return fmt.Errorf("%w: room_id is empty", apperr.ErrInvalidMessage)
	}

	targets := roomSvc.BroadcastTargets(roomID)
	if targets == nil {
		return fmt.Errorf("%w: %q", apperr.ErrRoomNotFound, roomID)
	}
	for _, target := range targets {
		_ = target
	}

	return nil
}

func handleHeartbeat(msg protocol.Message, sess *session.Session) error {
	sess.LastSeenAt = time.Now()
	sess.Activity = session.ActivityActive
	return nil
}

func handleLeave(msg protocol.Message, sess *session.Session, roomSvc *room.Service) error {
	roomID := msg.RoomID
	if roomID == "" {
		roomID = sess.RoomID
	}
	if roomID == "" {
		return fmt.Errorf("%w: room_id is empty", apperr.ErrInvalidMessage)
	}

	if err := roomSvc.Leave(roomID, sess.ConnID); err != nil {
		return err
	}

	sess.RoomID = ""
	sess.Lifecycle = session.LifecycleAuthenticated
	return nil
}

func validateAuth(msg protocol.Message, sess *session.Session) error {
	claims, err := auth.ValidateToken(msg.Token)
	if err != nil {
		return err
	}

	if msg.UserID != "" && msg.UserID != claims.UserID {
		return fmt.Errorf("%w: token user mismatch %q", apperr.ErrUnauthorized, msg.UserID)
	}

	sess.Token = msg.Token
	sess.UserID = claims.UserID
	sess.ExpiresAt = claims.ExpiresAt
	sess.Lifecycle = session.LifecycleAuthenticated
	if sess.Activity == "" {
		sess.Activity = session.ActivityInactive
	}
	return nil
}
