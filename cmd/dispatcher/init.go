package dispatcher

import (
	"github.com/spf13/cobra"
	"justQit/conn"
	"justQit/types"
	"justQit/utils"
)

var Cmd = &cobra.Command{
	Use:   "dispatcher",
	Short: "Commands related to the dispatcher",
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the dispatcher",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt16("port")
		json, _ := cmd.Flags().GetString("json")

		var config types.DispatcherConfig
		if json == "" {
			config = utils.DispatcherInitPrompt()
		} else {
			config = utils.DispatcherReadJSON(json)
		}
		conn.Initialize(port, config)
	},
}

func init() {
	initCmd.Flags().Int16P("port", "p", 7777, "Port for running justQit")
	initCmd.Flags().StringP("json", "j","", "Initialize from JSON file")
	Cmd.AddCommand(initCmd)
}
