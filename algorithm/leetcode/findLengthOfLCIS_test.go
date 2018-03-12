package leetcode

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{nums: []int{}}, want: 0},
		{name: "", args: args{nums: []int{1, 2, 3, 4}}, want: 4},
		{name: "", args: args{nums: []int{1, 1, 1, 1}}, want: 1},
		{name: "", args: args{nums: []int{1, 2, 1, 1}}, want: 2},
		{name: "", args: args{nums: []int{1, 2, 4, 3}}, want: 3},
		{name: "", args: args{nums: []int{4, 3, 2, 1}}, want: 1},
		{name: "", args: args{nums: []int{1, 2, 3, 1, 2, 3, 4, 5}}, want: 5},
		{name: "", args: args{nums: []int{1, 2, 3, 4, 5, 1, 2, 3}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
