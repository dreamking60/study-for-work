package server

import (
	"edge-gateway/internal/apperr"
	"edge-gateway/internal/protocol"
	"edge-gateway/internal/room"
	"edge-gateway/internal/session"
	"errors"
	"testing"
)

func TestRouteJoinMessage(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if err := Route(msg, sess, roomSvc); err != nil {
		t.Fatalf("route join message: %v", err)
	}
}

func TestRouteJoinWithValidToken(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if err := Route(msg, sess, roomSvc); err != nil {
		t.Fatalf("route join with valid token: %v", err)
	}

	if sess.UserID != "USER-1" {
		t.Fatalf("expected session user id to be USER-1, got %q", sess.UserID)
	}

	if sess.RoomID != "ROOM-1" {
		t.Fatalf("expected session room id to be ROOM-1, got %q", sess.RoomID)
	}

	if sess.Lifecycle != session.LifecycleJoined {
		t.Fatalf("expected lifecycle to be joined, got %q", sess.Lifecycle)
	}

	if sess.Activity != session.ActivityInactive {
		t.Fatalf("expected activity to be inactive after join, got %q", sess.Activity)
	}

	count, err := roomSvc.MemberCount("ROOM-1")
	if err != nil {
		t.Fatalf("member count for joined room: %v", err)
	}

	if count != 1 {
		t.Fatalf("expected room member count to be 1, got %d", count)
	}
}

func TestRouteJoinWithInvalidTokenDoesNotWriteRoomState(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "bad-token",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	err := Route(msg, sess, roomSvc)
	if err == nil {
		t.Fatal("expected invalid token error, got nil")
	}

	if !errors.Is(err, apperr.ErrUnauthorized) {
		t.Fatalf("expected unauthorized error, got %v", err)
	}

	if sess.UserID != "" {
		t.Fatalf("expected session user id to remain empty, got %q", sess.UserID)
	}

	if sess.RoomID != "" {
		t.Fatalf("expected session room id to remain empty, got %q", sess.RoomID)
	}

	if _, err := roomSvc.MemberCount("ROOM-1"); !errors.Is(err, apperr.ErrRoomNotFound) {
		t.Fatalf("expected room not found after rejected join, got %v", err)
	}
}

func TestRouteChatMessage(t *testing.T) {
	msg := protocol.Message{
		Type:    protocol.MessageTypeChat,
		RoomID:  "ROOM-1",
		UserID:  "USER-1",
		Content: "Hello",
		Token:   "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()
	roomSvc.Join(msg.RoomID, sess)

	if err := Route(msg, sess, roomSvc); err != nil {
		t.Fatalf("route chat message: %v", err)
	}
}

func TestRouteChatMessageUsesSessionRoomID(t *testing.T) {
	msg := protocol.Message{
		Type:    protocol.MessageTypeChat,
		UserID:  "USER-1",
		Content: "Hello",
		Token:   "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
		RoomID: "ROOM-1",
	}
	roomSvc := room.NewService()
	roomSvc.Join(sess.RoomID, sess)

	if err := Route(msg, sess, roomSvc); err != nil {
		t.Fatalf("route chat message with session room id: %v", err)
	}
}

func TestRouteUnknownMessage(t *testing.T) {
	msg := protocol.Message{
		Type:  "unknown",
		Token: "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if err := Route(msg, sess, roomSvc); err == nil {
		t.Fatalf("expected unknown message type error, got nil")
	}
}

func TestRouteRejectsInvalidToken(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "bad-token",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if err := Route(msg, sess, roomSvc); err == nil {
		t.Fatal("expected invalid token error, got nil")
	}
}

func TestRouteRejectsExpiredToken(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "expired-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if err := Route(msg, sess, roomSvc); err == nil {
		t.Fatal("expected expired token error, got nil")
	}
}

func TestRouteHeartbeatUpdatesLastSeenAt(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeHeartbeat,
		UserID: "USER-1",
		Token:  "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if !sess.LastSeenAt.IsZero() {
		t.Fatalf("expected zero LastSeenAt before heartbeat, got %v", sess.LastSeenAt)
	}

	if err := Route(msg, sess, roomSvc); err != nil {
		t.Fatalf("route heartbeat message: %v", err)
	}

	if sess.LastSeenAt.IsZero() {
		t.Fatal("expected LastSeenAt to be updated after heartbeat")
	}

	if sess.Activity != session.ActivityActive {
		t.Fatalf("expected activity to be active after heartbeat, got %q", sess.Activity)
	}
}

