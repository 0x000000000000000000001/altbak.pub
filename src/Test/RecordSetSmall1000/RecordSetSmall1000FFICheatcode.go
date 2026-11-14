// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_RecordSetSmall1000FFICheatcode

type recordSetSmall1000FFICheatcode_record struct{ a, b int }

func recordSetSmall1000FFICheatcode_sumRecord(value recordSetSmall1000FFICheatcode_record) int {
	return value.a + value.b
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := recordSetSmall1000FFICheatcode_record{a: 1, b: 2}
		return func(count int) int {
			result := fixture
			for i := 0; i < count; i++ {
				result.b++
			}
			return recordSetSmall1000FFICheatcode_sumRecord(result) + recordSetSmall1000FFICheatcode_sumRecord(fixture)
		}
	}
}
