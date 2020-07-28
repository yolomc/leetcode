/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	l := len(digits)

	if digits[l-1] != 9 {
		digits[l-1]++
		return digits
	}

	i := l - 2
	for ; i >= 0; i-- {
		digits[i+1] = 0
		if digits[i] != 9 {
			digits[i]++
			break
		}
	}

	if i < 0 {
		digits[0] = 0
		return append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end
