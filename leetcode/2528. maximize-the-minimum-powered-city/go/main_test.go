package leetcode2528

import "testing"

func TestMaxPower(t *testing.T) {
	tests := []struct {
		stations []int
		r        int
		k        int
		want     int64
	}{
		{
			stations: []int{1, 2, 4, 5, 0},
			r:        1,
			k:        2,
			want:     5,
		},
		{
			stations: []int{4, 4, 4, 4},
			r:        0,
			k:        3,
			want:     4,
		},
	}

	for _, tt := range tests {
		if got := maxPower(tt.stations, tt.r, tt.k); got != tt.want {
			t.Errorf("maxPower(%v, %d, %d) = %v, want %v", tt.stations, tt.r, tt.k, got, tt.want)
		}
	}
}
