/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	if s == nil || len(s) == 1 {
		return
	}
	doReverse(s, 0, len(s)-1)
}

func doReverse(s []byte, lIdx int, rIdx int) {
	if lIdx >= rIdx {
		return
	}
	s[lIdx], s[rIdx] = s[rIdx], s[lIdx]
	doReverse(s, lIdx+1, rIdx-1)
}

// @lc code=end
