package config

import "testing"

func TestDefault(t *testing.T) {
	cfg := Default()

	if cfg.AppName != "demo" {
		t.Errorf("expected AppName %q, got %q", "demo", cfg.AppName)
	}

	if cfg.Env != "local" {
		t.Errorf("expected Env %q, got %q", "local", cfg.Env)
	}

	if cfg.Port != 8765 {
		t.Errorf("expected Port %d, got %d", 8765, cfg.Port)
	}
}
