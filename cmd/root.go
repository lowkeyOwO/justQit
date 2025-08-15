package cmd

import (
	"github.com/spf13/cobra"
	"justQit/cmd/dispatcher_cmd"
)

var rootCmd = &cobra.Command{
	Use:   "justQit",
	Short: "A job queue dispatcher and worker tool",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(dispatcher_cmd.Cmd)
}
