package apperr

import "errors"

var (
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvalidMessage  = errors.New("invalid_message")
	ErrRoomNotFound    = errors.New("room_not_found")
	ErrSessionNotFound = errors.New("session_not_found")
)
