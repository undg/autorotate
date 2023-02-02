package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/undg/autorotate/sys"
)

const APP_NAME = "autorotate"


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "autorotate",
	Short: "Rotate screen and digitizer.",
	Long:  fmt.Sprintf(`Rotate screen and all inputs for 2in1 laptops with touch screen, stylus devices and other inputs. In Xorg you need to rotate them independently with xrandr and xinput. %s will rotate screen and all input devices at once.`, APP_NAME),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(sys.InAccel())
	fmt.Println(sys.GetOrientation(sys.InAccel()))
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.autorotate.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
