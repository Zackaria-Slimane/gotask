package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type task struct {
	ID        int
	Value     string
	Completed bool
}

func main() {
	dayIs := time.Now().Day()
	monthIs := time.Now().Month()
	yearIs := time.Now().Year()
	todayIs := fmt.Sprintf("%d-%d-%d", dayIs, monthIs, yearIs)
	fmt.Println("\nYour todo list for today: ", todayIs)
	fmt.Println("======================================")

	const filename = "clinotes.json"
	loadedTasks, err := loadTasks(filename)
	emptyTasks := len(loadedTasks)

	if err != nil && emptyTasks == 0 {
		fmt.Println("No Tasks added yet")
	} else if err != nil && emptyTasks > 1 {
		fmt.Println("Error loading tasks:", err)
	}

	Tasks := loadedTasks
	printTasks(Tasks)
	fmt.Println("======================================")

	for {
		fmt.Print("\nChoose an option (add/list/cross/remove/exit): ")
		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')

		switch option {
		case "add\n":
			fmt.Print("Enter task name: ")
			taskInput, _ := reader.ReadString('\n')
			Tasks = addTasks(Tasks, taskInput)
			fmt.Println("Task successfully added")

		case "list\n":
			printTasks(Tasks)

		case "cross\n":
			fmt.Print("Enter the ID of the task you want to mark as completed: ")
			var taskID int
			_, err := fmt.Scanf("%d", &taskID)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task ID.")
				continue
			}
			Tasks = markDone(Tasks, taskID)

		case "remove\n":
			fmt.Print("Enter the ID of the task you want to remove: ")
			var taskID int
			_, err := fmt.Scanf("%d", &taskID)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid task ID.")
			}
			Tasks = deleteTask(Tasks, taskID)

		case "exit\n":
			fmt.Println("Saving your tasks, Goodbye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid option. Try again.")
		}

		err = saveTasks(Tasks, filename)
		if err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	}
}

func saveTasks(tasks []task, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}

func loadTasks(filename string) ([]task, error) {
	var tasks []task
	file, err := os.Open(filename)
	if err != nil {
		return tasks, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func addTasks(tasks []task, taskName string) []task {
	newTask := task{
		ID:        len(tasks) + 1,
		Value:     taskName,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	return tasks
}

func printTasks(tasks []task) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	if len(tasks) == 0 {
		fmt.Println("No tasks added yet")
	}
	for _, task := range tasks {
		status := ""
		taskName := task.Value
		if task.Completed {
			status = green("(Done)")
		} else {
			status = red("(Not Done)")
		}
		fmt.Printf("%d: %s%s\n", task.ID, taskName, status)
	}
}

func deleteTask(tasks []task, id int) []task {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Removing task ID: ", id)
			return tasks
		}
	}
	return tasks
}

func markDone(tasks []task, id int) []task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task ID: ", id, "marked as done")
			return tasks
		}
	}
	return tasks
}
