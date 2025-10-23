package leetcode3461

func hasSameDigits(s string) bool {
	n := len(s)
	arr := []byte(s)

	for i := 1; i <= n-2; i++ {
		for j := 0; j <= n-1-i; j++ {
			a := arr[j] - '0'
			b := arr[j+1] - '0'
			c := (a + b) % 10

			arr[j] = c
		}
	}

	return arr[0] == arr[1]
}
