# Week 3 Report: edge-gateway

## Current Progress

I continued the `edge-gateway` project after the minimal Week 2 version and completed the work from Week 3 Day 1 to Day 6.

### Day 1: Minimal Auth Plan

I chose the simplest authentication plan for the current stage:

- use static tokens instead of JWT
- validate token at the router entrypoint
- store authentication result in `Session`

The goal of this step was not to build a full auth system, but to make the gateway reject unauthorized business messages before they enter handlers.

### Day 2: Static Token Validation

I added a minimal auth module:

- `internal/auth/auth.go`
- `ValidateToken(token string) (Claims, error)`

The current static token table includes:

- `token-user-1 -> USER-1`
- `token-user-2 -> USER-2`
- `expired-user-1 -> USER-1` as an expired token example

I also extended the session model:

```go
type Session struct {
	ConnID     string
	UserID     string
	RoomID     string
	Token      string
	ExpiresAt  time.Time
	LastSeenAt time.Time
}
```

Current authentication flow:

```text
message carries token
-> Route(...)
-> validateAuth(...)
-> auth.ValidateToken(...)
-> update session auth state
-> dispatch to business handler
```

### Day 3: Additional Message Types

I extended the protocol and router with more business message types:

- `heartbeat`
- `leave`
- `room-info` as a minimal room query capability

Current protocol message types:

```go
const (
	MessageTypeJoin      MessageType = "join"
	MessageTypeChat      MessageType = "chat"
	MessageTypeHeartbeat MessageType = "heartbeat"
	MessageTypeLeave     MessageType = "leave"
	MessageTypeRoomInfo  MessageType = "room-info"
)
```

Current additions:

- `handleHeartbeat(...)` updates `sess.LastSeenAt = time.Now()`
- `room.Service.Leave(roomID, connID) error`
- `handleLeave(...)` removes the session from the room and clears `sess.RoomID`
- `room.Service.MemberCount(roomID) (int, error)` as a minimal room query capability

### Day 4: Store Abstraction

I moved room and session state behind dedicated stores instead of letting higher-level logic depend on raw maps.

Added:

- `internal/store/room_store.go`
- `internal/store/session_store.go`

Current structure:

- `room.Service` now depends on `RoomStore`
- `room.Service` also writes joined sessions into `SessionStore`
- `Disconnect(connID)` removes state from both room membership and session storage

This changes the service layer from direct map management to explicit state access through stores.

### Day 5: Unified Error Model

I introduced a minimal shared error package:

- `internal/apperr/errors.go`

Current stable error categories:

- `ErrUnauthorized`
- `ErrInvalidMessage`
- `ErrRoomNotFound`
- `ErrSessionNotFound`

I then updated auth, router, and room logic to return wrapped errors based on these categories instead of ad hoc strings.

This makes error behavior more stable and lets tests assert error class with `errors.Is(...)`.

### Day 6: Integration-Style Routing Checks

I added a small set of integration-style router tests to verify the current business flow is coherent across auth, routing, session updates, and room state.

Covered paths:

- `valid token -> join success`
- `invalid token -> join rejected and no room state write`
- `join -> heartbeat -> LastSeenAt updated and room state preserved`
- `join -> leave -> chat rejected because session is no longer in a room`

These tests are still in-process tests, not websocket-level end-to-end tests, but they already validate the current gateway core as a connected flow instead of isolated unit behavior.

## Tests

Auth and routing tests now include:

- `TestRouteJoinWithValidToken`
- `TestRouteJoinWithInvalidTokenDoesNotWriteRoomState`
- `TestRouteRejectsInvalidToken`
- `TestRouteRejectsExpiredToken`
- `TestRouteHeartbeatUpdatesLastSeenAt`
- `TestRouteHeartbeatAfterJoinPreservesRoomState`
- `TestRouteLeaveClearsSessionRoomID`
- `TestRouteChatAfterLeaveReturnsInvalidMessage`
- `TestRouteInvalidTokenReturnsUnauthorized`
- `TestRouteUnknownMessageReturnsInvalidMessage`
- `TestRouteChatWithoutRoomReturnsRoomNotFound`

Room service tests include:

- `TestJoinSavesSessionToSessionStore`
- `TestDisconnectRemovesSessionFromSessionStore`
- `TestLeaveRemovesSessionFromRoom`
- `TestLeaveRemovesEmptyRoom`
- `TestLeaveReturnsErrorForUnknownRoom`
- `TestMemberCountReturnsCount`
- `TestMemberCountReturnsErrorForUnknownRoom`

Verification command:

```bash
GOCACHE=../../../go-backend-learning/edge-gateway/.gocache go test ./...
```

Result on 2026-04-27:

```text
ok  	edge-gateway/internal/protocol
ok  	edge-gateway/internal/room
ok  	edge-gateway/internal/server
```

## Notes

This Week 3 version is still intentionally minimal:

- no JWT yet
- no websocket transport yet
- no real socket write-back yet
- no concurrency protection yet
- `room-info` is still only a service-level query, not a full response path

The main progress of this week is:

```text
the gateway now has a minimal auth layer,
more realistic message types,
store-backed state management,
stable error categories,
and basic integration-style flow verification
```

## Next Step

Week 4 should focus on connection lifecycle and runtime behavior:

- timeout and stale-session cleanup
- stronger disconnect semantics
- concurrency safety around shared state
- eventually connecting the current core to a real `/ws` transport path
