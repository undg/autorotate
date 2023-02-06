package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Detect screen orientation.",
	Long:  `Detect screen orientation from axis sensors.`,
	Run: func(cmd *cobra.Command, args []string) {
		isDaemon, _ := cmd.Flags().GetBool("daemon")

		if isDaemon {
			ticker := time.NewTicker(1 * time.Second)
			for {
				select {
				case <-ticker.C:
				sys.SetOrientation(DisplayName)
				}
			}
		} else {
			sys.SetOrientation(DisplayName)
		}
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)

	autoCmd.Flags().Bool("daemon", false, "Continuously check orientation in background (every 1sec by default).")
}
