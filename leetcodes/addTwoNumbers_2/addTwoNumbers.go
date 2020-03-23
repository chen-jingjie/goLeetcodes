package addTwoNumbers_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNode(l1, l2, 0)
}

func addNode(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	var listnode ListNode
	var sum int

	if l1 == nil {
		l1 = &ListNode{}
	}

	if l2 == nil {
		l2 = &ListNode{}
	}

	sum = l1.Val + l2.Val + carry
	carry = sum / 10
	listnode.Val = sum % 10

	if l1.Next != nil || l2.Next != nil {
		listnode.Next = addNode(l1.Next, l2.Next, carry)
	} else {
		if carry > 0 {
			listnode.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
		} else {
			listnode.Next = nil
		}
	}
	return &listnode
}
