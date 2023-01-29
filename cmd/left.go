package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var leftCmd = &cobra.Command{
	Use:   "left",
	Short: "Rotate 90deg left.",
	Long:  `Rotate 90deg left.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(sys.Left())
	},
}

func init() {
	rootCmd.AddCommand(leftCmd)
}
