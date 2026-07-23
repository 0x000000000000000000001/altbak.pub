package Control_Monad_Except

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var withExcept gopurs_runtime.Value
var once_withExcept sync.Once
func Get_withExcept() gopurs_runtime.Value {
	once_withExcept.Do(func() {
		withExcept = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), gopurs_runtime.RecordGet(v_1, "value0"))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Left"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_1, "value0")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return withExcept
}

var runExcept gopurs_runtime.Value
var once_runExcept sync.Once
func Get_runExcept() gopurs_runtime.Value {
	once_runExcept.Do(func() {
		runExcept = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return runExcept
}

var mapExcept gopurs_runtime.Value
var once_mapExcept sync.Once
func Get_mapExcept() gopurs_runtime.Value {
	once_mapExcept.Do(func() {
		mapExcept = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
	})
	return mapExcept
}


