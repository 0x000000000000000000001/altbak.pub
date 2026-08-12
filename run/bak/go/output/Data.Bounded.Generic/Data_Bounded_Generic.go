package Data_Bounded_Generic

import (
	pkg_Data_Generic_Rep "gopurs/output/Data.Generic.Rep"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genericTopNoArguments gopurs_runtime.Value
var once_genericTopNoArguments sync.Once
func Get_genericTopNoArguments() gopurs_runtime.Value {
	once_genericTopNoArguments.Do(func() {
		cache_genericTopNoArguments = gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_genericTopNoArguments
}

var cache_genericTopArgument gopurs_runtime.Value
var once_genericTopArgument sync.Once
func Get_genericTopArgument() gopurs_runtime.Value {
	once_genericTopArgument.Do(func() {
		cache_genericTopArgument = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopArgument(dictBounded_0_box)
})
	})
	return cache_genericTopArgument
}

var cache_genericTop_prime gopurs_runtime.Value
var once_genericTop_prime sync.Once
func Get_genericTop_prime() gopurs_runtime.Value {
	once_genericTop_prime.Do(func() {
		cache_genericTop_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTop_prime(dict_0_box)
})
	})
	return cache_genericTop_prime
}

var cache_genericTop_prime__gopurs_runtime_Value_1114003774 gopurs_runtime.Value
var once_genericTop_prime__gopurs_runtime_Value_1114003774 sync.Once
func Get_genericTop_prime__gopurs_runtime_Value_1114003774() gopurs_runtime.Value {
	once_genericTop_prime__gopurs_runtime_Value_1114003774.Do(func() {
		cache_genericTop_prime__gopurs_runtime_Value_1114003774 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTop_prime__gopurs_runtime_Value_1114003774(dict_0_box)
})
	})
	return cache_genericTop_prime__gopurs_runtime_Value_1114003774
}

var cache_genericTopConstructor gopurs_runtime.Value
var once_genericTopConstructor sync.Once
func Get_genericTopConstructor() gopurs_runtime.Value {
	once_genericTopConstructor.Do(func() {
		cache_genericTopConstructor = gopurs_runtime.Func(func(dictGenericTop_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopConstructor(dictGenericTop_0_box)
})
	})
	return cache_genericTopConstructor
}

var cache_genericTopProduct gopurs_runtime.Value
var once_genericTopProduct sync.Once
func Get_genericTopProduct() gopurs_runtime.Value {
	once_genericTopProduct.Do(func() {
		cache_genericTopProduct = gopurs_runtime.Func2(func(dictGenericTop_0_box gopurs_runtime.Value, dictGenericTop1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopProduct(dictGenericTop_0_box, dictGenericTop1_1_box)
})
	})
	return cache_genericTopProduct
}

var cache_genericTopSum gopurs_runtime.Value
var once_genericTopSum sync.Once
func Get_genericTopSum() gopurs_runtime.Value {
	once_genericTopSum.Do(func() {
		cache_genericTopSum = gopurs_runtime.Func(func(dictGenericTop_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTopSum(dictGenericTop_0_box)
})
	})
	return cache_genericTopSum
}

var cache_genericTop gopurs_runtime.Value
var once_genericTop sync.Once
func Get_genericTop() gopurs_runtime.Value {
	once_genericTop.Do(func() {
		cache_genericTop = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericTop_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericTop(dictGeneric_0_box, dictGenericTop_1_box)
})
	})
	return cache_genericTop
}

var cache_genericBottomNoArguments gopurs_runtime.Value
var once_genericBottomNoArguments sync.Once
func Get_genericBottomNoArguments() gopurs_runtime.Value {
	once_genericBottomNoArguments.Do(func() {
		cache_genericBottomNoArguments = gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Value{Type: 9, IntVal: 1454898258, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_genericBottomNoArguments
}

var cache_genericBottomArgument gopurs_runtime.Value
var once_genericBottomArgument sync.Once
func Get_genericBottomArgument() gopurs_runtime.Value {
	once_genericBottomArgument.Do(func() {
		cache_genericBottomArgument = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomArgument(dictBounded_0_box)
})
	})
	return cache_genericBottomArgument
}

var cache_genericBottom_prime gopurs_runtime.Value
var once_genericBottom_prime sync.Once
func Get_genericBottom_prime() gopurs_runtime.Value {
	once_genericBottom_prime.Do(func() {
		cache_genericBottom_prime = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottom_prime(dict_0_box)
})
	})
	return cache_genericBottom_prime
}

var cache_genericBottom_prime__gopurs_runtime_Value_2482077850 gopurs_runtime.Value
var once_genericBottom_prime__gopurs_runtime_Value_2482077850 sync.Once
func Get_genericBottom_prime__gopurs_runtime_Value_2482077850() gopurs_runtime.Value {
	once_genericBottom_prime__gopurs_runtime_Value_2482077850.Do(func() {
		cache_genericBottom_prime__gopurs_runtime_Value_2482077850 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottom_prime__gopurs_runtime_Value_2482077850(dict_0_box)
})
	})
	return cache_genericBottom_prime__gopurs_runtime_Value_2482077850
}

