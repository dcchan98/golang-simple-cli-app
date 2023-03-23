package cmd

import (
	"fmt"
	"github.com/dcchan98/golang-simple-cli-app/pkg/store"
)

func AddTask(s store.Store, name string) error {
	err := s.AddTask(name)
	if err != nil {
		return err
	}

	fmt.Println("Task added successfully.")
	return nil
}
