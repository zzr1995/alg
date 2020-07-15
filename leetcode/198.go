package leetcode

// f(1)=nums[1] //如果只有一家，肯定就是第一家肯定打劫
// f(2)=max(nums[1],nums[2]) // 如果2家 要么打劫第一家 要么打劫第二家 取两家的最大值
// f(3)=max(nums[3]+f(1),f(2)) // 如果打劫三家 ，方案1打劫了第二家 ，方案2 打劫了第一家和第三家 取 两个的最大值
// f(4)=max(nums[4]+f(2),f(3))

// 如果要打劫第n家：f(n-2)+nums(n),
// 如果不打劫第n家: f(n-1),
// 两者我们要去较大的那个，所以动态转移方程是：f(n)=max(nums[n]+f(n-2),f(n-1))

//func main() {
//	a := []int{1, 2, 3, 1, 4, 5, 7, 6}
//	max := rob(a)
//	fmt.Println(max)
//}

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 声明一个切片，代表前n家能抢到的最大钱
	p := make([]int, len(nums))
	p[0] = nums[0]
	p[1] = max(nums[0], nums[1])

	// 目前的最大值
	maxVal := p[1] // 2

	//从第三间房子开始
	for i := 2; i < len(nums); i++ {
		//a := []int{1, 2, 3, 1, 4, 5, 7, 6}
		p[i] = max(nums[i]+p[i-2], p[i-1])
		maxVal = max(maxVal, p[i])
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
