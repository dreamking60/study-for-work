# Week 2 Report: edge-gateway

## Current Progress

I restarted the `edge-gateway` project from Day 1 and completed the work from Day 1 to Day 6. Day 7 is used as the weekly summary for this first minimal version.

### Day 1: Project Skeleton

Completed:
- Created the `edge-gateway` Go module.
- Added the server entrypoint at `cmd/server/main.go`.
- Added the basic internal package layout:
  - `internal/protocol`
  - `internal/session`
  - `internal/room`
  - `internal/server`
- Implemented a minimal HTTP server using `net/http` and `http.NewServeMux`.

The current server starts with:

```go
fmt.Println("edge-gateway starting on :8080")
```

### Day 2: Message Protocol

I defined the first version of the client-to-server protocol.

Supported message types:
- `join`: user joins a room
- `chat`: user sends a chat message to a room

Current protocol model:

```go
type MessageType string

const (
	MessageTypeJoin MessageType = "join"
	MessageTypeChat MessageType = "chat"
)

type Message struct {
	Type    MessageType `json:"type"`
	RoomID  string      `json:"room_id"`
	UserID  string      `json:"user_id"`
	Content string      `json:"content"`
}
```

I also added `DecodeMessage(raw []byte) (Message, error)`.

It is responsible for:
- Decoding JSON into `Message`
- Rejecting unknown message types
- Returning a stable error when `type` is not `join` or `chat`

### Day 3: Router

I added a minimal message router in `internal/server/router.go`.

Current routing behavior:
- `join` -> `handleJoin`
- `chat` -> `handleChat`
- unknown type -> return error

This step established the business entrypoint after protocol decoding.

### Day 4: Room Join Flow

I added `room.Service` as the in-memory room manager.

Current service structure:

```go
type Service struct {
	rooms map[string]*Room
}
```

Current join flow:
- Find room by `room_id`
- Create room if it does not exist
- Set `sess.RoomID`
- Add the current session into `room.Members`

### Day 5: Broadcast Target Selection

I added the first in-memory broadcast target selection logic:

```go
func (s *Service) BroadcastTargets(roomID string) []*session.Session
```

Current behavior:
- Find the target room
- Return all online sessions in the room
- Do not send network messages yet

This is intentional. At this stage, the gateway only determines:

```text
who should receive the chat message
```

Actual socket or websocket write-back is not implemented yet.

I also updated `handleChat(...)` so that chat routing now depends on room state and room member selection, instead of being an empty placeholder.

### Day 6: Disconnect Cleanup

I added connection lifecycle cleanup to `room.Service`:

```go
func (s *Service) Disconnect(connID string)
```

Current behavior:
- remove the disconnected connection from room members
- delete empty rooms from the in-memory room registry

This completes the first minimal lifecycle loop:

```text
join -> chat target selection -> disconnect cleanup
```

## Tests

Added protocol tests:
- `TestDecodeJoinMessage`
- `TestDecodeUnknownMessageType`

Added router tests:
- `TestRouteJoinMessage`
- `TestRouteChatMessage`
- `TestRouteChatMessageUsesSessionRoomID`
- `TestRouteUnknownMessage`

Added room service tests:
- `TestJoinRoom`
- `TestJoinSetsSessionRoomID`
- `TestBroadcastTargetsMembersInSameRoom`
- `TestBroadcastDoesNotLeakToOtherRooms`
- `TestDisconnectRemovesSessionFromRoom`
- `TestEmptyRoomCleanup`

Verification command:

```bash
GOCACHE=../../../go-backend-learning/edge-gateway/.gocache go test ./...
```

Result:

```text
ok  	edge-gateway/internal/protocol
ok  	edge-gateway/internal/room
ok  	edge-gateway/internal/server
```

## Notes

I changed the tutorial wording from `message` to `chat` for the chat action because `chat` is clearer for this project protocol.

The current broadcast design is still the minimal version. It returns a slice of sessions as broadcast targets. This is enough for the current week, but it may be refactored later for large-room performance and lower allocation cost.

## Day 7 Summary

### What I borrowed from nano

I used `nano` mainly as a design reference, not as a framework dependency.

Borrowed ideas:
- use `session` as per-connection context
- use `room` as the boundary of group membership
- use typed message routing for business actions such as `join` and `chat`
- treat disconnect cleanup as part of the room/session lifecycle

### How my code maps to component / handler / session / group

- `room.Service` corresponds to the room management component
- `handleJoin` and `handleChat` correspond to message handlers
- `session.Session` corresponds to connection context
- `room.Members` corresponds to the group/member collection

### What this version intentionally does not include

- no real websocket or socket write-back yet
- no authentication
- no Redis or database
- no heartbeat or timeout cleanup
- no concurrency safety yet
- no cross-node routing

This version is intentionally kept as the minimum in-memory gateway so that protocol, routing, room membership, and disconnect cleanup are all easy to understand and verify.

## Next Step

Week 3 will move from the minimal gateway to a more business-capable version:
- add basic auth or session validation
- extend message types
- introduce store abstraction
- improve error and log structure
