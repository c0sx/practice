package leetcode3349

import "testing"

func TestHasIncreasingSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7},
				k:    5,
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{-15, 19},
				k:    1,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasIncreasingSubarrays(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("hasIncreasingSubarrays(%v, %d) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
			}
		})
	}
}
