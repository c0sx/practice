package leetcode2011

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	tests := []struct {
		operations []string
		want       int
	}{
		{[]string{"--X", "X++", "X++"}, 1},
		{[]string{"++X", "++X", "X++"}, 3},
		{[]string{"X++", "++X", "--X", "X--"}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := finalValueAfterOperations(tt.operations); got != tt.want {
				t.Errorf("finalValueAfterOperations(%v) = %v, want %v", tt.operations, got, tt.want)
			}
		})
	}
}
