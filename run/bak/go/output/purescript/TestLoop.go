package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_TestLoop_loop gopurs_runtime.Value
var once_TestLoop_loop sync.Once
func Get_TestLoop_loop() gopurs_runtime.Value {
	once_TestLoop_loop.Do(func() {
		cache_TestLoop_loop = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_TestLoop_loop(v_0_box.IntVal))
})
	})
	return cache_TestLoop_loop
}

func Call_TestLoop_loop(v_0_loop int64) int64 {
loop:
for {
if false { continue loop }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0) == (0) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
continue loop
__t0 = gopurs_runtime.Value{}.IntVal
}
end_branch_0:
return __t0
}
}


