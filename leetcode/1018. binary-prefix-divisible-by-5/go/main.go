package leetcode1018

func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	result := make([]bool, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			nums[i] = nums[i] % 10
		} else {
			nums[i] = (nums[i-1]*2 + nums[i]) % 10
		}
	}

	for i, num := range nums {
		if num == 0 || num == 5 {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
