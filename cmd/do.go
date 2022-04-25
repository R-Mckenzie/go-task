/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/R-Mckenzie/gotask-cli/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "complete a task from the list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		delId, err := strconv.ParseInt(args[0], 10, 0)
		if err != nil {
			fmt.Print("Please enter the number of a task from your list\n")
			return
		}
		err = db.Delete(int(delId))
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("completed task %d\n", delId)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
