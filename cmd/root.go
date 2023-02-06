package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const APP_NAME = "autorotate"

var Display string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "autorotate",
	Short: "Rotate screen and digitizer on 2in1 laptops.",
	Long: fmt.Sprintf(`Rotate screen and all inputs on 2in1 laptops with Xorg. Touch screen, stylus and other inputs devices,
need to be rotated independently.
%s will rotate screen and all input devices at once.`, APP_NAME),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Display, "display", "d", "eDP", "Display that will be rotated, for example eDP or LVDS. You can check it with `autorotate list` or `xrandr --listactivemonitors|awk '{print $NF}'`")
}
