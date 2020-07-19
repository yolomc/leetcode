/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if len(s) == 1 {
		return m[s[0]]
	}

	sum := 0
	curNum := 0
	preNum := 0
	for i := 0; i < len(s)-1; i++ {
		curNum = m[s[i]]
		preNum = m[s[i+1]]
		if curNum < preNum {
			sum -= curNum
		} else {
			sum += curNum
		}

		if i == len(s)-2 {
			sum += preNum
		}
	}

	return sum
}

// @lc code=end
