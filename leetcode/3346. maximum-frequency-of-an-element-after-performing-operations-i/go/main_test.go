package leetcode3346

import "testing"

func TestMaxFrequency(t *testing.T) {
	type args struct {
		nums          []int
		k             int
		numOperations int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums:          []int{1, 4, 5},
				k:             1,
				numOperations: 2,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums:          []int{5, 11, 20, 20},
				k:             5,
				numOperations: 1,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k, tt.args.numOperations); got != tt.want {
				t.Errorf("maxFrequency(%v, %d, %d) = %v, want %v", tt.args.nums, tt.args.k, tt.args.numOperations, got, tt.want)
			}
		})
	}
}
