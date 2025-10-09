package leetcode3494

import "testing"

func TestMinTime(t *testing.T) {
	tests := []struct {
		skills []int
		mana   []int
		want   int64
	}{
		{
			skills: []int{1, 5, 2, 4},
			mana:   []int{5, 1, 4, 2},
			want:   110,
		},
		{
			skills: []int{1, 1, 1},
			mana:   []int{1, 1, 1},
			want:   5,
		},
		{
			skills: []int{1, 2, 3, 4},
			mana:   []int{1, 2},
			want:   21,
		},
	}

	for _, tt := range tests {
		result := minTime(tt.skills, tt.mana)
		if result != tt.want {
			t.Errorf("minTime(%v, %v) = %d, want %d", tt.skills, tt.mana, result, tt.want)
		}
	}
}
