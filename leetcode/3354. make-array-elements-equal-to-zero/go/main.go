package leetcode3354

func countValidSelections(nums []int) int {
	count := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			continue
		}

		directions := []int{-1, 1}
		for _, dir := range directions {
			tmp := make([]int, len(nums))
			copy(tmp, nums)

			pos := i

			for {
				pos += dir
				if pos < 0 || pos >= len(nums) {
					break
				}

				if tmp[pos] > 0 {
					tmp[pos]--
					dir *= -1
				}
			}

			nonZeros := 0
			for _, val := range tmp {
				if val != 0 {
					nonZeros++
				}
			}

			if nonZeros == 0 {
				count++
			}
		}
	}

	return count
}
