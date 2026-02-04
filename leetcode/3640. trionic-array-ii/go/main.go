package lc

func maxSumTrionic(nums []int) int64 {
	n := len(nums)
	var maxSum int64 = -1e16

	for i := 1; i < n-2; {
		a, b := i, i

		sum := int64(nums[a])
		for b+1 < n && nums[b+1] < nums[b] {
			b++
			sum += int64(nums[b])
		}

		if b == a {
			i++
			continue
		}

		c := b
		var left, right int64
		lx, rx := int64(-1e16), int64(-1e16)
		for a-1 >= 0 && nums[a-1] < nums[a] {
			a--
			left += int64(nums[a])
			if left > lx {
				lx = left
			}
		}

		if a == i {
			i++
			continue
		}

		for b+1 < n && nums[b+1] > nums[b] {
			b++
			right += int64(nums[b])
			if right > rx {
				rx = right
			}
		}

		if b == c {
			i++
			continue
		}

		if lx+rx+sum > maxSum {
			maxSum = lx + rx + sum
		}

		i = b
	}

	return maxSum
}
