package lc

import "testing"

func TestFindDifferentBinaryString(t *testing.T) {
	tests := []struct {
		nums []string
		want string
	}{
		{
			nums: []string{"01", "10"},
			want: "11",
		},
		{
			nums: []string{"00", "01"},
			want: "10",
		},
		{
			nums: []string{"111", "011", "001"},
			want: "000",
		},
	}

	for _, tt := range tests {
		got := findDifferentBinaryString(tt.nums)
		if got != tt.want {
			t.Errorf("%v = %s, want %s", tt.nums, got, tt.want)
		}
	}
}
