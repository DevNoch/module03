package utils

func contains(x []string, a string) bool {
	for _, n := range x {
		if n == a {
			return true
		}
	}
	return false
}