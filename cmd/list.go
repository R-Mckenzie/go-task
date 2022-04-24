package cmd

import (
	"fmt"
	"log"

	"github.com/R-Mckenzie/gotask-cli/db"
	"github.com/spf13/cobra"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorWhite = "\033[37m"
const bold = "\033[1m"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all active tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.LoadTasks()
		if err != nil {
			log.Fatal("Could not load tasks")
		}

		if len(tasks) == 0 {
			fmt.Println("No active tasks")
		} else {
			for i, task := range tasks {
				fmt.Printf("%s%s %d) %s\n%s    - %s\n", bold, colorWhite, i+1, task.Title, colorReset, task.Desc)
			}
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
