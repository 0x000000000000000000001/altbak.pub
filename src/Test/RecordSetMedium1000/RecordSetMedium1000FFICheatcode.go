// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_RecordSetMedium1000FFICheatcode

type recordSetMedium1000FFICheatcode_record struct{ a, b, c, d, e, f, g, h int }

func recordSetMedium1000FFICheatcode_sumRecord(value recordSetMedium1000FFICheatcode_record) int {
	return value.a + value.b + value.c + value.d + value.e + value.f + value.g + value.h
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := recordSetMedium1000FFICheatcode_record{a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8}
		return func(count int) int {
			result := fixture
			for i := 0; i < count; i++ {
				result.h++
			}
			return recordSetMedium1000FFICheatcode_sumRecord(result) + recordSetMedium1000FFICheatcode_sumRecord(fixture)
		}
	}
}
