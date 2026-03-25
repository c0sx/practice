package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		want   []int
	}{
		{
			prices: []int{8, 4, 6, 2, 3},
			want:   []int{4, 2, 4, 2, 3},
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			prices: []int{10, 1, 1, 6},
			want:   []int{9, 0, 1, 6},
		},
	}

	for _, tt := range tests {
		got := finalPrices(tt.prices)
		if !equal(got, tt.want) {
			t.Errorf("func(%v) = %v, want %v", tt.prices, got, tt.want)
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
