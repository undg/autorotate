package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of connected devices",
	Long:  `List of connected devices`,
	Run: func(cmd *cobra.Command, args []string) {
		sys.LogXDisplays()
		sys.LogXInputs()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
