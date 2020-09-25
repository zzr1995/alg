package sort

import (
	"reflect"
	"testing"
)

func Test_selects(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name:"测试选择排序",args: args{[]int{3,4,6,9,2,7,8,1,5}},want: []int{1,2,3,4,5,6,7,8,9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selects(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selects() = %v, want %v", got, tt.want)
			}
		})
	}
}