var cache_genericBottomConstructor gopurs_runtime.Value
var once_genericBottomConstructor sync.Once
func Get_genericBottomConstructor() gopurs_runtime.Value {
	once_genericBottomConstructor.Do(func() {
		cache_genericBottomConstructor = gopurs_runtime.Func(func(dictGenericBottom_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomConstructor(dictGenericBottom_0_box)
})
	})
	return cache_genericBottomConstructor
}

var cache_genericBottomProduct gopurs_runtime.Value
var once_genericBottomProduct sync.Once
func Get_genericBottomProduct() gopurs_runtime.Value {
	once_genericBottomProduct.Do(func() {
		cache_genericBottomProduct = gopurs_runtime.Func2(func(dictGenericBottom_0_box gopurs_runtime.Value, dictGenericBottom1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomProduct(dictGenericBottom_0_box, dictGenericBottom1_1_box)
})
	})
	return cache_genericBottomProduct
}

var cache_genericBottomSum gopurs_runtime.Value
var once_genericBottomSum sync.Once
func Get_genericBottomSum() gopurs_runtime.Value {
	once_genericBottomSum.Do(func() {
		cache_genericBottomSum = gopurs_runtime.Func(func(dictGenericBottom_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottomSum(dictGenericBottom_0_box)
})
	})
	return cache_genericBottomSum
}

var cache_genericBottom gopurs_runtime.Value
var once_genericBottom sync.Once
func Get_genericBottom() gopurs_runtime.Value {
	once_genericBottom.Do(func() {
		cache_genericBottom = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, dictGenericBottom_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genericBottom(dictGeneric_0_box, dictGenericBottom_1_box)
})
	})
	return cache_genericBottom
}

type Constructor_GenericTop[T_a any] struct {
	Rc uint32
	V0 T_a
}


func init() {
	gopurs_runtime.StructGetters[190657853] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericTop[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericTop'": return c.V0
		default: panic("Key not found in dictionary Constructor_GenericTop: " + key)
		}
	}
}


type Constructor_GenericBottom[T_a any] struct {
	Rc uint32
	V0 T_a
}


func init() {
	gopurs_runtime.StructGetters[2049928665] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_GenericBottom[gopurs_runtime.Value])(ptr)
		switch key {
		case "genericBottom'": return c.V0
		default: panic("Key not found in dictionary Constructor_GenericBottom: " + key)
		}
	}
}


func Call_genericTopArgument(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.RecordGet(dictBounded_0, "top"))
}

func Call_genericTop_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTop'")
}

func Call_genericTop_prime__gopurs_runtime_Value_1114003774(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericTop'")
}

func Call_genericTopConstructor(dictGenericTop_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericTop_0 gopurs_runtime.Value = dictGenericTop_0_loop
_ = dictGenericTop_0
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.RecordGet(dictGenericTop_0, "genericTop'"))
}

func Call_genericTopProduct(dictGenericTop_0_loop gopurs_runtime.Value, dictGenericTop1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericTop_0 gopurs_runtime.Value = dictGenericTop_0_loop
_ = dictGenericTop_0
var dictGenericTop1_1 gopurs_runtime.Value = dictGenericTop1_1_loop
_ = dictGenericTop1_1
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericTop_0, "genericTop'"), gopurs_runtime.RecordGet(dictGenericTop1_1, "genericTop'")})}))})
}

func Call_genericTopSum(dictGenericTop_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericTop_0 gopurs_runtime.Value = dictGenericTop_0_loop
_ = dictGenericTop_0
return gopurs_runtime.RecordDict1("genericTop'", gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericTop_0, "genericTop'")})})
}

func Call_genericTop(dictGeneric_0_loop gopurs_runtime.Value, dictGenericTop_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericTop_1 gopurs_runtime.Value = dictGenericTop_1_loop
_ = dictGenericTop_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericTop_1, "genericTop'"))
}

func Call_genericBottomArgument(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.RecordGet(dictBounded_0, "bottom"))
}

func Call_genericBottom_prime(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericBottom'")
}

func Call_genericBottom_prime__gopurs_runtime_Value_2482077850(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "genericBottom'")
}

func Call_genericBottomConstructor(dictGenericBottom_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBottom_0 gopurs_runtime.Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.RecordGet(dictGenericBottom_0, "genericBottom'"))
}

func Call_genericBottomProduct(dictGenericBottom_0_loop gopurs_runtime.Value, dictGenericBottom1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBottom_0 gopurs_runtime.Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
var dictGenericBottom1_1 gopurs_runtime.Value = dictGenericBottom1_1_loop
_ = dictGenericBottom1_1
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericBottom_0, "genericBottom'"), gopurs_runtime.RecordGet(dictGenericBottom1_1, "genericBottom'")})}))})
}

func Call_genericBottomSum(dictGenericBottom_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericBottom_0 gopurs_runtime.Value = dictGenericBottom_0_loop
_ = dictGenericBottom_0
return gopurs_runtime.RecordDict1("genericBottom'", gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&pkg_Data_Generic_Rep.Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(dictGenericBottom_0, "genericBottom'")})})
}

func Call_genericBottom(dictGeneric_0_loop gopurs_runtime.Value, dictGenericBottom_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericBottom_1 gopurs_runtime.Value = dictGenericBottom_1_loop
_ = dictGenericBottom_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "to"), gopurs_runtime.RecordGet(dictGenericBottom_1, "genericBottom'"))
}


