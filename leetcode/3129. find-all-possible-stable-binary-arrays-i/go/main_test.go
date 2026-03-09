package lc

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		zero  int
		one   int
		limit int
		want  int
	}{
		{
			zero:  1,
			one:   1,
			limit: 2,
			want:  2,
		},
		{
			zero:  1,
			one:   2,
			limit: 1,
			want:  1,
		},
		{
			zero:  3,
			one:   3,
			limit: 2,
			want:  14,
		},
	}

	for _, tt := range tests {
		got := numberOfStableArrays(tt.zero, tt.one, tt.limit)
		if got != tt.want {
			t.Errorf("%d, %d, %d = %d, want %d", tt.zero, tt.one, tt.limit, got, tt.want)
		}
	}
}
