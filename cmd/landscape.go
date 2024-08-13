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
		display, _ := cmd.Flags().GetString("display")

		if isInvert {
			sys.RotateInvert(display)
		} else {
			sys.RotateNormal(display)
		}
	},
}

var normalCmd = &cobra.Command{
	Use:   "normal",
	Short: "Landscape screen orientation.",
	Long:  `Landscape screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		isInvert, _ := cmd.Flags().GetBool("invert")
		display, _ := cmd.Flags().GetString("display")

		if isInvert {
			sys.RotateInvert(display)
		} else {
			sys.RotateNormal(display)
		}
	},
}

var invertCmd = &cobra.Command{
	Use:   "invert",
	Short: "Rotate 180deg (upside down).",
	Long:  `Rotate 180deg (upside down).`,
	Run: func(cmd *cobra.Command, args []string) {
		display, _ := cmd.Flags().GetString("display")

		sys.RotateInvert(display)
	},
}

func init() {
	rootCmd.AddCommand(landscapeCmd)
	rootCmd.AddCommand(normalCmd)
	rootCmd.AddCommand(invertCmd)

	landscapeCmd.Flags().BoolP("invert", "i", false, "Invert landscape orientation.")
}
