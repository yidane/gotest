package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_TwoSum0(t *testing.T) {
	arr := []int{2, 7, 11, 15, 12, 31, 22, 12, 311, 123, 441, 41223, 3, 43, 1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10, 100000, 999}
	target := 100999
	fmt.Println(twoSum0(arr, target))
}

func Test_TwoSum(t *testing.T) {
	check := func(a []int, b []int, t *testing.T) {
		if a == nil || len(b) != 2 {
			t.Error("Result is Nil or len is not equal 2.")
			return
		}
		if b == nil || len(b) != 2 {
			t.Error("Expected Result is Nil or len is not equal 2.")
			return
		}
		if (a[0] != b[0] || a[1] != b[1]) && (a[0] != b[1] || a[1] == b[0]) {
			t.Error("Expected Result:", b, "Result:", a)
			return
		}
	}

	check(twoSum([]int{3, 3}, 6), []int{0, 0}, t)
	check(twoSum([]int{3, 2, 4}, 6), []int{1, 2}, t)
}

func Test_twoSum0(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum0(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
