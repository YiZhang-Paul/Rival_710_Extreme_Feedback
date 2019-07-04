package utils

// NextIndex finds next valid index in given collection.
// It prevents index out of bound by wrapping back to 0
func NextIndex(collection []string, index int) int {
	max := len(collection) - 1
	if index >= max {
		return 0
	}
	return index + 1
}
