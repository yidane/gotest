package algorithm

import (
	"fmt"
)

func findDaffodilNumber(min, max int) {
	fmt.Println("start")
	numberMap := make(map[int]int)
	for i := 1; i <= 9; i++ {
		numberMap[i] = i * i * i
	}

	for num := min; num < max; num++ {
		newNum := 0
		if newNum > num {
			break
		}

		str := string(num)
		for m := 0; m < len(str); m++ {
			fmt.Println(str[m])
			j := int(m)
			newNum += numberMap[j]
		}

		if newNum == num {
			fmt.Println(num)
		}
	}

	fmt.Println("complete")
}
