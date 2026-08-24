package Test_ArrayOpsFFICheatcode


func RunArrayOpsFFICheatcode(limit float64) float64 {
	n := int(limit)
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	return float64(sum)
}

