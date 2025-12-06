package leetcode3578

const MOD = 1e9 + 7

func countPartitions(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 1

	mx := make([]int, n)
	mn := make([]int, n)
	mxl, mxr := 0, 0
	mnl, mnr := 0, 0

	l := 0
	sum := 0

	for i := 0; i < n; i++ {
		for mxl < mxr && nums[mx[mxr-1]] <= nums[i] {
			mxr--
		}

		for mnl < mnr && nums[mn[mnr-1]] >= nums[i] {
			mnr--
		}

		mx[mxr] = i
		mxr++
		mn[mnr] = i
		mnr++

		for nums[mx[mxl]]-nums[mn[mnl]] > k {
			if mx[mxl] == l {
				mxl++
			}
			if mn[mnl] == l {
				mnl++
			}

			sum = (sum - dp[l] + MOD) % MOD
			l++
		}

		sum = (sum + dp[i]) % MOD
		dp[i+1] = sum
	}

	return dp[n]
}
