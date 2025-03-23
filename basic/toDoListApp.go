package main

import (
	"bufio"
	"fmt"
	"os"
)

func toDoListApp() {
	var tasks []string
	var choice int
	scanner := bufio.NewScanner(os.Stdin) // Initialize a new scanner for user input

	for {
		fmt.Println("\nTo-Do List Application:")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		fmt.Println()


		switch choice {
		case 1:
			fmt.Print("Enter the task: ")
			scanner.Scan()               // Read the entire line, including spaces
			task := scanner.Text()       // Get the task text
			tasks = append(tasks, task)  // Add the task to the list
			fmt.Println("Task added successfully.")
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks in the list.")
			} else {
				fmt.Println("Your tasks:")
				for i, t := range tasks {
					fmt.Printf("%d. %s\n", i+1, t)
				}

			}
		case 3:
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
			} else {
				fmt.Println("Your tasks:")
				for i, t := range tasks {
					fmt.Printf("%d. %s\n", i+1, t)
				}
				fmt.Print("Enter the task number to remove: ")
				var taskNum int
				fmt.Scanln(&taskNum)
				if taskNum > 0 && taskNum <= len(tasks) {
					tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
					fmt.Println("Task removed successfully.")
				} else {
					fmt.Println("Invalid task number!")
				}
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
