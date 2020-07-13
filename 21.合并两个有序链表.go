/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{0, l1} //记录链表头部节点
	index := head            //当前节点，当前节点next默认指向l1
	backup := l2             //记录备用链表头部节点，备用链表头节点默认为l2
	var temp *ListNode       //临时存放链表头节点

	for {
		if index.Next == nil {
			//如果当前链表节点已经处于最末端，则next直接指向备用链表后跳出循环
			index.Next = backup
			break
		} else if index.Next.Val > backup.Val {
			//如果当前节点next的值大于备用链表头节点的值，则链表互换，
			temp = index.Next
			index.Next = backup
			backup = temp
		} else {
			//如果当前节点不在末端，next节点值也不大于备用链表头节点的值，则当前节点后移
			index = index.Next
		}

	}

	return head.Next
}

// @lc code=end
