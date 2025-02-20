package cmd

import (
	"fmt"

	"github.com/cli-todo/internals/tasks"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new Task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := ""
		for _, word := range args {
			title += word + " "
		}
		title = title[:len(title)-1]
		if err := tasks.AddTask(title); err != nil {
			fmt.Println("Error adding task:", err)
		}
	},
}
