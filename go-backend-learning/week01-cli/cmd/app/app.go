package main

import (
	"fmt"
	"os"

	"week01-cli/internal/config"
	"week01-cli/internal/model"
)

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("missing subcommand: should be 'run' or 'task-summary'")
	}

	switch os.Args[1] {
	case "run":
		args, err := parseArgs()
		if err != nil {
			return err
		}
		return runCommand(args)
	case "task-summary":
		return taskSummaryCommand()
	default:
		return fmt.Errorf("unknown subcommand %q: should be 'run' or 'task-summary'", os.Args[1])
	}
}

func runCommand(args Args) error {
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

func taskSummaryCommand() error {
	task := model.Task{
		ID:     1,
		Name:   "learning-go-day2",
		Status: "running",
		Tags:   []string{"go"},
		Meta:   map[string]string{"owner": "dreamking60"},
	}

	fmt.Println(task.Summary())
	return nil
}
