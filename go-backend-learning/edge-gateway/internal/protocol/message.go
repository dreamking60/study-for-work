package protocol

import (
	"encoding/json"
	"fmt"
)

type MessageType string 
const (
	MessageTypeJoin	MessageType = "join"
	MessageTypeChat MessageType = "chat"
	MessageTypeHeartbeat MessageType = "heartbeat"
	MessageTypeLeave MessageType = "leave"
	MessageTypeRoomInfo MessageType = "room-info"
)

type Message struct {
	Type	MessageType	`json:"type"`
	RoomID	string		`json:"room_id"`
	UserID	string		`json:"user_id"`
	Content string		`json:"content"`
	Token	string		`json:"token"`
}

func DecodeMessage(raw []byte) (Message, error) {
	var msg Message
	if err := json.Unmarshal(raw, &msg); err != nil {
		return Message{}, err
	}
	if !ValidateMessageType(msg.Type) {
		return Message{}, fmt.Errorf("unknown message type: %q", msg.Type)
	}

	return msg, nil
}

func ValidateMessageType(msgType MessageType) bool {
	switch msgType {
	case MessageTypeJoin, MessageTypeChat, MessageTypeHeartbeat, MessageTypeLeave, MessageTypeRoomInfo:
		return true
	default:
		return false
	}
}
