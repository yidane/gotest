package main

import (
	"fmt"
)

//FindDaffodilNumber 查找水仙数
func FindDaffodilNumber(min, max int) {
	fmt.Println("start")
	numberMap := make(map[int]int)
	for i := 1; i <= 9; i++ {
		numberMap[i] = i * i * i
	}

	for num := min; num < max; num++ {
		newNum := 0
		currentNum := num

		for currentNum >= 10 {
			bnum := currentNum % 10
			newNum += numberMap[bnum]
			if newNum > num {
				break
			}
			currentNum = (currentNum - bnum) / 10
		}

		if num > 0 {
			newNum += numberMap[currentNum]
		}

		if newNum == num {
			fmt.Printf("%d\n", num)
		}
	}

	fmt.Println("complete")
}
