package longest_substring_without_repeating_characters

import "fmt"

func LengthOfLongestSubstring(s string) (r, la, lb int) {

	if len(s) > 1 {
		SlidingWindowA, SlidingWindowB := 0, 1
		for SlidingWindowA < SlidingWindowB && SlidingWindowB <= len(s) {
			ss := s[SlidingWindowA:SlidingWindowB]
			fmt.Println(ss)
			if Checkrepeat(ss) {
				SlidingWindowA++
			} else {
				if r < SlidingWindowB-SlidingWindowA {
					r = SlidingWindowB - SlidingWindowA
					la, lb = SlidingWindowA, SlidingWindowB
				}
				SlidingWindowB++
			}
		}
		return
	} else {
		return len(s), 0, len(s)
	}

}

func Checkrepeat(s string) bool {
	m := make(map[int32]byte)
	for _, c := range s {
		if _, exist := m[c]; exist {
			return true
		} else {
			m[c] = '-'
		}
	}
	return false
}
