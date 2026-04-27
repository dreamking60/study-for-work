package session

import (
	"time"
)

type Session struct {
	ConnID    	string
	UserID    	string
	RoomID    	string
	Token     	string
	ExpiresAt 	time.Time
	LastSeenAt 	time.Time
	Lifecycle  	LifecycleState
	Activity   	ActivityState
}
