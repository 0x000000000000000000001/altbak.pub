package Test_ArrayOpsFFI


func RunArrayOpsFFI(limit float64) float64 {
	n := int(limit)
	
	arr := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
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
	
	return float64(sum)
}
