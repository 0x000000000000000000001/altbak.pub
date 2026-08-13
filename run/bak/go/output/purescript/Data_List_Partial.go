package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_List_Partial_tail gopurs_runtime.Value
var once_Data_List_Partial_tail sync.Once
func Get_Data_List_Partial_tail() gopurs_runtime.Value {
	once_Data_List_Partial_tail.Do(func() {
		cache_Data_List_Partial_tail = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Partial_tail(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_box)))}
})
	})
	return cache_Data_List_Partial_tail
}

var cache_Data_List_Partial_last gopurs_runtime.Value
var once_Data_List_Partial_last sync.Once
func Get_Data_List_Partial_last() gopurs_runtime.Value {
	once_Data_List_Partial_last.Do(func() {
		cache_Data_List_Partial_last = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Partial_last(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_box))
})
	})
	return cache_Data_List_Partial_last
}

var cache_Data_List_Partial_init gopurs_runtime.Value
var once_Data_List_Partial_init sync.Once
func Get_Data_List_Partial_init() gopurs_runtime.Value {
	once_Data_List_Partial_init.Do(func() {
		cache_Data_List_Partial_init = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Partial_init(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_box)))}
})
	})
	return cache_Data_List_Partial_init
}

var cache_Data_List_Partial_head gopurs_runtime.Value
var once_Data_List_Partial_head sync.Once
func Get_Data_List_Partial_head() gopurs_runtime.Value {
	once_Data_List_Partial_head.Do(func() {
		cache_Data_List_Partial_head = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Partial_head(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_box))
})
	})
	return cache_Data_List_Partial_head
}

var cache_Data_List_Partial_init__1987719293 gopurs_runtime.Value
var once_Data_List_Partial_init__1987719293 sync.Once
func Get_Data_List_Partial_init__1987719293() gopurs_runtime.Value {
	once_Data_List_Partial_init__1987719293.Do(func() {
		cache_Data_List_Partial_init__1987719293 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_List_Partial_init__1987719293(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_box)))}
})
	})
	return cache_Data_List_Partial_init__1987719293
}

var cache_Data_List_Partial_last__2622911464 gopurs_runtime.Value
var once_Data_List_Partial_last__2622911464 sync.Once
func Get_Data_List_Partial_last__2622911464() gopurs_runtime.Value {
	once_Data_List_Partial_last__2622911464.Do(func() {
		cache_Data_List_Partial_last__2622911464 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_List_Partial_last__2622911464(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_1_box))
})
	})
	return cache_Data_List_Partial_last__2622911464
}

func Call_Data_List_Partial_tail(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var __t0 *Constructor_Data_List_Types_Cons
{
if (v_1 != nil) {
__t0 = (v_1).V1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_0:
return __t0
}

func Call_Data_List_Partial_last(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
last:
for {
if false { continue last }
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var __t2 gopurs_runtime.Value
{
if (v_1 != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = (v_1).V1
if (__t_tag_0 == nil) {
__t1 = (v_1).V0
goto end_branch_1
} else {

}
}
{
_dollar__unused_0_loop = gopurs_runtime.Value{}
v_1_loop = (v_1).V1
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

func Call_Data_List_Partial_init(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
init:
for {
if false { continue init }
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var __t2 *Constructor_Data_List_Types_Cons
{
if (v_1 != nil) {
var __t1 *Constructor_Data_List_Types_Cons
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = (v_1).V1
if (__t_tag_0 == nil) {
__t1 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Types_Cons{1, (v_1).V0, Call_Data_List_Partial_init(gopurs_runtime.Value{}, (v_1).V1)}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return __t2
}
}

func Call_Data_List_Partial_head(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (v_1 != nil) {
__t0 = (v_1).V0
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

func Call_Data_List_Partial_init__1987719293(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var __t6 *Constructor_Data_List_Types_Cons
{
if (v_1 != nil) {
var __t5 *Constructor_Data_List_Types_Cons
{
var __t_tag_4 *Constructor_Data_List_Types_Cons = (v_1).V1
if (__t_tag_4 == nil) {
__t5 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_5
} else {

}
}
{
var __t3 *Constructor_Data_List_Types_Cons
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = (v_1).V1
if (__t_tag_0 != nil) {
var __t2 *Constructor_Data_List_Types_Cons
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = ((v_1).V1).V1
if (__t_tag_1 == nil) {
__t2 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_List_Types_Cons{1, ((v_1).V1).V0, Call_Data_List_Partial_init(gopurs_runtime.Value{}, ((v_1).V1).V1)}
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
__t5 = &Constructor_Data_List_Types_Cons{1, (v_1).V0, __t3}
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return __t6
}

func Call_Data_List_Partial_last__2622911464(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *Constructor_Data_List_Types_Cons = v_1_loop
_ = v_1
var __t6 gopurs_runtime.Value
{
if (v_1 != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_List_Types_Cons = (v_1).V1
if (__t_tag_4 == nil) {
__t5 = (v_1).V0
goto end_branch_5
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = (v_1).V1
if (__t_tag_0 != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_List_Types_Cons = ((v_1).V1).V1
if (__t_tag_1 == nil) {
__t2 = ((v_1).V1).V0
goto end_branch_2
} else {

}
}
{
__t2 = Call_Data_List_Partial_last(gopurs_runtime.Value{}, ((v_1).V1).V1)
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t5 = __t3
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}


