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

			// @TODO (undg) 2023-02-27: dirty quick fix for external mouses. Should solve problem by config/flag (explicit/implicit/all)

			// don't rotate external device:
			switch device.Name {
			case "Razer Razer Basilisk X HyperSpeed": // mouse
				continue
			case "Virtual core XTEST pointer": // barrier/synergy
				continue
			}

			_, errXinput := exec.Command("sh", "-c", "xinput set-prop "+strconv.FormatUint(device.Id, 10)+" --type=float 'Coordinate Transformation Matrix' "+matrix).Output()

			if errXinput != nil {
				log.Fatal(errXinput)
			}
		}
	}
}

func RotateNormal(display string) {
	Rotate(display, "normal", "0 0 0 0 0 0 0 0 0")
}

func RotateInvert(display string) {
	Rotate(display, "inverted", "-1 0 1 0 -1 1 0 0 1")
}

func RotateRight(display string) {
	Rotate(display, "right", "0 1 0 -1 0 1 0 0 1")
}

func RotateLeft(display string) {
	Rotate(display, "left", "0 -1 1 1 0 0 0 0 1")
}
