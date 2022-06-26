package leetcode

import (
	"testing"
)

func Test0003(t *testing.T) {
	// boundary cases
	res := LengthOfLongestSubstring("")
	if res != 0 {
		t.Fail()
	}
	res = LengthOfLongestSubstring("a")
	if res != 1 {
		t.Fail()
	}

	// leetcode sample data
	res = LengthOfLongestSubstring("abcabcbb")
	if res != 3 {
		t.Fail()
	}
	res = LengthOfLongestSubstring("bbbbbbb")
	if res != 1 {
		t.Fail()
	}

	res = LengthOfLongestSubstring("pwwkew")
	if res != 3 {
		t.Fail()
	}

	// no repeating character case
	res = LengthOfLongestSubstring("abcdefghijklmnopqrstuvwxyz")
	if res != len("abcdefghijklmnopqrstuvwxyz") {
		t.Fail()
	}

	// longest starts from index 0
	res = LengthOfLongestSubstring("abcdefghaabbccddeeffgg")
	if res != len("abcdefgh") {
		t.Fail()
	}

	// longest ends at last letter
	res = LengthOfLongestSubstring("aabbccddeeffggabcdefhijk")
	if res != len("gabcdefhijk") {
		t.Fail()
	}

	// longest in the middle
	res = LengthOfLongestSubstring("aaaabcdefghijkaaa")
	if res != len("abcdefghijk") {
		t.Fail()
	}
}
