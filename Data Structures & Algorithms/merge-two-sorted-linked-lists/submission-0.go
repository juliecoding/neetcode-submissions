/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var tail = &ListNode{}
	current := tail

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		
		current.Next = list2
	}
	// The else is safe because if list1 == nil, then list2 is either:
		// non-nil, containing the remaining nodes, or
		// nil, in which case assigning current.Next = nil is harmless.
	return tail.Next
}

