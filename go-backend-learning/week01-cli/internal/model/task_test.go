package model

import "testing"

func TestTaskSummary(t *testing.T) {
	task := Task{
		ID:     1,
		Name:   "learning-go-day2",
		Status: "running",
		Tags:   []string{"go", "cli"},
		Meta:   map[string]string{"owner": "dreamking60"},
	}

	got := task.Summary()
	want := "Task[1] learning-go-day2 status=running tags=2"

	if got != want {
		t.Errorf("expected summary %q, got %q", want, got)
	}
}
