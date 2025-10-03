package leetcode407

import "testing"

func TestTrapRainWater(t *testing.T) {
	tests := []struct {
		input [][]int
		want  int
	}{
		{
			input: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}},
			want:  4,
		},
		{
			input: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}},
			want:  10,
		},
	}

	for _, tt := range tests {
		result := trapRainWater(tt.input)

		if result != tt.want {
			t.Errorf("trapRainWater() = %v, want %v", result, tt.want)
		}
	}
}
