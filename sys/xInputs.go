package sys

import (
	"fmt"

	"github.com/oxzi/go-xinput"
)

func GetXInputs() []xinput.XDeviceInfo {
	display := xinput.XOpenDisplay(nil)
	defer xinput.XCloseDisplay(display)

	return xinput.GetXDeviceInfos(display)
}

func LogXInputs() {
	fmt.Println("INPUTS:")
	fmt.Println("ID\tName")
	for _, device := range GetXInputs() {
		fmt.Printf("%v\t%v\n", device.Id, device.Name)
	}
}
