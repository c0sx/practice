package leetcode3315

func minBitwiseArray(nums []int) []int {
	for i, num := range nums {
		result := -1
		d := 1

		for (num & d) != 0 {
			result = num - d
			d <<= 1
		}

		nums[i] = result
	}

	return nums
}
