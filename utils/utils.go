// Package utils - Awesome comment
package utils

func Contains(x []string, a string) bool {
	for _, n := range x {
		if n == a {
			return true
		}
	}
	return false
}

func ContainsInt(x []int, a int) bool {
	for _, n := range x {
		if n == a {
			return true
		}
	}
	return false
}
