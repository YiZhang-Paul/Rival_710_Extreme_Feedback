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
