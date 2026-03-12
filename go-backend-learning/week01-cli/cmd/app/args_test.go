package main

import (
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
