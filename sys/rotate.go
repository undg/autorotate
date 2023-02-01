package sys

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/oxzi/go-xinput"
)

// Rotate screen and input devices. Uses xorg utilities xrandr and xinput.
//
// screen name can be checked in xrandr
//
// rotate screen "left" "right" "normal" "inverted"
//
// matrix for input orientation "0 0 0 0 0 0 0 0 0"
func Rotate(screen string, rotate string, matrix string) {
	_, errXrandr := exec.Command("xrandr", "--output", screen, "--rotate", rotate).Output()

	if errXrandr != nil {
		log.Fatal(errXrandr)
	}

	display := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(display)

	for _, device := range xinput.GetXDeviceInfos(display) {
		if device.Use == "slave pointer" {

			_, errXinput := exec.Command("sh", "-c", "xinput set-prop "+strconv.FormatUint(device.Id, 10)+" --type=float 'Coordinate Transformation Matrix' "+matrix).Output()

			if errXinput != nil {
				log.Fatal(errXinput)
			}
		}
	}
}
