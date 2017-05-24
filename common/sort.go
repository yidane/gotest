package main

import (
	"fmt"
	"time"
)

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

func bubbleSort1(arr []int) []int {
	beginTime := time.Now()
	if arr == nil {
		return nil
	}
	if len(arr) <= 1 {
		return arr
	}
	hasExchange := false
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-i; j++ {
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

func sort() {
	originalArr := []int{2, 123, 123, 45, 342, 5647, 822, 427, 417, 54, 92, 320, 502}
	arr := make([]int, len(originalArr))

	copy(arr, originalArr)
	fmt.Println(bubbleSort(arr))

	copy(arr, originalArr)
	fmt.Println(bubbleSort1(arr))
}
