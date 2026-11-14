// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_RecordSetLarge1000FFI

type recordSetLarge1000FFI_record struct{ a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p int }

func recordSetLarge1000FFI_sumRecord(value recordSetLarge1000FFI_record) int {
	return value.a + value.b + value.c + value.d + value.e + value.f + value.g + value.h + value.i + value.j + value.k + value.l + value.m + value.n + value.o + value.p
}

func recordSetLarge1000FFI_iterate(f func(recordSetLarge1000FFI_record) recordSetLarge1000FFI_record, count int, initial recordSetLarge1000FFI_record) recordSetLarge1000FFI_record {
	result := initial
	for i := 0; i < count; i++ {
		result = f(result)
	}
	return result
}

func recordSetLarge1000FFI_incrementLast(value recordSetLarge1000FFI_record) recordSetLarge1000FFI_record {
	result := value
	result.p++
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := recordSetLarge1000FFI_record{a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8, i: 9, j: 10, k: 11, l: 12, m: 13, n: 14, o: 15, p: 16}
		return func(count int) int {
			result := recordSetLarge1000FFI_iterate(recordSetLarge1000FFI_incrementLast, count, fixture)
			return recordSetLarge1000FFI_sumRecord(result) + recordSetLarge1000FFI_sumRecord(fixture)
		}
	}
}
