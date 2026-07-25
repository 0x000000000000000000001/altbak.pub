package Test_TCO

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
)

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Tail Call Optimization (100k calls):"))
	})
	return cache_describe
}

var cache_deepTailRec gopurs_runtime.Value
var once_deepTailRec sync.Once
func Get_deepTailRec() gopurs_runtime.Value {
	once_deepTailRec.Do(func() {
		cache_deepTailRec = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_deepTailRec(v_0_box, v1_1_box)
})
	})
	return cache_deepTailRec
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(100000))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), Call_deepTailRec(dummy_1_1, gopurs_runtime.Int(0)))), gopurs_runtime.Value{})
})
}()
	})
	return cache_act
}

func Call_deepTailRec(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
deepTailRec:
for {
if false { continue deepTailRec }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.IntVal) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = gopurs_runtime.Int((v_0.IntVal) - (1))
v1_1_loop = gopurs_runtime.Int((v1_1.IntVal) + (gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), v_0, gopurs_runtime.Int(3)).IntVal))
continue deepTailRec
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0
}
}


