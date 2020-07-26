package sort

import (
	"reflect"
	"testing"
)

func Test_bubble(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name:"测试冒泡排序",args: args{[]int{3,4,6,2,7,8,1,5,9}},want: []int{1,2,3,4,5,6,7,8,9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubble(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}