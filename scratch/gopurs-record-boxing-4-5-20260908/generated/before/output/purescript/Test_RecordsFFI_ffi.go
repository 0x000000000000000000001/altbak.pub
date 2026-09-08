package purescript

import "gopurs/output/gopurs_runtime"


type dictD struct { e int; f int }
type dictB struct { c int; d dictD }
type dictR struct { a int; b dictB }

func updateRec(n int, r dictR) dictR {
	if n == 0 { return r }
	newR := dictR{
		a: r.a + 1,
		b: dictB{
			c: r.b.c + 2,
			d: dictD{
				e: r.b.d.e + 3,
				f: r.b.d.f + (n % 5),
			},
		},
	}
	return updateRec(n - 1, newR)
}

func Test_RecordsFFI_RunRecordsFFI(limit int) int {
	dummy := limit
	initial := dictR{a: 0, b: dictB{c: 0, d: dictD{e: 0, f: 0}}}
	res := updateRec(dummy, initial)
	return (res.b.d.f)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_RecordsFFI_RunRecordsFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_RecordsFFI_RunRecordsFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})