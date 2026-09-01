package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Test_RowToList_RecordKeys_dollar_Dict gopurs_runtime.Value
var once_Test_RowToList_RecordKeys_dollar_Dict sync.Once

func Get_Test_RowToList_RecordKeys_dollar_Dict() gopurs_runtime.Value {
	once_Test_RowToList_RecordKeys_dollar_Dict.Do(func() {
		cache_Test_RowToList_RecordKeys_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_RowToList_RecordKeys_dollar_Dict(x_0_box)
		})
	})
	return cache_Test_RowToList_RecordKeys_dollar_Dict
}

var cache_Test_RowToList_keysNil gopurs_runtime.Value
var once_Test_RowToList_keysNil sync.Once

func Get_Test_RowToList_keysNil() gopurs_runtime.Value {
	once_Test_RowToList_keysNil.Do(func() {
		cache_Test_RowToList_keysNil = gopurs_runtime.Value{Type: 9, IntVal: 561174694, UnsafePtr: unsafe.Pointer((&Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(0)
		})}))}
	})
	return cache_Test_RowToList_keysNil
}

var cache_Test_RowToList_keysImpl gopurs_runtime.Value
var once_Test_RowToList_keysImpl sync.Once

func Get_Test_RowToList_keysImpl() gopurs_runtime.Value {
	once_Test_RowToList_keysImpl.Do(func() {
		cache_Test_RowToList_keysImpl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_RowToList_keysImpl(gopurs_runtime.CoerceToStruct[Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value]](dict_0_box))
		})
	})
	return cache_Test_RowToList_keysImpl
}

var cache_Test_RowToList_keysCons gopurs_runtime.Value
var once_Test_RowToList_keysCons sync.Once

func Get_Test_RowToList_keysCons() gopurs_runtime.Value {
	once_Test_RowToList_keysCons.Do(func() {
		cache_Test_RowToList_keysCons = gopurs_runtime.Func(func(dictRecordKeys_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_RowToList_keysCons(dictRecordKeys_0_box)
		})
	})
	return cache_Test_RowToList_keysCons
}

var cache_Test_RowToList_keysCons1 gopurs_runtime.Value
var once_Test_RowToList_keysCons1 sync.Once

func Get_Test_RowToList_keysCons1() gopurs_runtime.Value {
	once_Test_RowToList_keysCons1.Do(func() {
		cache_Test_RowToList_keysCons1 = gopurs_runtime.Value{Type: 9, IntVal: 561174694, UnsafePtr: unsafe.Pointer((&Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(5)
		})}))}
	})
	return cache_Test_RowToList_keysCons1
}

var cache_Test_RowToList_keys gopurs_runtime.Value
var once_Test_RowToList_keys sync.Once

func Get_Test_RowToList_keys() gopurs_runtime.Value {
	once_Test_RowToList_keys.Do(func() {
		cache_Test_RowToList_keys = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, dictRecordKeys_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_RowToList_keys(_dollar___unused_0_box, gopurs_runtime.CoerceToStruct[Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value]](dictRecordKeys_1_box), v_2_box))
		})
	})
	return cache_Test_RowToList_keys
}

var cache_Test_RowToList_describe gopurs_runtime.Value
var once_Test_RowToList_describe sync.Once

func Get_Test_RowToList_describe() gopurs_runtime.Value {
	once_Test_RowToList_describe.Do(func() {
		cache_Test_RowToList_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("RowToList (Keys Count):"))
	})
	return cache_Test_RowToList_describe
}

var cache_Test_RowToList_act gopurs_runtime.Value
var once_Test_RowToList_act sync.Once

func Get_Test_RowToList_act() gopurs_runtime.Value {
	once_Test_RowToList_act.Do(func() {
		cache_Test_RowToList_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10000))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(5)).StrVal())
		})
	})
	return cache_Test_RowToList_act
}

type Constructor_Test_RowToList_RecordKeys[T_rl any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}

func init() {
	gopurs_runtime.StructGetters[561174694] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Test_RowToList_RecordKeys[any])(ptr)
		_ = c
		switch key {
		case "keysImpl":
			return gopurs_runtime.Box(c.V0)
		default:
			panic("Key not found in dictionary Constructor_Test_RowToList_RecordKeys: " + key)
		}
	}
}

func Call_Test_RowToList_RecordKeys_dollar_Dict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Test_RowToList_keysImpl(dict_0_loop *Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value]) gopurs_runtime.Value {
	var dict_0 *Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value] = dict_0_loop
	_ = dict_0
	return gopurs_runtime.Box(dict_0.V0)
}

func Call_Test_RowToList_keysCons(dictRecordKeys_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var dictRecordKeys_0 gopurs_runtime.Value = dictRecordKeys_0_loop
	_ = dictRecordKeys_0
	return gopurs_runtime.Value{Type: 9, IntVal: 561174694, UnsafePtr: unsafe.Pointer((&Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Int((1) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRecordKeys_0, "keysImpl"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).IntVal))
	})}))}
}

func Call_Test_RowToList_keys(_dollar___unused_0_loop gopurs_runtime.Value, dictRecordKeys_1_loop *Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value], v_2_loop gopurs_runtime.Value) int64 {
	var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
	_ = _dollar___unused_0
	var dictRecordKeys_1 *Constructor_Test_RowToList_RecordKeys[gopurs_runtime.Value] = dictRecordKeys_1_loop
	_ = dictRecordKeys_1
	var v_2 gopurs_runtime.Value = v_2_loop
	_ = v_2
	return gopurs_runtime.Apply(gopurs_runtime.Box(dictRecordKeys_1.V0), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).IntVal
}
