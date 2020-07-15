package leetcode

import (
	"testing"
	"time"
)

func TestGetDayGap(t *testing.T) {
	type args struct {
		now    time.Time
		before time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ name:"测试用例一",args: args{time.Now(),time.Now()},want:0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDayGap(tt.args.now, tt.args.before); got != tt.want {
				t.Errorf("GetDayGap() = %v, want %v", got, tt.want)
			}
		})
	}
}