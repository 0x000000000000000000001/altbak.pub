// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArrayAppend1000FFICheatcode

func arrayAppend1000FFICheatcode_weightedChecksum(values []int) int {
	total := 0
	for index, value := range values {
		total += (1 + index%2) * value
	}
	return total
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		left := make([]int, n)
		right := make([]int, n)
		for i := range left {
			left[i], right[i] = i+1, n+i+1
		}
		return func(count int) int {
			result := make([]int, count*2)
			copy(result, left[:count])
			copy(result[count:], right[:count])
			return arrayAppend1000FFICheatcode_weightedChecksum(result)
		}
	}
}
