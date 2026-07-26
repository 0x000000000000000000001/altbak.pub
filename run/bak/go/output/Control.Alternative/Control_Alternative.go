package Control_Alternative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		cache_guard = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard(dictAlternative_0_box)
})
	})
	return cache_guard
}

var cache_guard__gopurs_runtime_Value_489569669 gopurs_runtime.Value
var once_guard__gopurs_runtime_Value_489569669 sync.Once
func Get_guard__gopurs_runtime_Value_489569669() gopurs_runtime.Value {
	once_guard__gopurs_runtime_Value_489569669.Do(func() {
		cache_guard__gopurs_runtime_Value_489569669 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard__gopurs_runtime_Value_489569669(dictAlternative_0_box)
})
	})
	return cache_guard__gopurs_runtime_Value_489569669
}

var cache_alternativeArray gopurs_runtime.Value
var once_alternativeArray sync.Once
func Get_alternativeArray() gopurs_runtime.Value {
	once_alternativeArray.Do(func() {
		cache_alternativeArray = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
}))
	})
	return cache_alternativeArray
}

func Call_guard(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
empty_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "empty")
_ = empty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) != (0) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
__t1 = empty_1_0
}
end_branch_1:
return __t1
})
}

func Call_guard__gopurs_runtime_Value_489569669(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
empty_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "empty")
_ = empty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) != (0) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
__t1 = empty_1_0
}
end_branch_1:
return __t1
})
}


