/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var temp *ListNode
	preNode := head
	idxNode := head

	for idxNode != nil {
		idxNode = idxNode.Next
		preNode.Next = temp
		temp = preNode
		preNode = idxNode
	}

	return temp
}

// @lc code=end
