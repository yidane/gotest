package main

func output0() {
	A := make(chan int)
	B := make(chan int)
	C := make(chan int)
	D := make(chan int)

	i := 1
	for {
		j := i % 4
		switch j {
		case 0:
			write1(A)
			write2(B)
			write3(C)
			write4(D)
		case 1:
			write1(B)
			write2(C)
			write3(D)
			write4(A)
		case 2:
			write1(C)
			write2(D)
			write3(A)
			write4(B)
		case 3:
			write1(D)
			write2(A)
			write3(B)
			write4(C)
		}
	}

}

func write1(ch chan int) {
	ch <- 1
}

func write2(ch chan int) {
	ch <- 2
}

func write3(ch chan int) {
	ch <- 3
}

func write4(ch chan int) {
	ch <- 4
}
