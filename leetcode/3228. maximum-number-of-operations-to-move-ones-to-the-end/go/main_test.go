package leetcode3228

import "testing"

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{
			s:      "1001101",
			expect: 4,
		},
		{
			s:      "00111",
			expect: 0,
		},
	}

	for _, test := range tests {
		if output := maxOperations(test.s); output != test.expect {
			t.Errorf("maxOperations(%v) = %v, expect %v", test.s, output, test.expect)
		}
	}
}
