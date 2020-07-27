package sort

// 3,4,6,2,7,8,1,5,9
// j是要插的数,反正是从第二个开始
func insert(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j > 0 && data[j-1] > data[j]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
	return data
}
