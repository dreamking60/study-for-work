package session

type LifecycleState string

const (
	LifecycleConnected     LifecycleState = "connected"
	LifecycleAuthenticated LifecycleState = "authenticated"
	LifecycleJoined        LifecycleState = "joined"
	LifecycleDisconnected  LifecycleState = "disconnected"
)

type ActivityState string

const (
	ActivityActive   ActivityState = "active"
	ActivityInactive ActivityState = "inactive"
)
