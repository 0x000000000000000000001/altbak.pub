package Data_Monoid_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	unsafe "unsafe"
)

var cache_genericMonoidNoArguments gopurs_runtime.Value
var once_genericMonoidNoArguments sync.Once
func Get_genericMonoidNoArguments() gopurs_runtime.Value {
	once_genericMonoidNoArguments.Do(func() {
		cache_genericMonoidNoArguments = gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: nil})
	})
	return cache_genericMonoidNoArguments
}

var cache_genericMonoidArgument gopurs_runtime.Value
var once_genericMonoidArgument sync.Once
func Get_genericMonoidArgument() gopurs_runtime.Value {
	once_genericMonoidArgument.Do(func() {
		cache_genericMonoidArgument = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMonoidArgument(dictMonoid_0_box)
})
	})
	return cache_genericMonoidArgument
}

var cache_genericMempty_prime gopurs_runtime.Value
var once_genericMempty_prime sync.Once
func Get_genericMempty_prime() gopurs_runtime.Value {
	once_genericMempty_prime.Do(func() {
		cache_genericMempty_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMempty_prime(dict_0_box)
})
	})
	return cache_genericMempty_prime
}

var cache_genericMempty_prime__gopurs_runtime_Value_3900442342 gopurs_runtime.Value
var once_genericMempty_prime__gopurs_runtime_Value_3900442342 sync.Once
func Get_genericMempty_prime__gopurs_runtime_Value_3900442342() gopurs_runtime.Value {
	once_genericMempty_prime__gopurs_runtime_Value_3900442342.Do(func() {
		cache_genericMempty_prime__gopurs_runtime_Value_3900442342 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMempty_prime__gopurs_runtime_Value_3900442342(dict_0_box)
})
	})
	return cache_genericMempty_prime__gopurs_runtime_Value_3900442342
}

var cache_genericMonoidConstructor gopurs_runtime.Value
var once_genericMonoidConstructor sync.Once
func Get_genericMonoidConstructor() gopurs_runtime.Value {
	once_genericMonoidConstructor.Do(func() {
		cache_genericMonoidConstructor = gopurs_runtime.Func(func(dictGenericMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMonoidConstructor(dictGenericMonoid_0_box)
})
	})
	return cache_genericMonoidConstructor
}

var cache_genericMonoidProduct gopurs_runtime.Value
var once_genericMonoidProduct sync.Once
func Get_genericMonoidProduct() gopurs_runtime.Value {
	once_genericMonoidProduct.Do(func() {
		cache_genericMonoidProduct = gopurs_runtime.Func(func(dictGenericMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMonoidProduct(dictGenericMonoid_0_box)
})
	})
	return cache_genericMonoidProduct
}

var cache_genericMempty gopurs_runtime.Value
var once_genericMempty sync.Once
func Get_genericMempty() gopurs_runtime.Value {
	once_genericMempty.Do(func() {
		cache_genericMempty = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericMempty(dictGeneric_0_box, dictGenericMonoid_1_box)
})
	})
	return cache_genericMempty
}

func Call_genericMonoidArgument(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
}

func Call_genericMempty_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericMempty'")
}

func Call_genericMempty_prime__gopurs_runtime_Value_3900442342(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericMempty'")
}

func Call_genericMonoidConstructor(dictGenericMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
_ = dictGenericMonoid_0
return gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'"))
}

func Call_genericMonoidProduct(dictGenericMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
_ = dictGenericMonoid_0
genericMempty_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'")
_ = genericMempty_prime1_1_0
return gopurs_runtime.Func(func(dictGenericMonoid1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericMempty'", gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, genericMempty_prime1_1_0, gopurs_runtime.RecordGet(dictGenericMonoid1_2, "genericMempty'")})})
})
}

func Call_genericMempty(dictGeneric_0_loop gopurs_runtime.Value, dictGenericMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericMonoid_1 gopurs_runtime.Value = dictGenericMonoid_1_loop
_ = dictGenericMonoid_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericMonoid_1, "genericMempty'"))
}


