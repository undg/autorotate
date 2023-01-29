/*
Copyright © 2023 Bartek Laskowski <dev@undg.xyz>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// landscapeCmd represents the landscape command
var landscapeCmd = &cobra.Command{
	Use:   "landscape",
	Short: "Landscape screen orientation.",
	Long:  `Landscape screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("landscape called")
	},
}

func init() {
	rootCmd.AddCommand(landscapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// landscapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	landscapeCmd.Flags().BoolP("invert", "i", false, "Invert portrait orientation.")
}
