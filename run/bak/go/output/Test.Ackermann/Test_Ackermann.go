package Test_Ackermann

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Ackermann (3, 4):"))
	})
	return describe
}

var ackermann gopurs_runtime.Value
var once_ackermann sync.Once
func Get_ackermann() gopurs_runtime.Value {
	once_ackermann.Do(func() {
		ackermann = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ackermann(v_0_box, v1_1_box)
})
	})
	return ackermann
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(3))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), Call_ackermann(dummy_1_1, gopurs_runtime.Int(4)))), gopurs_runtime.Value{})
})
}()
	})
	return act
}

func Call_ackermann(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
ackermann:
for {
if false { continue ackermann }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if v_0.IntVal == 0 {
__t0 = gopurs_runtime.Int(v1_1.IntVal + 1)
goto end_branch_0
} else {

}
}
{
if v1_1.IntVal == 0 {
v_0_loop = gopurs_runtime.Int(v_0.IntVal - 1)
v1_1_loop = gopurs_runtime.Int(1)
continue ackermann
__t0 = gopurs_runtime.Value{}
goto end_branch_0
} else {

}
}
{
v_0_loop = gopurs_runtime.Int(v_0.IntVal - 1)
v1_1_loop = Call_ackermann(v_0, gopurs_runtime.Int(v1_1.IntVal - 1))
continue ackermann
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0
}
}


