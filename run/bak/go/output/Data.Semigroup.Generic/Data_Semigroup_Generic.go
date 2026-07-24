package Data_Semigroup_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericSemigroupNoConstructors gopurs_runtime.Value
var once_genericSemigroupNoConstructors sync.Once
func Get_genericSemigroupNoConstructors() gopurs_runtime.Value {
	once_genericSemigroupNoConstructors.Do(func() {
		genericSemigroupNoConstructors = gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
	})
	return genericSemigroupNoConstructors
}

var genericSemigroupNoArguments gopurs_runtime.Value
var once_genericSemigroupNoArguments sync.Once
func Get_genericSemigroupNoArguments() gopurs_runtime.Value {
	once_genericSemigroupNoArguments.Do(func() {
		genericSemigroupNoArguments = gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
	})
	return genericSemigroupNoArguments
}

var genericSemigroupArgument gopurs_runtime.Value
var once_genericSemigroupArgument sync.Once
func Get_genericSemigroupArgument() gopurs_runtime.Value {
	once_genericSemigroupArgument.Do(func() {
		genericSemigroupArgument = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), v_1, v1_2)
}))
})
	})
	return genericSemigroupArgument
}

var genericAppend_prime gopurs_runtime.Value
var once_genericAppend_prime sync.Once
func Get_genericAppend_prime() gopurs_runtime.Value {
	once_genericAppend_prime.Do(func() {
		genericAppend_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericAppend'")
})
	})
	return genericAppend_prime
}

var genericSemigroupConstructor gopurs_runtime.Value
var once_genericSemigroupConstructor sync.Once
func Get_genericSemigroupConstructor() gopurs_runtime.Value {
	once_genericSemigroupConstructor.Do(func() {
		genericSemigroupConstructor = gopurs_runtime.Func(func(dictGenericSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup_0, "genericAppend'"), v_1, v1_2)
}))
})
	})
	return genericSemigroupConstructor
}

var genericSemigroupProduct gopurs_runtime.Value
var once_genericSemigroupProduct sync.Once
func Get_genericSemigroupProduct() gopurs_runtime.Value {
	once_genericSemigroupProduct.Do(func() {
		genericSemigroupProduct = gopurs_runtime.Func2(func(dictGenericSemigroup_0 gopurs_runtime.Value, dictGenericSemigroup1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericAppend'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup_0, "genericAppend'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup1_1, "genericAppend'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]))
}))
})
	})
	return genericSemigroupProduct
}

var genericAppend gopurs_runtime.Value
var once_genericAppend sync.Once
func Get_genericAppend() gopurs_runtime.Value {
	once_genericAppend.Do(func() {
		genericAppend = gopurs_runtime.Func4(func(dictGeneric_0 gopurs_runtime.Value, dictGenericSemigroup_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericSemigroup_1, "genericAppend'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3)))
})
	})
	return genericAppend
}


