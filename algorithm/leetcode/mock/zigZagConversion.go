package mock

func convert(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}

	arrs := make([]string, numRows)

	cr := 0
	cc := 0
	flag := true
	for _, r := range s {
		arrs[cr] = arrs[cr] + string(r)
		if cr == 0 {
			flag = true
		}
		if cr == numRows-1 {
			flag = false
		}
		if cr < numRows-1 {
			if flag {
				cr++
			} else {
				cr--
			}
		} else {
			cc++
			cr--
		}
	}

	result := ""
	for i := 0; i < len(arrs); i++ {
		result += arrs[i]
	}

	return result
}
