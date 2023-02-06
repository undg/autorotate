package sys

import (
	"fmt"
)

func GetVersion() string {
	return "0.0.1"
}

func LogVersion() {
	fmt.Println("Version:")
	fmt.Println(GetVersion())
}
