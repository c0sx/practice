package leetcode3539

const MOD = 1_000_000_007

func magicalSum(m int, k int, nums []int) int {
	n := len(nums)

	c := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		c[i] = make([]int, m+1)
	}

	for i := 0; i <= m; i++ {
		c[i][0] = 1
		c[i][i] = 1

		for j := 1; j < i; j++ {
			c[i][j] = (c[i-1][j-1] + c[i-1][j]) % MOD
		}
	}

	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		p[i][0] = 1

		a := nums[i] % MOD
		for t := 1; t <= m; t++ {
			p[i][t] = (p[i][t-1]) * a % MOD
		}
	}

	cur := make([][][]int64, m+1)
	for i := 0; i <= m; i++ {
		cur[i] = make([][]int64, m+1)

		for j := 0; j <= m; j++ {
			cur[i][j] = make([]int64, m+1)
		}
	}

	cur[m][0][0] = 1

	for i := 0; i < n; i++ {
		next := make([][][]int64, m+1)
		for r := 0; r <= m; r++ {
			next[r] = make([][]int64, m+1)
			for c := 0; c <= m; c++ {
				next[r][c] = make([]int64, m+1)
			}
		}

		for r := 0; r <= m; r++ {
			for carry := 0; carry <= m; carry++ {
				for ones := 0; ones <= m; ones++ {
					val := cur[r][carry][ones]
					if val == 0 {
						continue
					}

					for t := 0; t <= r; t++ {
						newr := r - t
						s := carry + t
						bit := s & 1
						newones := ones + bit

						if newones > m {
							continue
						}

						newcarry := s >> 1
						mult := (int64(c[r][t]) * int64(p[i][t])) % MOD
						add := (val * mult) % MOD
						next[newr][newcarry][newones] = (next[newr][newcarry][newones] + add) % MOD
					}
				}
			}
		}

		cur = next
	}

	var ans int = 0
	for carry := 0; carry <= m; carry++ {
		for ones := 0; ones <= m; ones++ {
			val := cur[0][carry][ones]
			if val == 0 {
				continue
			}

			extra := popCount(carry)
			if ones+extra == k {
				ans = (ans + int(val)) % MOD
			}
		}
	}

	return ans
}

func popCount(x int) int {
	cnt := 0
	for x > 0 {
		if x&1 == 1 {
			cnt++
		}
		x >>= 1
	}
	return cnt
}
