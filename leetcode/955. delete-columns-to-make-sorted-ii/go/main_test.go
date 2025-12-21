package leetcode955

import "testing"

func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		strs     []string
		expected int
	}{
		{[]string{"ca", "bb", "ac"}, 1},
		{[]string{"xc", "yb", "za"}, 0},
		{[]string{"zyx", "wvu", "tsr"}, 3},
		{[]string{"cccczcaccz", "cdccccbccz", "cdccccbccx"}, 1},
	}

	for _, test := range tests {
		if result := minDeletionSize(test.strs); result != test.expected {
			t.Errorf("For strs %v, expected %d but got %d", test.strs, test.expected, result)
		}
	}
}
