package Data_List_Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	unsafe "unsafe"
)

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_tail(_dollar__unused_0_box, (*pkg_Data_List_Types.Constructor_Cons[interface{}])(v_1_box.UnsafePtr)))}
})
	})
	return cache_tail
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_last(_dollar__unused_0_box, (*pkg_Data_List_Types.Constructor_Cons[interface{}])(v_1_box.UnsafePtr)))
})
	})
	return cache_last
}

var cache_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731 gopurs_runtime.Value
var once_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731 sync.Once
func Get_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731() gopurs_runtime.Value {
	once_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731.Do(func() {
		cache_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731(_dollar__unused_0_box, (*pkg_Data_List_Types.Constructor_Cons[interface{}])(v_1_box.UnsafePtr)))
})
	})
	return cache_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731
}

var cache_init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		cache_init_ = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_init_(_dollar__unused_0_box, (*pkg_Data_List_Types.Constructor_Cons[interface{}])(v_1_box.UnsafePtr)))}
})
	})
	return cache_init_
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_head(_dollar__unused_0_box, (*pkg_Data_List_Types.Constructor_Cons[interface{}])(v_1_box.UnsafePtr)))
})
	})
	return cache_head
}

func Call_tail(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[interface{}]) *pkg_Data_List_Types.Constructor_Cons[interface{}] {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[interface{}] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (*pkg_Data_List_Types.Constructor_Cons[interface{}])(__t0.UnsafePtr)
}

func Call_last(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[interface{}]) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[interface{}] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(Call_last(gopurs_runtime.Value{}, (*pkg_Data_List_Types.Constructor_Cons[interface{}])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}.UnsafePtr)))
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.UnboxAny(__t0)
}

func Call_last__func_gopurs_runtime_Value__ptrData_List_Types_Constructor_Cons[interface__]__interface___3843077731(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[interface{}]) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[interface{}] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Any((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(Call_last(gopurs_runtime.Value{}, (*pkg_Data_List_Types.Constructor_Cons[interface{}])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}.UnsafePtr)))
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.UnboxAny(__t0)
}

func Call_init_(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[interface{}]) *pkg_Data_List_Types.Constructor_Cons[interface{}] {
init_:
for {
if false { continue init_ }
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[interface{}] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0)), (*pkg_Data_List_Types.Constructor_Cons[interface{}])(gopurs_runtime.Apply2(Get_init_(), gopurs_runtime.Value{}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}).UnsafePtr)})})
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (*pkg_Data_List_Types.Constructor_Cons[interface{}])(__t0.UnsafePtr)
}
}

func Call_head(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[interface{}]) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[interface{}] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Any((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.UnboxAny(__t0)
}
