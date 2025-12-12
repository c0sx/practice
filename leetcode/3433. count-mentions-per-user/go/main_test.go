package leetcode3433

import "testing"

func TestCountMentions(t *testing.T) {
	tests := []struct {
		numberOfUsers int
		events        [][]string
		expected      []int
	}{
		{
			numberOfUsers: 2,
			events:        [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "71", "HERE"}},
			expected:      []int{2, 2},
		},
		{
			numberOfUsers: 2,
			events:        [][]string{{"MESSAGE", "10", "id1 id0"}, {"OFFLINE", "11", "0"}, {"MESSAGE", "12", "ALL"}},
			expected:      []int{2, 2},
		},
		{
			numberOfUsers: 2,
			events:        [][]string{{"OFFLINE", "10", "0"}, {"MESSAGE", "12", "HERE"}},
			expected:      []int{0, 1},
		},
	}

	for _, test := range tests {
		result := countMentions(test.numberOfUsers, test.events)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("For numberOfUsers=%d and events=%v, expected %v but got %v", test.numberOfUsers, test.events, test.expected, result)
				break
			}
		}
	}
}
