package purescript

import "gopurs/output/gopurs_runtime"



func Test_PrimesFFICheatcode_RunPrimesFFICheatcode(limit int) int {
	n := limit
	if n < 2 {
		return 0
	}
	sieve := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		sieve[i] = true
	}
	
	for p := 2; p*p <= n; p++ {
		if sieve[p] {
			for i := p * p; i <= n; i += p {
				sieve[i] = false
			}
		}
	}
	
	sum := 0
	for p := 2; p <= n; p++ {
		if sieve[p] {
			sum += p
		}
	}
	return sum
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Test_PrimesFFICheatcode_RunPrimesFFICheatcode = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := Test_PrimesFFICheatcode_RunPrimesFFICheatcode(go_arg0)
	return gopurs_runtime.Int(int64(go_res))
})