package leetcode1292

import "testing"

func TestMaxSideLength(t *testing.T) {
	type args struct {
		mat       [][]int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				mat: [][]int{
					{1, 1, 3, 2, 4, 3, 2},
					{1, 1, 3, 2, 4, 3, 2},
					{1, 1, 3, 2, 4, 3, 2},
				},
				threshold: 4,
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				mat: [][]int{
					{2, 2, 2, 2, 2},
					{2, 2, 2, 2, 2},
					{2, 2, 2, 2, 2},
					{2, 2, 2, 2, 2},
					{2, 2, 2, 2, 2},
				},
				threshold: 1,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		if got := maxSideLength(tt.args.mat, tt.args.threshold); got != tt.want {
			t.Errorf("maxSideLength(%v) = %v, want %v", tt.args.mat, got, tt.want)
		}
	}
}
