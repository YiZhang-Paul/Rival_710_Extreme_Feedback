package utils

// MaxInt takes random number of values and returns the max value
func MaxInt(values ...int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// MaxUint8 takes random number of values and returns the max value
func MaxUint8(values ...uint8) uint8 {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

// MinInt takes random number of values and returns the min value
func MinInt(values ...int) int {
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}
