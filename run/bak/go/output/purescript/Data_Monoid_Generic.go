package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_Generic_GenericMonoid_dollarDict gopurs_runtime.Value
var once_Data_Monoid_Generic_GenericMonoid_dollarDict sync.Once
func Get_Data_Monoid_Generic_GenericMonoid_dollarDict() gopurs_runtime.Value {
	once_Data_Monoid_Generic_GenericMonoid_dollarDict.Do(func() {
		cache_Data_Monoid_Generic_GenericMonoid_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Generic_GenericMonoid_dollarDict(x_0_box)
})
	})
	return cache_Data_Monoid_Generic_GenericMonoid_dollarDict
}

var cache_Data_Monoid_Generic_genericMonoidNoArguments gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidNoArguments sync.Once
func Get_Data_Monoid_Generic_genericMonoidNoArguments() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidNoArguments.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidNoArguments = gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Generic_GenericMonoid{1, gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}})}
	})
	return cache_Data_Monoid_Generic_genericMonoidNoArguments
}

var cache_Data_Monoid_Generic_genericMonoidArgument gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidArgument sync.Once
func Get_Data_Monoid_Generic_genericMonoidArgument() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidArgument.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidArgument = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Generic_genericMonoidArgument(dictMonoid_0_box)
})
	})
	return cache_Data_Monoid_Generic_genericMonoidArgument
}

var cache_Data_Monoid_Generic_genericMempty_prime gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMempty_prime sync.Once
func Get_Data_Monoid_Generic_genericMempty_prime() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMempty_prime.Do(func() {
		cache_Data_Monoid_Generic_genericMempty_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Generic_genericMempty_prime(dict_0_box)
})
	})
	return cache_Data_Monoid_Generic_genericMempty_prime
}

var cache_Data_Monoid_Generic_genericMonoidConstructor gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidConstructor sync.Once
func Get_Data_Monoid_Generic_genericMonoidConstructor() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidConstructor.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidConstructor = gopurs_runtime.Func(func(dictGenericMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Generic_genericMonoidConstructor(dictGenericMonoid_0_box)
})
	})
	return cache_Data_Monoid_Generic_genericMonoidConstructor
}

var cache_Data_Monoid_Generic_genericMonoidProduct gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMonoidProduct sync.Once
func Get_Data_Monoid_Generic_genericMonoidProduct() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMonoidProduct.Do(func() {
		cache_Data_Monoid_Generic_genericMonoidProduct = gopurs_runtime.Func2(func(dictGenericMonoid_0_box gopurs_runtime.Value, dictGenericMonoid1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Generic_genericMonoidProduct(dictGenericMonoid_0_box, dictGenericMonoid1_1_box)
})
	})
	return cache_Data_Monoid_Generic_genericMonoidProduct
}

var cache_Data_Monoid_Generic_genericMempty gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMempty sync.Once
func Get_Data_Monoid_Generic_genericMempty() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMempty.Do(func() {
		cache_Data_Monoid_Generic_genericMempty = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Generic_genericMempty(dictGeneric_0_box, dictGenericMonoid_1_box)
})
	})
	return cache_Data_Monoid_Generic_genericMempty
}

var cache_Data_Monoid_Generic_genericMempty_prime__3900442342 gopurs_runtime.Value
var once_Data_Monoid_Generic_genericMempty_prime__3900442342 sync.Once
func Get_Data_Monoid_Generic_genericMempty_prime__3900442342() gopurs_runtime.Value {
	once_Data_Monoid_Generic_genericMempty_prime__3900442342.Do(func() {
		cache_Data_Monoid_Generic_genericMempty_prime__3900442342 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Generic_genericMempty_prime__3900442342(dict_0_box)
})
	})
	return cache_Data_Monoid_Generic_genericMempty_prime__3900442342
}

type Constructor_Data_Monoid_Generic_GenericMonoid struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2569012965] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_Generic_GenericMonoid)(ptr)
		_ = c
		switch key {
		case "genericMempty'": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Monoid_Generic_GenericMonoid: " + key)
		}
	}
}


func Call_Data_Monoid_Generic_GenericMonoid_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Generic_genericMonoidArgument(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Generic_GenericMonoid{1, gopurs_runtime.RecordGet(dictMonoid_0, "mempty")})}
}

func Call_Data_Monoid_Generic_genericMempty_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericMempty'")
}

func Call_Data_Monoid_Generic_genericMonoidConstructor(dictGenericMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
_ = dictGenericMonoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Generic_GenericMonoid{1, gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'")})}
}

func Call_Data_Monoid_Generic_genericMonoidProduct(dictGenericMonoid_0_loop gopurs_runtime.Value, dictGenericMonoid1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericMonoid_0 gopurs_runtime.Value = dictGenericMonoid_0_loop
_ = dictGenericMonoid_0
var dictGenericMonoid1_1 gopurs_runtime.Value = dictGenericMonoid1_1_loop
_ = dictGenericMonoid1_1
return gopurs_runtime.Value{Type: 9, IntVal: 2569012965, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Generic_GenericMonoid{1, gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, gopurs_runtime.RecordGet(dictGenericMonoid_0, "genericMempty'"), gopurs_runtime.RecordGet(dictGenericMonoid1_1, "genericMempty'")})}})}
}

func Call_Data_Monoid_Generic_genericMempty(dictGeneric_0_loop gopurs_runtime.Value, dictGenericMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericMonoid_1 gopurs_runtime.Value = dictGenericMonoid_1_loop
_ = dictGenericMonoid_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericMonoid_1, "genericMempty'"))
}

func Call_Data_Monoid_Generic_genericMempty_prime__3900442342(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericMempty'")
}


