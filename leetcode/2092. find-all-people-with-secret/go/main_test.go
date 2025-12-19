package leetcode2092

import "testing"

func TestFindAllPeople(t *testing.T) {
	tests := []struct {
		n           int
		meetings    [][]int
		firstPerson int
		want        []int
	}{
		{
			n:           6,
			meetings:    [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3, 5},
		},
		{
			n:           4,
			meetings:    [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}},
			firstPerson: 3,
			want:        []int{0, 1, 3},
		},
		{
			n:           5,
			meetings:    [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		if got := findAllPeople(tt.n, tt.meetings, tt.firstPerson); !equalSlices(got, tt.want) {
			t.Errorf("findAllPeople() = %v, want %v", got, tt.want)
		}
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		if m[v] == 0 {
			return false
		}

		m[v]--
	}

	return true
}
