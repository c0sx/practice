package leetcode33

func search(nums []int, target int) int {
	n := len(nums)
	lo := 0
	hi := n - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		}

		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
