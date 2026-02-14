package lc

import "testing"

func TestChampagneTower(t *testing.T) {
	tests := []struct {
		poured      int
		query_row   int
		query_glass int
		want        float64
	}{
		{
			poured:      1,
			query_row:   1,
			query_glass: 1,
			want:        0.0,
		},
		{
			poured:      2,
			query_row:   1,
			query_glass: 1,
			want:        0.5,
		},
		{
			poured:      100000009,
			query_row:   33,
			query_glass: 17,
			want:        1.00000,
		},
	}

	for _, tt := range tests {
		got := champagneTower(tt.poured, tt.query_row, tt.query_glass)
		if got != tt.want {
			t.Errorf("champagneTower(%d, %d, %d) = %f, want %f", tt.poured, tt.query_row, tt.query_glass, got, tt.want)
		}
	}
}
