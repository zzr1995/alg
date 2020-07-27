package sort

// 在后面选最小的排到前面
func selects(data []int) []int {
	for i := 0; i < len(data); i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[min] > data[j] {
				data[min], data[j] = data[j], data[min]
			}
		}
	}
	return data
}
