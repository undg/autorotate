/*
Copyright © 2023 Bartek Laskowski <dev@undg.xyz>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// portraitCmd represents the portrait command
var portraitCmd = &cobra.Command{
	Use:   "portrait",
	Short: "Portrait screen orientation.",
	Long:  `Portrait screen orientation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("portrait called")
	},
}

func init() {
	rootCmd.AddCommand(portraitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portraitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	portraitCmd.Flags().BoolP("invert", "i", false, "Invert portrait orientation.")
}
