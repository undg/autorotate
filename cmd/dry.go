package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dryCmd = &cobra.Command{
	Use: "dry",
	Short: "Dry run",
	Long: "Dry run, aka no noting",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Do nothing")
	},
}

func init() {
	rootCmd.AddCommand(dryCmd)
}
