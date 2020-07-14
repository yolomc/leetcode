/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	lIdx := 0
	rIdx := len(nums) - 1
	mIdx := rIdx / 2
	mNum := nums[mIdx]

	for {
		if mNum == target {
			mNum = mIdx
			break
		} else if mNum > target {
			rIdx = mIdx - 1
		} else {
			lIdx = mIdx + 1
		}

		if lIdx > rIdx {
			mNum = -1
			break
		}

		mIdx = (rIdx + lIdx + 1) / 2
		mNum = nums[mIdx]
	}

	return mNum
}

// @lc code=end
