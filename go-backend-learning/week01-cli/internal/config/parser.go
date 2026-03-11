package config

import (
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func parseByExt(ext string, data []byte) (Config, error) {
	var cfg Config

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse config as json: %w", err)
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse config as yaml: %w", err)
		}
	case ".toml":
		if err := toml.Unmarshal(data, &cfg); err != nil {
			return Config{}, fmt.Errorf("parse config as toml: %w", err)
		}
	default:
		return Config{}, fmt.Errorf("parse config as %s: unsupported type", ext)
	}

	return cfg, nil
}
