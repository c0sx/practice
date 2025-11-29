package leetcode3512

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{3, 9, 7},
				k:    5,
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{4, 1, 3},
				k:    4,
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{3, 2},
				k:    6,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
			t.Errorf("minOperations(%v, %d) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
		}
	}
}
