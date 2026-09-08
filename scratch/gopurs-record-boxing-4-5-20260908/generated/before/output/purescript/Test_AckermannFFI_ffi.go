package purescript

import "gopurs/output/gopurs_runtime"



func ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if m > 0 && n == 0 {
		return ack(m-1, 1)
	}
	return ack(m-1, ack(m, n-1))
}

func Test_AckermannFFI_RunAckermannFFI(limit int) int {
	return (ack(3, 4))
}



// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_AckermannFFI_RunAckermannFFI = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_AckermannFFI_RunAckermannFFI(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})