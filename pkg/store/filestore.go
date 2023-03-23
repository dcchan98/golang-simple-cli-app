package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FileStore struct {
	filename string
}

func NewFileStore(filename string) *FileStore {
	fs := &FileStore{filename: filename}
	if _, err := os.Stat(fs.filename); os.IsNotExist(err) {
		// Create an empty tasks.json file if it doesn't exist
		err := ioutil.WriteFile(fs.filename, []byte("[]"), 0644)
		if err != nil {
			fmt.Printf("Error creating %s: %v\n", fs.filename, err)
		}
	}
	return fs
}

func (s *FileStore) GetTasks() ([]Task, error) {
	data, err := ioutil.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *FileStore) AddTask(name string) error {
	tasks, err := s.GetTasks()
	if err != nil {
		return err
	}

	id := 1
	for _, task := range tasks {
		if task.ID >= id {
			id = task.ID + 1
		}
	}

	newTask := Task{ID: id, Name: name, Completed: false}
	tasks = append(tasks, newTask)

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.filename, data, 0644)
}

func (s *FileStore) CompleteTask(id int) error {
	tasks, err := s.GetTasks()
	if err != nil {
		return err
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return os.ErrNotExist
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.filename, data, 0644)
}
