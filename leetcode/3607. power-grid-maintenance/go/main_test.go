package leetcode3607

import "testing"

func TestProcessQueries(t *testing.T) {
	type args struct {
		c           int
		connections [][]int
		queries     [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				c: 5,
				connections: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 5},
				},
				queries: [][]int{
					{1, 3},
					{2, 1},
					{1, 1},
					{2, 2},
					{1, 2},
				},
			},
			want: []int{3, 2, 3},
		},
		{
			name: "example 2",
			args: args{
				c:           3,
				connections: [][]int{},
				queries: [][]int{
					{1, 1}, {2, 1}, {1, 1},
				},
			},
			want: []int{1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processQueries(tt.args.c, tt.args.connections, tt.args.queries); !equal(got, tt.want) {
				t.Errorf("processQueries(%d, %v, %v) = %v, want %v", tt.args.c, tt.args.connections, tt.args.queries, got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
