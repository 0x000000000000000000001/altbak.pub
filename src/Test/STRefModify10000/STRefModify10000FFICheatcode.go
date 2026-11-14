// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
package Test_STRefModify10000FFICheatcode

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		// A reference belongs to one invocation, never to the prepared fixture.
		return func(count int) int {
			ref := new(int)
			for i := 1; i <= count; i++ {
				*ref = *ref + 1
			}
			return *ref
		}
	}
}
