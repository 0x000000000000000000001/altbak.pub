package Data_HeytingAlgebra_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericTT_prime gopurs_runtime.Value
var once_genericTT_prime sync.Once
func Get_genericTT_prime() gopurs_runtime.Value {
	once_genericTT_prime.Do(func() {
		genericTT_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericTT'")
})
	})
	return genericTT_prime
}

var genericTT gopurs_runtime.Value
var once_genericTT sync.Once
func Get_genericTT() gopurs_runtime.Value {
	once_genericTT.Do(func() {
		genericTT = gopurs_runtime.Func2(func(dictGeneric_0 gopurs_runtime.Value, dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericTT'"))
})
	})
	return genericTT
}

var genericNot_prime gopurs_runtime.Value
var once_genericNot_prime sync.Once
func Get_genericNot_prime() gopurs_runtime.Value {
	once_genericNot_prime.Do(func() {
		genericNot_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericNot'")
})
	})
	return genericNot_prime
}

var genericNot gopurs_runtime.Value
var once_genericNot sync.Once
func Get_genericNot() gopurs_runtime.Value {
	once_genericNot.Do(func() {
		genericNot = gopurs_runtime.Func3(func(dictGeneric_0 gopurs_runtime.Value, dictGenericHeytingAlgebra_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericNot'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2)))
})
	})
	return genericNot
}

var genericImplies_prime gopurs_runtime.Value
var once_genericImplies_prime sync.Once
func Get_genericImplies_prime() gopurs_runtime.Value {
	once_genericImplies_prime.Do(func() {
		genericImplies_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericImplies'")
})
	})
	return genericImplies_prime
}

var genericImplies gopurs_runtime.Value
var once_genericImplies sync.Once
func Get_genericImplies() gopurs_runtime.Value {
	once_genericImplies.Do(func() {
		genericImplies = gopurs_runtime.Func4(func(dictGeneric_0 gopurs_runtime.Value, dictGenericHeytingAlgebra_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericImplies'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
})
	})
	return genericImplies
}

var genericHeytingAlgebraNoArguments gopurs_runtime.Value
var once_genericHeytingAlgebraNoArguments sync.Once
func Get_genericHeytingAlgebraNoArguments() gopurs_runtime.Value {
	once_genericHeytingAlgebraNoArguments.Do(func() {
		genericHeytingAlgebraNoArguments = gopurs_runtime.RecordDict([]string{"genericFF'", "genericTT'", "genericImplies'", "genericConj'", "genericDisj'", "genericNot'"}, []gopurs_runtime.Value{gopurs_runtime.Constructor0("NoArguments"), gopurs_runtime.Constructor0("NoArguments"), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("NoArguments")
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("NoArguments")
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("NoArguments")
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("NoArguments")
})})
	})
	return genericHeytingAlgebraNoArguments
}

var genericHeytingAlgebraArgument gopurs_runtime.Value
var once_genericHeytingAlgebraArgument sync.Once
func Get_genericHeytingAlgebraArgument() gopurs_runtime.Value {
	once_genericHeytingAlgebraArgument.Do(func() {
		genericHeytingAlgebraArgument = gopurs_runtime.Func(func(dictHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"genericFF'", "genericTT'", "genericImplies'", "genericConj'", "genericDisj'", "genericNot'"}, []gopurs_runtime.Value{gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff"), gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), v_1, v1_2)
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), v_1)
})})
})
	})
	return genericHeytingAlgebraArgument
}

var genericFF_prime gopurs_runtime.Value
var once_genericFF_prime sync.Once
func Get_genericFF_prime() gopurs_runtime.Value {
	once_genericFF_prime.Do(func() {
		genericFF_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericFF'")
})
	})
	return genericFF_prime
}

var genericFF gopurs_runtime.Value
var once_genericFF sync.Once
func Get_genericFF() gopurs_runtime.Value {
	once_genericFF.Do(func() {
		genericFF = gopurs_runtime.Func2(func(dictGeneric_0 gopurs_runtime.Value, dictGenericHeytingAlgebra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericFF'"))
})
	})
	return genericFF
}

