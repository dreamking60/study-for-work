package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadBytes_Success(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	content := []byte(`{"AppName":"demo","Env":"local","Port":8080}`)

	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	data, ext, err := loadBytes(path)
	if err != nil {
		t.Fatalf("expected no error, got %q", err.Error())
	}

	if string(data) != string(content) {
		t.Errorf("expected data %q, got %q", string(content), string(data))
	}

	if ext != ".json" {
		t.Errorf("expected extension %q, got %q", ".json", ext)
	}
}

func TestLoadBytes_EmptyPath(t *testing.T) {
	_, _, err := loadBytes("")
	expect := "empty path"
	if err == nil {
		t.Fatal("expected error for empty path, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestLoadBytes_FileNotFound(t *testing.T) {
	_, _, err := loadBytes("/tmp/definitely-missing-config.json")
	expect := "load config"
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestLoadBytes_MissingExtension(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config")

	if err := os.WriteFile(path, []byte("demo"), 0o644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	_, _, err := loadBytes(path)
	expect := "missing file extension"
	if err == nil {
		t.Fatal("expected error for missing extension, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestLoadConfig_Success_JSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	content := []byte(`{"AppName":"demo","Env":"local","Port":8080}`)

	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	cfg, err := LoadConfig(path)
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

func TestLoadConfig_InvalidPath(t *testing.T) {
	_, err := LoadConfig("/tmp/definitely-missing-config.json")
	expect := "load config"
	if err == nil {
		t.Fatal("expected error for invalid path, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestLoadConfig_InvalidFormat(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	content := []byte(`{"AppName":"demo","Env":"local","Port":}`)

	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	_, err := LoadConfig(path)
	expect := "parse config as json"
	if err == nil {
		t.Fatal("expected error for invalid format, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}

func TestLoadConfig_InvalidField(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	content := []byte(`{"AppName":"demo","Env":"stage","Port":8080}`)

	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	_, err := LoadConfig(path)
	expect := "invalid Env"
	if err == nil {
		t.Fatal("expected error for invalid field, got nil")
	}

	if !strings.Contains(err.Error(), expect) {
		t.Errorf("expected error containing %q, got %q", expect, err.Error())
	}
}
