package ex15

func max(values ...int) (int, bool) {
	if len(values) == 0 {
		return 0, false
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max, true
}

func min(values ...int) (int, bool) {
	if len(values) == 0 {
		return 0, false
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min, true
}

func max2(v int, values ...int) int {
	for i := 0; i < len(values); i++ {
		if values[i] > v {
			v = values[i]
		}
	}
	return v
}

func min2(v int, values ...int) int {
	for i := 0; i < len(values); i++ {
		if values[i] < v {
			v = values[i]
		}
	}
	return v
}
