package Data_List_Partial

import (
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_tail(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_tail
}

var cache_last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		cache_last = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_last(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_last
}

var cache_init gopurs_runtime.Value
var once_init sync.Once
func Get_init() gopurs_runtime.Value {
	once_init.Do(func() {
		cache_init = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_init(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_init
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_head
}

var cache_init__1987719293 gopurs_runtime.Value
var once_init__1987719293 sync.Once
func Get_init__1987719293() gopurs_runtime.Value {
	once_init__1987719293.Do(func() {
		cache_init__1987719293 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_init__1987719293(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_init__1987719293
}

var cache_last__2622911464 gopurs_runtime.Value
var once_last__2622911464 sync.Once
func Get_last__2622911464() gopurs_runtime.Value {
	once_last__2622911464.Do(func() {
		cache_last__2622911464 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_last__2622911464(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_last__2622911464
}

func Call_tail(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
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
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_last(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
last:
for {
if false { continue last }
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t1 = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
_dollar__unused_0_loop = gopurs_runtime.Value{}
v_1_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1
continue last
__t1 = gopurs_runtime.Value{}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}

func Call_init(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
init:
for {
if false { continue init }
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_init(gopurs_runtime.Value{}, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2)
}
}

func Call_head(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_init__1987719293(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_init(gopurs_runtime.Value{}, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1))})})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2)
}

func Call_last__2622911464(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)}
if (__t_tag_0.Type == 9 && __t_tag_0.IntVal == 1358893437 && __t_tag_0.UnsafePtr == nil) {
__t1 = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = Call_last(gopurs_runtime.Value{}, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}


