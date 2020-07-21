/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s) == 1 {
		return false
	}

	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	ss := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(ss) > 0 && s[i] == ss[len(ss)-1] {
			ss = ss[:len(ss)-1]
			continue
		}

		sv, ok := m[s[i]]
		if !ok {
			return false
		}
		ss = append(ss, sv)
	}
	return len(ss) == 0
}

// @lc code=end
