/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0] //暂时保存第一个字符串为前缀

	//循环剩余所有字符串
	for i := 1; i < len(strs); i++ {

		if strs[i] == "" {
			return ""
		}

		minlen := len(prefix)  //最小循环长度（防止循环下标溢出），暂时记录最小长度为前缀长度
		curStrLenLess := false //记录是否当前字符串比当前前缀更短，如果更短标记为true，并且更新最小循环长度为当前字符串长度
		if len(strs[i]) < minlen {
			curStrLenLess = true
			minlen = len(strs[i])
		}

		//依次对比当前字符串和当前前缀的每个字符，遇到不同的字符就更新当前前缀并跳出
		j := 0
		for ; j < minlen; j++ {
			if strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}

		//如果上面的循环没有跳出，且当前字符串长度更短，说明当前字符串与当前前缀完全匹配，更换当前字符串成为新的前缀
		if minlen == j && curStrLenLess {
			prefix = strs[i]
		}
	}
	return prefix
}

// @lc code=end
