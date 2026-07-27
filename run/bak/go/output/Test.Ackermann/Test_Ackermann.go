package Test_Ackermann

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Ackermann (3, 4):")), nil)
}()
})
	})
	return cache_describe
}

var cache_ackermann gopurs_runtime.Value
var once_ackermann sync.Once
func Get_ackermann() gopurs_runtime.Value {
	once_ackermann.Do(func() {
		cache_ackermann = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_ackermann(v_0_box.IntVal, v1_1_box.IntVal))
})
	})
	return cache_ackermann
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(3)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_ackermann(dummy_0.IntVal, 4))))
})), nil)
}()
})
	})
	return cache_act
}

func Call_ackermann(v_0_loop int64, v1_1_loop int64) int64 {
ackermann:
for {
if false { continue ackermann }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Int((v1_1) + (1))
goto end_branch_0
} else {

}
}
{
if (v1_1) == (0) {
v_0_loop = (v_0) - (1)
v1_1_loop = 1
continue ackermann
__t0 = gopurs_runtime.Value{}
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = gopurs_runtime.Int(Call_ackermann(v_0, (v1_1) - (1))).IntVal
continue ackermann
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0.IntVal
}
}
