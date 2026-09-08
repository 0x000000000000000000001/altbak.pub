package purescript

import "gopurs/output/gopurs_runtime"


type dictD_cc struct { e int; f int }
type dictB_cc struct { c int; d dictD_cc }
type dictR_cc struct { a int; b dictB_cc }

func Test_RecordsFFICheatcode_RunRecordsFFICheatcode(limit int) int {
	n := int(limit)
	r := &dictR_cc{a: 0, b: dictB_cc{c: 0, d: dictD_cc{e: 0, f: 0}}}
	for n > 0 {
		r.a += 1
		r.b.c += 2
		r.b.d.e += 3
		r.b.d.f += (n % 5)
		n--
	}
	return r.b.d.f
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_RecordsFFICheatcode_RunRecordsFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_RecordsFFICheatcode_RunRecordsFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})