package sort

import (
	"fmt"
	"math"
)

// j开始为i，递减gap插入，arr[j-gap]>arr[i] ,则arr[j] = arr[j-gap]
func ShellSort(arr []int) []int{
	fmt.Println("原始数组",arr)
	length := len(arr)
	gap := int(math.Floor(float64(length / 2)))
	for gap >= 1 {
		fmt.Println("间隔gap:",gap)
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i
			for j > gap-1 && arr[j-gap] > temp {

				fmt.Printf("交换%v和%v:",arr[j],arr[j-gap])
				arr[j] = arr[j-gap]
				fmt.Println("---数组",arr,"j-gap")
				j -= gap
			}
			arr[j] = temp
			fmt.Println("---i 完成一轮 arr[j] = temp",arr)
		}
		gap = int(math.Floor(float64(gap / 2)))
	}
	return arr
}