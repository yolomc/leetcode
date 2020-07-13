/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	ret := 0
	if needle == "" {
		return ret
	}

	ret = -1
	//循环haystack，不需要到最后一位
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if ret != -1 {
			break
		}

		ret = i
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				ret = -1
				break
			}
		}
	}

	return ret
}

// @lc code=end
