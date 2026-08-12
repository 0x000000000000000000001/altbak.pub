package Data_HeytingAlgebra_Generic

import (
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genericTT_prime gopurs_runtime.Value
var once_genericTT_prime sync.Once
func Get_genericTT_prime() gopurs_runtime.Value {
	once_genericTT_prime.Do(func() {
		cache_genericTT_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTT_prime(dict_0_box)
})
	})
	return cache_genericTT_prime
}

var cache_genericTT gopurs_runtime.Value
var once_genericTT sync.Once
func Get_genericTT() gopurs_runtime.Value {
	once_genericTT.Do(func() {
		cache_genericTT = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTT(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box)
})
	})
	return cache_genericTT
}

var cache_genericNot_prime gopurs_runtime.Value
var once_genericNot_prime sync.Once
func Get_genericNot_prime() gopurs_runtime.Value {
	once_genericNot_prime.Do(func() {
		cache_genericNot_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericNot_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericNot_prime
}

var cache_genericNot gopurs_runtime.Value
var once_genericNot sync.Once
func Get_genericNot() gopurs_runtime.Value {
	once_genericNot.Do(func() {
		cache_genericNot = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericNot(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box)
})
	})
	return cache_genericNot
}

var cache_genericImplies_prime gopurs_runtime.Value
var once_genericImplies_prime sync.Once
func Get_genericImplies_prime() gopurs_runtime.Value {
	once_genericImplies_prime.Do(func() {
		cache_genericImplies_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericImplies_prime
}

var cache_genericImplies gopurs_runtime.Value
var once_genericImplies sync.Once
func Get_genericImplies() gopurs_runtime.Value {
	once_genericImplies.Do(func() {
		cache_genericImplies = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_genericImplies
}

var cache_genericHeytingAlgebraNoArguments gopurs_runtime.Value
var once_genericHeytingAlgebraNoArguments sync.Once
func Get_genericHeytingAlgebraNoArguments() gopurs_runtime.Value {
	once_genericHeytingAlgebraNoArguments.Do(func() {
		cache_genericHeytingAlgebraNoArguments = gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_genericHeytingAlgebraNoArguments
}

var cache_genericHeytingAlgebraArgument gopurs_runtime.Value
var once_genericHeytingAlgebraArgument sync.Once
func Get_genericHeytingAlgebraArgument() gopurs_runtime.Value {
	once_genericHeytingAlgebraArgument.Do(func() {
		cache_genericHeytingAlgebraArgument = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericHeytingAlgebraArgument(dictHeytingAlgebra_0_box)
})
	})
	return cache_genericHeytingAlgebraArgument
}

var cache_genericFF_prime gopurs_runtime.Value
var once_genericFF_prime sync.Once
func Get_genericFF_prime() gopurs_runtime.Value {
	once_genericFF_prime.Do(func() {
		cache_genericFF_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFF_prime(dict_0_box)
})
	})
	return cache_genericFF_prime
}

var cache_genericFF gopurs_runtime.Value
var once_genericFF sync.Once
func Get_genericFF() gopurs_runtime.Value {
	once_genericFF.Do(func() {
		cache_genericFF = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFF(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box)
})
	})
	return cache_genericFF
}