var genericDisj_prime gopurs_runtime.Value
var once_genericDisj_prime sync.Once
func Get_genericDisj_prime() gopurs_runtime.Value {
	once_genericDisj_prime.Do(func() {
		genericDisj_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericDisj'")
})
	})
	return genericDisj_prime
}

var genericDisj gopurs_runtime.Value
var once_genericDisj sync.Once
func Get_genericDisj() gopurs_runtime.Value {
	once_genericDisj.Do(func() {
		genericDisj = gopurs_runtime.Func4(func(dictGeneric_0 gopurs_runtime.Value, dictGenericHeytingAlgebra_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericDisj'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
})
	})
	return genericDisj
}

var genericConj_prime gopurs_runtime.Value
var once_genericConj_prime sync.Once
func Get_genericConj_prime() gopurs_runtime.Value {
	once_genericConj_prime.Do(func() {
		genericConj_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericConj'")
})
	})
	return genericConj_prime
}

var genericHeytingAlgebraConstructor gopurs_runtime.Value
var once_genericHeytingAlgebraConstructor sync.Once
func Get_genericHeytingAlgebraConstructor() gopurs_runtime.Value {
	once_genericHeytingAlgebraConstructor.Do(func() {
		genericHeytingAlgebraConstructor = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"genericFF'", "genericTT'", "genericImplies'", "genericConj'", "genericDisj'", "genericNot'"}, []gopurs_runtime.Value{gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'"), gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), v_1, v1_2)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), v_1, v1_2)
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), v_1)
})})
})
	})
	return genericHeytingAlgebraConstructor
}

var genericHeytingAlgebraProduct gopurs_runtime.Value
var once_genericHeytingAlgebraProduct sync.Once
func Get_genericHeytingAlgebraProduct() gopurs_runtime.Value {
	once_genericHeytingAlgebraProduct.Do(func() {
		genericHeytingAlgebraProduct = gopurs_runtime.Func(func(dictGenericHeytingAlgebra_0 gopurs_runtime.Value) gopurs_runtime.Value {
genericFF_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericFF'")
_ = genericFF_prime1_1_0
genericTT_prime1_2_1 := gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericTT'")
_ = genericTT_prime1_2_1
return gopurs_runtime.Func(func(dictGenericHeytingAlgebra1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"genericFF'", "genericTT'", "genericImplies'", "genericConj'", "genericDisj'", "genericNot'"}, []gopurs_runtime.Value{gopurs_runtime.Constructor2("Product", genericFF_prime1_1_0, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericFF'")), gopurs_runtime.Constructor2("Product", genericTT_prime1_2_1, gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericTT'")), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericImplies'"), gopurs_runtime.ConstructorGet(v_4, 0), gopurs_runtime.ConstructorGet(v1_5, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericImplies'"), gopurs_runtime.ConstructorGet(v_4, 1), gopurs_runtime.ConstructorGet(v1_5, 1)))
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericConj'"), gopurs_runtime.ConstructorGet(v_4, 0), gopurs_runtime.ConstructorGet(v1_5, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericConj'"), gopurs_runtime.ConstructorGet(v_4, 1), gopurs_runtime.ConstructorGet(v1_5, 1)))
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericDisj'"), gopurs_runtime.ConstructorGet(v_4, 0), gopurs_runtime.ConstructorGet(v1_5, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericDisj'"), gopurs_runtime.ConstructorGet(v_4, 1), gopurs_runtime.ConstructorGet(v1_5, 1)))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_0, "genericNot'"), gopurs_runtime.ConstructorGet(v_4, 0)), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra1_3, "genericNot'"), gopurs_runtime.ConstructorGet(v_4, 1)))
})})
})
})
	})
	return genericHeytingAlgebraProduct
}

var genericConj gopurs_runtime.Value
var once_genericConj sync.Once
func Get_genericConj() gopurs_runtime.Value {
	once_genericConj.Do(func() {
		genericConj = gopurs_runtime.Func4(func(dictGeneric_0 gopurs_runtime.Value, dictGenericHeytingAlgebra_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericHeytingAlgebra_1, "genericConj'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
})
	})
	return genericConj
}


