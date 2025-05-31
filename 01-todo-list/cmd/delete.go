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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "tasks delete <task-id>",
	Long: `to delete multiple tasks: 
	tasks delete <taskid1,taskid2,taskid3>
	might support 'delete all' in future`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		tasksList := strings.Split(args[0], `,`)
		deleteTasks(tasksList)
	},
}

func deleteTasks(tasks []string) {
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
	var freeIds []string

	for {
		entry, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if m[entry[0]] {
			freeIds = append(freeIds, entry[0])
			continue
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

	updateFreeIds(freeIds)
}

func updateFreeIds(freeIds []string) {
	file, err := os.OpenFile("free-ids.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer file.Close()

	for _, id := range freeIds {
		file.WriteString(id + "\n")
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
