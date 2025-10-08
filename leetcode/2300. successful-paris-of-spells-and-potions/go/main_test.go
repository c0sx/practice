package leetcode2300

import "testing"

func TestSuccessfulPairs(t *testing.T) {
	tests := []struct {
		spells  []int
		potions []int
		success int64
		want    []int
	}{
		{
			spells:  []int{5, 1, 3},
			potions: []int{1, 2, 3, 4, 5},
			success: 7,
			want:    []int{4, 0, 3},
		},
		{
			spells:  []int{3, 1, 2},
			potions: []int{8, 5, 8},
			success: 16,
			want:    []int{2, 0, 2},
		},
	}

	for _, tt := range tests {
		result := successfulPairs(tt.spells, tt.potions, tt.success)
		if !equal(result, tt.want) {
			t.Errorf("successfulPairs(%v, %v) = %v, want = %v", tt.spells, tt.potions, result, tt.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
