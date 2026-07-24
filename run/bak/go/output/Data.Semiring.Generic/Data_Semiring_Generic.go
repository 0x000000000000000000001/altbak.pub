package Data_Semiring_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericZero_prime gopurs_runtime.Value
var once_genericZero_prime sync.Once
func Get_genericZero_prime() gopurs_runtime.Value {
	once_genericZero_prime.Do(func() {
		genericZero_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericZero'")
}()
})
	})
	return genericZero_prime
}

var genericZero gopurs_runtime.Value
var once_genericZero sync.Once
func Get_genericZero() gopurs_runtime.Value {
	once_genericZero.Do(func() {
		genericZero = gopurs_runtime.Func2(Call_genericZero)
	})
	return genericZero
}

var genericSemiringNoArguments gopurs_runtime.Value
var once_genericSemiringNoArguments sync.Once
func Get_genericSemiringNoArguments() gopurs_runtime.Value {
	once_genericSemiringNoArguments.Do(func() {
		genericSemiringNoArguments = gopurs_runtime.RecordDict4("genericAdd'", "genericZero'", "genericMul'", "genericOne'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("NoArguments")
}), gopurs_runtime.Constructor0("NoArguments"), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("NoArguments")
}), gopurs_runtime.Constructor0("NoArguments"))
	})
	return genericSemiringNoArguments
}

var genericSemiringArgument gopurs_runtime.Value
var once_genericSemiringArgument sync.Once
func Get_genericSemiringArgument() gopurs_runtime.Value {
	once_genericSemiringArgument.Do(func() {
		genericSemiringArgument = gopurs_runtime.Func(func(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict4("genericAdd'", "genericZero'", "genericMul'", "genericOne'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0_loop, "add"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictSemiring_0_loop, "zero"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0_loop, "mul"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictSemiring_0_loop, "one"))
}()
})
	})
	return genericSemiringArgument
}

var genericOne_prime gopurs_runtime.Value
var once_genericOne_prime sync.Once
func Get_genericOne_prime() gopurs_runtime.Value {
	once_genericOne_prime.Do(func() {
		genericOne_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericOne'")
}()
})
	})
	return genericOne_prime
}

var genericOne gopurs_runtime.Value
var once_genericOne sync.Once
func Get_genericOne() gopurs_runtime.Value {
	once_genericOne.Do(func() {
		genericOne = gopurs_runtime.Func2(Call_genericOne)
	})
	return genericOne
}

var genericMul_prime gopurs_runtime.Value
var once_genericMul_prime sync.Once
func Get_genericMul_prime() gopurs_runtime.Value {
	once_genericMul_prime.Do(func() {
		genericMul_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericMul'")
}()
})
	})
	return genericMul_prime
}

var genericMul gopurs_runtime.Value
var once_genericMul sync.Once
func Get_genericMul() gopurs_runtime.Value {
	once_genericMul.Do(func() {
		genericMul = gopurs_runtime.Func4(Call_genericMul)
	})
	return genericMul
}

var genericAdd_prime gopurs_runtime.Value
var once_genericAdd_prime sync.Once
func Get_genericAdd_prime() gopurs_runtime.Value {
	once_genericAdd_prime.Do(func() {
		genericAdd_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericAdd'")
}()
})
	})
	return genericAdd_prime
}

var genericSemiringConstructor gopurs_runtime.Value
var once_genericSemiringConstructor sync.Once
func Get_genericSemiringConstructor() gopurs_runtime.Value {
	once_genericSemiringConstructor.Do(func() {
		genericSemiringConstructor = gopurs_runtime.Func(func(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
return gopurs_runtime.RecordDict4("genericAdd'", "genericZero'", "genericMul'", "genericOne'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericAdd'"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericZero'"), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericMul'"), v_1, v1_2)
}), gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericOne'"))
}()
})
	})
	return genericSemiringConstructor
}

var genericSemiringProduct gopurs_runtime.Value
var once_genericSemiringProduct sync.Once
func Get_genericSemiringProduct() gopurs_runtime.Value {
	once_genericSemiringProduct.Do(func() {
		genericSemiringProduct = gopurs_runtime.Func(func(dictGenericSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericSemiring_0 gopurs_runtime.Value = dictGenericSemiring_0_loop
_ = dictGenericSemiring_0
genericZero_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericZero'")
_ = genericZero_prime1_1_0
genericOne_prime1_2_1 := gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericOne'")
_ = genericOne_prime1_2_1
return gopurs_runtime.Func(func(dictGenericSemiring1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("genericAdd'", "genericZero'", "genericMul'", "genericOne'", gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericAdd'"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericAdd'"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]))
}), gopurs_runtime.Constructor2("Product", genericZero_prime1_1_0, gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericZero'")), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_0_loop, "genericMul'"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericMul'"), (*[1024]gopurs_runtime.Value)(v_4.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_5.UnsafePtr)[1]))
}), gopurs_runtime.Constructor2("Product", genericOne_prime1_2_1, gopurs_runtime.RecordGet(dictGenericSemiring1_3, "genericOne'")))
})
}()
})
	})
	return genericSemiringProduct
}

var genericAdd gopurs_runtime.Value
var once_genericAdd sync.Once
func Get_genericAdd() gopurs_runtime.Value {
	once_genericAdd.Do(func() {
		genericAdd = gopurs_runtime.Func4(Call_genericAdd)
	})
	return genericAdd
}

func Call_genericZero(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1_loop, "genericZero'"))
}

func Call_genericOne(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "to"), gopurs_runtime.RecordGet(dictGenericSemiring_1_loop, "genericOne'"))
}

func Call_genericMul(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_1_loop, "genericMul'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), x_2_loop), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), y_3_loop)))
}

func Call_genericAdd(dictGeneric_0_loop gopurs_runtime.Value, dictGenericSemiring_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericSemiring_1 gopurs_runtime.Value = dictGenericSemiring_1_loop
_ = dictGenericSemiring_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemiring_1_loop, "genericAdd'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), x_2_loop), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), y_3_loop)))
}


