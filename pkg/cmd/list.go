package cmd

import (
	"fmt"
	"github.com/dcchan98/golang-simple-cli-app/pkg/store"
)

func ListTasks(s store.Store) error {
	tasks, err := s.GetTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	for _, task := range tasks {
		completed := " "
		if task.Completed {
			completed = "x"
		}
		fmt.Printf("[%s] %d: %s\n", completed, task.ID, task.Name)
	}

	return nil
}
