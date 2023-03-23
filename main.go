package main

import (
	"bufio"
	"fmt"
	"github.com/dcchan98/cli-app/pkg/cmd"
	"github.com/dcchan98/cli-app/pkg/store"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := store.NewFileStore("tasks.json")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		if len(parts) == 0 {
			continue
		}

		command := parts[0]
		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Please provide a task name")
				continue
			}
			name := parts[1]
			err := cmd.AddTask(s, name)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "list":
			err := cmd.ListTasks(s)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "complete":
			if len(parts) < 2 {
				fmt.Println("Please provide a task ID")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid task ID")
				continue
			}
			err = cmd.CompleteTask(s, id)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "exit":
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}
