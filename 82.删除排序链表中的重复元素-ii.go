/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
	if head == nil {
		return head
	}
	temp := &ListNode{Val: 0}
	temp.Next = head
	idxNode := temp

	var v int

	for idxNode.Next != nil && idxNode.Next.Next != nil {
		if idxNode.Next.Val == idxNode.Next.Next.Val {
			v = idxNode.Next.Val
			for idxNode.Next != nil && idxNode.Next.Val == v {
				idxNode.Next = idxNode.Next.Next
			}
		} else {
			idxNode = idxNode.Next
		}
	}

	return temp.Next
}

// @lc code=end
