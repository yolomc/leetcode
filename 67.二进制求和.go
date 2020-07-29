import "strconv"

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	// ai, _ := new(big.Int).SetString(a, 2)
	// bi, _ := new(big.Int).SetString(b, 2)
	// ai.Add(ai, bi)
	// return ai.Text(2)

	res := []byte("0")
	la := len(a)
	lb := len(b)
	tmp := 0
	for i := 1; i <= la || i <= lb; i++ {
		if res[0] == '1' {
			tmp++
		}
		if i <= la && a[la-i] == '1' {
			tmp++
		}
		if i <= lb && b[lb-i] == '1' {
			tmp++
		}
		res[0] = strconv.Itoa(tmp % 2)[0]
		res = append([]byte(strconv.Itoa(tmp/2)), res...)
		tmp = 0
	}

	if res[0] == '0' {
		res = res[1:]
	}

	return string(res)
}

// @lc code=end
