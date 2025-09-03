package leetcode3027

import "testing"

func TestNumberOfPairs(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			want:   0,
		},
		{
			points: [][]int{{6, 2}, {4, 4}, {2, 6}},
			want:   2,
		},
		{
			points: [][]int{{3, 1}, {1, 1}, {1, 3}},
			want:   2,
		},
	}
	for _, tt := range tests {
		if got := numberOfPairs(tt.points); got != tt.want {
			t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
		}
	}
}
