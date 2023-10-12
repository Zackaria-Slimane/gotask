package main

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	tasks := []task{{1, "Task 1", false}, {2, "Task 2", false}}
	newTaskName := "Task 3"

	updatedTasks := addTasks(tasks, newTaskName)

	// Check if the new task is added
	found := false
	for _, task := range updatedTasks {
		if task.Value == newTaskName {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected new task to be added, but it's not in the task list.")
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []task{{1, "Task 1", false}, {2, "Task 2", false}}
	updatedTasks := deleteTask(tasks, 1)

	for _, task := range updatedTasks {
		if task.ID == 1 {
			t.Errorf("Expected task with ID 1 to be removed, but it's still present.")
		}
	}
}

func TestMarkDone(t *testing.T) {
	tasks := []task{{1, "Task 1", false}, {2, "Task 2", false}}
	taskID := 1

	updatedTasks := markDone(tasks, taskID)

	// Check if the task with the specified ID is marked as done
	for _, task := range updatedTasks {
		if task.ID == taskID && !task.Completed {
			t.Errorf("Expected task with ID %d to be marked as done, but it's still not done.", taskID)
		}
	}
}
