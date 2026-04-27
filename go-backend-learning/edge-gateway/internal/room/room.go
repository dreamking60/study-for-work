package room

import "edge-gateway/internal/session"

type Room struct {
	ID	string
	Members	map[string]*session.Session
}