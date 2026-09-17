/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	l := dummy
	r := head

	for range n {
		r = r.Next
	}

	for r != nil {
		r = r.Next
		l = l.Next
	}

	l.Next = l.Next.Next

	return dummy.Next
}
