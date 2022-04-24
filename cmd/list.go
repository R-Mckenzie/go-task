package cmd

import (
	"fmt"
	"log"

	"github.com/R-Mckenzie/gotask-cli/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all active tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.LoadTasks()
		if err != nil {
			log.Fatal("Could not load tasks")
		}

		for i, task := range tasks {
			fmt.Printf("%d) %s\n", i+1, task.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
