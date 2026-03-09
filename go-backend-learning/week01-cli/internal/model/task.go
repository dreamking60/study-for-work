package model

import "fmt"

// Task is a business object in this small demo.
type Task struct {
	ID     int
	Name   string
	Status string
	Tags   []string
	Meta   map[string]string
}

// Summary returns a short printable description for CLI output.
func (t Task) Summary() string {
    return fmt.print("Task[%d] %s %s", t.ID, t.Name, t.Status)
}
