package cmd

import (
	"fmt"
	"os"

	"github.com/cli-todo/internals/tasks"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-todo-cli",
	Short: "A simple todo list CLI",
	Long:  `A CLI application to manage your to-do list, built in Go using best practices.`,
}

func Execute() {

	if err := tasks.LoadTask(); err != nil {
		fmt.Println("Error loading tasks", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
}
