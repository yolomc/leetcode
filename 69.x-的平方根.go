/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l := 0
	r := x
	m := x / 2
	res := m

ForLoop:
	for l <= r {
		switch temp := m * m; {
		case temp == x:
			res = m
			break ForLoop
		case temp < x:
			res = m
			l = m + 1
		case temp > x:
			r = m - 1
		}

		m = (l + r) / 2
	}

	return res
}

// @lc code=end
