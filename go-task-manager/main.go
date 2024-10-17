package main

import (
	"fmt"
	"os"
)

func main() {
	tasks := []string{}

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Delete a task")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var task string
			fmt.Print("Enter task: ")
			fmt.Scan(&task)
			tasks = addTask(tasks, task)
		case 2:
			listTasks(tasks)
		case 3:
			var index int
			fmt.Print("Enter task number to delete: ")
			fmt.Scan(&index)
			tasks = deleteTask(tasks, index-1)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addTask(tasks []string, task string) []string {
	return append(tasks, task)
}

func listTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func deleteTask(tasks []string, index int) []string {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task number.")
		return tasks
	}
	return append(tasks[:index], tasks[index+1:]...)
}
