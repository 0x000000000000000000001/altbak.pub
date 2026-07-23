package Test_Ackermann

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
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
		ackermann = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
ackermann:
for {
if false { continue ackermann }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Int(v1_1.IntVal + gopurs_runtime.Int(1).IntVal)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v1_1.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(Get_ackermann(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Int(1))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Get_ackermann(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Apply2(Get_ackermann(), v_0, gopurs_runtime.Int(v1_1.IntVal - gopurs_runtime.Int(1).IntVal)))
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return ackermann
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply2(Get_ackermann(), gopurs_runtime.Int(3), gopurs_runtime.Int(4))))
	})
	return act
}


