// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_EffectRefWrite10000FFI

type effectRefWrite10000FFI_reference struct{ value int }

func effectRefWrite10000FFI_newRef(value int) *effectRefWrite10000FFI_reference {
	return &effectRefWrite10000FFI_reference{value}
}
func effectRefWrite10000FFI_readRef(ref *effectRefWrite10000FFI_reference) int { return ref.value }
func effectRefWrite10000FFI_writeRef(value int, ref *effectRefWrite10000FFI_reference) {
	ref.value = value
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		// A reference belongs to one invocation, never to the prepared fixture.
		return func(count int) int {
			ref := effectRefWrite10000FFI_newRef(0)
			for i := 1; i <= count; i++ {
				effectRefWrite10000FFI_writeRef(i, ref)
			}
			return effectRefWrite10000FFI_readRef(ref)
		}
	}
}
