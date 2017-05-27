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
	if month < 1 || month > 12 {
		panic("month must more than 0 and less than 13!")
	}

	if day < 1 || day > 31 {
		panic("day must more than 0 and less than 32!")
	}

	if month == 1 {
		if day < 1 || day > 31 {
			panic("January has 31 days,you cannot input a number more than 31 or less than 1!")
		}
		return day
	}
	monthDay := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		monthDay[1] = 29
	}
	if month == 2 {
		if day > monthDay[1] {
			panic(fmt.Sprintf("Febrary of year %d has %d days,you cannot input a number more %d as day number!", year, monthDay[1], monthDay[1]))
		}
	}
	for i := 1; i < month; i++ {
		result += monthDay[i-1]
	}

	return result + day
}

//有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
//1 2，2 2，3 4，4 6，5 10
func sum4(num int) {
	i := 1
	arr := make(map[int]int)
	for {
		if i > num {
			break
		}

		if i < 3 {
			arr[i-1] = 2
		} else {
			arr[i-1] = arr[i-2] + arr[i-3]
		}
		i++
	}

	for i := 0; i < num; i++ {
		fmt.Printf("第%d个月%d\n", i+1, arr[i])
	}
}
