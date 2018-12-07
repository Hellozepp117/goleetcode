package functions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	base := head
	var count int
	for count = 1; ; count++ {
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	// delete
	if count == n {
		return base.Next
	}
	num := count - n + 1

	head = base
	for i := 1; ; i++ {
		if i == num-1 {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}
	return base
}
