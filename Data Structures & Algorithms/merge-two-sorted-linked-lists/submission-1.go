/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // The sentinel or placeholder. Kind of like a container but it doesn't contain values. 
	tail := dummy // this is the tail node of the combined list; because these are linked lists, it is iterator-like

	for list1 != nil && list2 != nil { // Loop through until you reach the end of at least one list (until one list is exhausted)
		if list1.Val <= list2.Val { // compare what's smaller
			tail.Next = list1	// tack the target onto the "iterator" (the merged list's tail)
			list1 = list1.Next		// move the pointer forward by 1
		} else {
			tail.Next = list2	// Same deal: tack the smaller node onto the iterator 
			list2 = list2.Next		// Advance the pointer in the list we touched
		}

		tail = tail.Next		// THEN, advance the iterator so we can keep tacking things on to it. 
	}

	// Now we have to grab the remainder from whichever list we didn't exhaust (if any)
	if list1 != nil { // This means we finished list2 first
		tail.Next = list1 // add the rest of list1 to the iterator/value holder
	} else {
		// This else is safe because if list1 == nil, then list2 is either:
			// non-nil, containing the remaining nodes, or
			// nil, in which case assigning tail.Next = nil is harmless.
		tail.Next = list2
	}

	return dummy.Next // Where the list actually starts; the sentinel itself is not part of the result
}

