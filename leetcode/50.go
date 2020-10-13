package leetcode

// 求x的n次方
func f(x, n int) int {
	if n == 0 {
		return 1
	}
	t := f(x, n/2) // 2^7   t=x^3
	if n%2 == 1 {
		return t * t * x
	}
	return t * t
}
