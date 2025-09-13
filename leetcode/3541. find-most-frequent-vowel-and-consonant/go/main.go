package leetcode3541

func maxFreqSum(s string) int {
	vowels := [26]bool{}
	vowels['a'-'a'] = true
	vowels['e'-'a'] = true
	vowels['i'-'a'] = true
	vowels['o'-'a'] = true
	vowels['u'-'a'] = true

	var counter [26]int
	v := 0
	c := 0

	for _, r := range s {
		index := r - 'a'
		counter[index]++

		if vowels[index] {
			if counter[index] > v {
				v = counter[index]
			}
		} else {
			if counter[index] > c {
				c = counter[index]
			}
		}
	}

	return c + v
}
