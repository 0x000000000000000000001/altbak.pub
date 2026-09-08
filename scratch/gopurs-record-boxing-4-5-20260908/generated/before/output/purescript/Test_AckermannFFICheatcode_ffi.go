package purescript

import "gopurs/output/gopurs_runtime"



func ack_cheatcode(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if m > 0 && n == 0 {
		return ack_cheatcode(m-1, 1)
	}
	return ack_cheatcode(m-1, ack_cheatcode(m, n-1))
}

func Test_AckermannFFICheatcode_RunAckermannFFICheatcode(limit int) int {
	return (ack_cheatcode(3, 4))
}



// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_AckermannFFICheatcode_RunAckermannFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_AckermannFFICheatcode_RunAckermannFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})