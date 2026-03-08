package lc

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	ans := make([]byte, n)

	for i := 0; i < n; i++ {
		if nums[i][i] == '0' {
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
	}

	return string(ans)
}
