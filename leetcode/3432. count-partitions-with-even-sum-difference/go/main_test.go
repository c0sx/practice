package leetcode3432

import "testing"

func TestCountPartitions(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{10, 10, 3, 7, 6}},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2, 2}},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{nums: []int{2, 4, 6, 8}},
			want: 3,
		},
	}

	for _, tt := range tests {
		if got := countPartitions(tt.args.nums); got != tt.want {
			t.Errorf("countPartitions(%v) = %v, want %v", tt.args.nums, got, tt.want)
		}
	}
}
