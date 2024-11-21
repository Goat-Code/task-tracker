package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tasks := []Task{}
	tasks = add("Test Task", tasks)

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Name != "Test Task" {
		t.Errorf("Expected task name 'Test Task', got %s", tasks[0].Name)
	}
}

func TestUpdate(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task"},
	}
	tasks = update(1, tasks, "Updated Task")

	if tasks[0].Name == "Test Task" {
		t.Errorf("Expected task name to be updated")
	}
}

func TestRemove(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task"},
	}
	tasks = remove(1, tasks)

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}
