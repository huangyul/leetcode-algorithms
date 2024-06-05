package removeelement

import linklist "github.com/huangyul/leetcode-algorithms/link-list"

// removeelement 203.移除链表元素
func removeelement(head *linklist.ListNode, val int) *linklist.ListNode {
	dummy := &linklist.ListNode{Next: head}
	prev, cur := dummy, head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
