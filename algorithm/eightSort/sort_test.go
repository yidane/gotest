package eightSort

import "testing"

func checkSort(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func testSort(t *testing.T, sort func([]int) []int) {
	for i, arr := range preSortedArr {
		if !checkSort(sort(arr)) {
			t.Fail()
			t.Log(i)
			break
		}
	}
}

func Test_CheckSort(t *testing.T) {
	var sortedArr = [][]int{
		{1, 2, 3},
		{1, 3, 4},
		{1, 3, 2},
		{3, 2, 1},
		{2, 3, 4, 5, 1, 6},
		{2, 3},
		{-1, 9},
		{1},
	}
	if !checkSort(sortedArr[0]) {
		t.Fail()
	}
	if !checkSort(sortedArr[1]) {
		t.Fail()
	}
	if checkSort(sortedArr[2]) {
		t.Fail()
	}
	if checkSort(sortedArr[3]) {
		t.Fail()
	}
	if checkSort(sortedArr[4]) {
		t.Fail()
	}
	if !checkSort(sortedArr[5]) {
		t.Fail()
	}
	if !checkSort(sortedArr[6]) {
		t.Fail()
	}
	if !checkSort(sortedArr[7]) {
		t.Fail()
	}
}

func Test_BubbleSort(t *testing.T) {
	testSort(t, bubbleSort)
}

func Test_BubbleSort1(t *testing.T) {
	testSort(t, bubbleSort1)
}

func Test_BubbleSort2(t *testing.T) {
	testSort(t, bubbleSort2)
}

func Test_InsertSort(t *testing.T) {
	testSort(t, insertionSort)
}

func Test_SelectionSort(t *testing.T) {
	testSort(t, selectionSort)
}

func Test_SelectionSort1(t *testing.T) {
	testSort(t, selectionSort1)
}
