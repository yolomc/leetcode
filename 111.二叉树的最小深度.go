/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	list := []*TreeNode{root}
	level := 1
ForLoop:
	for {
		for _, node := range list {
			list = list[1:]
			if node.Left == nil && node.Right == nil {
				break ForLoop
			}

			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		level++
	}

	return level
}

// @lc code=end
