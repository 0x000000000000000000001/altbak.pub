package Data_Ring_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericSub_prime gopurs_runtime.Value
var once_genericSub_prime sync.Once
func Get_genericSub_prime() gopurs_runtime.Value {
	once_genericSub_prime.Do(func() {
		genericSub_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericSub'")
})
	})
	return genericSub_prime
}

var genericSub gopurs_runtime.Value
var once_genericSub sync.Once
func Get_genericSub() gopurs_runtime.Value {
	once_genericSub.Do(func() {
		genericSub = gopurs_runtime.Func4(func(dictGeneric_0 gopurs_runtime.Value, dictGenericRing_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_1, "genericSub'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
})
	})
	return genericSub
}

var genericRingProduct gopurs_runtime.Value
var once_genericRingProduct sync.Once
func Get_genericRingProduct() gopurs_runtime.Value {
	once_genericRingProduct.Do(func() {
		genericRingProduct = gopurs_runtime.Func2(func(dictGenericRing_0 gopurs_runtime.Value, dictGenericRing1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), gopurs_runtime.ConstructorGet(v_2, 0), gopurs_runtime.ConstructorGet(v1_3, 0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing1_1, "genericSub'"), gopurs_runtime.ConstructorGet(v_2, 1), gopurs_runtime.ConstructorGet(v1_3, 1)))
}))
})
	})
	return genericRingProduct
}

var genericRingNoArguments gopurs_runtime.Value
var once_genericRingNoArguments sync.Once
func Get_genericRingNoArguments() gopurs_runtime.Value {
	once_genericRingNoArguments.Do(func() {
		genericRingNoArguments = gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("NoArguments")
}))
	})
	return genericRingNoArguments
}

var genericRingConstructor gopurs_runtime.Value
var once_genericRingConstructor sync.Once
func Get_genericRingConstructor() gopurs_runtime.Value {
	once_genericRingConstructor.Do(func() {
		genericRingConstructor = gopurs_runtime.Func(func(dictGenericRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericRing_0, "genericSub'"), v_1, v1_2)
}))
})
	})
	return genericRingConstructor
}

var genericRingArgument gopurs_runtime.Value
var once_genericRingArgument sync.Once
func Get_genericRingArgument() gopurs_runtime.Value {
	once_genericRingArgument.Do(func() {
		genericRingArgument = gopurs_runtime.Func(func(dictRing_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericSub'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), v_1, v1_2)
}))
})
	})
	return genericRingArgument
}


