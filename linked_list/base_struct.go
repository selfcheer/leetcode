package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}



func SliceToListNode(nums []int) (head *ListNode) {
	var cur *ListNode
	for _, num := range nums {
		if head == nil {
			head = &ListNode{
				Val:  num,
				Next: nil,
			}
			cur = head
		}else {
			cur.Next = &ListNode{
				Val:  num,
				Next: nil,
			}
			cur = cur.Next
		}
	}
	return head
}

func ListNodeToSlice(head *ListNode) (nums []int) {
	nums = make([]int,0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}