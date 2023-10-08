package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/undg/autorotate/sys"
)

type DisplayConfig struct {
	Enable  string `mapstructure:"enable"`
	Disable string `mapstructure:"disable"`
}

type InputConfig struct {
	Enable  string `mapstructure:"enable"`
	Disable string `mapstructure:"disable"`
}

type KeyboardConfig struct {
	Name            string `mapstructure:"name"`
	Landscape       bool   `mapstructure:"landscape"`
	LandscapeInvert bool   `mapstructure:"landscape_invert"`
	Portrait        bool   `mapstructure:"portrait"`
	PortraitInvert  bool   `mapstructure:"portrait_invert"`
}

type Config struct {
	Display  DisplayConfig  `mapstructure:"display"`
	Input    InputConfig    `mapstructure:"input"`
	Keyboard KeyboardConfig `mapstructure:"keyboard"`
}

// @TODO (undg) 2023-02-06: move it to config
// Xorg monitor name
var displayName string

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "autorotate",
	Short: "Rotate screen and digitizer on 2in1 laptops.",
	Long: fmt.Sprintf(`Rotate screen and all inputs on 2in1 laptops with Xorg. Touch screen, stylus and other inputs devices,
need to be rotated independently.
%s will rotate screen and all input devices at once.`, sys.APP_NAME),
	Run: func(cmd *cobra.Command, args []string) {
		// isDaemon, _ := cmd.Flags().GetBool("daemon")

		sys.Autorotate(displayName, true)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initViper)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", xdg.ConfigHome+"/autorotate/", "config file (default XDG_CONFIG_HOME/autorotate/autorotate.toml)")
	rootCmd.PersistentFlags().StringVarP(&displayName, "display", "d", "eDP", "X11 Display that will be rotated, for example eDP or LVDS. You can check monitor names with command `autorotate list` or `xrandr --listactivemonitors|awk '{print $NF}'`")

	rootCmd.Flags().Bool("daemon", false, "Continuously check orientation in background (every 1sec by default).")
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func initViper() {
	cfg, err := LoadConfig(cfgFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Display)

	fmt.Println("Config file used in viper:", viper.ConfigFileUsed())
	fmt.Println("Viper config variables:", cfg.Display)
}
