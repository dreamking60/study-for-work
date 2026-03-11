package config

import (
	"fmt"
)

func validate(cfg Config) error {
	if cfg.AppName == "" {
		return fmt.Errorf("validate config: AppName is required")
	}

	switch cfg.Env {
	case "local", "dev", "test", "prod":
	default:
		return fmt.Errorf("validate config: invalid Env %q", cfg.Env)
	}

	if cfg.Port < 1 || cfg.Port > 65535 {
		return fmt.Errorf("validate config: invalid Port %d", cfg.Port)
	}

	return nil
}
