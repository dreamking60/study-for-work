package main

import (
	"fmt"

	"week01-cli/internal/config"
	"week01-cli/internal/model"
)

func main() {
	cfg := config.Default()

	task := model.Task{
		ID:     1,
		Name:   "learning-go-day2",
		Status: "running",
		Tags:   []string{"go"},
		Meta:   map[string]string{"owner": "dreamking60"},
	}

	fmt.Printf("App=%s, Env=%s, Port=%d\n", cfg.AppName, cfg.Env, cfg.Port)
	fmt.Println(task.Summary())

}
