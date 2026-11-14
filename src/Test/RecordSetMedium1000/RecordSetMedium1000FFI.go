// Native functional reference. Fixtures are prepared by the factory Effect.
// Every returned action computes and consumes a fresh result; no result is cached.
// The action count matches the fixture size supplied by the PureScript wrapper.
// Private names are module-prefixed because the Go FFI linker merges their scope.
package Test_RecordSetMedium1000FFI

type recordSetMedium1000FFI_record struct{ a, b, c, d, e, f, g, h int }

func recordSetMedium1000FFI_sumRecord(value recordSetMedium1000FFI_record) int {
	return value.a + value.b + value.c + value.d + value.e + value.f + value.g + value.h
}

func recordSetMedium1000FFI_iterate(f func(recordSetMedium1000FFI_record) recordSetMedium1000FFI_record, count int, initial recordSetMedium1000FFI_record) recordSetMedium1000FFI_record {
	result := initial
	for i := 0; i < count; i++ {
		result = f(result)
	}
	return result
}

func recordSetMedium1000FFI_incrementLast(value recordSetMedium1000FFI_record) recordSetMedium1000FFI_record {
	result := value
	result.h++
	return result
}

func MakeAction(n int) func() func(int) int {
	return func() func(int) int {
		fixture := recordSetMedium1000FFI_record{a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8}
		return func(count int) int {
			result := recordSetMedium1000FFI_iterate(recordSetMedium1000FFI_incrementLast, count, fixture)
			return recordSetMedium1000FFI_sumRecord(result) + recordSetMedium1000FFI_sumRecord(fixture)
		}
	}
}
