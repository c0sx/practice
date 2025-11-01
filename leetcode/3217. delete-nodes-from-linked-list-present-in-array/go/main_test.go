package leetcode3217

import "testing"

func TestModifiedList(t *testing.T) {
	tests := []struct {
		nums []int
		head *ListNode
		want *ListNode
	}{
		{
			nums: []int{1, 2, 3},
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5}},
		},
		{
			nums: []int{1},
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}},
		},
		{
			nums: []int{5},
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}

	for _, tt := range tests {
		got := modifiedList(tt.nums, tt.head)
		if !equal(got, tt.want) {
			t.Errorf("modifiedList(%v, %v) = %v; want %v", tt.nums, tt.head, got, tt.want)
		}
	}
}

func equal(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return a == nil && b == nil
}
