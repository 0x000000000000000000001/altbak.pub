package Data_String_NonEmpty_CodeUnits

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	unsafe "unsafe"
)

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_snoc(c_0_box, s_1_box))
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), x_0)
}()
})
	})
	return cache_singleton
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(f_0_box, x_1_box)
})
	})
	return cache_takeWhile
}

var cache_lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		cache_lastIndexOf_prime = gopurs_runtime.Func(func(pat_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf_prime(), pat_0)
}()
})
	})
	return cache_lastIndexOf_prime
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf(), x_0)
}()
})
	})
	return cache_lastIndexOf
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Func(func(pat_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf_prime(), pat_0)
}()
})
	})
	return cache_indexOf_prime
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf(), x_0)
}()
})
	})
	return cache_indexOf
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), x_0)
}()
})
	})
	return cache_length
}

var cache_splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		cache_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitAt(i_0_box, nes_1_box)
})
	})
	return cache_splitAt
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take(i_0_box, nes_1_box)
})
	})
	return cache_take
}

var cache_takeRight gopurs_runtime.Value
var once_takeRight sync.Once
func Get_takeRight() gopurs_runtime.Value {
	once_takeRight.Do(func() {
		cache_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeRight(i_0_box, nes_1_box)
})
	})
	return cache_takeRight
}

var cache_toChar gopurs_runtime.Value
var once_toChar sync.Once
func Get_toChar() gopurs_runtime.Value {
	once_toChar.Do(func() {
		cache_toChar = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toChar(), x_0)
}()
})
	})
	return cache_toChar
}

var cache_toCharArray gopurs_runtime.Value
var once_toCharArray sync.Once
func Get_toCharArray() gopurs_runtime.Value {
	once_toCharArray.Do(func() {
		cache_toCharArray = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toCharArray(), x_0)
}()
})
	})
	return cache_toCharArray
}

var cache_toNonEmptyCharArray gopurs_runtime.Value
var once_toNonEmptyCharArray sync.Once
func Get_toNonEmptyCharArray() gopurs_runtime.Value {
	once_toNonEmptyCharArray.Do(func() {
		cache_toNonEmptyCharArray = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toCharArray(), x_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_1_0))).IntVal) > (0) {
__t1 = __local_var_1_0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}()
})
	})
	return cache_toNonEmptyCharArray
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(nes_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var nes_0 gopurs_runtime.Value = nes_0_loop
_ = nes_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(1), nes_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{__local_var_1_0})}
}
end_branch_1:
return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), nes_0), __t1)
}()
})
	})
	return cache_uncons
}

var cache_fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		cache_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_0, "foldMap1"), pkg_Data_Semigroup.Get_semigroupString(), Get_singleton())
}()
})
	})
	return cache_fromFoldable1
}

var cache_fromCharArray gopurs_runtime.Value
var once_fromCharArray sync.Once
func Get_fromCharArray() gopurs_runtime.Value {
	once_fromCharArray.Do(func() {
		cache_fromCharArray = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(v_0))).IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_fromCharArray(), v_0)})}
}
end_branch_0:
return __t0
}()
})
	})
	return cache_fromCharArray
}

var cache_fromNonEmptyCharArray gopurs_runtime.Value
var once_fromNonEmptyCharArray sync.Once
func Get_fromNonEmptyCharArray() gopurs_runtime.Value {
	once_fromNonEmptyCharArray.Do(func() {
		cache_fromNonEmptyCharArray = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(x_0))).IntVal) == (0) {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_fromCharArray(), x_0)
}
end_branch_0:
return __t0
}()
})
	})
	return cache_fromNonEmptyCharArray
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(f_0_box, x_1_box)
})
	})
	return cache_dropWhile
}

var cache_dropRight gopurs_runtime.Value
var once_dropRight sync.Once
func Get_dropRight() gopurs_runtime.Value {
	once_dropRight.Do(func() {
		cache_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropRight(i_0_box, nes_1_box)
})
	})
	return cache_dropRight
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop(i_0_box, nes_1_box)
})
	})
	return cache_drop
}

var cache_countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		cache_countPrefix = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_countPrefix(), x_0)
}()
})
	})
	return cache_countPrefix
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_cons(c_0_box, s_1_box))
})
	})
	return cache_cons
}

var cache_charAt gopurs_runtime.Value
var once_charAt sync.Once
func Get_charAt() gopurs_runtime.Value {
	once_charAt.Do(func() {
		cache_charAt = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_charAt(), x_0)
}()
})
	})
	return cache_charAt
}

func Call_snoc(c_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) string {
var c_0 gopurs_runtime.Value = c_0_loop
_ = c_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (s_1.StrVal()) + (gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), c_0).StrVal())
}

func Call_takeWhile(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_countPrefix(), f_0, x_1), x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{__local_var_2_0})}
}
end_branch_1:
return __t1
}

func Call_splitAt(i_0_loop gopurs_runtime.Value, nes_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 gopurs_runtime.Value = i_0_loop
_ = i_0
var nes_1 gopurs_runtime.Value = nes_1_loop
_ = nes_1
v_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_splitAt(), i_0, nes_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "before").StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordGet(v_2_0, "before")})}
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "after").StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.RecordGet(v_2_0, "after")})}
}
end_branch_2:
return gopurs_runtime.RecordDict2("before", "after", __t1, __t2)
}

func Call_take(i_0_loop gopurs_runtime.Value, nes_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 gopurs_runtime.Value = i_0_loop
_ = i_0
var nes_1 gopurs_runtime.Value = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0.IntVal) < (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), i_0, nes_1)})}
}
end_branch_0:
return __t0
}

func Call_takeRight(i_0_loop gopurs_runtime.Value, nes_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 gopurs_runtime.Value = i_0_loop
_ = i_0
var nes_1 gopurs_runtime.Value = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0.IntVal) < (1) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int((gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal) - (i_0.IntVal)), nes_1)})}
}
end_branch_0:
return __t0
}

func Call_dropWhile(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_countPrefix(), f_0, x_1), x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{__local_var_2_0})}
}
end_branch_1:
return __t1
}

func Call_dropRight(i_0_loop gopurs_runtime.Value, nes_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 gopurs_runtime.Value = i_0_loop
_ = i_0
var nes_1 gopurs_runtime.Value = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0.IntVal) >= (gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int((gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal) - (i_0.IntVal)), nes_1)})}
}
end_branch_0:
return __t0
}

func Call_drop(i_0_loop gopurs_runtime.Value, nes_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var i_0 gopurs_runtime.Value = i_0_loop
_ = i_0
var nes_1 gopurs_runtime.Value = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0.IntVal) >= (gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), nes_1).IntVal) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), i_0, nes_1)})}
}
end_branch_0:
return __t0
}

func Call_cons(c_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) string {
var c_0 gopurs_runtime.Value = c_0_loop
_ = c_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), c_0).StrVal()) + (s_1.StrVal())
}


