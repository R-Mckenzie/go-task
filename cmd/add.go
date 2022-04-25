package cmd

import (
	"fmt"

	"github.com/R-Mckenzie/gotask-cli/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add \"task title\" \"task desctiption\"",
	Short: "add a task to the task list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please enter a task")
			return
		}

		if len(args) == 1 {
			args = append(args, "No description provided")
		}

		task := db.Task{Title: args[0], Desc: args[1]}
		db.Save(task)
		fmt.Printf("added: %q\n", task.Title)
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
