package __longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	var s = "sdsfsddcc"
	fmt.Printf("length of longest substring of '%s' is %v", s, lengthOfLongestSubstring(s))
}
