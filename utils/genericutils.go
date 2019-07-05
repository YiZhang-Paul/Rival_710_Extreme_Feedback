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

// TernaryBool is a simple implementation of ternary operator.
// Only use this for trivial assignments as all expressions will be evaluated
func TernaryBool(condition, ifTrue, ifFalse bool) bool {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// TernaryInt is a simple implementation of ternary operator.
// Only use this for trivial assignments as all expressions will be evaluated
func TernaryInt(condition bool, ifTrue, ifFalse int) int {
	if condition {
		return ifTrue
	}
	return ifFalse
}
