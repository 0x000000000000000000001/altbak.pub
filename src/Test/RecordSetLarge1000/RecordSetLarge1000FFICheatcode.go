// Native imperative reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_RecordSetLarge1000FFICheatcode

type recordSetLarge1000FFICheatcode_record struct{ a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p int }

func recordSetLarge1000FFICheatcode_sumRecord(value recordSetLarge1000FFICheatcode_record) int {
	return value.a + value.b + value.c + value.d + value.e + value.f + value.g + value.h + value.i + value.j + value.k + value.l + value.m + value.n + value.o + value.p
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := recordSetLarge1000FFICheatcode_record{a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8, i: 9, j: 10, k: 11, l: 12, m: 13, n: 14, o: 15, p: 16}
		return func(count int) int {
			result := fixture
			for i := 0; i < count; i++ {
				result.p++
			}
			return recordSetLarge1000FFICheatcode_sumRecord(result) + recordSetLarge1000FFICheatcode_sumRecord(fixture)
		}
	}
}
