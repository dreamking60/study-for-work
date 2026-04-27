package room

import (
	"edge-gateway/internal/session"
	"testing"
)

func TestJoinRoom(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}

	rm := sv.Join(roomID, sess)

	if rm == nil {
		t.Fatal("expected room, got nil")
	}

	if rm.ID != roomID {
		t.Fatalf("unexpected room id: got %q", rm.ID)
	}

	got, ok := rm.Members[sess.ConnID]
	if !ok {
		t.Fatalf("expected member %q in room", sess.ConnID)
	}

	if got != sess {
		t.Fatalf("unexpected session stored in room")
	}
}

func TestJoinSetsSessionRoomID(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess := &session.Session{
		ConnID: "CONN-2",
		UserID: "USER-2",
	}

	sv.Join(roomID, sess)

	if sess.RoomID != roomID {
		t.Fatalf("unexpected session room id: got %q", sess.RoomID)
	}

	if sess.Activity != session.ActivityInactive {
		t.Fatalf("expected session activity to default to inactive, got %q", sess.Activity)
	}
}

func TestJoinSavesSessionToSessionStore(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess := &session.Session{
		ConnID: "CONN-3",
		UserID: "USER-3",
	}

	sv.Join(roomID, sess)

	got, ok := sv.sessionStore.Get(sess.ConnID)
	if !ok {
		t.Fatalf("expected session %q in session store", sess.ConnID)
	}

	if got != sess {
		t.Fatalf("unexpected session stored in session store")
	}
}

func TestBroadcastTargetsMembersInSameRoom(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess1 := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}
	sess2 := &session.Session{
		ConnID: "CONN-2",
		UserID: "USER-2",
	}

	sv.Join(roomID, sess1)
	sv.Join(roomID, sess2)

	targets := sv.BroadcastTargets(roomID)
	if len(targets) != 2 {
		t.Fatalf("unexpected target count: got %d", len(targets))
	}

	got := make(map[string]bool, len(targets))
	for _, sess := range targets {
		got[sess.ConnID] = true
	}

	if !got[sess1.ConnID] {
		t.Fatalf("expected target %q in broadcast targets", sess1.ConnID)
	}

	if !got[sess2.ConnID] {
		t.Fatalf("expected target %q in broadcast targets", sess2.ConnID)
	}
}

func TestBroadcastDoesNotLeakToOtherRooms(t *testing.T) {
	sv := NewService()
	room1 := "ROOM-1"
	room2 := "ROOM-2"
	sess1 := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}
	sess2 := &session.Session{
		ConnID: "CONN-2",
		UserID: "USER-2",
	}

	sv.Join(room1, sess1)
	sv.Join(room2, sess2)

	targets := sv.BroadcastTargets(room1)
	if len(targets) != 1 {
		t.Fatalf("unexpected target count: got %d", len(targets))
	}

	if targets[0].ConnID != sess1.ConnID {
		t.Fatalf("unexpected broadcast target: got %q", targets[0].ConnID)
	}
}

func TestDisconnectRemovesSessionFromRoom(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess1 := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}
	sess2 := &session.Session{
		ConnID: "CONN-2",
		UserID: "USER-2",
	}

	sv.Join(roomID, sess1)
	sv.Join(roomID, sess2)

	sv.Disconnect(sess1.ConnID)

	rm, ok := sv.store.Get(roomID)
	if !ok {
		t.Fatal("expected room to still exist")
	}

	if _, ok := rm.Members[sess1.ConnID]; ok {
		t.Fatalf("expected session %q to be removed", sess1.ConnID)
	}

	if _, ok := rm.Members[sess2.ConnID]; !ok {
		t.Fatalf("expected session %q to remain", sess2.ConnID)
	}
}

func TestEmptyRoomCleanup(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}

	sv.Join(roomID, sess)
	sv.Disconnect(sess.ConnID)

	if _, ok := sv.store.Get(roomID); ok {
		t.Fatalf("expected empty room %q to be deleted", roomID)
	}

}

func TestDisconnectRemovesSessionFromSessionStore(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess := &session.Session{
		ConnID: "CONN-3",
		UserID: "USER-3",
	}

	sv.Join(roomID, sess)
	sv.Disconnect(sess.ConnID)

	if _, ok := sv.sessionStore.Get(sess.ConnID); ok {
		t.Fatalf("expected session %q to be deleted from session store", sess.ConnID)
	}

	if sess.Lifecycle != session.LifecycleDisconnected {
		t.Fatalf("expected lifecycle to be disconnected, got %q", sess.Lifecycle)
	}

	if sess.Activity != session.ActivityInactive {
		t.Fatalf("expected activity to be inactive after disconnect, got %q", sess.Activity)
	}
}

func TestLeaveRemovesSessionFromRoom(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess1 := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}
	sess2 := &session.Session{
		ConnID: "CONN-2",
		UserID: "USER-2",
	}

	sv.Join(roomID, sess1)
	sv.Join(roomID, sess2)

	if err := sv.Leave(roomID, sess1.ConnID); err != nil {
		t.Fatalf("leave room: %v", err)
	}

	rm, ok := sv.store.Get(roomID)
	if !ok {
		t.Fatal("expected room to still exist")
	}

	if _, ok := rm.Members[sess1.ConnID]; ok {
		t.Fatalf("expected session %q to be removed", sess1.ConnID)
	}

	if _, ok := rm.Members[sess2.ConnID]; !ok {
		t.Fatalf("expected session %q to remain", sess2.ConnID)
	}
}

func TestLeaveRemovesEmptyRoom(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}

	sv.Join(roomID, sess)

	if err := sv.Leave(roomID, sess.ConnID); err != nil {
		t.Fatalf("leave room: %v", err)
	}

	if _, ok := sv.store.Get(roomID); ok {
		t.Fatalf("expected empty room %q to be deleted", roomID)
	}
}

func TestLeaveReturnsErrorForUnknownRoom(t *testing.T) {
	sv := NewService()

	if err := sv.Leave("ROOM-404", "CONN-1"); err == nil {
		t.Fatal("expected leave error for unknown room, got nil")
	}
}

func TestMemberCountReturnsCount(t *testing.T) {
	sv := NewService()
	roomID := "ROOM-1"
	sess1 := &session.Session{
		ConnID: "CONN-1",
		UserID: "USER-1",
	}
	sess2 := &session.Session{
		ConnID: "CONN-2",
		UserID: "USER-2",
	}

	sv.Join(roomID, sess1)
	sv.Join(roomID, sess2)

	got, err := sv.MemberCount(roomID)
	if err != nil {
		t.Fatalf("member count: %v", err)
	}

	if got != 2 {
		t.Fatalf("unexpected member count: got %d", got)
	}
}

func TestMemberCountReturnsErrorForUnknownRoom(t *testing.T) {
	sv := NewService()

	if _, err := sv.MemberCount("ROOM-404"); err == nil {
		t.Fatal("expected member count error for unknown room, got nil")
	}
}
