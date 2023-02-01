package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var invertCmd = &cobra.Command{
	Use:   "invert",
	Short: "Rotate 180deg (upside down).",
	Long:  `Rotate 180deg (upside down).`,
	Run: func(cmd *cobra.Command, args []string) {
		sys.Rotate(sys.Display(), "inverted", "-1 0 1 0 -1 1 0 0 1")
	},
}

func init() {
	rootCmd.AddCommand(invertCmd)
}
