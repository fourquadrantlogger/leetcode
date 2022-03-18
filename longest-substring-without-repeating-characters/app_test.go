package longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func TestCheckrepeat(t *testing.T) {
	fmt.Println(Checkrepeat("abcde"))
	fmt.Println(Checkrepeat("abcdea"))
}
func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(LengthOfLongestSubstring("abcabcbb"))
	fmt.Println(LengthOfLongestSubstring("bbbbb"))
	fmt.Println(LengthOfLongestSubstring("pwwkew"))

}

func TestLengthOfLongestSubstring2(t *testing.T) {

	fmt.Println(LengthOfLongestSubstring("au"))
}
