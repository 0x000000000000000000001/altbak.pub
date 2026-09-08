package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString = Get_Data_String_NonEmpty_Internal_NonEmptyString()
	})
	return cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString
}

var cache_Data_String_NonEmpty_CodeUnits_snoc gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_snoc sync.Once
func Get_Data_String_NonEmpty_CodeUnits_snoc() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_snoc.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_snoc(c_0_box.StrVal(), s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_snoc
}

var cache_Data_String_NonEmpty_CodeUnits_singleton gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_singleton sync.Once
func Get_Data_String_NonEmpty_CodeUnits_singleton() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_singleton.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_singleton(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_singleton
}

var cache_Data_String_NonEmpty_CodeUnits_liftS gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_liftS(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS
}

var cache_Data_String_NonEmpty_CodeUnits_liftS__1645621157 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS__1645621157 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS__1645621157() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS__1645621157.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS__1645621157 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_liftS__1645621157(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS__1645621157
}

var cache_Data_String_NonEmpty_CodeUnits_takeWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_takeWhile sync.Once
func Get_Data_String_NonEmpty_CodeUnits_takeWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_takeWhile.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_takeWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_takeWhile(f_0_box, x_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_takeWhile
}

var cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_ sync.Once
func Get_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_ = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_(pat_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_
}

var cache_Data_String_NonEmpty_CodeUnits_lastIndexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_lastIndexOf sync.Once
func Get_Data_String_NonEmpty_CodeUnits_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_lastIndexOf.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_lastIndexOf = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_lastIndexOf(x_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_lastIndexOf
}

var cache_Data_String_NonEmpty_CodeUnits_indexOf_prime_ gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_indexOf_prime_ sync.Once
func Get_Data_String_NonEmpty_CodeUnits_indexOf_prime_() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_indexOf_prime_.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_indexOf_prime_ = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_indexOf_prime_(pat_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_indexOf_prime_
}

var cache_Data_String_NonEmpty_CodeUnits_indexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_indexOf sync.Once
func Get_Data_String_NonEmpty_CodeUnits_indexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_indexOf.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_indexOf = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_indexOf(x_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_indexOf
}

var cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(v_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString
}

var cache_Data_String_NonEmpty_CodeUnits_length gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_length sync.Once
func Get_Data_String_NonEmpty_CodeUnits_length() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_length.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_NonEmpty_CodeUnits_length(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_length
}

var cache_Data_String_NonEmpty_CodeUnits_splitAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_splitAt sync.Once
func Get_Data_String_NonEmpty_CodeUnits_splitAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_splitAt.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodeUnits_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"after", "before"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(orig.after))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(orig.before))}})
				}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_splitAt
}

var cache_Data_String_NonEmpty_CodeUnits_take gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_take sync.Once
func Get_Data_String_NonEmpty_CodeUnits_take() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_take.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_take(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_take
}

var cache_Data_String_NonEmpty_CodeUnits_takeRight gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_takeRight sync.Once
func Get_Data_String_NonEmpty_CodeUnits_takeRight() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_takeRight.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_takeRight(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_takeRight
}

var cache_Data_String_NonEmpty_CodeUnits_toChar gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toChar sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toChar() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toChar.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toChar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_toChar(x_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_toChar
}

var cache_Data_String_NonEmpty_CodeUnits_toCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toCharArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_String_NonEmpty_CodeUnits_toCharArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_toCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_uncons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_uncons sync.Once
func Get_Data_String_NonEmpty_CodeUnits_uncons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_uncons.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_String_NonEmpty_CodeUnits_uncons(nes_0_box.StrVal())
				_ = orig
				return gopurs_runtime.RecordDict([]string{"head", "tail"}, []gopurs_runtime.Value{gopurs_runtime.Str(orig.head), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(orig.tail))}})
				}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_uncons
}

var cache_Data_String_NonEmpty_CodeUnits_fromFoldable1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromFoldable1 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromFoldable1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromFoldable1.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromFoldable1
}

