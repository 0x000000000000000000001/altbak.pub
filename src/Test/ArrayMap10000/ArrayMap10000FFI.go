// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArrayMap10000FFI

func arrayMap10000FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func arrayMap10000FFI_weightedChecksum(values []int) int {
	type accumulator struct{ index, total int }
	result := arrayMap10000FFI_foldl(func(acc accumulator) func(int) accumulator {
		return func(value int) accumulator {
			return accumulator{acc.index + 1, acc.total + (1+acc.index%2)*value}
		}
	}, accumulator{}, values)
	return result.total
}

func arrayMap10000FFI_mapArray[A, B any](f func(A) B, values []A) []B {
	result := make([]B, len(values))
	for i, value := range values {
		result[i] = f(value)
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
			result := arrayMap10000FFI_mapArray(func(value int) int { return value + 1 }, values[:count])
			return arrayMap10000FFI_weightedChecksum(result)
		}
	}
}
