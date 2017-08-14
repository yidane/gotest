package mock

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	if n < 6 {
		return 2
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = i
	}

	return remove(m, true)
}

func remove(m map[int]int, f bool) int {
	if len(m) == 1 {
		return m[0]
	}
	if len(m) < 6 {
		return m[1]
	}
	result := make(map[int]int)
	if f {
		for i := len(m); i > 0; i-- {
			if i%2 == 0 {
				result[len(result)] = m[i]
			}
		}
	} else {
		for i := 0; i < len(m); i++ {
			if i%2 == 0 {
				result[len(result)] = m[i]
			}
		}
	}

	return remove(result, !f)
}
