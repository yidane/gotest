package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

const LargestNumber string = `123`

func largestNumber() {

	nums := make(map[int]int, 10)

	for i := 0; i < 10; i++ {
		nums[i] = 0
	}

	for _, v := range LargestNumber {
		vi, _ := strconv.Atoi(string(v))
		nums[vi] = nums[vi] + 1
	}

	fmt.Println(nums)
}

func largestNumber1() {
	totalNums := make(map[int]int, 10)

	for i := 0; i < 10; i++ {
		totalNums[i] = 0
	}

	cpunums := runtime.NumCPU()
	runtime.GOMAXPROCS(cpunums)

	wg := sync.WaitGroup{}
	wg.Add(cpunums + 1)

	numsChan := make(chan map[int]int)

	subFunc := func(step int, c chan map[int]int) {
		nums := make(map[int]int, 10)

		for i := 0; i < 10; i++ {
			nums[i] = 0
		}
		for step < len(LargestNumber) {

		}

		c <- nums

		wg.Done()
	}

	for i := 0; i < cpunums; i++ {
		go subFunc(i, numsChan)
	}

	go func() {
		for nums := range numsChan {
			for _, v := range totalNums {
				totalNums[v] += nums[v]
			}
		}

		wg.Done()
	}()

	wg.Wait()

	fmt.Println(totalNums)
}
