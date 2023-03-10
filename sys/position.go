package sys

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

type X = int
type Y = int
type Z = int

func getInAccel() (X, Y, Z) {
	file_x, _ := filepath.Glob("/sys/bus/iio/devices/*/in_accel_x_raw")
	file_y, _ := filepath.Glob("/sys/bus/iio/devices/*/in_accel_y_raw")
	file_z, _ := filepath.Glob("/sys/bus/iio/devices/*/in_accel_z_raw")

	read_x, _ := ioutil.ReadFile(file_x[0])
	read_y, _ := ioutil.ReadFile(file_y[0])
	read_z, _ := ioutil.ReadFile(file_z[0])

	var x int
	var y int
	var z int
	fmt.Sscan(string(read_x), &x)
	fmt.Sscan(string(read_y), &y)
	fmt.Sscan(string(read_z), &z)

	return x, y, z
}

type Orientation string

const (
	OrientationLeft    Orientation = "left"
	OrientationRight               = "right"
	OrientationInvert              = "invert"
	OrientationNormal              = "normal"
	OrientationUnknown             = "unknown"
)

func Autorotate(display string, isDaemon bool) {
	if isDaemon {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				SetOrientation(display)
			}
		}
	} else {
		SetOrientation(display)
	}
}
func SetOrientation(display string) Orientation {
	x, y, z := getInAccel()

	var orientation Orientation

	if isLandscape(x, y, z) {
		orientation = OrientationNormal
		RotateNormal(display)
		// enable keyboard
	} else if isInvert(x, y, z) {
		orientation = OrientationInvert
		RotateInvert(display)
		// disable keyboard
	} else if isLeft(x, y, z) {
		orientation = OrientationLeft
		RotateLeft(display)
		// disable keyboard
	} else if isRight(x, y, z) {
		orientation = OrientationRight
		RotateRight(display)
		// disable keyboard
	} else {
		orientation = OrientationUnknown
	}

	return Orientation(orientation)
}

func isLandscape(x int, y int, z int) bool {
	xInRange := -3 <= x && x <= 3
	yInRange := y <= -6 // up to -9 ish and bellow

	return xInRange && yInRange
}

func isInvert(x int, y int, z int) bool {
	xInRange := -3 <= x && x <= 3
	yInRange := 6 <= y // up to 9 ish and above

	return xInRange && yInRange
}

func isLeft(x int, y int, z int) bool {
	xInRange := 6 <= x // up to 9 ish and above
	yInRange := -3 <= y && y <= 3

	return xInRange && yInRange
}

func isRight(x int, y int, z int) bool {
	xInRange := x <= -3 // up to -9 ish and bellow
	yInRange := -3 <= y && y <= 3

	return xInRange && yInRange
}
