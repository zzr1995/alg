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
	fmt.Println(3 % 1000000007)
}

func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	// dfs里面要用所以提前声明
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
