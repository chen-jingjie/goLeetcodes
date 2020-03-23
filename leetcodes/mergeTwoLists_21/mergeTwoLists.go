package mergeTwoLists_21

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{-1, nil}
	prev := prehead

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return prehead.Next
}

// 优化空间
//func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//
//	if l1.Val > l2.Val {
//		l2.Next = MergeTwoLists(l1, l2.Next)
//		return l2
//	}
//	l1.Next = MergeTwoLists(l1.Next, l2)
//	return l1
//}
