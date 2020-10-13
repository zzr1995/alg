package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	fmt.Println(printNumbers(3))
}

func printNumbers(n int) []int {
	max := int(math.Pow10(3)) - 1
	ans := make([]int, max)
	for i := 0; i < max; i++ {
		ans[i] = i + 1
	}
	return ans
}
