package linklist

type ListNode struct {
	Val  any
	Next *ListNode
}

// NewLinkedListFromSlice 从数组构造链表
func NewLinkedListFromSlice(arr []int) *ListNode {
	var head, prev *ListNode

	for _, val := range arr {
		node := &ListNode{Val: val}

		if head == nil {
			head = node
		} else {
			prev.Next = node
		}

		prev = node
	}

	return head
}

// ValuesFromLinkedList 将链表转换为切片
func ValuesFromLinkedList(head *ListNode) []int {
	values := []int{}

	current := head
	for current != nil {
		values = append(values, current.Val.(int))
		current = current.Next
	}

	return values
}
