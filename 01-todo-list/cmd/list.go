/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

const (
	timeLayout = "2006-01-02 15:04:05.999999999 -0700 MST m=+0.000000000"
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

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 4, 4, 1, ' ', 0)
	defer w.Flush()

	if all {
		fmt.Fprintln(w, "ID\tTASK\tCREATED\tDONE")
	} else {
		fmt.Fprintln(w, "ID\tTASK\tCREATED")
	}

	for _, v := range list {
		if all {
			fmt.Fprintf(w, "%s\t%s\t%v\t%s\n", v[0], v[1], humanize(v[2]), v[3])
		} else {
			fmt.Fprintf(w, "%s\t%s\t%v\n", v[0], v[1], humanize(v[2]))
		}
	}

	fmt.Fprintln(w)

}

func humanize(t string) string {
	tm, err := time.Parse(timeLayout, t)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return timediff.TimeDiff(tm)

}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("all", "a", false, "list uncompleted and completed tasks")
}
