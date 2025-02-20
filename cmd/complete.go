package cmd

import (
	"fmt"
	"strconv"

	"github.com/cli-todo/internals/tasks"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		if err := tasks.CompleteTask(id); err != nil {
			fmt.Println("Error:", err)
		}
	},
}
