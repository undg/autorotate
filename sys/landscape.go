package sys

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/oxzi/go-xinput"
)


func Landscape() string {
	screen := "eDP"
	message := "Landscape: placeholder: " + screen

	outXrandr, errXrandr := exec.Command("xrandr", "--output", screen, "--rotate", "normal").Output()

	if errXrandr != nil {
		log.Fatal(errXrandr)
	}

	fmt.Println(string(outXrandr))

	display := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(display)

	for _, device := range xinput.GetXDeviceInfos(display) {
		if device.Use == "slave pointer" {

			outXinput, errXinput := exec.Command("xinput", "set-prop", strconv.FormatUint(device.Id, 10), "--type=foat", "Coordinate Transformation Matrix", "0 0 0 0 0 0 0 0 0").Output()
			
			if errXinput != nil {
				log.Fatal(errXinput)
			}

			fmt.Println(string(outXinput))
		}
	}

	return message
}

