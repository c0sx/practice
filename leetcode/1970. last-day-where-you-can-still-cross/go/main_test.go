package leetcode1970

import "testing"

func TestLatestDayToCross(t *testing.T) {
	type args struct {
		row   int
		col   int
		cells [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				row:   2,
				col:   2,
				cells: [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}},
			},
			want: 2,
		},
		{
			name: "test case 2",
			args: args{
				row:   2,
				col:   2,
				cells: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{
				row:   3,
				col:   3,
				cells: [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		if got := latestDayToCross(tt.args.row, tt.args.col, tt.args.cells); got != tt.want {
			t.Errorf("latestDayToCross() = %v, want %v", got, tt.want)
		}
	}
}
