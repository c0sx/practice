package leetcode2943

import "testing"

func TestMaximizeSquareHoleArea(t *testing.T) {
	tests := []struct {
		n, m    int
		hBars   []int
		vBars   []int
		wantAns int
	}{
		{
			n:       2,
			m:       1,
			hBars:   []int{2, 3},
			vBars:   []int{2},
			wantAns: 4,
		},
		{
			n:       1,
			m:       1,
			hBars:   []int{2},
			vBars:   []int{2},
			wantAns: 4,
		},
		{
			n:       2,
			m:       3,
			hBars:   []int{2, 3},
			vBars:   []int{2, 4},
			wantAns: 4,
		},
	}

	for _, tt := range tests {
		if ans := maximizeSquareHoleArea(tt.n, tt.m, tt.hBars, tt.vBars); ans != tt.wantAns {
			t.Errorf("got %v, want %v", ans, tt.wantAns)
		}
	}
}
