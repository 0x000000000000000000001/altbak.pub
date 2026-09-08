package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_NonEmpty_CodePoints_toNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodePoints_toNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toNonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_toNonEmptyString(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_toNonEmptyString
}

var cache_Data_String_NonEmpty_CodePoints_snoc gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_snoc sync.Once
func Get_Data_String_NonEmpty_CodePoints_snoc() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_snoc.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_snoc(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_snoc
}

var cache_Data_String_NonEmpty_CodePoints_singleton gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_singleton sync.Once
func Get_Data_String_NonEmpty_CodePoints_singleton() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_singleton.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_singleton(x_0_box.IntVal))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_singleton
}

var cache_Data_String_NonEmpty_CodePoints_liftS gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS sync.Once
func Get_Data_String_NonEmpty_CodePoints_liftS() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_liftS(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS
}

var cache_Data_String_NonEmpty_CodePoints_liftS__1645621157 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS__1645621157 sync.Once
func Get_Data_String_NonEmpty_CodePoints_liftS__1645621157() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS__1645621157.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS__1645621157 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_liftS__1645621157(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS__1645621157
}

var cache_Data_String_NonEmpty_CodePoints_takeWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_takeWhile sync.Once
func Get_Data_String_NonEmpty_CodePoints_takeWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_takeWhile.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_takeWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_takeWhile(f_0_box, x_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_takeWhile
}

var cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_ sync.Once
func Get_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_ = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_
}

var cache_Data_String_NonEmpty_CodePoints_lastIndexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_lastIndexOf sync.Once
func Get_Data_String_NonEmpty_CodePoints_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_lastIndexOf.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_lastIndexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_lastIndexOf(x_0_box.StrVal(), v_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_lastIndexOf
}

var cache_Data_String_NonEmpty_CodePoints_indexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_indexOf_prime_ sync.Once
func Get_Data_String_NonEmpty_CodePoints_indexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_indexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_indexOf_prime_ = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_indexOf_prime_(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_indexOf_prime_
}

var cache_Data_String_NonEmpty_CodePoints_indexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_indexOf sync.Once
func Get_Data_String_NonEmpty_CodePoints_indexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_indexOf.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_indexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_indexOf(x_0_box.StrVal(), v_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_indexOf
}

var cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_fromNonEmptyString(v_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString
}

var cache_Data_String_NonEmpty_CodePoints_length gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_length sync.Once
func Get_Data_String_NonEmpty_CodePoints_length() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_length.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_NonEmpty_CodePoints_length(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_length
}

var cache_Data_String_NonEmpty_CodePoints_splitAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_splitAt sync.Once
func Get_Data_String_NonEmpty_CodePoints_splitAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_splitAt.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodePoints_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(orig.after))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(orig.before))})
				}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_splitAt
}

var cache_Data_String_NonEmpty_CodePoints_take gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_take sync.Once
func Get_Data_String_NonEmpty_CodePoints_take() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_take.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_take(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_take
}

var cache_Data_String_NonEmpty_CodePoints_toCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_toCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_String_NonEmpty_CodePoints_toCodePointArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_toCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_uncons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_uncons sync.Once
func Get_Data_String_NonEmpty_CodePoints_uncons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_uncons.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodePoints_uncons(nes_0_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(orig.head), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(orig.tail))})
				}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_uncons
}

var cache_Data_String_NonEmpty_CodePoints_fromFoldable1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromFoldable1 sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromFoldable1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromFoldable1.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromFoldable1
}

var cache_Data_String_NonEmpty_CodePoints_fromCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromCodePointArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_fromCodePointArray(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_dropWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_dropWhile sync.Once
func Get_Data_String_NonEmpty_CodePoints_dropWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_dropWhile.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_dropWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_dropWhile(f_0_box, x_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_dropWhile
}

var cache_Data_String_NonEmpty_CodePoints_drop gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_drop sync.Once
func Get_Data_String_NonEmpty_CodePoints_drop() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_drop.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_drop(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_drop
}

var cache_Data_String_NonEmpty_CodePoints_countPrefix gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_countPrefix sync.Once
func Get_Data_String_NonEmpty_CodePoints_countPrefix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_countPrefix.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_countPrefix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_countPrefix(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodePoints_countPrefix
}

var cache_Data_String_NonEmpty_CodePoints_cons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_cons sync.Once
func Get_Data_String_NonEmpty_CodePoints_cons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_cons.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_cons(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_cons
}

var cache_Data_String_NonEmpty_CodePoints_codePointAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_codePointAt sync.Once
func Get_Data_String_NonEmpty_CodePoints_codePointAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_codePointAt.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_codePointAt = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodePoints_codePointAt(x_0_box.IntVal, v_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_codePointAt
}

func Call_Data_String_NonEmpty_CodePoints_toNonEmptyString(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_NonEmpty_CodePoints_snoc(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (s_1) + (gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(c_0)).StrVal())
}

func Call_Data_String_NonEmpty_CodePoints_singleton(x_0_loop int64) string {
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(x_0)).StrVal()).StrVal()
}

func Call_Data_String_NonEmpty_CodePoints_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_CodePoints_liftS__1645621157(f_0_loop gopurs_runtime.Value, v_1_loop string) string {
liftS__1645621157:
for {
if false { continue liftS__1645621157 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).StrVal()
}
}

