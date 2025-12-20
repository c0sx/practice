package leetcode944

import "testing"

func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		strs []string
		want int
	}{
		{
			strs: []string{"cba", "daf", "ghi"},
			want: 1,
		},
		{
			strs: []string{"a", "b"},
			want: 0,
		},
		{
			strs: []string{"zyx", "wvu", "tsr"},
			want: 3,
		},
		{
			strs: []string{"rrjk", "furt", "guzm"},
			want: 2,
		},
	}

	for _, tt := range tests {
		if got := minDeletionSize(tt.strs); got != tt.want {
			t.Errorf("minDeletionSize(%v) = %v, want %v", tt.strs, got, tt.want)
		}
	}
}
