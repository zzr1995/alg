package leetcode

import "testing"

func TestRob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "测试一",args:args{[]int{1, 2, 3, 1, 4, 5, 7, 6}},want:15 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rob(tt.args.nums); got != tt.want {
				t.Errorf("Rob() = %v, want %v", got, tt.want)
			}
		})
	}
}