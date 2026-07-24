package Control_Alternative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Plus "gopurs/output/Control.Plus"
)

var guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		guard = gopurs_runtime.Func(func(dictAlternative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
empty_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0_loop, "Plus1"), gopurs_runtime.Value{}), "empty")
_ = empty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if v_2.IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0_loop, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
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
}()
})
	})
	return guard
}

var alternativeArray gopurs_runtime.Value
var once_alternativeArray sync.Once
func Get_alternativeArray() gopurs_runtime.Value {
	once_alternativeArray.Do(func() {
		alternativeArray = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Plus.Get_plusArray()
}))
	})
	return alternativeArray
}




