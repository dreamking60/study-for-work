package model

import "fmt"

// Task is a business object in this small demo.
type Task struct {
	ID int // Unique identifier of the task.
	// Human-readable task name shown in logs/CLI.
	Name string
	// Current lifecycle state, e.g. pending/running/success/failed.
	Status string
	// Category labels used for filtering and grouping.
	Tags []string
	// Extensible key-value metadata, e.g. owner/source/priority.
	Meta map[string]string
}

// Summary returns a short printable description for CLI output.
func (t Task) Summary() string {
	return fmt.Sprintf("Task[%d] %s status=%s tags=%d",
		t.ID, t.Name, t.Status, len(t.Tags))
}
