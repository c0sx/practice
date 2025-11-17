package leetcode1437

func kLengthApart(nums []int, k int) bool {
	pos := -k - 1
	for i, num := range nums {
		if num == 1 {
			if i-pos-1 < k {
				return false
			}

			pos = i
		}
	}

	return true
}
