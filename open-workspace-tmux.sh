#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="/home/dreamking/study-for-work"
SESSION_NAME="${1:-study-for-work}"
WINDOW_NAME="workspace"
LEFT_DIR="$ROOT_DIR"
RIGHT_DIR="$ROOT_DIR/go-backend-learning"

if ! command -v tmux >/dev/null 2>&1; then
  echo "tmux is not installed or not in PATH." >&2
  exit 1
fi

if [ ! -d "$RIGHT_DIR" ]; then
  RIGHT_DIR="$ROOT_DIR"
fi

ensure_layout() {
  if ! tmux list-windows -t "$SESSION_NAME" -F '#{window_name}' | grep -qx "$WINDOW_NAME"; then
    tmux new-window -t "$SESSION_NAME" -n "$WINDOW_NAME" -c "$LEFT_DIR"
  fi

  local pane_count
  pane_count="$(tmux list-panes -t "${SESSION_NAME}:${WINDOW_NAME}" | wc -l | tr -d ' ')"
  if [ "$pane_count" -lt 2 ]; then
    tmux split-window -h -t "${SESSION_NAME}:${WINDOW_NAME}" -c "$RIGHT_DIR"
  fi

  tmux select-layout -t "${SESSION_NAME}:${WINDOW_NAME}" even-horizontal
  tmux select-pane -t "${SESSION_NAME}:${WINDOW_NAME}.1"
}

if tmux has-session -t "$SESSION_NAME" 2>/dev/null; then
  ensure_layout
  if [ -n "${TMUX:-}" ]; then
    tmux switch-client -t "$SESSION_NAME"
  else
    tmux attach-session -t "$SESSION_NAME"
  fi
  exit 0
fi

tmux new-session -d -s "$SESSION_NAME" -n "$WINDOW_NAME" -c "$LEFT_DIR"
ensure_layout

if [ -n "${TMUX:-}" ]; then
  tmux switch-client -t "$SESSION_NAME"
else
  tmux attach-session -t "$SESSION_NAME"
fi
