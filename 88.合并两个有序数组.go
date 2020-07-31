/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1=append(nums1[:m],nums2...)
	// sort.Ints(nums1)

	if len(nums1) == 0 || len(nums2) == 0 || n == 0 {
		return
	}

	for i := m + n - 1; i >= 0; i-- {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
			continue
		}

		nums1[i] = nums2[n-1]
		n--
		if n == 0 {
			break
		}
	}
}

// @lc code=end
