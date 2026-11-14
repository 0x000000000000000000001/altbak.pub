// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
package Test_ArrayFoldlSum10000FFICheatcode

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		values := make([]int, n)
		for i := range values {
			values[i] = i + 1
		}
		return func(count int) int {
			total := 0
			for _, value := range values[:count] {
				total += value
			}
			return total
		}
	}
}
