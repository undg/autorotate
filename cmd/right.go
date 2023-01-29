/*
Copyright © 2023 Bartek Laskowski <dev@undg.xyz>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rightCmd represents the right command
var rightCmd = &cobra.Command{
	Use:   "right",
	Short: "Rotate 90deg right.",
	Long: `Rotate 90deg right.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("right called")
	},
}

func init() {
	rootCmd.AddCommand(rightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rightCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rightCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
