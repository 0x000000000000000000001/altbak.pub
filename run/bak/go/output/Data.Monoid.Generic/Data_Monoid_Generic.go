package Data_Monoid_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericMonoidNoArguments gopurs_runtime.Value
var once_genericMonoidNoArguments sync.Once
func Get_genericMonoidNoArguments() gopurs_runtime.Value {
	once_genericMonoidNoArguments.Do(func() {
		genericMonoidNoArguments = gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.Constructor0("NoArguments"))
	})
	return genericMonoidNoArguments
}

var genericMonoidArgument gopurs_runtime.Value
var once_genericMonoidArgument sync.Once
func Get_genericMonoidArgument() gopurs_runtime.Value {
	once_genericMonoidArgument.Do(func() {
		genericMonoidArgument = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
}()
})
	})
	return genericMonoidArgument
}

var genericMempty_prime gopurs_runtime.Value
var once_genericMempty_prime sync.Once
func Get_genericMempty_prime() gopurs_runtime.Value {
	once_genericMempty_prime.Do(func() {
		genericMempty_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericMempty'")
}()
})
	})
	return genericMempty_prime
}

var genericMonoidConstructor gopurs_runtime.Value
var once_genericMonoidConstructor sync.Once
func Get_genericMonoidConstructor() gopurs_runtime.Value {
	once_genericMonoidConstructor.Do(func() {
		genericMonoidConstructor = gopurs_runtime.Func(func(dictGenericMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
_ = dictGenericMonoid_0
return gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'"))
}()
})
	})
	return genericMonoidConstructor
}

var genericMonoidProduct gopurs_runtime.Value
var once_genericMonoidProduct sync.Once
func Get_genericMonoidProduct() gopurs_runtime.Value {
	once_genericMonoidProduct.Do(func() {
		genericMonoidProduct = gopurs_runtime.Func(func(dictGenericMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
_ = dictGenericMonoid_0
genericMempty_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'")
_ = genericMempty_prime1_1_0
return gopurs_runtime.Func(func(dictGenericMonoid1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.Constructor2("Product", genericMempty_prime1_1_0, gopurs_runtime.RecordGet(dictGenericMonoid1_2, "genericMempty'")))
})
}()
})
	})
	return genericMonoidProduct
}

var genericMempty gopurs_runtime.Value
var once_genericMempty sync.Once
func Get_genericMempty() gopurs_runtime.Value {
	once_genericMempty.Do(func() {
		genericMempty = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMempty(dictGeneric_0_box, dictGenericMonoid_1_box)
})
	})
	return genericMempty
}

func Call_genericMempty(dictGeneric_0_loop gopurs_runtime.Value, dictGenericMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericMonoid_1 gopurs_runtime.Value = dictGenericMonoid_1_loop
_ = dictGenericMonoid_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericMonoid_1, "genericMempty'"))
}


