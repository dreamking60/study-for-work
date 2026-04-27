# Session State Model

## Purpose

This document defines the current session lifecycle and activity model for `edge-gateway`.

The session state is intentionally split into two dimensions:

- lifecycle state
- activity state

This avoids forcing `active` and `inactive` into the same linear chain as connection lifecycle events.

## Lifecycle State

Lifecycle state describes where a session is in the connection and room flow.

Current lifecycle states:

- `connected`
- `authenticated`
- `joined`
- `disconnected`

### Lifecycle Meaning

- `connected`
  A session exists for a connection, but authentication has not completed yet.

- `authenticated`
  Token validation has succeeded, but the session is not currently in a room.

- `joined`
  The session has successfully joined a room.

- `disconnected`
  The connection has been closed or explicitly cleaned up.

## Activity State

Activity state describes whether the session is currently considered active.

Current activity states:

- `active`
- `inactive`

### Activity Meaning

- `active`
  The session has received a recent heartbeat or other activity signal.

- `inactive`
  The session currently has no recent activity signal, or it has already been disconnected.

## Current Transition Rules

### Lifecycle transitions

```text
new session                    -> connected
validateAuth success           -> authenticated
join success                   -> joined
leave success                  -> authenticated
disconnect / cleanup           -> disconnected
```

### Activity transitions

```text
new session                    -> inactive
heartbeat                      -> active
disconnect / cleanup           -> inactive
```

## Current Combined Paths

Typical path:

```text
connected + inactive
-> authenticated + inactive
-> joined + inactive
-> joined + active
```

Leave path:

```text
joined + active|inactive
-> authenticated + active|inactive
```

Disconnect path:

```text
connected|authenticated|joined + active|inactive
-> disconnected + inactive
```

## Notes

- `active` and `inactive` are not treated as lifecycle phases.
- `LastSeenAt` is still kept as the time-based evidence behind activity behavior.
- timeout-driven inactive transitions are not fully implemented yet; this document defines the target model for Week 4 runtime work.
