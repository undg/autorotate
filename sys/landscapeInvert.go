package sys

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/oxzi/go-xinput"
)

func LandscapeInvert() string {
	screen := "eDP"
	message := "Landscape: placeholder: " + screen

	cmdXrandr := exec.Command("xrandr", "--output", screen, "--rotate", "inverted")
	errXrandr := cmdXrandr.Run()
	if errXrandr != nil {
		log.Fatal(errXrandr)
	}

	display := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(display)
	for _, device := range xinput.GetXDeviceInfos(display) {
		if device.Use == "slave pointer" {
			log.Println()
			cmdXinput := exec.Command("xinput", "set-prop", strconv.FormatUint(device.Id, 10), "--type=foat", "Coordinate Transformation Matrix", "-1 0 1 0 -1 1 0 0 1")
			errXinput := cmdXinput.Run()
			if errXinput != nil {
				log.Fatal(errXinput)
			}
		}
	}

	return message
}
