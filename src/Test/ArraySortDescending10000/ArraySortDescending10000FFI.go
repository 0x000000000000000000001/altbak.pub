// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArraySortDescending10000FFI

import (
	"sort"
)

func arraySortDescending10000FFI_foldl[A, B any](f func(B) func(A) B, initial B, values []A) B {
	result := initial
	for _, value := range values {
		result = f(result)(value)
	}
	return result
}

func arraySortDescending10000FFI_weightedChecksum(values []int) int {
	type accumulator struct{ index, total int }
	result := arraySortDescending10000FFI_foldl(func(acc accumulator) func(int) accumulator {
		return func(value int) accumulator {
			return accumulator{acc.index + 1, acc.total + (1+acc.index%2)*value}
		}
	}, accumulator{}, values)
	return result.total
}

func arraySortDescending10000FFI_sortArray(less func(int) func(int) bool, values []int) []int {
	result := append([]int(nil), values...)
	sort.SliceStable(result, func(i, j int) bool { return less(result[i])(result[j]) })
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		values := make([]int, n)
		for i := range values {
			values[i] = n - i
		}
		return func(count int) int {
			result := arraySortDescending10000FFI_sortArray(func(left int) func(int) bool {
				return func(right int) bool { return left < right }
			}, values[:count])
			return arraySortDescending10000FFI_weightedChecksum(result)
		}
	}
}
