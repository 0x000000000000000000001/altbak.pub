package Test_FibFFICheatcode


func fib_cheatcode(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib_cheatcode(n-1) + fib_cheatcode(n-2)
}

func RunFibFFICheatcode(limit float64) float64 {
	dummy := int(limit)
	return float64(fib_cheatcode(dummy))
}
