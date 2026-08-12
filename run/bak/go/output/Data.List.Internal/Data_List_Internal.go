package Data_List_Internal

import (
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		cache_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Leaf
}

var cache_Two gopurs_runtime.Value
var once_Two sync.Once
func Get_Two() gopurs_runtime.Value {
	once_Two.Do(func() {
		cache_Two = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, value0, value1, value2})}
})
})
})
	})
	return cache_Two
}

var cache_Three gopurs_runtime.Value
var once_Three sync.Once
func Get_Three() gopurs_runtime.Value {
	once_Three.Do(func() {
		cache_Three = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, value0, value1, value2, value3, value4})}
})
})
})
})
})
	})
	return cache_Three
}

var cache_TwoLeft gopurs_runtime.Value
var once_TwoLeft sync.Once
func Get_TwoLeft() gopurs_runtime.Value {
	once_TwoLeft.Do(func() {
		cache_TwoLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&Constructor_TwoLeft[gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_TwoLeft
}

var cache_TwoRight gopurs_runtime.Value
var once_TwoRight sync.Once
func Get_TwoRight() gopurs_runtime.Value {
	once_TwoRight.Do(func() {
		cache_TwoRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&Constructor_TwoRight[gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_TwoRight
}

var cache_ThreeLeft gopurs_runtime.Value
var once_ThreeLeft sync.Once
func Get_ThreeLeft() gopurs_runtime.Value {
	once_ThreeLeft.Do(func() {
		cache_ThreeLeft = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&Constructor_ThreeLeft[gopurs_runtime.Value]{1, value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_ThreeLeft
}

var cache_ThreeMiddle gopurs_runtime.Value
var once_ThreeMiddle sync.Once
func Get_ThreeMiddle() gopurs_runtime.Value {
	once_ThreeMiddle.Do(func() {
		cache_ThreeMiddle = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&Constructor_ThreeMiddle[gopurs_runtime.Value]{1, value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_ThreeMiddle
}

var cache_ThreeRight gopurs_runtime.Value
var once_ThreeRight sync.Once
func Get_ThreeRight() gopurs_runtime.Value {
	once_ThreeRight.Do(func() {
		cache_ThreeRight = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&Constructor_ThreeRight[gopurs_runtime.Value]{1, value0, value1, value2, value3})}
})
})
})
})
	})
	return cache_ThreeRight
}

var cache_KickUp gopurs_runtime.Value
var once_KickUp sync.Once
func Get_KickUp() gopurs_runtime.Value {
	once_KickUp.Do(func() {
		cache_KickUp = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, value0, value1, value2})}
})
})
})
	})
	return cache_KickUp
}

var cache_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any gopurs_runtime.Value
var once_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any sync.Once
func Get_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any() gopurs_runtime.Value {
	once_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any.Do(func() {
		cache_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)
})
	})
	return cache_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any
}

var cache_fromZipper gopurs_runtime.Value
var once_fromZipper sync.Once
func Get_fromZipper() gopurs_runtime.Value {
	once_fromZipper.Do(func() {
		cache_fromZipper = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)
})
	})
	return cache_fromZipper
}

var cache_fromZipper__gopurs_runtime_Value_1019554324 gopurs_runtime.Value
var once_fromZipper__gopurs_runtime_Value_1019554324 sync.Once
func Get_fromZipper__gopurs_runtime_Value_1019554324() gopurs_runtime.Value {
	once_fromZipper__gopurs_runtime_Value_1019554324.Do(func() {
		cache_fromZipper__gopurs_runtime_Value_1019554324 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromZipper__gopurs_runtime_Value_1019554324(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), v1_1_box)
})
	})
	return cache_fromZipper__gopurs_runtime_Value_1019554324
}

var cache_insertAndLookupBy gopurs_runtime.Value
var once_insertAndLookupBy sync.Once
func Get_insertAndLookupBy() gopurs_runtime.Value {
	once_insertAndLookupBy.Do(func() {
		cache_insertAndLookupBy = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, orig_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAndLookupBy(comp_0_box, k_1_box, orig_2_box)
})
	})
	return cache_insertAndLookupBy
}

