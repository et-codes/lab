/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/et-codes/lab/task/storage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task_description]",
	Short: "Add a new task to the task list",
	Long: `The task add command adds a new task with the provided information to the database.

Example:
  task add "Get groceries for the week"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No task description provided.")
			os.Exit(1)
		}
		request := storage.AddTaskRequest{
			Description: args[0],
		}
		db, err := storage.OpenDB(dbPath)
		if err != nil {
			fmt.Printf("Error opening database: %v\n", err)
			os.Exit(1)
		}
		id, err := db.AddTask(request)
		if err != nil {
			fmt.Printf("Error creating task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task created with ID %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