func Call_Data_String_NonEmpty_CodePoints_takeWhile(f_0_loop gopurs_runtime.Value, x_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=String
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodePoints_countPrefix(), f_0, gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())).IntVal), gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())).StrVal()
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_2_0)}))}
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_lastIndexOf_prime_(pat_0_loop string, x_1_loop int64, v_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_1170268447_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_String_CodePoints_lastIndexOf_prime_(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_lastIndexOf(x_0_loop string, v_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_1170268447_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_String_CodePoints_lastIndexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_indexOf_prime_(pat_0_loop string, x_1_loop int64, v_2_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_1170268447_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(Get_Data_String_CodePoints_indexOf_prime_(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_indexOf(x_0_loop string, v_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_1170268447_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_String_CodePoints_indexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_fromNonEmptyString(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return v_0
}

func Call_Data_String_NonEmpty_CodePoints_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(x_0))))).IntVal
}

func Call_Data_String_NonEmpty_CodePoints_splitAt(i_0_loop int64, nes_1_loop string) struct{
	after *Constructor_Data_Maybe_Just[string]
	before *Constructor_Data_Maybe_Just[string]
} {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
// TAST (Let): v_2_0 shape=App(Var) bindingType=(Record (Row [before: String, after: String] Any))
v_2_0 := func() struct{
	after string
	before string
} {
					orig := gopurs_runtime.Apply2(Get_Data_String_CodePoints_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
					_ = orig
					clone := struct{
	after string
	before string
}{}
					clone.after = gopurs_runtime.RecordGet(orig, "after").StrVal()
					clone.before = gopurs_runtime.RecordGet(orig, "before").StrVal()
					return clone
				}()
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (v_2_0.after) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(v_2_0.after)}))}
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (v_2_0.before) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(v_2_0.before)}))}
}
end_branch_2:
return struct{
	after *Constructor_Data_Maybe_Just[string]
	before *Constructor_Data_Maybe_Just[string]
}{Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1)), Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))}
}

func Call_Data_String_NonEmpty_CodePoints_take(i_0_loop int64, nes_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0) < (int64(1)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_toCodePointArray(x_0_loop string) []int64 {
var x_0 string = x_0_loop
_ = x_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(x_0)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray(x_0_loop string) []int64 {
var x_0 string = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(TypeVar d)
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal()))
_ = __local_var_1_0
var __t1 []int64
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_1_0))).IntVal) > (int64(0)) {
__t1 = func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
goto end_branch_1
} else {

}
}
{
__t1 = func() []int64 { panic("Failed pattern match") }()
}
end_branch_1:
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := __t1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_Data_String_NonEmpty_CodePoints_uncons(nes_0_loop string) struct{
	head int64
	tail *Constructor_Data_Maybe_Just[string]
} {
var nes_0 string = nes_0_loop
_ = nes_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply2(Get_Data_String_CodePoints_codePointAt(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Str(nes_0))
_ = __local_var_1_0
var __t1 int64
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 930809136 && __local_var_1_0.UnsafePtr != nil) {
__t1 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0.IntVal
goto end_branch_1
} else {

}
}
{
__t1 = func() int64 { panic("Failed pattern match") }()
}
end_branch_1:
// TAST (Let): __local_var_1_2 shape=App(Var) bindingType=Any
__local_var_1_2 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Str(nes_0)).StrVal())).IntVal), gopurs_runtime.Str(nes_0))
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (__local_var_1_2.StrVal()) == ("") {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_1_2.StrVal())}))}
}
end_branch_3:
return struct{
	head int64
	tail *Constructor_Data_Maybe_Just[string]
}{__t1, Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))}
}

func Call_Data_String_NonEmpty_CodePoints_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_443971153_4179793454(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[string]](Get_Data_Semigroup_semigroupString())))}, Get_Data_String_NonEmpty_CodePoints_singleton())
}

func Call_Data_String_NonEmpty_CodePoints_fromCodePointArray(v_0_loop []int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 []int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (int64(0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCodePointArray(), func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray(x_0_loop []int64) string {
var x_0 []int64 = x_0_loop
_ = x_0
var __t0 string
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) == (int64(0)) {
__t0 = func() string { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCodePointArray(), func() gopurs_runtime.Value {
					arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal()).StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0).StrVal()
}

func Call_Data_String_NonEmpty_CodePoints_dropWhile(f_0_loop gopurs_runtime.Value, x_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=String
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodePoints_dropWhile(), f_0, gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())).StrVal()
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_2_0)}))}
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_drop(i_0_loop int64, nes_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0) >= (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(nes_1))))).IntVal) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())).IntVal), gopurs_runtime.Str(nes_1)).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodePoints_countPrefix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [String] Int)
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodePoints_countPrefix(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_CodePoints_cons(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(c_0)).StrVal()) + (s_1)
}

func Call_Data_String_NonEmpty_CodePoints_codePointAt(x_0_loop int64, v_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodePoints_1170268447_3094389156(Rebox_Data_String_NonEmpty_CodePoints_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_String_CodePoints_codePointAt(), gopurs_runtime.Int(x_0), gopurs_runtime.Str(v_1))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Rebox_Data_String_NonEmpty_CodePoints_1170268447_3094389156(in *Constructor_Data_Maybe_Just[int64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
	return out
}

func Rebox_Data_String_NonEmpty_CodePoints_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_String_NonEmpty_CodePoints_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_String_NonEmpty_CodePoints_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_NonEmpty_CodePoints_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}


