package utils

import "math"

func Compare(a, b float32) int {
	if a == b {
		return 0
	}

	eps := 0.001
	a64, b64 := float64(a), float64(b)
	diff := a64 - b64
	abs := math.Abs(diff)

	if abs > eps {
		if diff > 0 {
			return 1
		}
		return -1
	}

	return 0
}
