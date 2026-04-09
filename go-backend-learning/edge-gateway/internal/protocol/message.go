package protocol

type MessageType string

const (
	MessageTypeJoin MessageType = "join"
	MessageTypeChat MessageType = "message"
)

type Message struct {
	Type    MessageType `json:"type"`
	RoomID  string      `json:"room_id"`
	UserID  string      `json:"user_id"`
	Content string      `json:"content"`
}
