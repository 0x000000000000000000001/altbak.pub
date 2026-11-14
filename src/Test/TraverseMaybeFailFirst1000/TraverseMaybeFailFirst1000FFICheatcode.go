// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_TraverseMaybeFailFirst1000FFICheatcode

func traverseMaybeFailFirst1000FFICheatcode_weightedChecksum(values []int) int {
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
			rejected := 1
			result := make([]int, count)
			valid := true
			for i, value := range values[:count] {
				if value == rejected {
					valid = false
				} else {
					result[i] = value + 1
				}
			}
			if !valid {
				return -1
			}
			return traverseMaybeFailFirst1000FFICheatcode_weightedChecksum(result)
		}
	}
}
