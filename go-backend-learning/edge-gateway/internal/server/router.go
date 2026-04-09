package server

import (
	"fmt"

	"edge-gateway/internal/protocol"
	"edge-gateway/internal/session"
)

func Route(msg protocol.Message, sess *session.Session) error {
	switch msg.Type {
	case protocol.MessageTypeJoin, protocol.MessageTypeChat:
		return nil
	default:
		return fmt.Errorf("unsupported message type: %s", msg.Type)
	}
}