var cache_genericDisj_prime gopurs_runtime.Value
var once_genericDisj_prime sync.Once
func Get_genericDisj_prime() gopurs_runtime.Value {
	once_genericDisj_prime.Do(func() {
		cache_genericDisj_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericDisj_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericDisj_prime
}

var cache_genericDisj gopurs_runtime.Value
var once_genericDisj sync.Once
func Get_genericDisj() gopurs_runtime.Value {
	once_genericDisj.Do(func() {
		cache_genericDisj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericDisj(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_genericDisj
}

var cache_genericConj_prime gopurs_runtime.Value
var once_genericConj_prime sync.Once
func Get_genericConj_prime() gopurs_runtime.Value {
	once_genericConj_prime.Do(func() {
		cache_genericConj_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj_prime(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericConj_prime
}

var cache_genericHeytingAlgebraConstructor gopurs_runtime.Value
var once_genericHeytingAlgebraConstructor sync.Once
func Get_genericHeytingAlgebraConstructor() gopurs_runtime.Value {
	once_genericHeytingAlgebraConstructor.Do(func() {
		cache_genericHeytingAlgebraConstructor = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericHeytingAlgebraConstructor(dictGenericHeytingAlgebra_0_box)
})
	})
	return cache_genericHeytingAlgebraConstructor
}

var cache_genericHeytingAlgebraProduct gopurs_runtime.Value
var once_genericHeytingAlgebraProduct sync.Once
func Get_genericHeytingAlgebraProduct() gopurs_runtime.Value {
	once_genericHeytingAlgebraProduct.Do(func() {
		cache_genericHeytingAlgebraProduct = gopurs_runtime.Func2(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_box, dictGenericHeytingAlgebra1_1_box)
})
	})
	return cache_genericHeytingAlgebraProduct
}

var cache_genericConj gopurs_runtime.Value
var once_genericConj sync.Once
func Get_genericConj() gopurs_runtime.Value {
	once_genericConj.Do(func() {
		cache_genericConj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dictGenericHeytingAlgebra_1_box), x_2_box, y_3_box)
})
	})
	return cache_genericConj
}

