package utils

func Contains(slice []rune, val rune) bool {
	for _, r := range slice {
		if r == val {
			return true
		}
	}
	return false
}
