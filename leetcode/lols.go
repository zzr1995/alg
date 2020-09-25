package leetcode

import (
	"fmt"
	"strings"
)

// 关键词:等距离跳跃
// the beginning and the end are always increasing cause we need the maximum distance
// 设置start和end，i慢走，出现重复时start跳到重复的后一位，end也等距离跳跃，i追上end的时候end再移动
func LengthOfLongestSubstring(s string) int {
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
