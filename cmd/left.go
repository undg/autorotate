/*
Copyright © 2023 Bartek Laskowski <dev@undg.xyz>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// leftCmd represents the left command
var leftCmd = &cobra.Command{
	Use:   "left",
	Short: "Rotate 90deg left.",
	Long:  `Rotate 90deg left.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("left called")
	},
}

func init() {
	rootCmd.AddCommand(leftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
