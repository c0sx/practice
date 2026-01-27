package leetcode3650

import "testing"

func TestMinCost(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:     4,
				edges: [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}},
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				n:     4,
				edges: [][]int{{0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 3}},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		if got := minCost(tt.args.n, tt.args.edges); got != tt.want {
			t.Errorf("minCost(%v, %d) = %v, want %v", tt.args.edges, tt.args.n, got, tt.want)
		}
	}
}
