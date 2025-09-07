package leetcode1304

import "testing"

func TestSumZero(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{5, []int{-7, -1, 1, 3, 4}},
		{3, []int{-1, 0, 1}},
		{1, []int{0}},
	}

	for _, test := range tests {
		result := sumZero(test.n)
		if len(result) != test.n {
			t.Errorf("sumZero(%d) returned %v; expected length %d", test.n, result, test.n)
			continue
		}

		sum := 0
		nums := make(map[int]bool)
		for _, num := range result {
			sum += num
			if nums[num] {
				t.Errorf("sumZero(%d) returned %v; contains duplicate %d", test.n, result, num)
			}
			nums[num] = true
		}

		if sum != 0 {
			t.Errorf("sumZero(%d) returned %v; sum is %d, expected 0", test.n, result, sum)
		}
	}
}
