package __longest_substring_without_repeating_characters

import "strings"

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	left := 0
	right := 0
	res := 0

	for i := 0; i < l; i++ {
		if strings.Contains(s[left:right], string(s[i])) {
			left += strings.Index(s[left:right], string(s[i])) + 1
		}
		right++

		if res < right-left {
			res = right - left
		}
	}

	return res
}
