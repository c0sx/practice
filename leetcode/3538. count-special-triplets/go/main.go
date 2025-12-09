package leetcode3538

const MOD = 1e9 + 7

func specialTriplets(nums []int) int {
	numCnt := make(map[int]int)
	numPartialCnt := make(map[int]int)

	for _, v := range nums {
		numCnt[v]++
	}

	ans := 0
	for _, v := range nums {
		target := v * 2
		lCnt := numPartialCnt[target]
		numPartialCnt[v]++
		rCnt := numCnt[target] - numPartialCnt[target]

		ans = (ans + lCnt*rCnt) % MOD
	}

	return ans
}
