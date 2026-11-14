// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_STRefWrite10000FFI

type sTRefWrite10000FFI_reference struct{ value int }

func sTRefWrite10000FFI_newRef(value int) *sTRefWrite10000FFI_reference {
	return &sTRefWrite10000FFI_reference{value}
}
func sTRefWrite10000FFI_readRef(ref *sTRefWrite10000FFI_reference) int         { return ref.value }
func sTRefWrite10000FFI_writeRef(value int, ref *sTRefWrite10000FFI_reference) { ref.value = value }

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		// A reference belongs to one invocation, never to the prepared fixture.
		return func(count int) int {
			ref := sTRefWrite10000FFI_newRef(0)
			for i := 1; i <= count; i++ {
				sTRefWrite10000FFI_writeRef(i, ref)
			}
			return sTRefWrite10000FFI_readRef(ref)
		}
	}
}
