import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := "1"
	for i := 2; i <= n; i++ {
		var temp strings.Builder
		count := 1
		current := s[0]

		for j := 1; j < len(s); j++ {
			if current == s[j] {
				count++
			} else {
				temp.WriteString(strconv.Itoa(count))
				temp.WriteByte(current)
				count = 1
				current = s[j]
			}
		}

		temp.WriteString(strconv.Itoa(count))
		temp.WriteByte(current)
		s = temp.String()
	}
	return s
}

// @lc code=end
