package main

import "fmt"

func sushu(num int) {
	arr := []int{}
	arr1 := []int{}
	for i := 2; i < num; i++ {
		arr = append(arr, i)
	}

	m := arr[0]
	for m < num/2 {
		fmt.Println(m)
		for _, v := range arr {
			if v%m != 0 {
				arr1 = append(arr1, v)
			}
		}

		arr = arr1
		arr1 = []int{}
		if len(arr) == 0 {
			break
		}
		m = arr[0]
	}

	fmt.Println(arr)
}

/*
http://tonybai.com/2017/04/20/go-coding-in-go-way/
【问题: 素数筛】

  问题描述：素数是一个自然数，它具有两个截然不同的自然数除数：1和它本身。 要找到小于或等于给定整数n的素数。针对这个问题，我们可以采用埃拉托斯特尼素数筛算法。
  算法描述：先用最小的素数2去筛，把2的倍数剔除掉；下一个未筛除的数就是素数(这里是3)。再用这个素数3去筛，筛除掉3的倍数... 这样不断重复下去，直到筛完为止。
*/

func generate(num int, ch chan<- int) {
	for i := 2; i < num; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src { // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
	}
}

func sieve(num int) {
	ch := make(chan int) // Create a new channel.
	go generate(num, ch) // Start generate() as a subprocess.
	for i := 2; i < num; i++ {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
