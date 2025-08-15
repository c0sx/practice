package leetcode342

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		n      int
		expect bool
	}{
		{1, true},
		{4, true},
		{16, true},
		{1024, true},
		{8, false},
		{5, false},
		{0, false},
	}

	for _, test := range tests {
		if got := isPowerOfFour(test.n); got != test.expect {
			t.Errorf("isPowerOfFour(%d) = %v; want %v", test.n, got, test.expect)
		}
	}
}
