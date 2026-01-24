package q1

import "testing"

func TestBuildArray(t *testing.T) {
	tests := []struct {
		target []int
		n      int
		want   []string
	}{
		{[]int{1, 3}, 3, []string{"Push", "Push", "Pop", "Push"}},
		{[]int{1, 2, 3}, 3, []string{"Push", "Push", "Push"}},
		{[]int{1, 2}, 4, []string{"Push", "Push"}},
	}

	for _, tt := range tests {
		if got := buildArray(tt.target, tt.n); !equalSlices(got, tt.want) {
			t.Errorf("buildArray(%v, %d) = %v, want %v", tt.target, tt.n, got, tt.want)
		}
	}
}

func equalSlices(a, b []string) bool {
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
