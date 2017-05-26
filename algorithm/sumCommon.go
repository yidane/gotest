package main

import (
	"fmt"
)

//有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13...求出这个数列的前20项之和。
func sum1() {
	var up float64 = 2
	var down float64 = 1
	var total float64
	for i := 1; i <= 20; i++ {
		total += up / down
		up, down = down, up
		up += down
	}

	fmt.Println(total)
}

//有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
func sum2() {
	arr := [4]int{1, 2, 3, 4}
	total := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if i == j || i == k || j == k {
					continue
				}
				total++
				fmt.Println(arr[i], arr[j], arr[k])
			}
		}
	}
	fmt.Println("总共组成", total, "个互不相同且无重复数字的三位数")
}

//输入某年某月某日，判断这一天是这一年的第几天？
func sum3(year, month, day int) (result int) {
	if month == 1 {
		return day
	}
	monthDay := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		monthDay[1] = 29
	}
	for i := 1; i < month; i++ {
		result += monthDay[i-1]
	}

	return result + day
}
