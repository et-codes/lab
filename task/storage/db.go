package storage

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

// OpenDB opens a new or existing DB
func OpenDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	newDB := &DB{db}

	if err := newDB.InitDB(); err != nil {
		return nil, err
	}

	return newDB, nil
}

// InitDB initializes the database
func (db *DB) InitDB() error {
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
func (db *DB) AddTask(request AddTaskRequest) (int64, error) {
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
func (db *DB) GetTasks() ([]Task, error) {
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
func (db *DB) DeleteTask(id int64) error {
	if id <= 0 {
		return fmt.Errorf("id must be a positive integer")
	}

	query := `DELETE FROM tasks WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("id %d not found", id)
	}

	return nil
}

// EditTask updates one or more fields of an existing task.
func (db *DB) EditTask(request EditTaskRequest) error {
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
