package sys

import (
	"fmt"
)

const VERSION = "v0.0.1"
const APP_NAME = "autorotate"

func LogVersion() {
	fmt.Println("Version:")
	fmt.Println(APP_NAME, VERSION)
}
