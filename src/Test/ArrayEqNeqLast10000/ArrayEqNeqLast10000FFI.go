// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_ArrayEqNeqLast10000FFI

func arrayEqNeqLast10000FFI_eqArray(equal func(int) func(int) bool, left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i, value := range left {
		if !equal(value)(right[i]) {
			return false
		}
	}
	return true
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		left := make([]int, n)
		right := make([]int, n)
		for i := range left {
			left[i], right[i] = i+1, i+1
		}
		if n > 0 {
			right[n-1] = n + 1
		}
		return func(count int) int {
			equal := arrayEqNeqLast10000FFI_eqArray(func(left int) func(int) bool {
				return func(right int) bool { return left == right }
			}, left[:count], right[:count])
			if equal {
				return 1
			}
			return 0
		}
	}
}
