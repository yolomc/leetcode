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

	lIdx := 0
	rIdx := len(nums) - 1
	mIdx := rIdx / 2
	mNum := nums[mIdx]
	res := 0
	for {
		switch {
		case mNum > target:
			rIdx = mIdx - 1
		case mNum < target:
			lIdx = mIdx + 1
		default:
			res = mIdx
			goto End
		}

		if lIdx > rIdx {
			res = lIdx
			goto End
		}
		mIdx = (lIdx + rIdx) / 2
		mNum = nums[mIdx]
	}

End:
	return res
}

// @lc code=end
