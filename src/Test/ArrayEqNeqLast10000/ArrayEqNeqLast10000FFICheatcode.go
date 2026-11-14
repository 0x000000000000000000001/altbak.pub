// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
package Test_ArrayEqNeqLast10000FFICheatcode

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
			for i, value := range left[:count] {
				if value != right[i] {
					return 0
				}
			}
			return 1
		}
	}
}
