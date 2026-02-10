package leecode3719

func longestBalanced(nums []int) int {
	maxLength := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		odd := make(map[int]int)
		even := make(map[int]int)

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				even[nums[j]]++
			} else {
				odd[nums[j]]++
			}

			if len(odd) == len(even) {
				if j-i+1 > maxLength {
					maxLength = j - i + 1
				}
			}
		}
	}

	return maxLength
}
