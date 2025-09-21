package leetcode1912

import "testing"

func TestMovieRentingSystem(t *testing.T) {
	mrs := Constructor(3, [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}})

	if res := mrs.Search(1); !equalSlice(res, []int{1, 0, 2}) {
		t.Fatalf("expect [1,0,2], got %v", res)
	}

	mrs.Rent(0, 1)
	mrs.Rent(1, 2)

	if res := mrs.Report(); !equal2DSlice(res, [][]int{{0, 1}, {1, 2}}) {
		t.Fatalf("expect [[0,1],[1,2]], got %v", res)
	}

	mrs.Drop(1, 2)

	if res := mrs.Search(2); !equalSlice(res, []int{0, 1}) {
		t.Fatalf("expect [0,1], got %v", res)
	}
}

func equalSlice(a, b []int) bool {
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

func equal2DSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !equalSlice(a[i], b[i]) {
			return false
		}
	}

	return true
}
