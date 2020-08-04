/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	ns := [][]int{}
	if root == nil {
		return ns
	}

	list := []*TreeNode{root} //初始化队列
	lenth := 1                //队列长度
	level := 0                //当前层

	for lenth > 0 {
		//循环当前层
		for i := 0; i < lenth; i++ {
			//出队
			node := list[0]
			list = list[1:]

			//追加节点值
			if len(ns) == level {
				ns = append(ns, []int{node.Val})
			} else {
				ns[level] = append(ns[level], node.Val)
			}

			//入队
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}

		lenth = len(list)
		level++
	}

	//链表反转
	nsLen := len(ns)
	for i := 0; i < nsLen/2; i++ {
		ns[i], ns[nsLen-1-i] = ns[nsLen-1-i], ns[i]
	}
	return ns
}

// @lc code=end
