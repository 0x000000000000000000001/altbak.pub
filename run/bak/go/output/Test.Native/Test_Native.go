package Test_Native

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_loopNative gopurs_runtime.Value
var once_loopNative sync.Once
func Get_loopNative() gopurs_runtime.Value {
	once_loopNative.Do(func() {
		cache_loopNative = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_loopNative(v_0_box.IntVal, v1_1_box.IntVal))
})
	})
	return cache_loopNative
}

func Call_loopNative(v_0_loop int64, v1_1_loop int64) int64 {
loopNative:
for {
if false { continue loopNative }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Int(v1_1)
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = (v1_1) + (1)
continue loopNative
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0.IntVal
}
}


