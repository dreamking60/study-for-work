package protocol

import (
	"testing"
)

func TestDecodeJoinMessage(t *testing.T) {
	raw := []byte(`{"type":"join","room_id":"ROOM-1","user_id":"USER-1"}`)

	msg, err := DecodeMessage(raw)

	if err != nil {
		t.Fatalf("decode join message: %v", err)
	}

	if msg.Type != MessageTypeJoin {
		t.Fatalf("unexpected type: got %q", msg.Type)
	}

	if msg.RoomID != "ROOM-1" {
		t.Fatalf("unexpected room id: got %q", msg.RoomID)
	}

	if msg.UserID != "USER-1" {
		t.Fatalf("unexpected user id: got %q", msg.UserID)
	}
}

func TestDecodeUnknownMessageType(t *testing.T) {
	raw := []byte(`{"type":"any","room_id":"ROOM-1","user_id":"USER-1"}`)

	msg, err := DecodeMessage(raw)

	if err == nil {
		t.Fatalf("expected unknown message type error, got nil with type %q", msg.Type)
	}

}