func TestRouteHeartbeatAfterJoinPreservesRoomState(t *testing.T) {
	joinMsg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "token-user-1",
	}
	heartbeatMsg := protocol.Message{
		Type:   protocol.MessageTypeHeartbeat,
		UserID: "USER-1",
		Token:  "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if err := Route(joinMsg, sess, roomSvc); err != nil {
		t.Fatalf("route join before heartbeat: %v", err)
	}

	if !sess.LastSeenAt.IsZero() {
		t.Fatalf("expected zero LastSeenAt before heartbeat, got %v", sess.LastSeenAt)
	}

	if err := Route(heartbeatMsg, sess, roomSvc); err != nil {
		t.Fatalf("route heartbeat after join: %v", err)
	}

	if sess.LastSeenAt.IsZero() {
		t.Fatal("expected LastSeenAt to be updated after heartbeat")
	}

	if sess.RoomID != "ROOM-1" {
		t.Fatalf("expected session room id to remain ROOM-1, got %q", sess.RoomID)
	}

	count, err := roomSvc.MemberCount("ROOM-1")
	if err != nil {
		t.Fatalf("member count after heartbeat: %v", err)
	}

	if count != 1 {
		t.Fatalf("expected room member count to remain 1, got %d", count)
	}
}

func TestRouteLeaveClearsSessionRoomID(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeLeave,
		UserID: "USER-1",
		Token:  "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
		RoomID: "ROOM-1",
	}
	roomSvc := room.NewService()
	roomSvc.Join(sess.RoomID, sess)

	if err := Route(msg, sess, roomSvc); err != nil {
		t.Fatalf("route leave message: %v", err)
	}

	if sess.RoomID != "" {
		t.Fatalf("expected session room id to be cleared, got %q", sess.RoomID)
	}

	if sess.Lifecycle != session.LifecycleAuthenticated {
		t.Fatalf("expected lifecycle to return to authenticated after leave, got %q", sess.Lifecycle)
	}
}

func TestRouteChatAfterLeaveReturnsInvalidMessage(t *testing.T) {
	joinMsg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "token-user-1",
	}
	leaveMsg := protocol.Message{
		Type:   protocol.MessageTypeLeave,
		UserID: "USER-1",
		Token:  "token-user-1",
	}
	chatMsg := protocol.Message{
		Type:    protocol.MessageTypeChat,
		UserID:  "USER-1",
		Content: "hello after leave",
		Token:   "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	if err := Route(joinMsg, sess, roomSvc); err != nil {
		t.Fatalf("route join before leave: %v", err)
	}

	if err := Route(leaveMsg, sess, roomSvc); err != nil {
		t.Fatalf("route leave before chat: %v", err)
	}

	err := Route(chatMsg, sess, roomSvc)
	if err == nil {
		t.Fatal("expected invalid message error after leave, got nil")
	}

	if !errors.Is(err, apperr.ErrInvalidMessage) {
		t.Fatalf("expected invalid_message error, got %v", err)
	}

	if sess.RoomID != "" {
		t.Fatalf("expected session room id to remain empty after leave, got %q", sess.RoomID)
	}
}

func TestRouteInvalidTokenReturnsUnauthorized(t *testing.T) {
	msg := protocol.Message{
		Type:   protocol.MessageTypeJoin,
		RoomID: "ROOM-1",
		UserID: "USER-1",
		Token:  "bad-token",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	err := Route(msg, sess, roomSvc)
	if err == nil {
		t.Fatal("expected invalid token error, got nil")
	}

	if !errors.Is(err, apperr.ErrUnauthorized) {
		t.Fatalf("expected unauthorized error, got %v", err)
	}
}

func TestRouteUnknownMessageReturnsInvalidMessage(t *testing.T) {
	msg := protocol.Message{
		Type:  "unknown",
		Token: "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	err := Route(msg, sess, roomSvc)
	if err == nil {
		t.Fatal("expected unknown message type error, got nil")
	}

	if !errors.Is(err, apperr.ErrInvalidMessage) {
		t.Fatalf("expected invalid_message error, got %v", err)
	}
}

func TestRouteChatWithoutRoomReturnsRoomNotFound(t *testing.T) {
	msg := protocol.Message{
		Type:    protocol.MessageTypeChat,
		RoomID:  "ROOM-404",
		UserID:  "USER-1",
		Content: "hello",
		Token:   "token-user-1",
	}

	sess := &session.Session{
		ConnID: "CONN-1",
	}
	roomSvc := room.NewService()

	err := Route(msg, sess, roomSvc)
	if err == nil {
		t.Fatal("expected room not found error, got nil")
	}

	if !errors.Is(err, apperr.ErrRoomNotFound) {
		t.Fatalf("expected room_not_found error, got %v", err)
	}
}
