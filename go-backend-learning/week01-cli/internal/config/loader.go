package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadConfig(path string) (Config, error) {
	data, ext, err := loadBytes(path)
	if err != nil {
		return Config{}, err
	}

	cfg, err := parseByExt(ext, data)
	if err != nil {
		return Config{}, err
	}

	if err := validate(cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func loadBytes(path string) ([]byte, string, error) {
	if path == "" {
		return nil, "", fmt.Errorf("load config: empty path")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("load config %q: %w", path, err)
	}

	ext := filepath.Ext(path) // ".json", ".yaml", ".toml"
	if ext == "" {
		return nil, "", fmt.Errorf("load config %q: missing file extension", path)
	}

	return data, ext, nil
}