var cache_insertAndLookupBy__gopurs_runtime_Value_3244745033 gopurs_runtime.Value
var once_insertAndLookupBy__gopurs_runtime_Value_3244745033 sync.Once
func Get_insertAndLookupBy__gopurs_runtime_Value_3244745033() gopurs_runtime.Value {
	once_insertAndLookupBy__gopurs_runtime_Value_3244745033.Do(func() {
		cache_insertAndLookupBy__gopurs_runtime_Value_3244745033 = gopurs_runtime.Func3(func(comp_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value, orig_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_insertAndLookupBy__gopurs_runtime_Value_3244745033(comp_0_box, k_1_box, orig_2_box)
})
	})
	return cache_insertAndLookupBy__gopurs_runtime_Value_3244745033
}

var cache_emptySet gopurs_runtime.Value
var once_emptySet sync.Once
func Get_emptySet() gopurs_runtime.Value {
	once_emptySet.Do(func() {
		cache_emptySet = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_emptySet
}

var cache_emptySet__gopurs_runtime_Value_2398681994 gopurs_runtime.Value
var once_emptySet__gopurs_runtime_Value_2398681994 sync.Once
func Get_emptySet__gopurs_runtime_Value_2398681994() gopurs_runtime.Value {
	once_emptySet__gopurs_runtime_Value_2398681994.Do(func() {
		cache_emptySet__gopurs_runtime_Value_2398681994 = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_emptySet__gopurs_runtime_Value_2398681994
}

type Constructor_Leaf[T_k any] struct {
	Rc uint32
}


type Constructor_Two[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
}


type Constructor_Three[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
	V3 T_k
	V4 gopurs_runtime.Value
}


type Constructor_TwoLeft[T_k any] struct {
	Rc uint32
	V0 T_k
	V1 gopurs_runtime.Value
}


type Constructor_TwoRight[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
}


type Constructor_ThreeLeft[T_k any] struct {
	Rc uint32
	V0 T_k
	V1 gopurs_runtime.Value
	V2 T_k
	V3 gopurs_runtime.Value
}


type Constructor_ThreeMiddle[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 T_k
	V3 gopurs_runtime.Value
}


type Constructor_ThreeRight[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
	V3 T_k
}


type Constructor_KickUp[T_k any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_k
	V2 gopurs_runtime.Value
}


func Call_fromZipper_ForAll_k_Func_ADT_Data_List_Types_List_ADT_Data_List_Internal_TreeContext_Any_ADT_Data_List_Internal_Set_Any_ADT_Data_List_Internal_Set_Any(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1304506903) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, v1_1, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2884341868) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1})})
goto end_branch_1
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2195694037) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, v1_1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_1
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1584522659) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_1
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3952671150) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3, v1_1})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
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
return __t0
}

func Call_fromZipper(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1304506903) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, v1_1, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2884341868) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1})})
goto end_branch_1
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2195694037) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, v1_1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_1
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1584522659) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_1
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3952671150) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3, v1_1})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
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
return __t0
}

func Call_fromZipper__gopurs_runtime_Value_1019554324(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1304506903) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, v1_1, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_1
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 2884341868) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1})})
goto end_branch_1
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2195694037) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, v1_1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_1
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 1584522659) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, v1_1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3})})
goto end_branch_1
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 3952671150) {
__t1 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0.UnsafePtr).V3, v1_1})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
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
return __t0
}

