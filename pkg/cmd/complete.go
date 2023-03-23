package cmd

import (
	"fmt"
	"github.com/dcchan98/golang-simple-cli-app/pkg/store"
)

func CompleteTask(s store.Store, id int) error {
	err := s.CompleteTask(id)
	if err != nil {
		return err
	}

	fmt.Println("Task marked as completed.")
	return nil
}
