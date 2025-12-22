package leetcode960

import "testing"

func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		strs     []string
		expected int
	}{
		{[]string{"babca", "bbazb"}, 3},
		{[]string{"edcba"}, 4},
		{[]string{"ghi", "def", "abc"}, 0},
		{[]string{"aabbaa", "baabab", "aaabab"}, 3},
	}

	for _, test := range tests {
		if result := minDeletionSize(test.strs); result != test.expected {
			t.Errorf("For strs %v, expected %d but got %d", test.strs, test.expected, result)
		}
	}
}
