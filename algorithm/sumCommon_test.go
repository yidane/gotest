package main

import "testing"

func Test_sum1(t *testing.T) {
	sum1()
}

func Test_sum2(t *testing.T) {
	sum2()
}

func Test_sum3(t *testing.T) {
	if sum3(2017, 1, 1) != 1 {
		t.Fail()
	}

	if sum3(2017, 2, 20) != 51 {
		t.Fail()
	}

	if sum3(2017, 3, 9) != 68 {
		t.Fail()
	}

	if sum3(2016, 3, 9) != 69 {
		t.Fail()
	}
}

func Test_sum4(t *testing.T) {
	sum4(20)
}
