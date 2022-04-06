package linked_list

import "fmt"

func reverseList(head *ListNode) *ListNode {
	var prev, next, curr *ListNode
	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}

func TestReverseList() {
	nums := []int{1,2,3,4,5}
	fmt.Println(nums)
	listNodes := SliceToListNodes(nums)
	reversed := reverseList(listNodes)
	reversedNums := ListNodesToSlice(reversed)
	fmt.Println(reversedNums)
}