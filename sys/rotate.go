package sys

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/oxzi/go-xinput"
)

func Rotate(screen string, rotate string, matrix string) string {
	outXrandr, errXrandr := exec.Command("xrandr", "--output", screen, "--rotate", rotate).Output()

	if errXrandr != nil {
		log.Fatal(errXrandr)
	}
	fmt.Println(string(outXrandr))

	display := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(display)

	for _, device := range xinput.GetXDeviceInfos(display) {
		if device.Use == "slave pointer" {

			outXinput, errXinput := exec.Command("sh", "-c", "xinput set-prop "+strconv.FormatUint(device.Id, 10)+" --type=float 'Coordinate Transformation Matrix' "+matrix).Output()

			if errXinput != nil {
				log.Fatal(errXinput)
			}
			fmt.Println(string(outXinput))

		}
	}

	return matrix + " " + rotate
}
