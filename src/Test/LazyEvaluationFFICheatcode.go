package Test_LazyEvaluationFFICheatcode


type Lazy_cheatcode func() int

func force_cheatcode(l Lazy_cheatcode) int {
	return l()
}

func deferFunc_cheatcode(f func() int) Lazy_cheatcode {
	return f
}

func buildThunks_cheatcode(depth int, acc Lazy_cheatcode) Lazy_cheatcode {
	if depth == 0 {
		return acc
	}
	return buildThunks_cheatcode(depth-1, deferFunc_cheatcode(func() int {
		return force_cheatcode(acc) + 1
	}))
}

func runManyTimes_cheatcode(times int, acc int) int {
	if times == 0 {
		return acc
	}
	return runManyTimes_cheatcode(times-1, acc+force_cheatcode(buildThunks_cheatcode(1000, deferFunc_cheatcode(func() int {
		return 0
	}))))
}

func RunLazyEvaluationFFICheatcode(limit float64) float64 {
	n := int(limit)
	return float64(runManyTimes_cheatcode(n, 0))
}
