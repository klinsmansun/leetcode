package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	// situation 1: start comparing with repeated character, ex: bb
	// situation 2: start comparing with left x 2 character, ex: ada
	fmt.Println(LongestPalindrome("abbd"))
	fmt.Println(LongestPalindrome(""))         // empty string
	fmt.Println(LongestPalindrome("abcba"))    // result is whole string
	fmt.Println(LongestPalindrome("xadddday")) // situation 1 longer
	fmt.Println(LongestPalindrome("ddddd"))    // situation 2 longer
	fmt.Println(LongestPalindrome("dddddd"))   // situation 1 longer
}
