package sys

import (
	"fmt"
	"log"
	"os/exec"
)

func GetXDisplays() string {
	// @TODO (undg) 2023-02-06: get that info directly from X11 C lib
	outXrandr, errXrandr := exec.Command("xrandr", "--listactivemonitors").CombinedOutput()
	if errXrandr != nil {
		fmt.Println(string(outXrandr))
		log.Fatal(errXrandr)
	}
	// @TODO (undg) 2023-02-06: parse it into array
	return string(outXrandr)
}

func LogXDisplays() {
	outXrandr, errXrandr := exec.Command("xrandr", "--listactivemonitors").CombinedOutput()
	if errXrandr != nil {
		fmt.Println(string(outXrandr))
		log.Fatal(errXrandr)
	}

	fmt.Println("DISPLAYS:")
	fmt.Println("ID\tState\tResolution\tName")
	fmt.Println(string(outXrandr))
}
