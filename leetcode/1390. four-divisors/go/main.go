package leetcode1390

func sumFourDivisors(nums []int) int {
	sum := 0

	for _, num := range nums {
		divisors := getDivisors(num)
		if len(divisors) == 4 {
			for _, d := range divisors {
				sum += d
			}
		}
	}

	return sum
}

func getDivisors(n int) []int {
	divisors := []int{}

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	return divisors
}
