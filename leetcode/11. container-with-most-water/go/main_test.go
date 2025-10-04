package leetcode11

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			height: []int{1, 1},
			want:   1,
		},
	}

	for _, tt := range tests {
		result := maxArea(tt.height)

		if result != tt.want {
			t.Errorf("maxArea() = %d, expected %d", result, tt.want)
		}
	}
}
