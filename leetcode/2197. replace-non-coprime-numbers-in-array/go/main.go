package leetcode2197

func replaceNonCoprimes(nums []int) []int {
	stack := []int{}

	for _, num := range nums {
		for len(stack) > 0 {
			g := gcd(stack[len(stack)-1], num)

			if g > 1 {
				num = lcm(stack[len(stack)-1], num)
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, num)
	}

	return stack
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
