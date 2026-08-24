package Test_FibFFI


func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func RunFibFFI(limit float64) float64 {
	return float64(fib(int(limit)))
}
