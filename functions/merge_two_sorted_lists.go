package functions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 暴力法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	targetHead := &ListNode{}
	pointer := targetHead
	for {
		if l1 == nil && l2 == nil {
			return targetHead.Next
		}
		if l1 == nil && l2 != nil {
			pointer.Next = &ListNode{}
			pointer.Next.Val = l2.Val
			l2 = l2.Next
		}
		if l1 != nil && l2 == nil {
			pointer.Next = &ListNode{}
			pointer.Next.Val = l1.Val
			l1 = l1.Next
		}
		if l1 != nil && l2 != nil {
			pointer.Next = &ListNode{}
			switch l1.Val < l2.Val {
			case true:
				pointer.Next.Val = l1.Val
				l1 = l1.Next
			default:
				pointer.Next.Val = l2.Val
				l2 = l2.Next
			}
		}
		pointer = pointer.Next
	}
}
