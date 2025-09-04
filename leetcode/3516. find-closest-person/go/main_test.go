package leetcode3516

import "testing"

func TestFindClosest(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		z    int
		want int
	}{
		{
			x:    2,
			y:    7,
			z:    4,
			want: 1,
		},
		{
			x:    2,
			y:    5,
			z:    6,
			want: 2,
		},
		{
			x:    1,
			y:    5,
			z:    3,
			want: 0,
		},
	}

	for _, tt := range tests {
		if got := findClosest(tt.x, tt.y, tt.z); got != tt.want {
			t.Errorf("findClosest() = %v, want %v", got, tt.want)
		}
	}
}
