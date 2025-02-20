package cmd

import (
	"github.com/cli-todo/internals/tasks"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available commands",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.ListTasks()
	},
}
