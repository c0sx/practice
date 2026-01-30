package lc

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		source   string
		target   string
		original []string
		changed  []string
		cost     []int
		output   int64
	}{
		{"abcd", "acbe", []string{"a", "b", "c", "c", "e", "d"}, []string{"b", "c", "b", "e", "b", "e"}, []int{2, 5, 5, 1, 2, 20}, 28},
		{"abcdefgh", "acdeeghh", []string{"bcd", "fgh", "thh"}, []string{"cde", "thh", "ghh"}, []int{1, 3, 5}, 9},
		{"abcdefgh", "addddddd", []string{"bcd", "defgh"}, []string{"ddd", "ddddd"}, []int{100, 1578}, -1},
	}

	for _, test := range tests {
		if got := minimumCost(test.source, test.target, test.original, test.changed, test.cost); got != test.output {
			t.Errorf("minimumCost(%v, %v, %v, %v, %v) = %v; want %v", test.source, test.target, test.original, test.changed, test.cost, got, test.output)
		}
	}
}
