package eightSort

/*
选择排序(Selection sort)也是一种简单直观的排序算法。
算法步骤：
	1）首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
	2）再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	3）重复第二步，直到所有元素均排序完毕。
*/
func selectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		t := i
		for j := i + 1; j < len(arr); j++ {
			if arr[t] > arr[j] {
				t = j
			}
		}
		if t != i {
			arr[i], arr[t] = arr[t], arr[i]
		}
	}

	return arr
}

//一次循环，找到最大和最小值，分别和两端交换
func selectionSort1(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-i; i++ {
		tMin := i
		tMax := l - 1 - i
		for j := i; j < l-i; j++ {
			if arr[tMin] > arr[j] {
				tMin = j
			}
			if arr[tMax] < arr[j] {
				tMax = j
			}
		}
		//如果此时最大最小恰好在对称位置，则这两次交换会回滚
		if tMin != i {
			arr[i], arr[tMin] = arr[tMin], arr[i]
		}
		if tMax != l-1-i {
			arr[l-1-i], arr[tMax] = arr[tMax], arr[l-1-i]
		}
	}

	return arr
}
