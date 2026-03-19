package main

import (
	"os"
	"strings"
	"testing"
)

func TestParseArgsFrom_Defaults(t *testing.T) {
	args, err := parseArgsFrom([]string{})
	if err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}

	if args.Env != "local" {
		t.Errorf("expected Env %q, got %q", "local", args.Env)
	}

	if args.Port != 8765 {
		t.Errorf("expected Port %d, got %d", 8765, args.Port)
	}

	if args.AppName != "demo" {
		t.Errorf("expected AppName %q, got %q", "demo", args.AppName)
	}
}

func TestParseArgsFrom_CustomValues(t *testing.T) {
	args, err := parseArgsFrom([]string{"--env", "dev", "--port", "8080", "--app", "hello"})
	if err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}

	if args.Env != "dev" {
		t.Errorf("expected Env %q, got %q", "dev", args.Env)
	}

	if args.Port != 8080 {
		t.Errorf("expected Port %d, got %d", 8080, args.Port)
	}

	if args.AppName != "hello" {
		t.Errorf("expected AppName %q, got %q", "hello", args.AppName)
	}
}

func TestParseArgsFrom_InvalidEnv(t *testing.T) {
	_, err := parseArgsFrom([]string{"--env", "stage"})
	expect := "invalid arg env"
	if err == nil {
		t.Fatal("expected error for invalid env, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestParseArgsFrom_InvalidPort(t *testing.T) {
	_, err := parseArgsFrom([]string{"--port", "70000"})
	expect := "invalid arg port"
	if err == nil {
		t.Fatal("expected error for invalid port, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestParseArgsFrom_ParseError(t *testing.T) {
	_, err := parseArgsFrom([]string{"--port", "abc"})
	expect := "parse args"
	if err == nil {
		t.Fatal("expected error for invalid flag value, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestParseArgs_UsesOSArgs(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() {
		os.Args = oldArgs
	})

	os.Args = []string{"app", "run", "--env", "test", "--port", "9000", "--app", "cli-demo"}

	args, err := parseArgs()
	if err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}

	if args.Env != "test" {
		t.Errorf("expected Env %q, got %q", "test", args.Env)
	}

	if args.Port != 9000 {
		t.Errorf("expected Port %d, got %d", 9000, args.Port)
	}

	if args.AppName != "cli-demo" {
		t.Errorf("expected AppName %q, got %q", "cli-demo", args.AppName)
	}
}
