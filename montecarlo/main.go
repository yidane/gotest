package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
1、 1个 1^2
2、 5个 5^2
3、 20个 20^2
4、 100个 100^2
5、 1000个 1000^2

πr^2/πR^2=1/5
*/

func price() int {
	var max float64 = math.Pow(10, 10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := r.Float64() * max
	y := r.Float64() * max
	n := x*x + y*y
	switch {
	case n >= 0 && n < 1:
		return 1
	case n >= 1 && n < 25:
		return 2
	case n >= 25 && n < 400:
		return 3
	case n >= 400 && n < 10000:
		return 4
	case n >= 10000 && n < 1000000:
		return 5
	default:
		return 6
	}
}

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	total := 0
	c := 0

	for i := 0; i < 1<<20; i++ {
		x := r.Float64()*2 - 1
		y := r.Float64()*2 - 1

		if x*x+y*y <= 1 {
			c++
		}
		total++
	}

	fmt.Printf("%.4f\r", float64(c)*4/float64(total))

	fmt.Println(price())
}
