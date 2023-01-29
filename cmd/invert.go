/*
Copyright © 2023 Bartek Laskowski <dev@undg.xyz>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// invertCmd represents the invert command
var invertCmd = &cobra.Command{
	Use:   "invert",
	Short: "Rotate 180deg (upside down).",
	Long:  `Rotate 180deg (upside down).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("invert called")
	},
}

func init() {
	rootCmd.AddCommand(invertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// invertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// invertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
