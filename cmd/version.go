package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version`,
	Run: func(cmd *cobra.Command, args []string) {
		isShort, _ := cmd.Flags().GetBool("short")
		sys.LogVersion(isShort)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	landscapeCmd.Flags().BoolP("short", "s", false, "Short version")
}
