package sys

import (
	"fmt"
	"log"
	"os/exec"
	// "github.com/oxzi/go-xinput"
)


func Rotate(screen string, rotate string, matrix string) string {
	outXrandr, errXrandr := exec.Command("xrandr", "--output", screen, "--rotate", rotate).Output()

	if errXrandr != nil {
		log.Fatal(errXrandr)
	}

	fmt.Println(string(outXrandr))


	outXinput, errXinput := exec.Command("xinput", "set-prop", "10", "--type=float", "\"Coordinate Transformation Matrix\"", matrix).Output()

	if errXinput != nil {
		log.Fatal(errXinput)
	}
	fmt.Println(outXinput)

	// display := xinput.XOpenDisplay(nil)
	// defer xinput.XCloseDisplay(display)
	//
	// for _, device := range xinput.GetXDeviceInfos(display) {
	// 	if device.Use == "slave pointer" {
	//
	// 		cmdXinput := exec.Command("xinput", "set-prop", "10", "--type=float", "\"Coordinate Transformation Matrix\"", matrix)
	//
	// 		errXinput := cmdXinput.Run()
	//
	// 		if errXinput != nil {
	// 			log.Fatal(errXinput)
	// 		}
	//
	// 	}
	// }

	return matrix + rotate
}
