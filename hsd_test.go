package hsd

import "testing"

func TestStringDistance(t *testing.T) {
	cases := []struct {
		lhs      string
		rhs      string
		expected int
	}{
		{lhs: "abc", rhs: "abc", expected: 0},
		{lhs: "abcd", rhs: "abef", expected: 2},
		{lhs: "abcd", rhs: "cdef", expected: 4},
		{lhs: "abc", rhs: "ab", expected: -1},
	}

	for _, c := range cases {
		result := StringDistance(c.lhs, c.rhs)
		if result != c.expected {
			t.Fatalf("expected: %v, result: %v", c.expected, result)
		}
	}
}
