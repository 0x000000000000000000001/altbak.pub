// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArrayAppend1000FFI

func arrayAppend1000FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func arrayAppend1000FFI_weightedChecksum(values []int) int {
	type accumulator struct{ index, total int }
	result := arrayAppend1000FFI_foldl(func(acc accumulator) func(int) accumulator {
		return func(value int) accumulator {
			return accumulator{acc.index + 1, acc.total + (1+acc.index%2)*value}
		}
	}, accumulator{}, values)
	return result.total
}

func arrayAppend1000FFI_appendArray[A any](left, right []A) []A {
	result := make([]A, len(left)+len(right))
	copy(result, left)
	copy(result[len(left):], right)
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		left := make([]int, n)
		right := make([]int, n)
		for i := range left {
			left[i], right[i] = i+1, n+i+1
		}
		return func(count int) int {
			result := arrayAppend1000FFI_appendArray(left[:count], right[:count])
			return arrayAppend1000FFI_weightedChecksum(result)
		}
	}
}
