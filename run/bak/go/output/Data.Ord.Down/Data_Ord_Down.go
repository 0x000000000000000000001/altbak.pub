package Data_Ord_Down

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	unsafe "unsafe"
)

var cache_Down gopurs_runtime.Value
var once_Down sync.Once
func Get_Down() gopurs_runtime.Value {
	once_Down.Do(func() {
		cache_Down = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Down
}

var cache_showDown gopurs_runtime.Value
var once_showDown sync.Once
func Get_showDown() gopurs_runtime.Value {
	once_showDown.Do(func() {
		cache_showDown = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Down ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showDown
}

var cache_newtypeDown gopurs_runtime.Value
var once_newtypeDown sync.Once
func Get_newtypeDown() gopurs_runtime.Value {
	once_newtypeDown.Do(func() {
		cache_newtypeDown = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeDown
}

var cache_eqDown gopurs_runtime.Value
var once_eqDown sync.Once
func Get_eqDown() gopurs_runtime.Value {
	once_eqDown.Do(func() {
		cache_eqDown = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return cache_eqDown
}

var cache_ordDown gopurs_runtime.Value
var once_ordDown sync.Once
func Get_ordDown() gopurs_runtime.Value {
	once_ordDown.Do(func() {
		cache_ordDown = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), v_2, v1_3)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 380165415) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 1527465420) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}))
}()
})
	})
	return cache_ordDown
}

var cache_boundedDown gopurs_runtime.Value
var once_boundedDown sync.Once
func Get_boundedDown() gopurs_runtime.Value {
	once_boundedDown.Do(func() {
		cache_boundedDown = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_2_1
ordDown1_3_2 := gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compare"), v_3, v1_4)
_ = __local_var_5_3
var __t4 gopurs_runtime.Value
{
if (__local_var_5_3.Type == 9 && __local_var_5_3.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}
goto end_branch_4
} else {

}
}
{
if (__local_var_5_3.Type == 9 && __local_var_5_3.IntVal == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
goto end_branch_4
} else {

}
}
{
if (__local_var_5_3.Type == 9 && __local_var_5_3.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
_ = ordDown1_3_2
return gopurs_runtime.RecordDict3("top", "bottom", "Ord0", gopurs_runtime.RecordGet(dictBounded_0, "bottom"), gopurs_runtime.RecordGet(dictBounded_0, "top"), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordDown1_3_2
}))
}()
})
	})
	return cache_boundedDown
}




