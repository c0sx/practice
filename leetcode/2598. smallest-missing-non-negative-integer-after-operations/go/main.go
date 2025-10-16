package leetcode2598

func findSmallestInteger(nums []int, value int) int {
	remainders := make([]int, value)

	for _, num := range nums {
		rem := (num%value + value) % value
		remainders[rem]++
	}

	mex := 0
	for remainders[mex%value] > 0 {
		remainders[mex%value]--
		mex++
	}

	return mex
}
