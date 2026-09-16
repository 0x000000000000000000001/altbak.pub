package Test_ArrayOpsFFICheatcode

func RunArrayOpsFFICheatcode(limit int) int {
	n := int(limit)
	sum := 0
	step := 1
	if n < 1 {
		step = -1
	}
	for i := 1; ; i += step {
		if i%2 == 0 {
			sum += i
		}
		if i == n {
			break
		}
	}
	return (sum)
}
