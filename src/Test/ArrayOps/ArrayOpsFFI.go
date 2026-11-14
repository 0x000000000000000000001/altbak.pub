package Test_ArrayOpsFFI

func arrayRange(start, end int) []int {
	step := 1
	if end < start {
		step = -1
	}
	arr := make([]int, 0)
	for i := start; ; i += step {
		arr = append(arr, i)
		if i == end {
			return arr
		}
	}
}

func arrayFilter[A any](predicate func(A) bool, arr []A) []A {
	result := make([]A, 0)
	for _, x := range arr {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

func arrayFoldl[A, B any](f func(B) func(A) B, initial B, arr []A) B {
	acc := initial
	for _, x := range arr {
		acc = f(acc)(x)
	}
	return acc
}

func RunArrayOpsFFI(limit int) int {
	values := arrayRange(1, limit)
	evens := arrayFilter(func(x int) bool { return x%2 == 0 }, values)
	return arrayFoldl(func(acc int) func(int) int {
		return func(x int) int { return acc + x }
	}, 0, evens)
}
