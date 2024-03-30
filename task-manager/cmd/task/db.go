package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

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

// OpenDB opens a new or existing DB
func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := InitDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

// InitDB initializes the database
func InitDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		description TEXT NOT NULL,
		priority TEXT NOT NULL,
		project TEXT,
		status TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// AddTask adds a new task to the database and returns the ID.
func AddTask(db *sql.DB, request AddTaskRequest) (int64, error) {
	query := `INSERT INTO tasks (
		created_at, 
		updated_at, 
		description, 
		priority, 
		project, 
		status
	)
	VALUES ($1, $2, $3, $4, $5, $6);`
	result, err := db.Exec(query,
		time.Now().UTC(),
		time.Now().UTC(),
		request.Description,
		request.Priority,
		request.Project,
		request.Status,
	)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetTasks returns a slice of all the tasks in the database.
func GetTasks(db *sql.DB) ([]Task, error) {
	query := `SELECT * FROM tasks;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	for rows.Next() {
		task := Task{}
		err := rows.Scan(
			&task.ID,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.Description,
			&task.Priority,
			&task.Project,
			&task.Status,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// DeleteTask deletes the task with the given ID, if it exists.
func DeleteTask(db *sql.DB, id int64) error {
	if id <= 0 {
		return fmt.Errorf("id must be a positive integer")
	}

	query := `DELETE FROM tasks WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// EditTask updates one or more fields of an existing task.
func EditTask(db *sql.DB, request EditTaskRequest) error {
	query := `UPDATE tasks SET
		description = $1,
		priority = $2,
		project = $3,
		status = $4,
		updated_at = $5
	WHERE id = $6;`

	_, err := db.Exec(query,
		request.Description,
		request.Priority,
		request.Project,
		request.Status,
		time.Now().UTC(),
		request.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
