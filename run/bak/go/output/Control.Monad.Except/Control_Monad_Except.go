package Control_Monad_Except

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var withExcept gopurs_runtime.Value
var once_withExcept sync.Once
func Get_withExcept() gopurs_runtime.Value {
	once_withExcept.Do(func() {
		withExcept = gopurs_runtime.Func2(Call_withExcept)
	})
	return withExcept
}

var runExcept gopurs_runtime.Value
var once_runExcept sync.Once
func Get_runExcept() gopurs_runtime.Value {
	once_runExcept.Do(func() {
		runExcept = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return runExcept
}

var mapExcept gopurs_runtime.Value
var once_mapExcept sync.Once
func Get_mapExcept() gopurs_runtime.Value {
	once_mapExcept.Do(func() {
		mapExcept = gopurs_runtime.Func2(Call_mapExcept)
	})
	return mapExcept
}

func Call_withExcept(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1_loop.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Right", (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0])
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_1_loop.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(f_0_loop, (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0]))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_mapExcept(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0_loop, v_1_loop)
}


