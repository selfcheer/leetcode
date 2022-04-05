package linked_list
/**
* 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
* https://leetcode-cn.com/problems/remove-linked-list-elements/
*
*/

func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val:-1,
		Next:head,
	}
	cur := dummy.Next
	prev:= dummy
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		}else {
			prev = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
