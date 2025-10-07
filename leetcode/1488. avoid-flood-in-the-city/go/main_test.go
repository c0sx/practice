package leetcode1488

import "testing"

func TestAvoidFlood(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3, 4},
			want:  []int{-1, -1, -1, -1},
		},
		{
			input: []int{1, 2, 0, 0, 2, 1},
			want:  []int{-1, -1, 2, 1, -1, -1},
		},
		{
			input: []int{1, 2, 0, 1, 2},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		result := avoidFlood(tt.input)

		if !equal(result, tt.want) {
			t.Errorf("avoidFlood(%v) = %v, want %v", tt.input, result, tt.want)
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
