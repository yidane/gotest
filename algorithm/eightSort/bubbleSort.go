package eightSort

/*
 冒泡排序（Bubble Sort）
 	也是一种简单直观的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。
	走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
	这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
算法步骤：
	1）比较相邻的元素。如果第一个比第二个大，就交换他们两个。
	2）对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
	3）针对所有的元素重复以上的步骤，除了最后一个。
	4）持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func bubbleSortError(arr []int) []int {
	hasChange := false
	for i := 0; i < len(arr); i++ {
		hasChange = false
		/*
			//这样从左到右，只能保证最有最大，不能保证最左最小；然后i+1之后，最左数就无法保证最小。
			//目标是最小一定要摆放在最左。
			//因此正确做法是将数向目标方向驱动。
			// for j := i; j < len(arr)-1; j++ {
			// 	if arr[j] > arr[j+1] {
			// 		arr[j+1], arr[j] = arr[j], arr[j+1]
			// 		hasChange = true
			// 	}
			// }
		*/
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				hasChange = true
			}
		}
		if !hasChange {
			break
		}
	}
	return arr
}

func bubbleSort1(arr []int) []int {
	hasChange := true
	for i := 0; i < len(arr) && hasChange; i++ {
		hasChange = false
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				hasChange = true
			}
		}
	}

	return arr
}

func bubbleSort2(arr []int) []int {
	hasChange := true
	for i := len(arr) - 1; i > 0 && hasChange; i-- {
		hasChange = false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				hasChange = true
			}
		}
	}
	return arr
}