var cache_from__1498760952 gopurs_runtime.Value
var once_from__1498760952 sync.Once
func Get_from__1498760952() gopurs_runtime.Value {
	once_from__1498760952.Do(func() {
		cache_from__1498760952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from__1498760952(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_from__1498760952
}

var cache_to__1498760952 gopurs_runtime.Value
var once_to__1498760952 sync.Once
func Get_to__1498760952() gopurs_runtime.Value {
	once_to__1498760952.Do(func() {
		cache_to__1498760952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_to__1498760952(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_to__1498760952
}

var cache_genericConj_prime__1368982791 gopurs_runtime.Value
var once_genericConj_prime__1368982791 sync.Once
func Get_genericConj_prime__1368982791() gopurs_runtime.Value {
	once_genericConj_prime__1368982791.Do(func() {
		cache_genericConj_prime__1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj_prime__1368982791(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericConj_prime__1368982791
}

var cache_genericDisj_prime__1368982791 gopurs_runtime.Value
var once_genericDisj_prime__1368982791 sync.Once
func Get_genericDisj_prime__1368982791() gopurs_runtime.Value {
	once_genericDisj_prime__1368982791.Do(func() {
		cache_genericDisj_prime__1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericDisj_prime__1368982791(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericDisj_prime__1368982791
}

var cache_genericFF_prime__265522150 gopurs_runtime.Value
var once_genericFF_prime__265522150 sync.Once
func Get_genericFF_prime__265522150() gopurs_runtime.Value {
	once_genericFF_prime__265522150.Do(func() {
		cache_genericFF_prime__265522150 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFF_prime__265522150(dict_0_box)
})
	})
	return cache_genericFF_prime__265522150
}

var cache_genericImplies_prime__1368982791 gopurs_runtime.Value
var once_genericImplies_prime__1368982791 sync.Once
func Get_genericImplies_prime__1368982791() gopurs_runtime.Value {
	once_genericImplies_prime__1368982791.Do(func() {
		cache_genericImplies_prime__1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies_prime__1368982791(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericImplies_prime__1368982791
}

var cache_genericNot_prime__4234376174 gopurs_runtime.Value
var once_genericNot_prime__4234376174 sync.Once
func Get_genericNot_prime__4234376174() gopurs_runtime.Value {
	once_genericNot_prime__4234376174.Do(func() {
		cache_genericNot_prime__4234376174 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericNot_prime__4234376174(gopurs_runtime.CoerceToStruct[Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_genericNot_prime__4234376174
}

var cache_genericTT_prime__265522150 gopurs_runtime.Value
var once_genericTT_prime__265522150 sync.Once
func Get_genericTT_prime__265522150() gopurs_runtime.Value {
	once_genericTT_prime__265522150.Do(func() {
		cache_genericTT_prime__265522150 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTT_prime__265522150(dict_0_box)
})
	})
	return cache_genericTT_prime__265522150
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

var cache_ff__2527024921 gopurs_runtime.Value
var once_ff__2527024921 sync.Once
func Get_ff__2527024921() gopurs_runtime.Value {
	once_ff__2527024921.Do(func() {
		cache_ff__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ff__2527024921(dict_0_box)
})
	})
	return cache_ff__2527024921
}

var cache_implies__3472268504 gopurs_runtime.Value
var once_implies__3472268504 sync.Once
func Get_implies__3472268504() gopurs_runtime.Value {
	once_implies__3472268504.Do(func() {
		cache_implies__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_implies__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_implies__3472268504
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

var cache_tt__2527024921 gopurs_runtime.Value
var once_tt__2527024921 sync.Once
func Get_tt__2527024921() gopurs_runtime.Value {
	once_tt__2527024921.Do(func() {
		cache_tt__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tt__2527024921(dict_0_box)
})
	})
	return cache_tt__2527024921
}

type Constructor_GenericHeytingAlgebra[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 T_a
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 T_a
}


func init() {
	gopurs_runtime.StructGetters[2831861733] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericHeytingAlgebra[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericConj'": return c.V0
		case "genericDisj'": return c.V1
		case "genericFF'": return c.V2
		case "genericImplies'": return c.V3
		case "genericNot'": return c.V4
		case "genericTT'": return c.V5
		default: panic("Key not found in dictionary Constructor_GenericHeytingAlgebra: " + key)
		}
	}
}


func Call_genericTT_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTT'")
}

func Call_genericTT(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericTT'"))
}

func Call_genericNot_prime(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_genericNot(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply2(Get_genericNot_prime__4234376174(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(dictGeneric_0.V0, x_2)))
}

func Call_genericImplies_prime(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_genericImplies(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply3(Get_genericImplies_prime__1368982791(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(dictGeneric_0.V0, x_2), gopurs_runtime.Apply(dictGeneric_0.V0, y_3)))
}

func Call_genericHeytingAlgebraArgument(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_genericFF_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericFF'")
}

func Call_genericFF(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericFF'"))
}

func Call_genericDisj_prime(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericDisj(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply3(Get_genericDisj_prime__1368982791(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(dictGeneric_0.V0, x_2), gopurs_runtime.Apply(dictGeneric_0.V0, y_3)))
}

func Call_genericConj_prime(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericHeytingAlgebraConstructor(dictGenericHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 gopurs_runtime.Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
var dictGenericHeytingAlgebra1_1 gopurs_runtime.Value = dictGenericHeytingAlgebra1_1_loop
_ = dictGenericHeytingAlgebra1_1
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericConj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericDisj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericFF'")})}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericImplies'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericNot'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_1, "genericTT'")})}})
}

func Call_genericConj(dictGeneric_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], dictGenericHeytingAlgebra_1_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(dictGeneric_0.V1, gopurs_runtime.Apply3(Get_genericConj_prime__1368982791(), gopurs_runtime.Value{Type: 9, IntVal: 2831861733, UnsafePtr: unsafe.Pointer(dictGenericHeytingAlgebra_1)}, gopurs_runtime.Apply(dictGeneric_0.V0, x_2), gopurs_runtime.Apply(dictGeneric_0.V0, y_3)))
}

func Call_from__1498760952(dict_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_to__1498760952(dict_0_loop *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Generic_Rep.Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericConj_prime__1368982791(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_genericDisj_prime__1368982791(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_genericFF_prime__265522150(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericFF'")
}

func Call_genericImplies_prime__1368982791(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_genericNot_prime__4234376174(dict_0_loop *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_GenericHeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_genericTT_prime__265522150(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTT'")
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

func Call_ff__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}

func Call_implies__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_tt__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}


