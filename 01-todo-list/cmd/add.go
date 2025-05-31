/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "tasks add <description>",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		task := strings.Join(args, " ")
		if err := addTask(task); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func addTask(task string) error {
	file, err := os.OpenFile("tstore.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	// check for any free IDs
	freeidsFile, err := os.OpenFile("free-ids.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	defer freeidsFile.Close()

	// BUG: read a free ID and DELETE it from the file
	scanner := bufio.NewScanner(freeidsFile)
	var taskID string
	for scanner.Scan() {
		taskID = scanner.Text()
		if taskID != "" {
			break
		}
	}

	if taskID == "" { // means we do not have a free ID
		// read ID from the id-counter file and update the counter
		fmt.Println("here in counter")
		taskID = readAndUpdateCounter("id-counter.txt")
	}

	taskTimestamp := time.Now()

	// write to tasks list
	csvEntry := []string{taskID, task, taskTimestamp.String(), "false"}
	err = w.Write(csvEntry)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	fmt.Printf("%s\t%s\n", taskID, task)
	fmt.Printf("created on %v\n", taskTimestamp.Format(time.UnixDate))

	return nil
}

func readAndUpdateCounter(filename string) string {
	n, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cur, _ := strconv.Atoi(string(n)) // ignore error
	next := cur + 1

	os.WriteFile(filename, []byte(strconv.Itoa(next)), 0644)

	return strconv.Itoa(cur)
}

func init() {
	rootCmd.AddCommand(addCmd)

}
