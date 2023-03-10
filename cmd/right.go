package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var rightCmd = &cobra.Command{
	Use:   "right",
	Short: "Rotate 90deg right.",
	Long:  `Rotate 90deg right.`,
	Run: func(cmd *cobra.Command, args []string) {
		sys.RotateRight(displayName)
	},
}

func init() {
	rootCmd.AddCommand(rightCmd)
}
