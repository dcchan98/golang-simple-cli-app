package store

type Task struct {
	ID        int
	Name      string
	Completed bool
}

type Store interface {
	GetTasks() ([]Task, error)
	AddTask(name string) error
	CompleteTask(id int) error
}
