/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete task by id",
	Long: `usage: tasks complete <task-id> 
	to complete multiple tasks: tasks complete <taskid1,taskid2,taskid3>`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		tasksList := strings.Split(args[0], `,`)
		completeTasks(tasksList)
	},
}

func completeTasks(tasks []string) {
	tasksFile, err := os.OpenFile("tstore.csv", os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var m = map[string]bool{}

	for _, task := range tasks {
		m[task] = true
	}

	r := csv.NewReader(tasksFile)

	var newEntries [][]string

	for {
		entry, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if m[entry[0]] {
			entry[3] = "true"
		}

		newEntries = append(newEntries, entry)

	}

	tasksFile.Close()

	if tasksFile, err = os.Create("tstore.csv"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w := csv.NewWriter(tasksFile)
	w.WriteAll(newEntries) // flush called internally

}

func init() {
	rootCmd.AddCommand(completeCmd)

}
