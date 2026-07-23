package Test_Fib

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var fib gopurs_runtime.Value
var once_fib sync.Once
func Get_fib() gopurs_runtime.Value {
	once_fib.Do(func() {
		fib = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
fib:
for {
if false { continue fib }
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Int(gopurs_runtime.Apply(Get_fib(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal)).IntVal + gopurs_runtime.Apply(Get_fib(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(2).IntVal)).IntVal)
}
end_branch_0:
return __t0
}
}()
})
	})
	return fib
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Fibonacci:"))
	})
	return describe
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_fib(), gopurs_runtime.Int(10))))
	})
	return act
}


