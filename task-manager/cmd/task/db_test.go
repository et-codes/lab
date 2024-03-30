package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

var newTask = AddTaskRequest{
	Description: "Test task",
	Priority:    "M",
	Project:     "test",
	Status:      "in progress",
}

func TestDB(t *testing.T) {
	t.Run("AddTask returns the ID", func(t *testing.T) {
		db, err := OpenDB(":memory:")
		assert.NoError(t, err)
		id, err := AddTask(db, newTask)
		assert.NoError(t, err)
		assert.Greater(t, id, int64(0))
	})

	t.Run("GetTasks returns the new task", func(t *testing.T) {
		db, _ := setupDB(t)
		tasks, err := GetTasks(db)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(tasks))
	})

	t.Run("DeleteTask deletes the task", func(t *testing.T) {
		db, id := setupDB(t)
		err := DeleteTask(db, id)
		assert.NoError(t, err)

		tasks, err := GetTasks(db)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(tasks))
	})

	t.Run("EditTask correctly updates task", func(t *testing.T) {
		db, id := setupDB(t)

		request := EditTaskRequest{
			ID:          id,
			Description: "Edited test task",
			Priority:    "H",
			Project:     "edited_test",
			Status:      "done",
		}

		err := EditTask(db, request)
		assert.NoError(t, err)
	})
}

// setupDB creates a new database and returns a pointer to it, along with the
// ID of the Task it was seeded with.
func setupDB(t *testing.T) (*sql.DB, int64) {
	t.Helper()
	db, err := OpenDB(":memory:")
	if err != nil {
		t.Fatal(err)
	}

	id, err := AddTask(db, newTask)
	if err != nil {
		t.Fatal(err)
	}

	return db, id
}
