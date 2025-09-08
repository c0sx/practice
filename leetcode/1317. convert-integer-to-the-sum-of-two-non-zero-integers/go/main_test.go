package leetcode1317

import (
	"testing"
)

func TestGetNoZeroIntegers(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{1, 1}},
		{11, []int{2, 9}},
		{10000, []int{1, 9999}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := getNoZeroIntegers(tt.n)

			if len(got) != 2 {
				t.Errorf("len(%d) = %v, want %v", tt.n, got, tt.want)
			}

			if !checkSum(got, tt.n) {
				t.Errorf("checkSum(%d) = %v, want %v", tt.n, got, tt.want)
			}

			if !checkZeros(got) {
				t.Errorf("checkZeros(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func checkSum(a []int, v int) bool {
	t := 0
	for i := range a {
		if a[i] == 0 {
			return false
		}

		t += a[i]
	}

	return t == v
}

func checkZeros(a []int) bool {
	for _, n := range a {
		if n == 0 {
			return false
		}

		for n > 0 {
			if n%10 == 0 {
				return false
			}
			n /= 10
		}
	}

	return true
}
