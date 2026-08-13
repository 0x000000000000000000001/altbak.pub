package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_TestLet_addOne gopurs_runtime.Value
var once_TestLet_addOne sync.Once
func Get_TestLet_addOne() gopurs_runtime.Value {
	once_TestLet_addOne.Do(func() {
		cache_TestLet_addOne = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_TestLet_addOne(x_0_box.IntVal))
})
	})
	return cache_TestLet_addOne
}

func Call_TestLet_addOne(x_0_loop int64) int64 {
var x_0 int64 = x_0_loop
_ = x_0
// TAST (Let): y_1_0 -> int64
y_1_0 := gopurs_runtime.Apply(Get_TestLet_opaque(), gopurs_runtime.Int(x_0)).IntVal
_ = y_1_0
// TAST (Let): z_2_1 -> int64
z_2_1 := (y_1_0) + (y_1_0)
_ = z_2_1
return (z_2_1) + (z_2_1)
}

func Get_TestLet_opaque() gopurs_runtime.Value {
	return _Gopurs_TestLet_Opaque
}
