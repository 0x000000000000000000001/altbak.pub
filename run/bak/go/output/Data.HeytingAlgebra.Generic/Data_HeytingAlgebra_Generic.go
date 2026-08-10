package Data_HeytingAlgebra_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
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

var cache_genericTT_prime__gopurs_runtime_Value_265522150 gopurs_runtime.Value
var once_genericTT_prime__gopurs_runtime_Value_265522150 sync.Once
func Get_genericTT_prime__gopurs_runtime_Value_265522150() gopurs_runtime.Value {
	once_genericTT_prime__gopurs_runtime_Value_265522150.Do(func() {
		cache_genericTT_prime__gopurs_runtime_Value_265522150 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTT_prime__gopurs_runtime_Value_265522150(dict_0_box)
})
	})
	return cache_genericTT_prime__gopurs_runtime_Value_265522150
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
return Call_genericNot_prime(dict_0_box)
})
	})
	return cache_genericNot_prime
}

var cache_genericNot_prime__gopurs_runtime_Value_4234376174 gopurs_runtime.Value
var once_genericNot_prime__gopurs_runtime_Value_4234376174 sync.Once
func Get_genericNot_prime__gopurs_runtime_Value_4234376174() gopurs_runtime.Value {
	once_genericNot_prime__gopurs_runtime_Value_4234376174.Do(func() {
		cache_genericNot_prime__gopurs_runtime_Value_4234376174 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericNot_prime__gopurs_runtime_Value_4234376174(dict_0_box)
})
	})
	return cache_genericNot_prime__gopurs_runtime_Value_4234376174
}

var cache_genericNot gopurs_runtime.Value
var once_genericNot sync.Once
func Get_genericNot() gopurs_runtime.Value {
	once_genericNot.Do(func() {
		cache_genericNot = gopurs_runtime.Func3(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericNot(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box, x_2_box)
})
	})
	return cache_genericNot
}

var cache_genericImplies_prime gopurs_runtime.Value
var once_genericImplies_prime sync.Once
func Get_genericImplies_prime() gopurs_runtime.Value {
	once_genericImplies_prime.Do(func() {
		cache_genericImplies_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies_prime(dict_0_box)
})
	})
	return cache_genericImplies_prime
}

var cache_genericImplies_prime__gopurs_runtime_Value_1368982791 gopurs_runtime.Value
var once_genericImplies_prime__gopurs_runtime_Value_1368982791 sync.Once
func Get_genericImplies_prime__gopurs_runtime_Value_1368982791() gopurs_runtime.Value {
	once_genericImplies_prime__gopurs_runtime_Value_1368982791.Do(func() {
		cache_genericImplies_prime__gopurs_runtime_Value_1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies_prime__gopurs_runtime_Value_1368982791(dict_0_box)
})
	})
	return cache_genericImplies_prime__gopurs_runtime_Value_1368982791
}

var cache_genericImplies gopurs_runtime.Value
var once_genericImplies sync.Once
func Get_genericImplies() gopurs_runtime.Value {
	once_genericImplies.Do(func() {
		cache_genericImplies = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericImplies(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box, x_2_box, y_3_box)
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
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}
}), gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil}})
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

var cache_genericFF_prime__gopurs_runtime_Value_265522150 gopurs_runtime.Value
var once_genericFF_prime__gopurs_runtime_Value_265522150 sync.Once
func Get_genericFF_prime__gopurs_runtime_Value_265522150() gopurs_runtime.Value {
	once_genericFF_prime__gopurs_runtime_Value_265522150.Do(func() {
		cache_genericFF_prime__gopurs_runtime_Value_265522150 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericFF_prime__gopurs_runtime_Value_265522150(dict_0_box)
})
	})
	return cache_genericFF_prime__gopurs_runtime_Value_265522150
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
return Call_genericDisj_prime(dict_0_box)
})
	})
	return cache_genericDisj_prime
}

