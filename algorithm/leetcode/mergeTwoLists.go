package leetcode

//*************************************************************************
/*
https://leetcode.com/problems/merge-two-sorted-lists/#/description

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
*/
//*************************************************************************

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := new(ListNode)
	next := result
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			next.Next = l1
			l1 = l1.Next
		} else {
			next.Next = l2
			l2 = l2.Next
		}
		next = next.Next
	}
	if l1 != nil {
		next.Next = l1
	} else {
		next.Next = l2
	}

	return result.Next
}

/*
 	var head *ListNode = new(ListNode)
	var next *ListNode = head
	for nil != l1 && nil != l2 {
		if l1.Val  < l2.Val {
			next.Next = l1
			l1 = l1.Next
		}else{
			next.Next = l2
			l2 = l2.Next
		}
		next = next.Next
	}
	if nil != l1{
		next.Next = l1
	}else{
		next.Next = l2
	}
	return head.Next
*/
