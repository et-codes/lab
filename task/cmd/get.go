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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list of all tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := storage.OpenDB(dbPath)
		if err != nil {
			fmt.Printf("Error opening database: %v\n", err)
			os.Exit(1)
		}
		tasks, err := db.GetTasks()
		if err != nil {
			fmt.Printf("Error retreiving tasks: %v\n", err)
			os.Exit(1)
		}
		for _, task := range tasks {
			fmt.Printf("%d - %s - %s - %s - %s\n",
				task.ID,
				task.Description,
				task.Priority,
				task.Project,
				task.Status,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
