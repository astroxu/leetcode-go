package __longest_palindromic_substring

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	var str = "weqewsdf"
	fmt.Printf("longest palindromic substring of \"%s\" is \"%s\"\n", str, longestPalindrome(str))
}
