package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("adbcbebajh"))
}

// 关键词，等距离跳跃
// 设置start和end，i慢走，start和end等距离跳跃，出现重复时start跳到重复的后一位，end也等距离跳跃，i追上end的时候end再移动
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
			fmt.Println("改变窗口", start, end)
		}
	}
	return end - start
}
