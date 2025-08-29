package leetcode3021

import "testing"

func TestFlowerGame(t *testing.T) {
	tests := []struct {
		args []int
		want int64
	}{
		{
			args: []int{3, 2},
			want: 3,
		},
		{
			args: []int{1, 1},
			want: 0,
		},
	}

	for _, tt := range tests {
		if got := flowerGame(tt.args[0], tt.args[1]); got != tt.want {
			t.Errorf("flowerGame() = %v, want %v", got, tt.want)
		}
	}
}
