package Test_LazyEvaluationFFI


type Lazy func() int

func force(l Lazy) int {
	return l()
}

func deferFunc(f func() int) Lazy {
	return f
}

func buildThunks(depth int, acc Lazy) Lazy {
	if depth == 0 {
		return acc
	}
	return buildThunks(depth-1, deferFunc(func() int {
		return force(acc) + 1
	}))
}

func runManyTimes(times int, acc int) int {
	if times == 0 {
		return acc
	}
	return runManyTimes(times-1, acc+force(buildThunks(1000, deferFunc(func() int {
		return 0
	}))))
}

func RunLazyEvaluationFFI(limit float64) float64 {
	n := int(limit)
	return float64(runManyTimes(n, 0))
}
