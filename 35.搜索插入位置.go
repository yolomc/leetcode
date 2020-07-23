/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}
	return i
}

// @lc code=end
