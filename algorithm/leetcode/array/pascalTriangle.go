package array

func pascalTriangle(l int) [][]int {
	result := [][]int{}
	for i := 0; i < l; i++ {
		arr := make([]int, i+1)
		arr[0] = 1
		for ii := 1; ii < i; ii++ {
			arr[ii] = result[i-1][ii-1] + result[i-1][ii]
		}
		arr[i] = 1
		result = append(result, arr)
	}

	return result
}
