package main

import (
	"fmt"

	"week01-cli/internal/config"
	"week01-cli/internal/model"
)

func run() error {
	args, err := parseArgs()
	if err != nil {
		return err
	}

	cfg := config.Default()
	cfg.Env = args.Env
	cfg.Port = args.Port
	cfg.AppName = args.AppName

	task := model.Task{
		ID:     1,
		Name:   "learning-go-day2",
		Status: "running",
		Tags:   []string{"go"},
		Meta:   map[string]string{"owner": "dreamking60"},
	}

	fmt.Printf("level=INFO msg=%q app=%s env=%s port=%d\n", "application started", cfg.AppName, cfg.Env, cfg.Port)
	fmt.Println(task.Summary())

	return nil
}
