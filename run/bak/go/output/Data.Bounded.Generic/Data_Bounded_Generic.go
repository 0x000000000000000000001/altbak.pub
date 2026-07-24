package Data_Bounded_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericTopNoArguments gopurs_runtime.Value
var once_genericTopNoArguments sync.Once
func Get_genericTopNoArguments() gopurs_runtime.Value {
	once_genericTopNoArguments.Do(func() {
		genericTopNoArguments = gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Constructor0("NoArguments"))
	})
	return genericTopNoArguments
}

var genericTopArgument gopurs_runtime.Value
var once_genericTopArgument sync.Once
func Get_genericTopArgument() gopurs_runtime.Value {
	once_genericTopArgument.Do(func() {
		genericTopArgument = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.RecordGet(dictBounded_0, "top"))
}()
})
	})
	return genericTopArgument
}

var genericTop_prime gopurs_runtime.Value
var once_genericTop_prime sync.Once
func Get_genericTop_prime() gopurs_runtime.Value {
	once_genericTop_prime.Do(func() {
		genericTop_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTop'")
}()
})
	})
	return genericTop_prime
}

var genericTopConstructor gopurs_runtime.Value
var once_genericTopConstructor sync.Once
func Get_genericTopConstructor() gopurs_runtime.Value {
	once_genericTopConstructor.Do(func() {
		genericTopConstructor = gopurs_runtime.Func(func(dictGenericTop_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericTop_0 gopurs_runtime.Value = dictGenericTop_0_loop
_ = dictGenericTop_0
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.RecordGet(dictGenericTop_0, "genericTop'"))
}()
})
	})
	return genericTopConstructor
}

var genericTopProduct gopurs_runtime.Value
var once_genericTopProduct sync.Once
func Get_genericTopProduct() gopurs_runtime.Value {
	once_genericTopProduct.Do(func() {
		genericTopProduct = gopurs_runtime.Func(func(dictGenericTop_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericTop_0 gopurs_runtime.Value = dictGenericTop_0_loop
_ = dictGenericTop_0
genericTop_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericTop_0, "genericTop'")
_ = genericTop_prime1_1_0
return gopurs_runtime.Func(func(dictGenericTop1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Constructor2("Product", genericTop_prime1_1_0, gopurs_runtime.RecordGet(dictGenericTop1_2, "genericTop'")))
})
}()
})
	})
	return genericTopProduct
}

var genericTopSum gopurs_runtime.Value
var once_genericTopSum sync.Once
func Get_genericTopSum() gopurs_runtime.Value {
	once_genericTopSum.Do(func() {
		genericTopSum = gopurs_runtime.Func(func(dictGenericTop_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericTop_0 gopurs_runtime.Value = dictGenericTop_0_loop
_ = dictGenericTop_0
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Constructor1("Inr", gopurs_runtime.RecordGet(dictGenericTop_0, "genericTop'")))
}()
})
	})
	return genericTopSum
}

var genericTop gopurs_runtime.Value
var once_genericTop sync.Once
func Get_genericTop() gopurs_runtime.Value {
	once_genericTop.Do(func() {
		genericTop = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTop(dictGeneric_0_box, dictGenericTop_1_box)
})
	})
	return genericTop
}

var genericBottomNoArguments gopurs_runtime.Value
var once_genericBottomNoArguments sync.Once
func Get_genericBottomNoArguments() gopurs_runtime.Value {
	once_genericBottomNoArguments.Do(func() {
		genericBottomNoArguments = gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Constructor0("NoArguments"))
	})
	return genericBottomNoArguments
}

var genericBottomArgument gopurs_runtime.Value
var once_genericBottomArgument sync.Once
func Get_genericBottomArgument() gopurs_runtime.Value {
	once_genericBottomArgument.Do(func() {
		genericBottomArgument = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.RecordGet(dictBounded_0, "bottom"))
}()
})
	})
	return genericBottomArgument
}

var genericBottom_prime gopurs_runtime.Value
var once_genericBottom_prime sync.Once
func Get_genericBottom_prime() gopurs_runtime.Value {
	once_genericBottom_prime.Do(func() {
		genericBottom_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericBottom'")
}()
})
	})
	return genericBottom_prime
}

var genericBottomConstructor gopurs_runtime.Value
var once_genericBottomConstructor sync.Once
func Get_genericBottomConstructor() gopurs_runtime.Value {
	once_genericBottomConstructor.Do(func() {
		genericBottomConstructor = gopurs_runtime.Func(func(dictGenericBottom_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericBottom_0 gopurs_runtime.Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.RecordGet(dictGenericBottom_0, "genericBottom'"))
}()
})
	})
	return genericBottomConstructor
}

var genericBottomProduct gopurs_runtime.Value
var once_genericBottomProduct sync.Once
func Get_genericBottomProduct() gopurs_runtime.Value {
	once_genericBottomProduct.Do(func() {
		genericBottomProduct = gopurs_runtime.Func(func(dictGenericBottom_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericBottom_0 gopurs_runtime.Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
genericBottom_prime1_1_0 := gopurs_runtime.RecordGet(dictGenericBottom_0, "genericBottom'")
_ = genericBottom_prime1_1_0
return gopurs_runtime.Func(func(dictGenericBottom1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Constructor2("Product", genericBottom_prime1_1_0, gopurs_runtime.RecordGet(dictGenericBottom1_2, "genericBottom'")))
})
}()
})
	})
	return genericBottomProduct
}

var genericBottomSum gopurs_runtime.Value
var once_genericBottomSum sync.Once
func Get_genericBottomSum() gopurs_runtime.Value {
	once_genericBottomSum.Do(func() {
		genericBottomSum = gopurs_runtime.Func(func(dictGenericBottom_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericBottom_0 gopurs_runtime.Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Constructor1("Inl", gopurs_runtime.RecordGet(dictGenericBottom_0, "genericBottom'")))
}()
})
	})
	return genericBottomSum
}

var genericBottom gopurs_runtime.Value
var once_genericBottom sync.Once
func Get_genericBottom() gopurs_runtime.Value {
	once_genericBottom.Do(func() {
		genericBottom = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBottom_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottom(dictGeneric_0_box, dictGenericBottom_1_box)
})
	})
	return genericBottom
}

func Call_genericTop(dictGeneric_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericTop_1 gopurs_runtime.Value = dictGenericTop_1_loop
_ = dictGenericTop_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'"))
}

func Call_genericBottom(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBottom_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBottom_1 gopurs_runtime.Value = dictGenericBottom_1_loop
_ = dictGenericBottom_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericBottom_1, "genericBottom'"))
}


