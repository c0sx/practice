package leetcode1733

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	hasCommonLanguages := func(a, b int) bool {
		for _, langA := range languages[a-1] {
			for _, langB := range languages[b-1] {
				if langA == langB {
					return true
				}
			}
		}

		return false
	}

	problematicUsers := make(map[int]bool)
	for _, friendship := range friendships {
		if !hasCommonLanguages(friendship[0], friendship[1]) {
			problematicUsers[friendship[0]] = true
			problematicUsers[friendship[1]] = true
		}
	}

	knownLanguages := make(map[int]int)
	for user := range problematicUsers {
		for _, lang := range languages[user-1] {
			knownLanguages[lang]++
		}
	}

	maxKnownLanguages := 0
	for _, count := range knownLanguages {
		if count > maxKnownLanguages {
			maxKnownLanguages = count
		}
	}

	result := len(problematicUsers) - maxKnownLanguages

	return result
}
