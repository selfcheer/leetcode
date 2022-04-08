package linked_list

import "fmt"

/**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
https://leetcode-cn.com/problems/swap-nodes-in-pairs/
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val : -1,
		Next : head,
	}
	pre := dummy
	for head != nil && head.Next != nil{
		pre.Next, head.Next, head.Next.Next = head.Next, head.Next.Next, head
		pre, head = head, head.Next
	}
	return dummy.Next
}


func TestSwapPairs() {
	nums := []int{1,2,3,4}
	fmt.Println(nums)
	listNodes := SliceToListNodes(nums)
	swapped := swapPairs(listNodes)
	swappedNums := ListNodesToSlice(swapped)
	fmt.Println(swappedNums)
}