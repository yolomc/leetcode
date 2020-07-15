import "strconv"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	for i := 0; i < len(xStr); i++ {
		if xStr[i] != xStr[len(xStr)-1-i] {
			return false
		}
	}
	return true
}

// @lc code=end
