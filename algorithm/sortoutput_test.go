package main

import "testing"

func Test_Sortoutput(t *testing.T) {
	sortOutput(1, 7, 9)
	sortOutput(1, 9, 7)
	sortOutput(7, 1, 9)
	sortOutput(9, 7, 1)
	sortOutput(9, 1, 7)
	sortOutput(7, 9, 1)
}
