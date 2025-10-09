package leetcode3494

func minTime(skill []int, mana []int) int64 {
	s := len(skill)
	m := len(mana)
	times := make([]int64, s)

	for i := 0; i < m; i++ {
		var current int64 = 0
		for j := 0; j < s; j++ {
			if current < times[j] {
				current = times[j]
			}

			current += int64(skill[j]) * int64(mana[i])
		}

		times[s-1] = current
		for j := s - 2; j >= 0; j-- {
			times[j] = times[j+1] - int64(skill[j+1])*int64(mana[i])
		}
	}

	return times[s-1]
}
