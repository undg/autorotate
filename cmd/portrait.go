package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var portraitCmd = &cobra.Command{
	Use:   "portrait",
	Short: "Portrait screen orientation.",
	Long:  `Portrait screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		isInvert, _ := cmd.Flags().GetBool("invert")

		var tmpOut string

		if isInvert {
			tmpOut = sys.PortraitInvert()
		} else {
			tmpOut = sys.Portrait()
		}

		fmt.Println(tmpOut)
	},
}

func init() {
	rootCmd.AddCommand(portraitCmd)

	portraitCmd.Flags().BoolP("invert", "i", false, "Invert portrait orientation.")
}
