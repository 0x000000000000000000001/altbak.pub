// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArrayFilter10000FFI

func arrayFilter10000FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func arrayFilter10000FFI_weightedChecksum(values []int) int {
	type accumulator struct{ index, total int }
	result := arrayFilter10000FFI_foldl(func(acc accumulator) func(int) accumulator {
		return func(value int) accumulator {
			return accumulator{acc.index + 1, acc.total + (1+acc.index%2)*value}
		}
	}, accumulator{}, values)
	return result.total
}

func arrayFilter10000FFI_filterArray[A any](predicate func(A) bool, values []A) []A {
	result := make([]A, 0)
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		values := make([]int, n)
		for i := range values {
			values[i] = i + 1
		}
		return func(count int) int {
			result := arrayFilter10000FFI_filterArray(func(value int) bool { return value%2 == 0 }, values[:count])
			return arrayFilter10000FFI_weightedChecksum(result)
		}
	}
}
