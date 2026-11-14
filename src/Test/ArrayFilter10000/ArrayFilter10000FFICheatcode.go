// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArrayFilter10000FFICheatcode

func arrayFilter10000FFICheatcode_weightedChecksum(values []int) int {
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
			result := make([]int, 0, count/2)
			for _, value := range values[:count] {
				if value%2 == 0 {
					result = append(result, value)
				}
			}
			return arrayFilter10000FFICheatcode_weightedChecksum(result)
		}
	}
}
