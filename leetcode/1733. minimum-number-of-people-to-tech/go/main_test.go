package leetcode1733

import "testing"

func TestMinimumTeachings(t *testing.T) {
	type args struct {
		n           int
		languages   [][]int
		friendships [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n:           2,
				languages:   [][]int{{1}, {2}, {1, 2}},
				friendships: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				n:           3,
				languages:   [][]int{{2}, {1, 3}, {1, 2}, {3}},
				friendships: [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}},
			},
			want: 2,
		},
		{
			name: "example 3",
			args: args{
				n:           2,
				languages:   [][]int{{1}, {1}, {1}},
				friendships: [][]int{{1, 2}, {2, 3}, {3, 1}},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTeachings(tt.args.n, tt.args.languages, tt.args.friendships); got != tt.want {
				t.Errorf("minimumTeachings() = %v, want %v", got, tt.want)
			}
		})
	}
}
