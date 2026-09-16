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

func RunFibFFICheatcode(limit int) int {
	dummy := limit
	return (fib_cheatcode(dummy))
}
