// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_EffectRefModify10000FFI

type effectRefModify10000FFI_reference struct{ value int }

func effectRefModify10000FFI_newRef(value int) *effectRefModify10000FFI_reference {
	return &effectRefModify10000FFI_reference{value}
}
func effectRefModify10000FFI_readRef(ref *effectRefModify10000FFI_reference) int { return ref.value }
func effectRefModify10000FFI_modifyRef(f func(int) int, ref *effectRefModify10000FFI_reference) {
	ref.value = f(ref.value)
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		// A reference belongs to one invocation, never to the prepared fixture.
		return func(count int) int {
			ref := effectRefModify10000FFI_newRef(0)
			for i := 1; i <= count; i++ {
				effectRefModify10000FFI_modifyRef(func(value int) int { return value + 1 }, ref)
			}
			return effectRefModify10000FFI_readRef(ref)
		}
	}
}
