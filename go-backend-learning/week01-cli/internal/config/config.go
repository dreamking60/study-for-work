package config

// Config stores runtime settings, not business data.
type Config struct {
	AppName string
	Env     string
	Port    int
}

// Default returns a baseline config for local development.
func Default() Config {
	return Config{
		AppName: "demo",
		Env:     "local",
		Port:    8765,
	}
}
