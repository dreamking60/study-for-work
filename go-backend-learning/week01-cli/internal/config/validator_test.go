package config

import (
	"strings"
	"testing"
)

func TestValidate_EmptyAppName(t *testing.T) {
	cfg := Config{
		AppName: "",
		Env:     "local",
		Port:    8080,
	}

	err := validate(cfg)
	expect := "AppName is required"

	if err == nil {
		t.Fatalf("expected error for empty AppName, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestValidate_Success(t *testing.T) {
	cfg := Config{
		AppName: "demo",
		Env:     "local",
		Port:    8080,
	}

	err := validate(cfg)
	if err != nil {
		t.Errorf("expected no error, got %q", err.Error())
	}
}

func TestValidate_InvalidEnv(t *testing.T) {
	cfg := Config{
		AppName: "demo",
		Env:     "app",
		Port:    8080,
	}

	err := validate(cfg)
	expect := "invalid Env"

	if err == nil {
		t.Fatalf("expected error for invalid env, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestValidate_InvalidPort_TooSmall(t *testing.T) {
	cfg := Config{
		AppName: "demo",
		Env:     "local",
		Port:    -1,
	}

	err := validate(cfg)
	expect := "invalid Port"

	if err == nil {
		t.Fatalf("expected error for invalid port, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestValidate_InvalidPort_TooLarge(t *testing.T) {
	cfg := Config{
		AppName: "demo",
		Env:     "local",
		Port:    65536,
	}

	err := validate(cfg)
	expect := "invalid Port"

	if err == nil {
		t.Fatalf("expected error for invalid port, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}
