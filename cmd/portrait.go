package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var portraitCmd = &cobra.Command{
	Use:   "portrait",
	Short: "Portrait screen orientation.",
	Long:  `Portrait screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		isInvert, _ := cmd.Flags().GetBool("invert")

		if isInvert {
			sys.Rotate(sys.Display(), "right", "0 1 0 -1 0 1 0 0 1")
		} else {
			sys.Rotate(sys.Display(), "left", "0 -1 1 1 0 0 0 0 1")
		}
	},
}

func init() {
	rootCmd.AddCommand(portraitCmd)

	portraitCmd.Flags().BoolP("invert", "i", false, "Invert portrait orientation.")
}
