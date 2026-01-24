package q3

import "testing"

func Test_exclusiveTime(t *testing.T) {
	tests := []struct {
		name string
		n    int
		logs []string
		want []int
	}{
		{
			name: "test1",
			n:    2,
			logs: []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
			want: []int{3, 4},
		},
		{
			name: "test2",
			n:    1,
			logs: []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
			want: []int{8},
		},
		{
			name: "test3",
			n:    2,
			logs: []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
			want: []int{7, 1},
		},
	}

	for _, tt := range tests {
		if got := exclusiveTime(tt.n, tt.logs); !equal(got, tt.want) {
			t.Errorf("exclusiveTime(%d, %v) = %v, want %v", tt.n, tt.logs, got, tt.want)
		}
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
