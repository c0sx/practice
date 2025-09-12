package leetcode3227

import "testing"

func TestDoesAliceWin(t *testing.T) {
	tests := []struct {
		s      string
		expect bool
	}{
		{
			s:      "leetcoder",
			expect: true,
		},
		{
			s:      "bbcd",
			expect: false,
		},
	}

	for _, test := range tests {
		if actual := doesAliceWin(test.s); actual != test.expect {
			t.Errorf("doesAliceWin(%v) = %v, expect %v", test.s, actual, test.expect)
		}
	}
}
