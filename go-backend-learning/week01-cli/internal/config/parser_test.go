package config

import (
	"strings"
	"testing"
)

func TestParseByExt_JSON_InvalidContent(t *testing.T) {
	ext := ".json"
	data := `{
		"AppName": "demo",
		"Env": "local",
		"Port": 8080,
	}`

	_, err := parseByExt(ext, []byte(data))
	expect := "parse config as json"
	if err == nil {
		t.Fatal("expected error for invalid JSON content, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestParseByExt_JSON_Success(t *testing.T) {
	ext := ".json"
	data := `{
		"AppName": "demo",
		"Env": "local",
		"Port": 8080
	}`

	cfg, err := parseByExt(ext, []byte(data))
	if err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}

	if cfg.AppName != "demo" {
		t.Errorf("expected AppName %q, got %q", "demo", cfg.AppName)
	}

	if cfg.Env != "local" {
		t.Errorf("expected Env %q, got %q", "local", cfg.Env)
	}

	if cfg.Port != 8080 {
		t.Errorf("expected Port %d, got %d", 8080, cfg.Port)
	}
}

func TestParseByExt_YAML_Success(t *testing.T) {
	ext := ".yaml"
	data := []byte("appname: demo\nenv: dev\nport: 9090\n")

	cfg, err := parseByExt(ext, data)
	if err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}

	if cfg.AppName != "demo" {
		t.Errorf("expected AppName %q, got %q", "demo", cfg.AppName)
	}

	if cfg.Env != "dev" {
		t.Errorf("expected Env %q, got %q", "dev", cfg.Env)
	}

	if cfg.Port != 9090 {
		t.Errorf("expected Port %d, got %d", 9090, cfg.Port)
	}
}

func TestParseByExt_TOML_Success(t *testing.T) {
	ext := ".toml"
	data := []byte("AppName = \"demo\"\nEnv = \"prod\"\nPort = 7070\n")

	cfg, err := parseByExt(ext, data)
	if err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}

	if cfg.AppName != "demo" {
		t.Errorf("expected AppName %q, got %q", "demo", cfg.AppName)
	}

	if cfg.Env != "prod" {
		t.Errorf("expected Env %q, got %q", "prod", cfg.Env)
	}

	if cfg.Port != 7070 {
		t.Errorf("expected Port %d, got %d", 7070, cfg.Port)
	}
}

func TestParseByExt_YAML_InvalidContent(t *testing.T) {
	_, err := parseByExt(".yaml", []byte("appname: [demo\n"))
	expect := "parse config as yaml"
	if err == nil {
		t.Fatal("expected error for invalid YAML content, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestParseByExt_TOML_InvalidContent(t *testing.T) {
	_, err := parseByExt(".toml", []byte("AppName = \"demo\nEnv = \"prod\"\n"))
	expect := "parse config as toml"
	if err == nil {
		t.Fatal("expected error for invalid TOML content, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestParseByExt_UnsupportedExt(t *testing.T) {
	_, err := parseByExt(".ini", []byte("appname=demo"))
	expect := "unsupported type"
	if err == nil {
		t.Fatal("expected error for unsupported extension, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}
