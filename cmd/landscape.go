package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var landscapeCmd = &cobra.Command{
	Use:   "landscape",
	Short: "Landscape screen orientation.",
	Long:  `Landscape screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		isInvert, _ := cmd.Flags().GetBool("invert")

		var tmpOut string

		if isInvert {
			tmpOut = sys.LandscapeInvert()
		} else {
			tmpOut = sys.Landscape()
		}

		fmt.Println(tmpOut)
	},
}

func init() {
	rootCmd.AddCommand(landscapeCmd)

	landscapeCmd.Flags().BoolP("invert", "i", false, "Invert portrait orientation.")
}
