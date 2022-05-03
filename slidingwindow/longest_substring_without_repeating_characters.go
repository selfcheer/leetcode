package slidingwindow

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := make(map[uint8]bool, 0)
	var left, max int
	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; ok {
			for m[s[i]] {
				delete(m, s[left])
				left++
			}
		}
		max = Max(max, i-left+1)

		m[s[i]] = true
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func TestLengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring(""))
}
