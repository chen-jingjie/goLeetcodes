package addTwoNumbers_2

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	num1 := initListNode([]int{2, 4, 3}, 0)
	num2 := initListNode([]int{5, 6, 4}, 0)
	printList(AddTwoNumbers(num1, num2))
}

func printList(l *ListNode) {
	fmt.Println(l.Val)
	if l.Next != nil {
		printList(l.Next)
	}
}

func initListNode(nums []int, target int) *ListNode {
	var listNode ListNode
	listNode.Val = nums[target]

	if len(nums) > target+1 {
		listNode.Next = initListNode(nums, target+1)
	}

	return &listNode
}
