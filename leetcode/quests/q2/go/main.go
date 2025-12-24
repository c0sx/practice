package leetcodeq2

func shuffle(nums []int, n int) []int {
	m := len(nums)
	result := make([]int, m)

	l := 0
	r := n
	for i := 0; i < m; i++ {
		if i%2 == 0 {
			result[i] = nums[l]
			l++
		} else {
			result[i] = nums[r]
			r++
		}
	}

	return result
}
