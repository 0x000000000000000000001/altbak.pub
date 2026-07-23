package Test_TCO

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Tail Call Optimization (100k calls):"))
	})
	return describe
}

var deepTailRec gopurs_runtime.Value
var once_deepTailRec sync.Once
func Get_deepTailRec() gopurs_runtime.Value {
	once_deepTailRec.Do(func() {
		deepTailRec = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
deepTailRec:
for {
if false { continue deepTailRec }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Get_deepTailRec(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Int(v1_1.IntVal + gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), v_0, gopurs_runtime.Int(3)).IntVal))
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return deepTailRec
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(100000))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply2(Get_deepTailRec(), dummy_1_1, gopurs_runtime.Int(0)))), gopurs_runtime.Value{})
})
}()
	})
	return act
}


