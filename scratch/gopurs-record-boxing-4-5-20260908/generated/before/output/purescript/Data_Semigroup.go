package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_SemigroupRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Semigroup_SemigroupRecord_dollar_Dict sync.Once
func Get_Data_Semigroup_SemigroupRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Semigroup_SemigroupRecord_dollar_Dict.Do(func() {
		cache_Data_Semigroup_SemigroupRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer(Call_Data_Semigroup_SemigroupRecord_dollar_Dict(func() struct{
	appendRecord gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	appendRecord gopurs_runtime.Value
}{}
					clone.appendRecord = gopurs_runtime.RecordGet(orig, "appendRecord")
					return clone
				}()))}
})
	})
	return cache_Data_Semigroup_SemigroupRecord_dollar_Dict
}

var cache_Data_Semigroup_Semigroup_dollar_Dict gopurs_runtime.Value
var once_Data_Semigroup_Semigroup_dollar_Dict sync.Once
func Get_Data_Semigroup_Semigroup_dollar_Dict() gopurs_runtime.Value {
	once_Data_Semigroup_Semigroup_dollar_Dict.Do(func() {
		cache_Data_Semigroup_Semigroup_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Call_Data_Semigroup_Semigroup_dollar_Dict(func() struct{
	go__append gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	go__append gopurs_runtime.Value
}{}
					clone.go__append = gopurs_runtime.RecordGet(orig, "append")
					return clone
				}()))}
})
	})
	return cache_Data_Semigroup_Semigroup_dollar_Dict
}

var cache_Data_Semigroup_semigroupVoid gopurs_runtime.Value
var once_Data_Semigroup_semigroupVoid sync.Once
func Get_Data_Semigroup_semigroupVoid() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupVoid.Do(func() {
		cache_Data_Semigroup_semigroupVoid = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Void_absurd()
})}))}
	})
	return cache_Data_Semigroup_semigroupVoid
}

var cache_Data_Semigroup_semigroupUnit gopurs_runtime.Value
var once_Data_Semigroup_semigroupUnit sync.Once
func Get_Data_Semigroup_semigroupUnit() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupUnit.Do(func() {
		cache_Data_Semigroup_semigroupUnit = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})}))}
	})
	return cache_Data_Semigroup_semigroupUnit
}

var cache_Data_Semigroup_semigroupString gopurs_runtime.Value
var once_Data_Semigroup_semigroupString sync.Once
func Get_Data_Semigroup_semigroupString() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupString.Do(func() {
		cache_Data_Semigroup_semigroupString = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_443971153_4179793454((&Constructor_Data_Semigroup_Semigroup[string]{1, Get_Data_Semigroup_concatString()})))}
	})
	return cache_Data_Semigroup_semigroupString
}

var cache_Data_Semigroup_semigroupRecordNil gopurs_runtime.Value
var once_Data_Semigroup_semigroupRecordNil sync.Once
func Get_Data_Semigroup_semigroupRecordNil() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupRecordNil.Do(func() {
		cache_Data_Semigroup_semigroupRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}()
})}))}
	})
	return cache_Data_Semigroup_semigroupRecordNil
}

var cache_Data_Semigroup_semigroupProxy gopurs_runtime.Value
var once_Data_Semigroup_semigroupProxy sync.Once
func Get_Data_Semigroup_semigroupProxy() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupProxy.Do(func() {
		cache_Data_Semigroup_semigroupProxy = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_1625289059_4179793454((&Constructor_Data_Semigroup_Semigroup[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Semigroup_semigroupProxy
}

var cache_Data_Semigroup_semigroupArray gopurs_runtime.Value
var once_Data_Semigroup_semigroupArray sync.Once
func Get_Data_Semigroup_semigroupArray() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupArray.Do(func() {
		cache_Data_Semigroup_semigroupArray = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_986430664_4179793454((&Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value]{1, Get_Data_Semigroup_concatArray()})))}
	})
	return cache_Data_Semigroup_semigroupArray
}

var cache_Data_Semigroup_appendRecord gopurs_runtime.Value
var once_Data_Semigroup_appendRecord sync.Once
func Get_Data_Semigroup_appendRecord() gopurs_runtime.Value {
	once_Data_Semigroup_appendRecord.Do(func() {
		cache_Data_Semigroup_appendRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_appendRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semigroup_appendRecord
}

var cache_Data_Semigroup_semigroupRecord gopurs_runtime.Value
var once_Data_Semigroup_semigroupRecord sync.Once
func Get_Data_Semigroup_semigroupRecord() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupRecord.Do(func() {
		cache_Data_Semigroup_semigroupRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictSemigroupRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_semigroupRecord(_dollar___unused_0_box, dictSemigroupRecord_1_box)
})
	})
	return cache_Data_Semigroup_semigroupRecord
}

