package Data_List_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	unsafe "unsafe"
)

var cache_Leaf gopurs_runtime.Value
var once_Leaf sync.Once
func Get_Leaf() gopurs_runtime.Value {
	once_Leaf.Do(func() {
		cache_Leaf = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Leaf{})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{value0, value1, value2})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{value0, value1, value2, value3, value4})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoLeft{value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoRight{value0, value1})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeLeft{value0, value1, value2, value3})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeMiddle{value0, value1, value2, value3})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeRight{value0, value1, value2, value3})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{value0, value1, value2})}
})
})
})
	})
	return cache_KickUp
}

var cache_fromZipper gopurs_runtime.Value
var once_fromZipper sync.Once
func Get_fromZipper() gopurs_runtime.Value {
	once_fromZipper.Do(func() {
		cache_fromZipper = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromZipper(v_0_box, v1_1_box)
})
	})
	return cache_fromZipper
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

var cache_emptySet gopurs_runtime.Value
var once_emptySet sync.Once
func Get_emptySet() gopurs_runtime.Value {
	once_emptySet.Do(func() {
		cache_emptySet = gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Leaf{})}
	})
	return cache_emptySet
}

type Data_Data_List_Internal_Leaf struct {
	
}
func Is_Data_Data_List_Internal_Leaf(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2764020654
}

type Data_Data_List_Internal_Two struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_Two(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1177901036
}

type Data_Data_List_Internal_Three struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_Three(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1064476974
}

type Data_Data_List_Internal_TwoLeft struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_TwoLeft(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1304506903
}

type Data_Data_List_Internal_TwoRight struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_TwoRight(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2884341868
}

type Data_Data_List_Internal_ThreeLeft struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_ThreeLeft(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2195694037
}

type Data_Data_List_Internal_ThreeMiddle struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_ThreeMiddle(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 1584522659
}

type Data_Data_List_Internal_ThreeRight struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_ThreeRight(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 3952671150
}

type Data_Data_List_Internal_KickUp struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Data_List_Internal_KickUp(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2023586927
}

func Call_fromZipper(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
fromZipper:
for {
if false { continue fromZipper }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 786377863) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437) {
var __t1 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.IntVal == 1304506903) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{v1_1, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.IntVal == 2884341868) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, v1_1})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.IntVal == 2195694037) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{v1_1, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V3})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.IntVal == 1584522659) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, v1_1, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V3})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.IntVal == 3952671150) {
v_0_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V1
v1_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_0.UnsafePtr).V0.UnsafePtr).V3, v1_1})}
continue fromZipper
__t1 = gopurs_runtime.Value{}
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
}

func Call_insertAndLookupBy(comp_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value, orig_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var orig_2 gopurs_runtime.Value = orig_2_loop
_ = orig_2
var up_3_0 gopurs_runtime.Value
up_3_0 = gopurs_runtime.Func(func(v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
up_3_0:
for {
if false { continue up_3_0 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 786377863) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})}
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1358893437) {
var __t2 gopurs_runtime.Value
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.IntVal == 1304506903) {
__t2 = Call_fromZipper((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1})})
goto end_branch_2
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.IntVal == 2884341868) {
__t2 = Call_fromZipper((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1064476974, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Three{(*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_TwoRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})})
goto end_branch_2
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.IntVal == 2195694037) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})}, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeLeft)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.IntVal == 1584522659) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0})}, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V2, (*Data_Data_List_Internal_ThreeMiddle)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V3})}})}
continue up_3_0
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
if ((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.Type == 9 && (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.IntVal == 3952671150) {
v_4_loop = (*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V1
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V0, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V1, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V2})}, (*Data_Data_List_Internal_ThreeRight)((*pkg_Data_List_Types.Data_Data_List_Types_Cons)(v_4.UnsafePtr).V0.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1177901036, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Two{(*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V0, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V1, (*Data_Data_List_Internal_KickUp)(v1_5.UnsafePtr).V2})}})}
continue up_3_0
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
var down_4_3 gopurs_runtime.Value
down_4_3 = gopurs_runtime.Func(func(v_5_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
down_4_3:
for {
if false { continue down_4_3 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 2764020654) {
__t4 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(false), gopurs_runtime.Apply2(up_3_0, v_5, gopurs_runtime.Value{Type: 9, IntVal: 2023586927, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_KickUp{gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Leaf{})}, k_1, gopurs_runtime.Value{Type: 9, IntVal: 2764020654, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_Leaf{})}})}))
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1177901036) {
v2_7_5 := gopurs_runtime.Apply2(comp_0, k_1, (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V1)
_ = v2_7_5
var __t6 gopurs_runtime.Value
{
if (v2_7_5.Type == 9 && v2_7_5.IntVal == 902936544) {
__t6 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_6
} else {

}
}
{
if (v2_7_5.Type == 9 && v2_7_5.IntVal == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1304506903, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoLeft{(*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V2})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V0
continue down_4_3
__t6 = gopurs_runtime.Value{}
goto end_branch_6
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 2884341868, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_TwoRight{(*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V0, (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V1})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Two)(v1_6.UnsafePtr).V2
continue down_4_3
__t6 = gopurs_runtime.Value{}
}
end_branch_6:
__t4 = __t6
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1064476974) {
v2_7_7 := gopurs_runtime.Apply2(comp_0, k_1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1)
_ = v2_7_7
var __t10 gopurs_runtime.Value
{
if (v2_7_7.Type == 9 && v2_7_7.IntVal == 902936544) {
__t10 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_10
} else {

}
}
{
v3_8_8 := gopurs_runtime.Apply2(comp_0, k_1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3)
_ = v3_8_8
var __t9 gopurs_runtime.Value
{
if (v3_8_8.Type == 9 && v3_8_8.IntVal == 902936544) {
__t9 = gopurs_runtime.RecordDict2("found", "result", gopurs_runtime.Bool(true), orig_2)
goto end_branch_9
} else {

}
}
{
if (v2_7_7.Type == 9 && v2_7_7.IntVal == 1527465420) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 2195694037, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeLeft{(*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V2, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V4})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V0
continue down_4_3
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
if ((v2_7_7.Type == 9 && v2_7_7.IntVal == 380165415)) && ((v3_8_8.Type == 9 && v3_8_8.IntVal == 1527465420)) {
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 1584522659, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeMiddle{(*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V0, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V4})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V2
continue down_4_3
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Cons{gopurs_runtime.Value{Type: 9, IntVal: 3952671150, UnsafePtr: unsafe.Pointer(&Data_Data_List_Internal_ThreeRight{(*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V0, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V1, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V2, (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V3})}, v_5})}
v1_6_loop = (*Data_Data_List_Internal_Three)(v1_6.UnsafePtr).V4
continue down_4_3
__t9 = gopurs_runtime.Value{}
}
end_branch_9:
__t10 = __t9
}
end_branch_10:
__t4 = __t10
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(down_4_3, gopurs_runtime.Value{Type: 9, IntVal: 786377863, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Data_Data_List_Types_Nil{})}, orig_2)
}


