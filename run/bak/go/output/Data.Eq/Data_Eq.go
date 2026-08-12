package Data_Eq

import (
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Symbol "gopurs/output/Data.Symbol"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_eqVoid gopurs_runtime.Value
var once_eqVoid sync.Once
func Get_eqVoid() gopurs_runtime.Value {
	once_eqVoid.Do(func() {
		cache_eqVoid = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_eqVoid
}

var cache_eqUnit gopurs_runtime.Value
var once_eqUnit sync.Once
func Get_eqUnit() gopurs_runtime.Value {
	once_eqUnit.Do(func() {
		cache_eqUnit = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_eqUnit
}

var cache_eqString gopurs_runtime.Value
var once_eqString sync.Once
func Get_eqString() gopurs_runtime.Value {
	once_eqString.Do(func() {
		cache_eqString = gopurs_runtime.RecordDict1("eq", Get_eqStringImpl())
	})
	return cache_eqString
}

var cache_eqRowNil gopurs_runtime.Value
var once_eqRowNil sync.Once
func Get_eqRowNil() gopurs_runtime.Value {
	once_eqRowNil.Do(func() {
		cache_eqRowNil = gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})
}))
	})
	return cache_eqRowNil
}

var cache_eqRecord gopurs_runtime.Value
var once_eqRecord sync.Once
func Get_eqRecord() gopurs_runtime.Value {
	once_eqRecord.Do(func() {
		cache_eqRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqRecord(gopurs_runtime.CoerceToStruct[Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eqRecord
}

var cache_eqRec gopurs_runtime.Value
var once_eqRec sync.Once
func Get_eqRec() gopurs_runtime.Value {
	once_eqRec.Do(func() {
		cache_eqRec = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictEqRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqRec(_dollar__unused_0_box, dictEqRecord_1_box)
})
	})
	return cache_eqRec
}

var cache_eqProxy gopurs_runtime.Value
var once_eqProxy sync.Once
func Get_eqProxy() gopurs_runtime.Value {
	once_eqProxy.Do(func() {
		cache_eqProxy = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_eqProxy
}

var cache_eqNumber gopurs_runtime.Value
var once_eqNumber sync.Once
func Get_eqNumber() gopurs_runtime.Value {
	once_eqNumber.Do(func() {
		cache_eqNumber = gopurs_runtime.RecordDict1("eq", Get_eqNumberImpl())
	})
	return cache_eqNumber
}

var cache_eqInt gopurs_runtime.Value
var once_eqInt sync.Once
func Get_eqInt() gopurs_runtime.Value {
	once_eqInt.Do(func() {
		cache_eqInt = gopurs_runtime.RecordDict1("eq", Get_eqIntImpl())
	})
	return cache_eqInt
}

var cache_eqChar gopurs_runtime.Value
var once_eqChar sync.Once
func Get_eqChar() gopurs_runtime.Value {
	once_eqChar.Do(func() {
		cache_eqChar = gopurs_runtime.RecordDict1("eq", Get_eqCharImpl())
	})
	return cache_eqChar
}

var cache_eqBoolean gopurs_runtime.Value
var once_eqBoolean sync.Once
func Get_eqBoolean() gopurs_runtime.Value {
	once_eqBoolean.Do(func() {
		cache_eqBoolean = gopurs_runtime.RecordDict1("eq", Get_eqBooleanImpl())
	})
	return cache_eqBoolean
}

var cache_eq1 gopurs_runtime.Value
var once_eq1 sync.Once
func Get_eq1() gopurs_runtime.Value {
	once_eq1.Do(func() {
		cache_eq1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1(gopurs_runtime.CoerceToStruct[Constructor_Eq1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq1
}

var cache_eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		cache_eq = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq(gopurs_runtime.CoerceToStruct[Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq
}

var cache_eqArray gopurs_runtime.Value
var once_eqArray sync.Once
func Get_eqArray() gopurs_runtime.Value {
	once_eqArray.Do(func() {
		cache_eqArray = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqArray(dictEq_0_box)
})
	})
	return cache_eqArray
}

var cache_eq1Array gopurs_runtime.Value
var once_eq1Array sync.Once
func Get_eq1Array() gopurs_runtime.Value {
	once_eq1Array.Do(func() {
		cache_eq1Array = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}))
	})
	return cache_eq1Array
}

var cache_eqRowCons gopurs_runtime.Value
var once_eqRowCons sync.Once
func Get_eqRowCons() gopurs_runtime.Value {
	once_eqRowCons.Do(func() {
		cache_eqRowCons = gopurs_runtime.Func4(func(dictEqRecord_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictIsSymbol_2_box gopurs_runtime.Value, dictEq_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqRowCons(dictEqRecord_0_box, _dollar__unused_1_box, dictIsSymbol_2_box, dictEq_3_box)
})
	})
	return cache_eqRowCons
}

var cache_notEq gopurs_runtime.Value
var once_notEq sync.Once
func Get_notEq() gopurs_runtime.Value {
	once_notEq.Do(func() {
		cache_notEq = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq(gopurs_runtime.CoerceToStruct[Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_notEq
}

var cache_notEq1 gopurs_runtime.Value
var once_notEq1 sync.Once
func Get_notEq1() gopurs_runtime.Value {
	once_notEq1.Do(func() {
		cache_notEq1 = gopurs_runtime.Func4(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq1(gopurs_runtime.CoerceToStruct[Constructor_Eq1[gopurs_runtime.Value]](dictEq1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Eq[gopurs_runtime.Value]](dictEq_1_box), x_2_box, y_3_box))
})
	})
	return cache_notEq1
}

var cache_eq__2276491096 gopurs_runtime.Value
var once_eq__2276491096 sync.Once
func Get_eq__2276491096() gopurs_runtime.Value {
	once_eq__2276491096.Do(func() {
		cache_eq__2276491096 = gopurs_runtime.RecordGet(Get_eqBoolean(), "eq")
	})
	return cache_eq__2276491096
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq1__1773593252 gopurs_runtime.Value
var once_eq1__1773593252 sync.Once
func Get_eq1__1773593252() gopurs_runtime.Value {
	once_eq1__1773593252.Do(func() {
		cache_eq1__1773593252 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__1773593252(gopurs_runtime.CoerceToStruct[Constructor_Eq1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq1__1773593252
}

var cache_eqRecord__1610867122 gopurs_runtime.Value
var once_eqRecord__1610867122 sync.Once
func Get_eqRecord__1610867122() gopurs_runtime.Value {
	once_eqRecord__1610867122.Do(func() {
		cache_eqRecord__1610867122 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqRecord__1610867122(gopurs_runtime.CoerceToStruct[Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eqRecord__1610867122
}

var cache_eqRecord__1747372340 gopurs_runtime.Value
var once_eqRecord__1747372340 sync.Once
func Get_eqRecord__1747372340() gopurs_runtime.Value {
	once_eqRecord__1747372340.Do(func() {
		cache_eqRecord__1747372340 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqRecord__1747372340(gopurs_runtime.CoerceToStruct[Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eqRecord__1747372340
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj")
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_reflectSymbol__3416619207 gopurs_runtime.Value
var once_reflectSymbol__3416619207 sync.Once
func Get_reflectSymbol__3416619207() gopurs_runtime.Value {
	once_reflectSymbol__3416619207.Do(func() {
		cache_reflectSymbol__3416619207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__3416619207(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__3416619207
}

var cache_reflectSymbol__1166932993 gopurs_runtime.Value
var once_reflectSymbol__1166932993 sync.Once
func Get_reflectSymbol__1166932993() gopurs_runtime.Value {
	once_reflectSymbol__1166932993.Do(func() {
		cache_reflectSymbol__1166932993 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__1166932993(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__1166932993
}

type Constructor_EqRecord[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1311326743] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "eqRecord": return c.V0
		default: panic("Key not found in dictionary Constructor_EqRecord: " + key)
		}
	}
}


type Constructor_Eq[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1012063514] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Eq[gopurs_runtime.Value])(ptr)
		switch key {
		case "eq": return c.V0
		default: panic("Key not found in dictionary Constructor_Eq: " + key)
		}
	}
}


type Constructor_Eq1[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1715248107] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Eq1[gopurs_runtime.Value])(ptr)
		switch key {
		case "eq1": return c.V0
		default: panic("Key not found in dictionary Constructor_Eq1: " + key)
		}
	}
}


func Call_eqRecord(dict_0_loop *Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eqRec(_dollar__unused_0_loop gopurs_runtime.Value, dictEqRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictEqRecord_1 gopurs_runtime.Value = dictEqRecord_1_loop
_ = dictEqRecord_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEqRecord_1, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_eq1(dict_0_loop *Constructor_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq(dict_0_loop *Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eqArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq")))
}

func Call_eqRowCons(dictEqRecord_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictIsSymbol_2_loop gopurs_runtime.Value, dictEq_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEqRecord_0 gopurs_runtime.Value = dictEqRecord_0_loop
_ = dictEqRecord_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictIsSymbol_2 gopurs_runtime.Value = dictIsSymbol_2_loop
_ = dictIsSymbol_2
var dictEq_3 gopurs_runtime.Value = dictEq_3_loop
_ = dictEq_3
return gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
get_7_0 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_2, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = get_7_0
return gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_conj__3676519832(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_3, "eq"), gopurs_runtime.Apply(get_7_0, ra_5), gopurs_runtime.Apply(get_7_0, rb_6)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEqRecord_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6)).IntVal) != (0))
})
})
}))
}

func Call_notEq(dictEq_0_loop *Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return (gopurs_runtime.Apply2(Get_eq__2276491096(), gopurs_runtime.Apply2(dictEq_0.V0, x_1, y_2), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_notEq1(dictEq1_0_loop *Constructor_Eq1[gopurs_runtime.Value], dictEq_1_loop *Constructor_Eq[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var dictEq1_0 *Constructor_Eq1[gopurs_runtime.Value] = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 *Constructor_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return (gopurs_runtime.Apply2(Get_eq__2276491096(), gopurs_runtime.Apply3(dictEq1_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_1)}, x_2, y_3), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_eq__2384498378(dict_0_loop *Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq1__1773593252(dict_0_loop *Constructor_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eqRecord__1610867122(dict_0_loop *Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eqRecord__1747372340(dict_0_loop *Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_reflectSymbol__3416619207(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__1166932993(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_eqArrayImpl() gopurs_runtime.Value {
	return _Gopurs_EqArrayImpl
}

func Get_eqBooleanImpl() gopurs_runtime.Value {
	return _Gopurs_EqBooleanImpl
}

func Get_eqCharImpl() gopurs_runtime.Value {
	return _Gopurs_EqCharImpl
}

func Get_eqIntImpl() gopurs_runtime.Value {
	return _Gopurs_EqIntImpl
}

func Get_eqNumberImpl() gopurs_runtime.Value {
	return _Gopurs_EqNumberImpl
}

func Get_eqStringImpl() gopurs_runtime.Value {
	return _Gopurs_EqStringImpl
}
