/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	idx := 0
	for i := 1; i < len(nums); i++ {
		if nums[idx] != nums[i] {
			idx++
			nums[idx] = nums[i]
		}

	}
	return idx + 1
}

// @lc code=end
