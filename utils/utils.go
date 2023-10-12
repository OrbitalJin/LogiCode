package utils

// Returns a if the condition is met, otherwise b will be returned
func Ternary(a, b, condition bool) any {
	if condition {
		return a
	} else {
		return b
	}
}