package main

import (
	"os"
	"strings"
	"testing"
)

func TestRunCommand_Success(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() {
		os.Args = oldArgs
	})

	os.Args = []string{"app", "run", "--env", "dev", "--port", "9000", "--app", "cli-demo"}

	if err := run(); err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}
}

func TestTaskSummaryCommand_Success(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() {
		os.Args = oldArgs
	})

	os.Args = []string{"app", "task-summary"}

	if err := run(); err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}
}

func TestRun_MissingSubcommand(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() {
		os.Args = oldArgs
	})

	os.Args = []string{"app"}

	err := run()
	if err == nil {
		t.Fatal("expected error for missing subcommand, got nil")
	}

	if !strings.Contains(err.Error(), "missing subcommand") {
		t.Fatalf("expected missing subcommand error, got %q", err.Error())
	}
}

func TestRun_UnknownSubcommand(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() {
		os.Args = oldArgs
	})

	os.Args = []string{"app", "unknown"}

	err := run()
	if err == nil {
		t.Fatal("expected error for unknown subcommand, got nil")
	}

	if !strings.Contains(err.Error(), "unknown subcommand") {
		t.Fatalf("expected unknown subcommand error, got %q", err.Error())
	}
}
