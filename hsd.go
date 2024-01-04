// Package hsd provides functions to calculate the distance between strings.
package hsd

// StringDistance calculates the distance between two strings, lhs and rhs.
// The distance is defined as the count of differing characters.
// Returns -1 if the lengths of the strings are different.
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// Distance calculates the distance between two rune slices, a and b.
// The distance is defined as the count of differing runes.
// Returns -1 if the lengths of the slices are different.
func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}
