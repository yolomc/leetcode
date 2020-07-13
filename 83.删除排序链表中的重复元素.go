/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	node := head

	for node != nil {
		for node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		}
		node = node.Next
	}

	return head
}

// @lc code=end
