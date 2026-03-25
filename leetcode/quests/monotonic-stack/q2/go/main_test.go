package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		temperatures []int
		want         []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
		{
			temperatures: []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70},
			want:         []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0},
		},
	}

	for _, tt := range tests {
		got := dailyTemperatures(tt.temperatures)
		if !equal(got, tt.want) {
			t.Errorf("func(%v) = %v, want %v", tt.temperatures, got, tt.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
