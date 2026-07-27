package Test_TCO

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
)

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Tail Call Optimization (100k calls):")), nil)
}()
})
	})
	return cache_describe
}

var cache_deepTailRec gopurs_runtime.Value
var once_deepTailRec sync.Once
func Get_deepTailRec() gopurs_runtime.Value {
	once_deepTailRec.Do(func() {
		cache_deepTailRec = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_deepTailRec(v_0_box.IntVal, v1_1_box.IntVal))
})
	})
	return cache_deepTailRec
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(100000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_deepTailRec(dummy_0.IntVal, 0))))
})), nil)
}()
})
	})
	return cache_act
}

func Call_deepTailRec(v_0_loop int64, v1_1_loop int64) int64 {
deepTailRec:
for {
if false { continue deepTailRec }
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
v1_1_loop = (v1_1) + (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), gopurs_runtime.Int(v_0), gopurs_runtime.Int(3)).IntVal)
continue deepTailRec
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0.IntVal
}
}
