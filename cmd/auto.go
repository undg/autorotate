package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Detect screen orientation.",
	Long:  `Detect screen orientation from axis sensors.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("auto called")
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	autoCmd.Flags().BoolP("daemon", "d", false, "Continuously check orientation in background (every 1sec by default).")
}
