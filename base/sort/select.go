package sort

import "fmt"

// 在后面选最小的排到前面
func selects(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i+1; j <= len(data)-1; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		fmt.Println(data)
		data[i], data[min] = data[min], data[i]
	}
	return data
}
