package leetcode837

import "testing"

func TestNew21Game(t *testing.T) {
	tests := []struct {
		n, k, maxPts int
		expected     float64
	}{
		{0, 0, 1, 1},
		{10, 1, 10, 1.0},
		{6, 1, 10, 0.6},
		{21, 17, 10, 0.732778},
	}

	for _, test := range tests {
		result := new21Game(test.n, test.k, test.maxPts)
		epsilon := 1e-6
		if abs(result-test.expected) > epsilon {
			t.Errorf("new21Game(%d, %d, %d) = %f; want %f", test.n, test.k, test.maxPts, result, test.expected)
		}
	}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
