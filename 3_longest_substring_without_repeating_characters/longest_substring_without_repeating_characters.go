package __longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	ret := 0
	left := 0
	// 保存遍历过的字符的最新下标
	store := make(map[uint8]int)

	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(s); i++ {
		if ci, ok := store[s[i]]; ok { // 如果i字符之前出现过
			left = max(left, ci)
		}
		ret = max(ret, (i - left + 1))
		// 存储下标时，以1为起始位置，为了保证上面的 left = max(left, ci) 的正确性
		store[s[i]] = i + 1
	}

	return ret
}