var cache_Data_String_NonEmpty_CodeUnits_fromCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromCharArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_fromCharArray(func() []string {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray(func() []string {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_dropWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_dropWhile sync.Once
func Get_Data_String_NonEmpty_CodeUnits_dropWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_dropWhile.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_dropWhile = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_dropWhile(f_0_box, x_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_dropWhile
}

var cache_Data_String_NonEmpty_CodeUnits_dropRight gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_dropRight sync.Once
func Get_Data_String_NonEmpty_CodeUnits_dropRight() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_dropRight.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_dropRight(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_dropRight
}

var cache_Data_String_NonEmpty_CodeUnits_drop gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_drop sync.Once
func Get_Data_String_NonEmpty_CodeUnits_drop() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_drop.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_drop(i_0_box.IntVal, nes_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_drop
}

var cache_Data_String_NonEmpty_CodeUnits_countPrefix gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_countPrefix sync.Once
func Get_Data_String_NonEmpty_CodeUnits_countPrefix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_countPrefix.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_countPrefix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_countPrefix(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_countPrefix
}

var cache_Data_String_NonEmpty_CodeUnits_cons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_cons sync.Once
func Get_Data_String_NonEmpty_CodeUnits_cons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_cons.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_cons(c_0_box.StrVal(), s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_cons
}

var cache_Data_String_NonEmpty_CodeUnits_charAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_charAt sync.Once
func Get_Data_String_NonEmpty_CodeUnits_charAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_charAt.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_charAt = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_CodeUnits_charAt(x_0_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_charAt
}

func Call_Data_String_NonEmpty_CodeUnits_snoc(c_0_loop string, s_1_loop string) string {
var c_0 string = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (s_1) + (gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Str(c_0)).StrVal())
}

func Call_Data_String_NonEmpty_CodeUnits_singleton(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Str(x_0)).StrVal()).StrVal()
}

func Call_Data_String_NonEmpty_CodeUnits_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_CodeUnits_liftS__1645621157(f_0_loop gopurs_runtime.Value, v_1_loop string) string {
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

func Call_Data_String_NonEmpty_CodeUnits_takeWhile(f_0_loop gopurs_runtime.Value, x_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=String
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_countPrefix(), f_0, gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())).IntVal), gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())).StrVal()
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime_(pat_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var pat_0 string = pat_0_loop
_ = pat_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_lastIndexOf_prime_(), gopurs_runtime.Str(pat_0))
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Str(v_4.StrVal()))
})
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_lastIndexOf(x_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 string = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [String] (ADT ["Data","Maybe","Maybe"] [Int]))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_lastIndexOf(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_indexOf_prime_(pat_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var pat_0 string = pat_0_loop
_ = pat_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf_prime_(), gopurs_runtime.Str(pat_0))
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Str(v_4.StrVal()))
})
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_indexOf(x_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 string = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [String] (ADT ["Data","Maybe","Maybe"] [Int]))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return v_0
}

func Call_Data_String_NonEmpty_CodeUnits_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(x_0)).IntVal
}

func Call_Data_String_NonEmpty_CodeUnits_splitAt(i_0_loop int64, nes_1_loop string) struct{
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
					orig := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
}{Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1)), Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t2))}
}

func Call_Data_String_NonEmpty_CodeUnits_take(i_0_loop int64, nes_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0) < (int64(1)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_takeRight(i_0_loop int64, nes_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0) < (int64(1)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) - (i_0)), gopurs_runtime.Str(nes_1)).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_toChar(x_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 string = x_0_loop
_ = x_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_String_CodeUnits_toChar(), gopurs_runtime.Str(x_0))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_toCharArray(x_0_loop string) []string {
var x_0 string = x_0_loop
_ = x_0
return func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_String_CodeUnits_toCharArray(), gopurs_runtime.Str(x_0)).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
}

func Call_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray(x_0_loop string) []string {
var x_0 string = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(TypeVar d)
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_toCharArray(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal()))
_ = __local_var_1_0
var __t1 []string
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_1_0))).IntVal) > (int64(0)) {
__t1 = func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
goto end_branch_1
} else {

}
}
{
__t1 = func() []string { panic("Failed pattern match") }()
}
end_branch_1:
return func() []string {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := __t1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
}

func Call_Data_String_NonEmpty_CodeUnits_uncons(nes_0_loop string) struct{
	head string
	tail *Constructor_Data_Maybe_Just[string]
} {
var nes_0 string = nes_0_loop
_ = nes_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Str(nes_0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_1_0.StrVal())}))}
}
end_branch_1:
return struct{
	head string
	tail *Constructor_Data_Maybe_Just[string]
}{gopurs_runtime.Apply2(Get_Data_String_Unsafe_charAt(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Str(nes_0)).StrVal(), Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))}
}

func Call_Data_String_NonEmpty_CodeUnits_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_443971153_4179793454(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[string]](Get_Data_Semigroup_semigroupString())))}, Get_Data_String_NonEmpty_CodeUnits_singleton())
}

func Call_Data_String_NonEmpty_CodeUnits_fromCharArray(v_0_loop []string) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 []string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (int64(0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodeUnits_fromCharArray(), func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray(x_0_loop []string) string {
var x_0 []string = x_0_loop
_ = x_0
var __t0 string
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()))).IntVal) == (int64(0)) {
__t0 = func() string { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodeUnits_fromCharArray(), func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := x_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal()).StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0).StrVal()
}

func Call_Data_String_NonEmpty_CodeUnits_dropWhile(f_0_loop gopurs_runtime.Value, x_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 string = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=String
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_countPrefix(), f_0, gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())).IntVal), gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())).StrVal()
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_dropRight(i_0_loop int64, nes_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0) >= (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) - (i_0)), gopurs_runtime.Str(nes_1)).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_drop(i_0_loop int64, nes_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (i_0) >= (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[string]](func() gopurs_runtime.Value {
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())}))}
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_CodeUnits_countPrefix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [String] Int)
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_countPrefix(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_CodeUnits_cons(c_0_loop string, s_1_loop string) string {
var c_0 string = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Str(c_0)).StrVal()) + (s_1)
}

func Call_Data_String_NonEmpty_CodeUnits_charAt(x_0_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var x_0 int64 = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=(Func [String] (ADT ["Data","Maybe","Maybe"] [Char]))
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_charAt(), gopurs_runtime.Int(x_0))
_ = __local_var_1_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Rebox_Data_String_NonEmpty_CodeUnits_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_String_NonEmpty_CodeUnits_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_NonEmpty_CodeUnits_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}


