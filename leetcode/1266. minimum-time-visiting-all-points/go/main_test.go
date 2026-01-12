package leetcode1266

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	tests := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			want:   7,
		},
		{
			points: [][]int{{3, 2}, {-2, 2}},
			want:   5,
		},
	}

	for _, tt := range tests {
		if got := minTimeToVisitAllPoints(tt.points); got != tt.want {
			t.Errorf("minTimeToVisitAllPoints(%v) = %v, want %v", tt.points, got, tt.want)
		}
	}
}
