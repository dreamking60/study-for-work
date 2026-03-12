package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Args struct {
	Env     string
	Port    int
	AppName string
}

func parseArgs() (Args, error) {
	return parseArgsFrom(os.Args[1:])
}

func parseArgsFrom(argv []string) (Args, error) {
	fs := flag.NewFlagSet("app", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	env := fs.String("env", "local", "Running Environment")
	port := fs.Int("port", 8765, "Port Number")
	appName := fs.String("app", "demo", "App Name")

	if err := fs.Parse(argv); err != nil {
		return Args{}, fmt.Errorf("parse args: %w", err)
	}

	if *port < 1 || *port > 65535 {
		return Args{}, fmt.Errorf("invalid arg port %d: should be from 1-65535", *port)
	}

	switch *env {
	case "local", "dev", "test", "prod":
	default:
		return Args{}, fmt.Errorf("invalid arg env %q: should be 'local', 'dev', 'test' or 'prod'", *env)
	}

	return Args{
		Env:     *env,
		Port:    *port,
		AppName: *appName,
	}, nil
}