var cache_genericDisj_prime__gopurs_runtime_Value_1368982791 gopurs_runtime.Value
var once_genericDisj_prime__gopurs_runtime_Value_1368982791 sync.Once
func Get_genericDisj_prime__gopurs_runtime_Value_1368982791() gopurs_runtime.Value {
	once_genericDisj_prime__gopurs_runtime_Value_1368982791.Do(func() {
		cache_genericDisj_prime__gopurs_runtime_Value_1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericDisj_prime__gopurs_runtime_Value_1368982791(dict_0_box)
})
	})
	return cache_genericDisj_prime__gopurs_runtime_Value_1368982791
}

var cache_genericDisj gopurs_runtime.Value
var once_genericDisj sync.Once
func Get_genericDisj() gopurs_runtime.Value {
	once_genericDisj.Do(func() {
		cache_genericDisj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericDisj(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box, x_2_box, y_3_box)
})
	})
	return cache_genericDisj
}

var cache_genericConj_prime gopurs_runtime.Value
var once_genericConj_prime sync.Once
func Get_genericConj_prime() gopurs_runtime.Value {
	once_genericConj_prime.Do(func() {
		cache_genericConj_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj_prime(dict_0_box)
})
	})
	return cache_genericConj_prime
}

var cache_genericConj_prime__gopurs_runtime_Value_1368982791 gopurs_runtime.Value
var once_genericConj_prime__gopurs_runtime_Value_1368982791 sync.Once
func Get_genericConj_prime__gopurs_runtime_Value_1368982791() gopurs_runtime.Value {
	once_genericConj_prime__gopurs_runtime_Value_1368982791.Do(func() {
		cache_genericConj_prime__gopurs_runtime_Value_1368982791 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj_prime__gopurs_runtime_Value_1368982791(dict_0_box)
})
	})
	return cache_genericConj_prime__gopurs_runtime_Value_1368982791
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
		cache_genericHeytingAlgebraProduct = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_box)
})
	})
	return cache_genericHeytingAlgebraProduct
}

var cache_genericConj gopurs_runtime.Value
var once_genericConj sync.Once
func Get_genericConj() gopurs_runtime.Value {
	once_genericConj.Do(func() {
		cache_genericConj = gopurs_runtime.Func4(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericHeytingAlgebra_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericConj(dictGeneric_0_box, dictGenericHeytingAlgebra_1_box, x_2_box, y_3_box)
})
	})
	return cache_genericConj
}

func Call_genericTT_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTT'")
}

func Call_genericTT_prime__gopurs_runtime_Value_265522150(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_genericNot_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericNot'")
}

func Call_genericNot_prime__gopurs_runtime_Value_4234376174(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericNot'")
}

func Call_genericNot(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericNot'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2)))
}

func Call_genericImplies_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericImplies'")
}

func Call_genericImplies_prime__gopurs_runtime_Value_1368982791(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericImplies'")
}

func Call_genericImplies(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericImplies'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
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

func Call_genericFF_prime__gopurs_runtime_Value_265522150(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_genericDisj_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericDisj'")
}

func Call_genericDisj_prime__gopurs_runtime_Value_1368982791(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericDisj'")
}

func Call_genericDisj(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericDisj'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
}

func Call_genericConj_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericConj'")
}

func Call_genericConj_prime__gopurs_runtime_Value_1368982791(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericConj'")
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

func Call_genericHeytingAlgebraProduct(dictGenericHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericHeytingAlgebra_0 gopurs_runtime.Value = dictGenericHeytingAlgebra_0_loop
_ = dictGenericHeytingAlgebra_0
genericFF_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'")
_ = genericFF_prime1_1_0
genericTT_prime1_2_1 := gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'")
_ = genericTT_prime1_2_1
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"genericConj'", "genericDisj'", "genericFF'", "genericImplies'", "genericNot'", "genericTT'"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericConj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericDisj'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericFF_prime1_1_0, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericFF'")})}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericImplies'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v1_5.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericNot'"), (*pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
}), gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericTT_prime1_2_1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericTT'")})}})
})
}

func Call_genericConj(dictGeneric_0_loop gopurs_runtime.Value, dictGenericHeytingAlgebra_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericHeytingAlgebra_1 gopurs_runtime.Value = dictGenericHeytingAlgebra_1_loop
_ = dictGenericHeytingAlgebra_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericConj'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
}


