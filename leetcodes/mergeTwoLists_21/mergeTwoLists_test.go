package mergeTwoLists_21

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	result := MergeTwoLists(initListNode([]int{1, 2, 4}, 0), initListNode([]int{1, 3, 4}, 0))
	printList(result)
}

func initListNode(nums []int, target int) *ListNode {
	var listNode ListNode
	listNode.Val = nums[target]

	if len(nums) > target+1 {
		listNode.Next = initListNode(nums, target+1)
	}

	return &listNode
}

func printList(l *ListNode) {
	fmt.Println(l.Val)
	if l.Next != nil {
		printList(l.Next)
	}
}
