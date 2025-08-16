package leetcode1323

import "testing"

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		num    int
		expect int
	}{
		{9669, 9969},
		{9996, 9999},
		{9999, 9999},
		{6969, 9969},
	}

	for _, test := range tests {
		if got := maximum69Number(test.num); got != test.expect {
			t.Errorf("maximum69Number(%d) = %d; want %d", test.num, got, test.expect)
		}
	}
}
