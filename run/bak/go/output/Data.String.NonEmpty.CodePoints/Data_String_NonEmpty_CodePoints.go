package Data_String_NonEmpty_CodePoints

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_CodePoints "gopurs/output/Data.String.CodePoints"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Array_NonEmpty "gopurs/output/Data.Array.NonEmpty"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	unsafe "unsafe"
)

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420))
})
}()
	})
	return cache_lessThan
}

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool(((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1527465420)) != (true))
})
}()
	})
	return cache_greaterThanOrEq
}

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_snoc(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_singleton(x_0_box.IntVal))
})
	})
	return cache_singleton
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_takeWhile(f_0_box, x_1_box.StrVal()))}
})
	})
	return cache_takeWhile
}

var cache_lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		cache_lastIndexOf_prime = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_lastIndexOf_prime(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal()))}
})
	})
	return cache_lastIndexOf_prime
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_lastIndexOf(x_0_box.StrVal(), v_1_box.StrVal()))}
})
	})
	return cache_lastIndexOf
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_indexOf_prime(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal()))}
})
	})
	return cache_indexOf_prime
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_indexOf(x_0_box.StrVal(), v_1_box.StrVal()))}
})
	})
	return cache_indexOf
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_length(x_0_box.StrVal()))
})
	})
	return cache_length
}

var cache_splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		cache_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
})
	})
	return cache_splitAt
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_take(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_take
}

var cache_toCodePointArray gopurs_runtime.Value
var once_toCodePointArray sync.Once
func Get_toCodePointArray() gopurs_runtime.Value {
	once_toCodePointArray.Do(func() {
		cache_toCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toCodePointArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toCodePointArray
}

var cache_toNonEmptyCodePointArray gopurs_runtime.Value
var once_toNonEmptyCodePointArray sync.Once
func Get_toNonEmptyCodePointArray() gopurs_runtime.Value {
	once_toNonEmptyCodePointArray.Do(func() {
		cache_toNonEmptyCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toNonEmptyCodePointArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toNonEmptyCodePointArray
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(nes_0_box.StrVal())
})
	})
	return cache_uncons
}

var cache_fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		cache_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable1(dictFoldable1_0_box)
})
	})
	return cache_fromFoldable1
}

var cache_fromCodePointArray gopurs_runtime.Value
var once_fromCodePointArray sync.Once
func Get_fromCodePointArray() gopurs_runtime.Value {
	once_fromCodePointArray.Do(func() {
		cache_fromCodePointArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromCodePointArray(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))}
})
	})
	return cache_fromCodePointArray
}

var cache_fromNonEmptyCodePointArray gopurs_runtime.Value
var once_fromNonEmptyCodePointArray sync.Once
func Get_fromNonEmptyCodePointArray() gopurs_runtime.Value {
	once_fromNonEmptyCodePointArray.Do(func() {
		cache_fromNonEmptyCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_fromNonEmptyCodePointArray(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))
})
	})
	return cache_fromNonEmptyCodePointArray
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_dropWhile(f_0_box, x_1_box.StrVal()))}
})
	})
	return cache_dropWhile
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_drop(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_drop
}

var cache_countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		cache_countPrefix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_countPrefix(x_0_box)
})
	})
	return cache_countPrefix
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_cons(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_cons
}

var cache_codePointAt gopurs_runtime.Value
var once_codePointAt sync.Once
func Get_codePointAt() gopurs_runtime.Value {
	once_codePointAt.Do(func() {
		cache_codePointAt = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_codePointAt(x_0_box.IntVal, v_1_box.StrVal()))}
})
	})
	return cache_codePointAt
}

func Call_snoc(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(s_1), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), gopurs_runtime.Int(c_0))).StrVal()
}

func Call_singleton(x_0_loop int64) string {
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), gopurs_runtime.Int(x_0)).StrVal()
}

func Call_takeWhile(f_0_loop gopurs_runtime.Value, x_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_countPrefix(), f_0, gopurs_runtime.Str(x_1)), gopurs_runtime.Str(x_1))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{__local_var_2_0})}
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t1.UnsafePtr)
}

func Call_lastIndexOf_prime(pat_0_loop string, x_1_loop int64, v_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply3(pkg_Data_String_CodePoints.Get_lastIndexOf_prime(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2)).UnsafePtr)
}

func Call_lastIndexOf(x_0_loop string, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_lastIndexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1)).UnsafePtr)
}

func Call_indexOf_prime(pat_0_loop string, x_1_loop int64, v_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply3(pkg_Data_String_CodePoints.Get_indexOf_prime(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2)).UnsafePtr)
}

func Call_indexOf(x_0_loop string, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_indexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1)).UnsafePtr)
}

func Call_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), gopurs_runtime.Str(x_0))))).IntVal
}

func Call_splitAt(i_0_loop int64, nes_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
v_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "after").StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordGet(v_2_0, "after")})}
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "before").StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordGet(v_2_0, "before")})}
}
end_branch_2:
return gopurs_runtime.RecordDict2("after", "before", __t1, __t2)
}

func Call_take(i_0_loop int64, nes_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))})}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t0.UnsafePtr)
}

func Call_toCodePointArray(x_0_loop string) []int64 {
var x_0 string = x_0_loop
_ = x_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), gopurs_runtime.Str(x_0)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_toNonEmptyCodePointArray(x_0_loop string) []int64 {
var x_0 string = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(pkg_Data_Array_NonEmpty.Get_greaterThan(), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_1_0))), gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = __local_var_1_0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__t1.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_uncons(nes_0_loop string) gopurs_runtime.Value {
var nes_0 string = nes_0_loop
_ = nes_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_codePointAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(nes_0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
__local_var_1_2 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(1), gopurs_runtime.Str(nes_0))), gopurs_runtime.Str(nes_0))
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (__local_var_1_2.StrVal()) == ("") {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{__local_var_1_2})}
}
end_branch_3:
return gopurs_runtime.RecordDict2("head", "tail", __t1, __t3)
}

func Call_fromFoldable1(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), pkg_Data_Semigroup.Get_semigroupString(), Get_singleton())
}

func Call_fromCodePointArray(v_0_loop []int64) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 []int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_fromCodePointArray(), func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}())})}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t0.UnsafePtr)
}

func Call_fromNonEmptyCodePointArray(x_0_loop []int64) string {
var x_0 []int64 = x_0_loop
_ = x_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) == (0) {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_fromCodePointArray(), func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}())
}
end_branch_0:
return __t0.StrVal()
}

func Call_dropWhile(f_0_loop gopurs_runtime.Value, x_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_dropWhile(), f_0, gopurs_runtime.Str(x_1))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{__local_var_2_0})}
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t1.UnsafePtr)
}

func Call_drop(i_0_loop int64, nes_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThanOrEq(), gopurs_runtime.Int(i_0), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), gopurs_runtime.Str(nes_1)))))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))), gopurs_runtime.Str(nes_1))})}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t0.UnsafePtr)
}

func Call_countPrefix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_countPrefix(), x_0)
}

func Call_cons(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), gopurs_runtime.Int(c_0)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_codePointAt(x_0_loop int64, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_codePointAt(), gopurs_runtime.Int(x_0), gopurs_runtime.Str(v_1)).UnsafePtr)
}


