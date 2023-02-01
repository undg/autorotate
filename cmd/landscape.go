package cmd

import (
	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var landscapeCmd = &cobra.Command{
	Use:   "landscape",
	Short: "Landscape screen orientation.",
	Long:  `Landscape screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		isInvert, _ := cmd.Flags().GetBool("invert")

		if isInvert {
			sys.Rotate(sys.Display(), "inverted", "-1 0 1 0 -1 1 0 0 1")
		} else {
			sys.Rotate(sys.Display(), "normal", "0 0 0 0 0 0 0 0 0")
		}
	},
}

func init() {
	rootCmd.AddCommand(landscapeCmd)

	landscapeCmd.Flags().BoolP("invert", "i", false, "Invert landscape orientation.")
}
