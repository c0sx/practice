package leetcode2749

import "testing"

func TestMakeTheIntegerZero(t *testing.T) {
	tests := []struct {
		num1 int
		num2 int
		want int
	}{
		{
			num1: 3,
			num2: -2,
			want: 3,
		},
		{
			num1: 5,
			num2: 7,
			want: -1,
		},
		{
			num1: 10,
			num2: 10,
			want: -1,
		},
	}

	for _, tt := range tests {
		if got := makeTheIntegerZero(tt.num1, tt.num2); got != tt.want {
			t.Errorf("makeTheIntegerZero() = %v, want %v", got, tt.want)
		}
	}
}
