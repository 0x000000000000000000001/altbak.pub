// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArrayMap10000FFICheatcode

func arrayMap10000FFICheatcode_weightedChecksum(values []int) int {
	total := 0
	for index, value := range values {
		total += (1 + index%2) * value
	}
	return total
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		values := make([]int, n)
		for i := range values {
			values[i] = i + 1
		}
		return func(count int) int {
			result := make([]int, count)
			for i, value := range values[:count] {
				result[i] = value + 1
			}
			return arrayMap10000FFICheatcode_weightedChecksum(result)
		}
	}
}
