package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = getRootCmd()

func getRootCmd() *cobra.Command {

	Use := "tasks <cmd1> <cmd2> <...args>"
	Short := "root command to add / delete / complete tasks"
	Long := "find full specification in README (insert link)"

	RunE := func

}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
