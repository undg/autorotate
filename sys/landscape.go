package sys

import (
	"log"
	"os/exec"
	"strconv"
)

// SCREEN="eDP"
// TRANSFORMATION="Coordinate Transformation Matrix"

// input_devices=( 10 12 13 20 16 )

// xrandr --output $SCREEN --rotate left
// xinput set-prop $i --type=float "$TRANSFORMATION" 1 0 0 0 1 0 0 0 1

func Landscape() string {

	screen := "eDP"
	devicesIds := [6]int{10, 12, 13, 16, 20, 21}
	message := "Landscape: placeholder: " + screen


	cmdXrandr := exec.Command("xrandr --output", screen, "--rotate left")
	errXrandr := cmdXrandr.Run()
	if errXrandr != nil {  
		log.Fatal(errXrandr)
	}
	

	for _, deviceId := range devicesIds {
		message += strconv.Itoa(deviceId)	

		cmdXinput := exec.Command("xinput set-prop", strconv.Itoa(deviceId), "--type=foat Coordinate Transformation Matrix", "1 0 0 0 1 0 0 0 1")
		errXinput := cmdXinput.Run()
		if errXinput != nil {
			log.Fatal(errXinput)
		}
	}
	return message
}
