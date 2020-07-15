import "strconv"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	res := true
	for i := 0; ; {
		if i >= len(xStr) {
			break
		}

		if xStr[i] != xStr[len(xStr)-1-i] {
			res = false
			break
		}
		i++
	}
	return res
}

// @lc code=end
