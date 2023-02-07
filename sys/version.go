package sys

import (
	"fmt"
)

const VERSION = "v0.0.3"
const APP_NAME = "autorotate"

func LogVersion(short bool) {
	if short {
		fmt.Println(VERSION)
	} else {
		fmt.Println("Version:")
		fmt.Println(APP_NAME, VERSION)
	}
}