var cache_Data_Semigroup_go__append gopurs_runtime.Value
var once_Data_Semigroup_go__append sync.Once
func Get_Data_Semigroup_go__append() gopurs_runtime.Value {
	once_Data_Semigroup_go__append.Do(func() {
		cache_Data_Semigroup_go__append = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_go__append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semigroup_go__append
}

var cache_Data_Semigroup_append__2496145349 gopurs_runtime.Value
var once_Data_Semigroup_append__2496145349 sync.Once
func Get_Data_Semigroup_append__2496145349() gopurs_runtime.Value {
	once_Data_Semigroup_append__2496145349.Do(func() {
		cache_Data_Semigroup_append__2496145349 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2496145349(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Semigroup_append__2496145349
}

var cache_Data_Semigroup_append__142987400 gopurs_runtime.Value
var once_Data_Semigroup_append__142987400 sync.Once
func Get_Data_Semigroup_append__142987400() gopurs_runtime.Value {
	once_Data_Semigroup_append__142987400.Do(func() {
		cache_Data_Semigroup_append__142987400 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__142987400(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Semigroup_append__142987400
}

var cache_Data_Semigroup_append__294757560 gopurs_runtime.Value
var once_Data_Semigroup_append__294757560 sync.Once
func Get_Data_Semigroup_append__294757560() gopurs_runtime.Value {
	once_Data_Semigroup_append__294757560.Do(func() {
		cache_Data_Semigroup_append__294757560 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__294757560(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Semigroup_append__294757560
}

var cache_Data_Semigroup_append__278170894 gopurs_runtime.Value
var once_Data_Semigroup_append__278170894 sync.Once
func Get_Data_Semigroup_append__278170894() gopurs_runtime.Value {
	once_Data_Semigroup_append__278170894.Do(func() {
		cache_Data_Semigroup_append__278170894 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__278170894(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Semigroup_append__278170894
}

var cache_Data_Semigroup_semigroupFn gopurs_runtime.Value
var once_Data_Semigroup_semigroupFn sync.Once
func Get_Data_Semigroup_semigroupFn() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupFn.Do(func() {
		cache_Data_Semigroup_semigroupFn = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_semigroupFn(dictSemigroup_0_box)
})
	})
	return cache_Data_Semigroup_semigroupFn
}

var cache_Data_Semigroup_semigroupRecordCons gopurs_runtime.Value
var once_Data_Semigroup_semigroupRecordCons sync.Once
func Get_Data_Semigroup_semigroupRecordCons() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupRecordCons.Do(func() {
		cache_Data_Semigroup_semigroupRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictSemigroupRecord_2_box gopurs_runtime.Value, dictSemigroup_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_semigroupRecordCons(dictIsSymbol_0_box, _dollar___unused_1_box, dictSemigroupRecord_2_box, dictSemigroup_3_box)
})
	})
	return cache_Data_Semigroup_semigroupRecordCons
}

type Constructor_Data_Semigroup_SemigroupRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3847494007] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_SemigroupRecord[any, any, any])(ptr)
		_ = c
		switch key {
		case "appendRecord": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Semigroup_SemigroupRecord: " + key)
		}
	}
}


type Constructor_Data_Semigroup_Semigroup[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2053112122] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Semigroup[any])(ptr)
		_ = c
		switch key {
		case "append": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Semigroup_Semigroup: " + key)
		}
	}
}


func Call_Data_Semigroup_SemigroupRecord_dollar_Dict(x_0_loop struct{
	appendRecord gopurs_runtime.Value
}) *Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	appendRecord gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"appendRecord"}, []gopurs_runtime.Value{orig.appendRecord})
				}())
}

func Call_Data_Semigroup_Semigroup_dollar_Dict(x_0_loop struct{
	go__append gopurs_runtime.Value
}) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
var x_0 struct{
	go__append gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"append"}, []gopurs_runtime.Value{orig.go__append})
				}())
}

func Call_Data_Semigroup_appendRecord(dict_0_loop *Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_semigroupRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictSemigroupRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictSemigroupRecord_1 gopurs_runtime.Value = dictSemigroupRecord_1_loop
_ = dictSemigroupRecord_1
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroupRecord_1, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Semigroup_go__append(dict_0_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__2496145349(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
append__2496145349:
for {
if false { continue append__2496145349 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t3 uint32
{
var __t_tag_0 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_0) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
var __t_tag_1 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_1) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
var __t_tag_2 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_2) == 902936544) {
__t3 = uint32(__eta_norm_0_1.IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
}
}

func Call_Data_Semigroup_append__142987400(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
append__142987400:
for {
if false { continue append__142987400 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), __eta_norm_1_0, __eta_norm_0_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}
}

func Call_Data_Semigroup_append__294757560(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
append__294757560:
for {
if false { continue append__294757560 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Float((__eta_norm_1_0.FloatVal()) + (__eta_norm_0_1.FloatVal()))
}
}

func Call_Data_Semigroup_append__278170894(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
append__278170894:
for {
if false { continue append__278170894 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Str((__eta_norm_1_0.StrVal()) + (__eta_norm_0_1.StrVal()))
}
}

func Call_Data_Semigroup_semigroupFn(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})}))}
}

func Call_Data_Semigroup_semigroupRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictSemigroupRecord_2_loop gopurs_runtime.Value, dictSemigroup_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
_ = _dollar___unused_1
var dictSemigroupRecord_2 gopurs_runtime.Value = dictSemigroupRecord_2_loop
_ = dictSemigroupRecord_2
var dictSemigroup_3 gopurs_runtime.Value = dictSemigroup_3_loop
_ = dictSemigroup_3
return gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, ra_5 gopurs_runtime.Value, rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_7_0 shape=App(Other) bindingType=String
key_7_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
_ = key_7_0
// TAST (Let): get__3905987523_8_1 shape=App(Var) bindingType=(Func [(Record (Row [] (TypeVar row)))] (TypeVar focus))
get__3905987523_8_1 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_7_0))
_ = get__3905987523_8_1
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_7_0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_3, "append"), gopurs_runtime.Apply(get__3905987523_8_1, ra_5), gopurs_runtime.Apply(get__3905987523_8_1, rb_6)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemigroupRecord_2, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6))
})}))}
}

func Rebox_Data_Semigroup_1625289059_4179793454(in *Constructor_Data_Semigroup_Semigroup[uint32]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Semigroup_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Semigroup_986430664_4179793454(in *Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Get_Data_Semigroup_concatArray() gopurs_runtime.Value {
	return _Gopurs_Data_Semigroup_ConcatArray
}

func Get_Data_Semigroup_concatString() gopurs_runtime.Value {
	return _Gopurs_Data_Semigroup_ConcatString
}
