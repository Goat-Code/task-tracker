package main

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"
)

func TestAddEmptyTask(t *testing.T) {
	tasks := []Task{}
	tasks = add("", tasks)

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Name != "" {
		t.Errorf("Expected task name to be empty, got %s", tasks[0].Name)
	}
}

func TestUpdateNonExistentTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task"},
	}
	tasks = update(2, tasks, "Updated Task")

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Name != "Test Task" {
		t.Errorf("Task should not be updated")
	}
}

func TestRemoveNonExistentTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task"},
	}
	tasks = remove(2, tasks)

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestMarkDoneAlreadyDone(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task", Status: "done"},
	}
	tasks = markDone(tasks, 1)

	if tasks[0].Status != "done" {
		t.Errorf("Expected task status to remain 'done', got %s", tasks[0].Status)
	}
}

func TestMarkInProgressAlreadyInProgress(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task", Status: "in-progress"},
	}
	tasks = markInProgress(tasks, 1)

	if tasks[0].Status != "in-progress" {
		t.Errorf("Expected task status to remain 'in-progress', got %s", tasks[0].Status)
	}
}

func TestGetLastIdAfterRemoveAll(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task"},
	}
	tasks = remove(1, tasks)
	lastId := getLastId(tasks)

	if lastId != 0 {
		t.Errorf("Expected last id 0, got %d", lastId)
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

func TestMarkInProgress(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task", Status: "to-do"},
	}
	tasks = markInProgress(tasks, 1)

	if tasks[0].Status != "in-progress" {
		t.Errorf("Expected task status 'in-progress', got %s", tasks[0].Status)
	}
}

func TestMarkDone(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task", Status: "in-progress"},
	}
	tasks = markDone(tasks, 1)

	if tasks[0].Status != "done" {
		t.Errorf("Expected task status 'done', got %s", tasks[0].Status)
	}
}

func TestGetLastId(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task"},
	}
	lastId := getLastId(tasks)

	if lastId != 1 {
		t.Errorf("Expected last id 1, got %d", lastId)
	}
}

func TestGetLastIdEmpty(t *testing.T) {
	tasks := []Task{}
	lastId := getLastId(tasks)

	if lastId != 0 {
		t.Errorf("Expected last id 0, got %d", lastId)
	}
}

func TestGetLastIdMultiple(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Test Task"},
		{ID: 2, Name: "Test Task 2"},
	}
	lastId := getLastId(tasks)

	if lastId != 2 {
		t.Errorf("Expected last id 2, got %d", lastId)
	}
}


func TestListEmptyTasks(t *testing.T) {
	tasks := []Task{}
	output := captureOutput(func() {
		list(tasks)
	})
	expected := "No items in the list\n"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestListAllTasks(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Task 1", Status: "todo", CreatedAt: time.Now().Format("RCF1123")},
		{ID: 2, Name: "Task 2", Status: "done", CreatedAt: time.Now().Format("RCF1123")},
	}
	os.Args = []string{"cmd", "list"}
	output := captureOutput(func() {
		list(tasks)
	})
	if !bytes.Contains([]byte(output), []byte("Task 1")) || !bytes.Contains([]byte(output), []byte("Task 2")) {
		t.Errorf("Expected all tasks to be listed")
	}
}

func TestListFilteredDone(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Task 1", Status: "done", CreatedAt: time.Now().Format("RCF1123")},
		{ID: 2, Name: "Task 2", Status: "todo", CreatedAt: time.Now().Format("RCF1123")},
	}
	os.Args = []string{"cmd", "list", "done"}
	output := captureOutput(func() {
		list(tasks)
	})
	if !bytes.Contains([]byte(output), []byte("Task 1")) || bytes.Contains([]byte(output), []byte("Task 2")) {
		t.Errorf("Expected only done tasks to be listed")
	}
}

func TestListFilteredTodo(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Task 1", Status: "done", CreatedAt: time.Now().Format("RCF1123")},
		{ID: 2, Name: "Task 2", Status: "todo", CreatedAt: time.Now().Format("RCF1123")},
	}
	os.Args = []string{"cmd", "list", "todo"}
	output := captureOutput(func() {
		list(tasks)
	})
	if !bytes.Contains([]byte(output), []byte("Task 2")) || bytes.Contains([]byte(output), []byte("Task 1")) {
		t.Errorf("Expected only todo tasks to be listed")
	}
}

func TestListFilteredInProgress(t *testing.T) {
	tasks := []Task{
		{ID: 1, Name: "Task 1", Status: "in-progress", CreatedAt: time.Now().Format("RCF1123")},
		{ID: 2, Name: "Task 2", Status: "todo", CreatedAt: time.Now().Format("RCF1123")},
	}
	os.Args = []string{"cmd", "list", "in-progress"}
	output := captureOutput(func() {
		list(tasks)
	})
	if !bytes.Contains([]byte(output), []byte("Task 1")) || bytes.Contains([]byte(output), []byte("Task 2")) {
		t.Errorf("Expected only in-progress tasks to be listed")
	}
}
// captureOutput captures the output of a function that writes to stdout
func captureOutput(f func()) string {
	// Backup the original stdout
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore stdout after capture

	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Close the writer and read the captured output
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
