package main

import (
	"fmt"
	"time"
)

/*
冒泡排序
冒泡排序算法的运作如下：（从后往前）
1、比较相邻的元素。如果第一个比第二个大，就交换他们两个。
2、对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
3、针对所有的元素重复以上的步骤，除了最后一个。
4、持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/

//普通版本，遍历每一个没有计算过的数，判断大小
func bubbleSort(arr []int) []int {
	beginTime := time.Now()
	if arr == nil {
		return nil
	}
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println(time.Now().Sub(beginTime).String())
	return arr
}

//普通版本的升级版本，当某一次遍历不再有数据交换，则排序成功
func bubbleSort1(arr []int) []int {
	beginTime := time.Now()
	if arr == nil {
		return nil
	}
	if len(arr) <= 1 {
		return arr
	}
	//判断，当某次循环没有数据交换，则排序结束
	hasExchange := false
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				hasExchange = true
			}
		}

		if !hasExchange {
			break
		}
		hasExchange = false
	}

	fmt.Println(time.Now().Sub(beginTime).String())
	return arr
}

//简单排序，每一次排序都获取当前数组的最小值，然后当前数组最左交换。排除已经放置的数，剩下继续此操作。
func selectSort(arr []int) []int {
	beginTime := time.Now()
	for i := 0; i < len(arr); i++ {
		t := arr[i]
		l := 0
		b := false
		for j := i; j < len(arr); j++ {
			if t > arr[j] {
				l = j
				t = arr[j]
				b = true
			}
		}
		if b {
			arr[i], arr[l] = arr[l], arr[i]
		}
	}

	fmt.Println(time.Now().Sub(beginTime).String())
	return arr
}

//简单排序升级版本，采用双向指针，每次获取当前数组的最大和最小值，分别和当前数组的两端交换，剩下的数组继续此操作
func selectSort1(arr []int) []int {
	beginTime := time.Now()
	if len(arr) < 2 {
		return arr
	}
	l := len(arr)
	for i := 0; i < l-i; i++ {
		//在一次循环中同时找出当前被循环区域的最大值和最小值，并依次继续下去
		min := arr[i]
		minIndex := i
		max := arr[l-i-1]
		maxIndex := l - i - 1
		for j := i; j < l-i; j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
		if maxIndex != l-i-1 {
			arr[l-i-1], arr[maxIndex] = arr[maxIndex], arr[l-i-1]
		}
	}

	fmt.Println(time.Now().Sub(beginTime).String())
	return arr
}

//快速排序算法
func quickSort(arr []int, l, r int) []int {
	quickSort := func(arr *[]int, begin, end int) int {
		i, j := begin, end
		x := (*arr)[begin]
		for i < j {
			for i < j && (*arr)[j] >= x {
				j--
			}
			if i < j {
				(*arr)[i] = (*arr)[j]
				i++
			}
			for i < j && (*arr)[i] <= x {
				i++
			}
			if i < j {
				(*arr)[j] = (*arr)[i]
				j--
			}
		}
		(*arr)[i] = x
		return i
	}
	begin, end := 0, len(arr)
	quickSort(&arr, begin, end)

	return arr
}