func Call_insertAndLookupBy(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, orig_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var orig_2 gopurs_runtime.Value = orig_2_loop
_ = orig_2
var up_3_0_0 gopurs_runtime.Value
up_3_0_0 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
up_3_0_0:
for {
if false { continue up_3_0_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1304506903) {
__t2 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2884341868) {
__t2 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})})
goto end_branch_2
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2195694037) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 1584522659) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0})}, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 3952671150) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2})}, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}})}
continue up_3_0_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
})
var down_4_8_1 gopurs_runtime.Value
down_4_8_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
down_4_8_1:
for {
if false { continue down_4_8_1 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t9 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 2764020654) {
__t9 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0_0, v_5, gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}})}))
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1177901036) {
v2_7_10 := gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_10
var __t11 gopurs_runtime.Value
{
if (uint32(v2_7_10.IntVal) == 902936544) {
__t11 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_11
} else {

}
}
{
if (uint32(v2_7_10.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&Constructor_TwoLeft[gopurs_runtime.Value]{1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_1
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&Constructor_TwoRight[gopurs_runtime.Value]{1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_1
__t11 = gopurs_runtime.Value{}
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1064476974) {
v2_7_12 := gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_12
var __t15 gopurs_runtime.Value
{
if (uint32(v2_7_12.IntVal) == 902936544) {
__t15 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_15
} else {

}
}
{
v3_8_13 := gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3)
_ = v3_8_13
var __t14 gopurs_runtime.Value
{
if (uint32(v3_8_13.IntVal) == 902936544) {
__t14 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_14
} else {

}
}
{
if (uint32(v2_7_12.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&Constructor_ThreeLeft[gopurs_runtime.Value]{1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_1
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
if ((uint32(v2_7_12.IntVal) == 380165415)) && ((uint32(v3_8_13.IntVal) == 1527465420)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&Constructor_ThreeMiddle[gopurs_runtime.Value]{1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_1
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&Constructor_ThreeRight[gopurs_runtime.Value]{1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4
continue down_4_8_1
__t14 = gopurs_runtime.Value{}
}
end_branch_14:
__t15 = __t14
}
end_branch_15:
__t9 = __t15
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
return gopurs_runtime.Apply2(down_4_8_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, orig_2)
}

func Call_insertAndLookupBy__gopurs_runtime_Value_3244745033(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, orig_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var orig_2 gopurs_runtime.Value = orig_2_loop
_ = orig_2
var up_3_0_2 gopurs_runtime.Value
up_3_0_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
up_3_0_2:
for {
if false { continue up_3_0_2 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437 && v_4.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1304506903) {
__t2 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_2
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2884341868) {
__t2 = Call_fromZipper(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}), gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Constructor_Three[gopurs_runtime.Value]{1, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_TwoRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})})
goto end_branch_2
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 2195694037) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeLeft[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_2
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 1584522659) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0})}, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Constructor_ThreeMiddle[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0_2
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0
if (__t_tag_7.Type == 9 && __t_tag_7.IntVal == 3952671150) {
v_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V2})}, (*Constructor_ThreeRight[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_4.UnsafePtr).V0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Constructor_Two[gopurs_runtime.Value]{1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V1, (*Constructor_KickUp[gopurs_runtime.Value])(v1_5.UnsafePtr).V2})}})}
continue up_3_0_2
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
})
var down_4_8_3 gopurs_runtime.Value
down_4_8_3 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
down_4_8_3:
for {
if false { continue down_4_8_3 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t9 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 2764020654) {
__t9 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0_2, v_5, gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Constructor_KickUp[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(nil)}})}))
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1177901036) {
v2_7_10 := gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_10
var __t11 gopurs_runtime.Value
{
if (uint32(v2_7_10.IntVal) == 902936544) {
__t11 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_11
} else {

}
}
{
if (uint32(v2_7_10.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&Constructor_TwoLeft[gopurs_runtime.Value]{1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_3
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&Constructor_TwoRight[gopurs_runtime.Value]{1, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V1})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Two[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_3
__t11 = gopurs_runtime.Value{}
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1064476974) {
v2_7_12 := gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)
_ = v2_7_12
var __t15 gopurs_runtime.Value
{
if (uint32(v2_7_12.IntVal) == 902936544) {
__t15 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_15
} else {

}
}
{
v3_8_13 := gopurs_runtime.Apply2(comp_0, k_1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3)
_ = v3_8_13
var __t14 gopurs_runtime.Value
{
if (uint32(v3_8_13.IntVal) == 902936544) {
__t14 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_14
} else {

}
}
{
if (uint32(v2_7_12.IntVal) == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&Constructor_ThreeLeft[gopurs_runtime.Value]{1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
continue down_4_8_3
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
if ((uint32(v2_7_12.IntVal) == 380165415)) && ((uint32(v3_8_13.IntVal) == 1527465420)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&Constructor_ThreeMiddle[gopurs_runtime.Value]{1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2
continue down_4_8_3
__t14 = gopurs_runtime.Value{}
goto end_branch_14
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&Constructor_ThreeRight[gopurs_runtime.Value]{1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V0, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V1, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V2, (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V3})}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_5)})}
v1_6_loop = (*Constructor_Three[gopurs_runtime.Value])(v1_6.UnsafePtr).V4
continue down_4_8_3
__t14 = gopurs_runtime.Value{}
}
end_branch_14:
__t15 = __t14
}
end_branch_15:
__t9 = __t15
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}
}()
})
})
return gopurs_runtime.Apply2(down_4_8_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, orig_2)
}


