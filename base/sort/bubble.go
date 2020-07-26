package sort

import "fmt"



// i代表最后已经排好的数
func bubble(arr []int) []int{
	for i:=0;i<=len(arr)-1;i++ {
		for j:=0;j<len(arr)-1-i;j++{
			if arr[j] > arr [j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
			fmt.Printf("j:%v j+1:%v 结果：%v\n",j,j+1,arr)
		}
	}
	return arr
}
