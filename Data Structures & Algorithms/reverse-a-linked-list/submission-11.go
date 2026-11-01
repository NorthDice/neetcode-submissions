/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
	0 1 2 3
	dummmy = head = 0
	prev = nil
	head = 0
	temp = 1
	head.next = nil
	prev = 0
	head = 1

	prev = 0
	head = 1
	temp = 2 
	head.Next = 0
	prev = 1
	head = 2
 0
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}	

	return prev
}
