package linked_list

/**
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val和next。val是当前节点的值，next是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性prev以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第index个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为val的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第index个节点之前添加值为val 的节点。如果index等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引index 有效，则删除链表中的第index 个节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type MyLinkedList struct {
	Head *ListNode
	Size int
}


func Constructor() MyLinkedList {
	l := MyLinkedList{
		Head: nil,
		Size: 0,
	}
	return l
}


func (l *MyLinkedList) Get(index int) int {
	if index > l.Size - 1 {
		return -1
	}
	curIndex := 0
	cur := l.Head
	for cur != nil {
		if curIndex == index {
			return cur.Val
		}

		cur = cur.Next
		curIndex ++
	}
	return -1
}


func (l *MyLinkedList) AddAtHead(val int)  {
	l.Head =  &ListNode{
		Val:  val,
		Next: l.Head,
	}
	l.Size ++
}


func (l *MyLinkedList) AddAtTail(val int)  {
	cur := l.Head
	tail :=  &ListNode{
		Val:  val,
		Next: nil,
	}
	l.Size ++
	if cur == nil {
		l.Head = tail
		return
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = tail

}


func (l *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > l.Size {
		return
	}
	if index <= 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.Size {
		l.AddAtTail(val)
		return
	}
	curIndex := 0

	for cur := l.Head; cur != nil; cur = cur.Next{
		if curIndex + 1 == index {
			cur.Next = &ListNode{
				Val:  val,
				Next: cur.Next,
			}
			break
		}
		curIndex ++

	}
	l.Size ++
}


func (l *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= l.Size {
		return
	}
	curIndex := -1
	prev := &ListNode{
		Val:  -1,
		Next: l.Head,
	}
	dummy := prev
	for prev != nil {
		if curIndex + 1 == index {
			prev.Next = prev.Next.Next
			break
		}
		curIndex ++
		prev = prev.Next
	}
	l.Head = dummy.Next
	l.Size --
}