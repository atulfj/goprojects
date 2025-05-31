/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "tasks list (lists all incomplete tasks)",
	Long:  `tasks list -a or tasks list --all lists all tasks whether complete or incomplete`,

	Run: func(cmd *cobra.Command, args []string) {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		listTasks(all)
	},
}

func listTasks(all bool) {
	tasksFile, err := os.OpenFile("tstore.csv", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer tasksFile.Close()

	r := csv.NewReader(tasksFile)

	var list [][]string

	for {
		entry, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if all {
			list = append(list, entry)
		} else if entry[3] == "false" {
			list = append(list, entry)
		}
	}

	fmt.Println(list)
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("all", "a", false, "list uncompleted and completed tasks")
}
