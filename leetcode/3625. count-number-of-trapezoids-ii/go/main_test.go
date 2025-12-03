package leetcode3625

import "testing"

func TestCountTrapezoids(t *testing.T) {
	type args struct {
		points [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				points: [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		if got := countTrapezoids(tt.args.points); got != tt.want {
			t.Errorf("countTrapezoids() = %v, want %v", got, tt.want)
		}
	}
}
