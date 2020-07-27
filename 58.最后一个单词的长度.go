import "strings"

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	sSlice := strings.Split(strings.TrimSpace(s), " ")
	return len([]rune(sSlice[len(sSlice)-1]))
}

// @lc code=end
