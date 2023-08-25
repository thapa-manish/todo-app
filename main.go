package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tasks := []string{}

	for {
		fmt.Println("\n\nTODO APPLICATION\nEnter your choice:")
		fmt.Println("1: to add\n2: to remove:\n3: to view tasks:\n4: to exit")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Enter task:")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tasks = append(tasks, task)
			fmt.Println("task added:")
		case "2":
			fmt.Println("Tasks: ")
			for i, task := range tasks {
				fmt.Printf("%d: %s\n", i+1, task)
			}
			fmt.Println("Enter task number to remove:")

			var index int
			fmt.Scan(&index)
			if index > 0 && index <= len(tasks) {
				tasks = append(tasks[:index-1], tasks[index:]...)
				fmt.Println("task removed:")
			} else {
				fmt.Println("invalid number:")
			}
		case "3":
			fmt.Println("Tasks: ")
			for i, task := range tasks {
				fmt.Printf("%d: %s\n", i+1, task)
			}
		case "4":
			fmt.Println("exiting program.........")
			os.Exit(1)
		default:
			fmt.Println("Invalid choice. Please select the valid option:")
		}
	}
}
