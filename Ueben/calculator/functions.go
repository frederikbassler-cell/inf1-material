package calculator

import "math"

func Abs(value int) int {

	if value >= 0 {
		return value

	}

	return -value

}

func IsSquareNumber(value int) bool {

	sqrt := math.Sqrt(float64(value))

	if sqrt != float64(int(sqrt)) {

		return false
	}

	return true
}
