package storage

import "time"

type Task struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
	Priority    string
	Project     string
	Status      string
}

type AddTaskRequest struct {
	Description string
	Priority    string
	Project     string
	Status      string
}

type EditTaskRequest struct {
	ID          int64
	Description string
	Priority    string
	Project     string
	Status      string
}
