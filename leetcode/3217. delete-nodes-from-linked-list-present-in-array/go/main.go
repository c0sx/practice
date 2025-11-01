package leetcode3217

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	for head != nil {
		if !numSet[head.Val] {
			break
		}

		head = head.Next
	}

	cursor := head
	for cursor != nil {
		next := cursor.Next

		if next != nil && numSet[next.Val] {
			cursor.Next = next.Next
		} else {
			cursor = cursor.Next
		}
	}

	return head
}
