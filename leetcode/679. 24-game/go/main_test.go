package leetcode679

import "testing"

func TestJudgePoint24(t *testing.T) {
	tests := []struct {
		cards []int
		want  bool
	}{
		{[]int{4, 1, 8, 7}, true},
		{[]int{1, 2, 1, 2}, false},
	}

	for _, tt := range tests {
		if got := judgePoint24(tt.cards); got != tt.want {
			t.Errorf("judgePoint24(%v) = %v; want %v", tt.cards, got, tt.want)
		}
	}
}
