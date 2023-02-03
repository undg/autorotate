package cmd

import (
	"fmt"
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
		isDebug, _ := cmd.Flags().GetBool("debug")

		if isDaemon {
			ticker := time.NewTicker(1 * time.Second)
			for {
				select {
				case <-ticker.C:
					if isDebug {
						fmt.Println(sys.SetOrientation())
					} else {
						sys.SetOrientation()
					}
				}
			}
		} else {
			if isDebug {
				fmt.Println(sys.SetOrientation())
			} else {
				sys.SetOrientation()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)

	autoCmd.Flags().BoolP("daemon", "d", false, "Continuously check orientation in background (every 1sec by default).")
	autoCmd.Flags().BoolP("debug", "", false, "Debug")
}
