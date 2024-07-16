package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	count := 0
	set := make([]int, 256)
	for i, j := 0, 0; i < len(s); i++ {
		set[s[i]]++
		for set[s[i]] > 1 {
			set[s[i]]--
			j++
		}
		if i-j+1 > count {
			count = i - j + 1
		}
	}

	return count
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
