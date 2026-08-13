package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict sync.Once
func Get_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict.Do(func() {
		cache_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict(x_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict
}

var cache_Data_HeytingAlgebra_Generic_genericTT_prime gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericTT_prime sync.Once
func Get_Data_HeytingAlgebra_Generic_genericTT_prime() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericTT_prime.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericTT_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericTT_prime(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericTT_prime
}

var cache_Data_HeytingAlgebra_Generic_genericTT gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericTT sync.Once
func Get_Data_HeytingAlgebra_Generic_genericTT() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericTT.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericTT = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericTT(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericTT
}

var cache_Data_HeytingAlgebra_Generic_genericNot_prime gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericNot_prime sync.Once
func Get_Data_HeytingAlgebra_Generic_genericNot_prime() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericNot_prime.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericNot_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericNot_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericNot_prime
}

var cache_Data_HeytingAlgebra_Generic_genericNot gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericNot sync.Once
func Get_Data_HeytingAlgebra_Generic_genericNot() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericNot.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericNot = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericNot(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dictGenericHeytingAlgebra_1_box), x_2_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericNot
}

var cache_Data_HeytingAlgebra_Generic_genericImplies_prime gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericImplies_prime sync.Once
func Get_Data_HeytingAlgebra_Generic_genericImplies_prime() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericImplies_prime.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericImplies_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericImplies_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericImplies_prime
}

var cache_Data_HeytingAlgebra_Generic_genericImplies gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericImplies sync.Once
func Get_Data_HeytingAlgebra_Generic_genericImplies() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericImplies.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericImplies = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericImplies(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericImplies
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments = gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
}), gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}})
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraNoArguments
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument
}

var cache_Data_HeytingAlgebra_Generic_genericFF_prime gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericFF_prime sync.Once
func Get_Data_HeytingAlgebra_Generic_genericFF_prime() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericFF_prime.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericFF_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericFF_prime(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericFF_prime
}

var cache_Data_HeytingAlgebra_Generic_genericFF gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericFF sync.Once
func Get_Data_HeytingAlgebra_Generic_genericFF() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericFF.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericFF = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericFF(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericFF
}

var cache_Data_HeytingAlgebra_Generic_genericDisj_prime gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericDisj_prime sync.Once
func Get_Data_HeytingAlgebra_Generic_genericDisj_prime() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericDisj_prime.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericDisj_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericDisj_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericDisj_prime
}

var cache_Data_HeytingAlgebra_Generic_genericDisj gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericDisj sync.Once
func Get_Data_HeytingAlgebra_Generic_genericDisj() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericDisj.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericDisj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericDisj(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericDisj
}

var cache_Data_HeytingAlgebra_Generic_genericConj_prime gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericConj_prime sync.Once
func Get_Data_HeytingAlgebra_Generic_genericConj_prime() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericConj_prime.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericConj_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericConj_prime(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericConj_prime
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor(dictGenericHeytingAlgebra_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor
}

var cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct sync.Once
func Get_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct = gopurs_runtime.Func2(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_box, dictGenericHeytingAlgebra1_1_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct
}

var cache_Data_HeytingAlgebra_Generic_genericConj gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericConj sync.Once
func Get_Data_HeytingAlgebra_Generic_genericConj() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericConj.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericConj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericConj(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericConj
}

var cache_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791 gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791 sync.Once
func Get_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791
}

var cache_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791 gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791 sync.Once
func Get_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791
}

var cache_Data_HeytingAlgebra_Generic_genericFF_prime__265522150 gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericFF_prime__265522150 sync.Once
func Get_Data_HeytingAlgebra_Generic_genericFF_prime__265522150() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericFF_prime__265522150.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericFF_prime__265522150 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericFF_prime__265522150(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericFF_prime__265522150
}

var cache_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791 gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791 sync.Once
func Get_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791
}

var cache_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174 gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174 sync.Once
func Get_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174
}

var cache_Data_HeytingAlgebra_Generic_genericTT_prime__265522150 gopurs_runtime.Value
var once_Data_HeytingAlgebra_Generic_genericTT_prime__265522150 sync.Once
func Get_Data_HeytingAlgebra_Generic_genericTT_prime__265522150() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_Generic_genericTT_prime__265522150.Do(func() {
		cache_Data_HeytingAlgebra_Generic_genericTT_prime__265522150 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_Generic_genericTT_prime__265522150(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_Generic_genericTT_prime__265522150
}

type Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2831861733] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra)(ptr)
		_ = c
		switch key {
		case "genericConj'": return gopurs_runtime.Box(c.V0)
		case "genericDisj'": return gopurs_runtime.Box(c.V1)
		case "genericFF'": return gopurs_runtime.Box(c.V2)
		case "genericImplies'": return gopurs_runtime.Box(c.V3)
		case "genericNot'": return gopurs_runtime.Box(c.V4)
		case "genericTT'": return gopurs_runtime.Box(c.V5)
		default: panic("Key not found in dictionary Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra: " + key)
		}
	}
}


func Call_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_HeytingAlgebra_Generic_genericTT_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTT'")
}

func Call_Data_HeytingAlgebra_Generic_genericTT(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericTT'"))
}

func Call_Data_HeytingAlgebra_Generic_genericNot_prime(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_Generic_genericNot(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply2(Get_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2)))
}

func Call_Data_HeytingAlgebra_Generic_genericImplies_prime(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_Generic_genericImplies(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply3(Get_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraArgument(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), v_1)
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")})
}

func Call_Data_HeytingAlgebra_Generic_genericFF_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericFF'")
}

func Call_Data_HeytingAlgebra_Generic_genericFF(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericFF'"))
}

func Call_Data_HeytingAlgebra_Generic_genericDisj_prime(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_Generic_genericDisj(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply3(Get_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_HeytingAlgebra_Generic_genericConj_prime(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraConstructor(dictGenericHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 gopurs_runtime.Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), v_1, v1_2)
})
}), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), v_1, v1_2)
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), v_1)
}), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'")})
}

func Call_Data_HeytingAlgebra_Generic_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 gopurs_runtime.Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
var dictGenericHeytingAlgebra1_1 gopurs_runtime.Value = dictGenericHeytingAlgebra1_1_loop
_ = dictGenericHeytingAlgebra1_1
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericConj'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericDisj'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericFF'")})}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericImplies'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1, (*Constructor_Data_Generic_Rep_Product)(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericNot'"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericTT'")})}})
}

func Call_Data_HeytingAlgebra_Generic_genericConj(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, dictGenericHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V1), gopurs_runtime.Apply3(Get_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), x_2), gopurs_runtime.Apply(gopurs_runtime.Box(dictGeneric_0.V0), y_3)))
}

func Call_Data_HeytingAlgebra_Generic_genericConj_prime__1368982791(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_Generic_genericDisj_prime__1368982791(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_Generic_genericFF_prime__265522150(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericFF'")
}

func Call_Data_HeytingAlgebra_Generic_genericImplies_prime__1368982791(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_Generic_genericNot_prime__4234376174(dict_0_loop *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_Generic_GenericHeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_Generic_genericTT_prime__265522150(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTT'")
}


