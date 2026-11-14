// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArraySortDescending10000FFICheatcode

import (
	"sort"
)

func arraySortDescending10000FFICheatcode_weightedChecksum(values []int) int {
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
			values[i] = n - i
		}
		return func(count int) int {
			result := append([]int(nil), values[:count]...)
			sort.Ints(result)
			return arraySortDescending10000FFICheatcode_weightedChecksum(result)
		}
	}
}
