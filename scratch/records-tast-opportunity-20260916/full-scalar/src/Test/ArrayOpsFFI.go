package Test_ArrayOpsFFI

func RunArrayOpsFFI(limit int) int {
	n := int(limit)

	arr := make([]int, 0)
	step := 1
	if n < 1 {
		step = -1
	}
	for i := 1; ; i += step {
		arr = append(arr, i)
		if i == n {
			break
		}
	}

	evens := make([]int, 0)
	for _, x := range arr {
		if x%2 == 0 {
			evens = append(evens, x)
		}
	}

	sum := 0
	for _, x := range evens {
		sum += x
	}

	return (sum)
}
