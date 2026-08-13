package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Generic_Rep_Inl gopurs_runtime.Value
var once_Data_Generic_Rep_Inl sync.Once
func Get_Data_Generic_Rep_Inl() gopurs_runtime.Value {
	once_Data_Generic_Rep_Inl.Do(func() {
		cache_Data_Generic_Rep_Inl = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inl{1, value0})}
})
	})
	return cache_Data_Generic_Rep_Inl
}

var cache_Data_Generic_Rep_Inr gopurs_runtime.Value
var once_Data_Generic_Rep_Inr sync.Once
func Get_Data_Generic_Rep_Inr() gopurs_runtime.Value {
	once_Data_Generic_Rep_Inr.Do(func() {
		cache_Data_Generic_Rep_Inr = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Inr{1, value0})}
})
	})
	return cache_Data_Generic_Rep_Inr
}

var cache_Data_Generic_Rep_Product gopurs_runtime.Value
var once_Data_Generic_Rep_Product sync.Once
func Get_Data_Generic_Rep_Product() gopurs_runtime.Value {
	once_Data_Generic_Rep_Product.Do(func() {
		cache_Data_Generic_Rep_Product = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Data_Generic_Rep_Product{1, value0, value1})}
})
})
	})
	return cache_Data_Generic_Rep_Product
}

var cache_Data_Generic_Rep_NoConstructors gopurs_runtime.Value
var once_Data_Generic_Rep_NoConstructors sync.Once
func Get_Data_Generic_Rep_NoConstructors() gopurs_runtime.Value {
	once_Data_Generic_Rep_NoConstructors.Do(func() {
		cache_Data_Generic_Rep_NoConstructors = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_NoConstructors(x_0_box)
})
	})
	return cache_Data_Generic_Rep_NoConstructors
}

var cache_Data_Generic_Rep_NoArguments gopurs_runtime.Value
var once_Data_Generic_Rep_NoArguments sync.Once
func Get_Data_Generic_Rep_NoArguments() gopurs_runtime.Value {
	once_Data_Generic_Rep_NoArguments.Do(func() {
		cache_Data_Generic_Rep_NoArguments = gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
	})
	return cache_Data_Generic_Rep_NoArguments
}

var cache_Data_Generic_Rep_Generic_dollarDict gopurs_runtime.Value
var once_Data_Generic_Rep_Generic_dollarDict sync.Once
func Get_Data_Generic_Rep_Generic_dollarDict() gopurs_runtime.Value {
	once_Data_Generic_Rep_Generic_dollarDict.Do(func() {
		cache_Data_Generic_Rep_Generic_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_Generic_dollarDict(x_0_box)
})
	})
	return cache_Data_Generic_Rep_Generic_dollarDict
}

var cache_Data_Generic_Rep_Constructor gopurs_runtime.Value
var once_Data_Generic_Rep_Constructor sync.Once
func Get_Data_Generic_Rep_Constructor() gopurs_runtime.Value {
	once_Data_Generic_Rep_Constructor.Do(func() {
		cache_Data_Generic_Rep_Constructor = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_Constructor(x_0_box)
})
	})
	return cache_Data_Generic_Rep_Constructor
}

var cache_Data_Generic_Rep_Argument gopurs_runtime.Value
var once_Data_Generic_Rep_Argument sync.Once
func Get_Data_Generic_Rep_Argument() gopurs_runtime.Value {
	once_Data_Generic_Rep_Argument.Do(func() {
		cache_Data_Generic_Rep_Argument = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_Argument(x_0_box)
})
	})
	return cache_Data_Generic_Rep_Argument
}

var cache_Data_Generic_Rep_to gopurs_runtime.Value
var once_Data_Generic_Rep_to sync.Once
func Get_Data_Generic_Rep_to() gopurs_runtime.Value {
	once_Data_Generic_Rep_to.Do(func() {
		cache_Data_Generic_Rep_to = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_to(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dict_0_box))
})
	})
	return cache_Data_Generic_Rep_to
}

var cache_Data_Generic_Rep_showSum gopurs_runtime.Value
var once_Data_Generic_Rep_showSum sync.Once
func Get_Data_Generic_Rep_showSum() gopurs_runtime.Value {
	once_Data_Generic_Rep_showSum.Do(func() {
		cache_Data_Generic_Rep_showSum = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_showSum(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Generic_Rep_showSum
}

var cache_Data_Generic_Rep_showProduct gopurs_runtime.Value
var once_Data_Generic_Rep_showProduct sync.Once
func Get_Data_Generic_Rep_showProduct() gopurs_runtime.Value {
	once_Data_Generic_Rep_showProduct.Do(func() {
		cache_Data_Generic_Rep_showProduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_showProduct(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Generic_Rep_showProduct
}

var cache_Data_Generic_Rep_showNoArguments gopurs_runtime.Value
var once_Data_Generic_Rep_showNoArguments sync.Once
func Get_Data_Generic_Rep_showNoArguments() gopurs_runtime.Value {
	once_Data_Generic_Rep_showNoArguments.Do(func() {
		cache_Data_Generic_Rep_showNoArguments = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("NoArguments")
}))
	})
	return cache_Data_Generic_Rep_showNoArguments
}

var cache_Data_Generic_Rep_showConstructor gopurs_runtime.Value
var once_Data_Generic_Rep_showConstructor sync.Once
func Get_Data_Generic_Rep_showConstructor() gopurs_runtime.Value {
	once_Data_Generic_Rep_showConstructor.Do(func() {
		cache_Data_Generic_Rep_showConstructor = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_showConstructor(dictIsSymbol_0_box, dictShow_1_box)
})
	})
	return cache_Data_Generic_Rep_showConstructor
}

var cache_Data_Generic_Rep_showArgument gopurs_runtime.Value
var once_Data_Generic_Rep_showArgument sync.Once
func Get_Data_Generic_Rep_showArgument() gopurs_runtime.Value {
	once_Data_Generic_Rep_showArgument.Do(func() {
		cache_Data_Generic_Rep_showArgument = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_showArgument(dictShow_0_box)
})
	})
	return cache_Data_Generic_Rep_showArgument
}

var cache_Data_Generic_Rep_repOf gopurs_runtime.Value
var once_Data_Generic_Rep_repOf sync.Once
func Get_Data_Generic_Rep_repOf() gopurs_runtime.Value {
	once_Data_Generic_Rep_repOf.Do(func() {
		cache_Data_Generic_Rep_repOf = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Generic_Rep_repOf(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dictGeneric_0_box), uint32(v_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_Data_Generic_Rep_repOf
}

var cache_Data_Generic_Rep_from gopurs_runtime.Value
var once_Data_Generic_Rep_from sync.Once
func Get_Data_Generic_Rep_from() gopurs_runtime.Value {
	once_Data_Generic_Rep_from.Do(func() {
		cache_Data_Generic_Rep_from = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_from(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dict_0_box))
})
	})
	return cache_Data_Generic_Rep_from
}

var cache_Data_Generic_Rep_from__1498760952 gopurs_runtime.Value
var once_Data_Generic_Rep_from__1498760952 sync.Once
func Get_Data_Generic_Rep_from__1498760952() gopurs_runtime.Value {
	once_Data_Generic_Rep_from__1498760952.Do(func() {
		cache_Data_Generic_Rep_from__1498760952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_from__1498760952(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dict_0_box))
})
	})
	return cache_Data_Generic_Rep_from__1498760952
}

var cache_Data_Generic_Rep_to__1498760952 gopurs_runtime.Value
var once_Data_Generic_Rep_to__1498760952 sync.Once
func Get_Data_Generic_Rep_to__1498760952() gopurs_runtime.Value {
	once_Data_Generic_Rep_to__1498760952.Do(func() {
		cache_Data_Generic_Rep_to__1498760952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Generic_Rep_to__1498760952(gopurs_runtime.CoerceToStruct[Constructor_Data_Generic_Rep_Generic](dict_0_box))
})
	})
	return cache_Data_Generic_Rep_to__1498760952
}

type Constructor_Data_Generic_Rep_Inl struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Data_Generic_Rep_Inr struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Data_Generic_Rep_Product struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Data_Generic_Rep_NoArguments struct {
	Rc uint32
}


type Constructor_Data_Generic_Rep_Generic struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1921946594] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Generic_Rep_Generic)(ptr)
		_ = c
		switch key {
		case "from": return gopurs_runtime.Box(c.V0)
		case "to": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Generic_Rep_Generic: " + key)
		}
	}
}


func Call_Data_Generic_Rep_NoConstructors(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Generic_Rep_Generic_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Generic_Rep_Constructor(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Generic_Rep_Argument(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Generic_Rep_to(dict_0_loop *Constructor_Data_Generic_Rep_Generic) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Generic_Rep_Generic = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Generic_Rep_showSum(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
__t0 = (("(Inl ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Generic_Rep_Inl)(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 492034566) {
__t0 = (("(Inr ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Generic_Rep_Inr)(v_2.UnsafePtr).V0).StrVal())) + (")")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
}))
}

func Call_Data_Generic_Rep_showProduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Product ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V0).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Generic_Rep_Product)(v_2.UnsafePtr).V1).StrVal())) + (")"))
}))
}

func Call_Data_Generic_Rep_showConstructor(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((((("(Constructor @") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal())).StrVal())) + (" ")) + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), v_2).StrVal())) + (")"))
}))
}

func Call_Data_Generic_Rep_showArgument(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Argument ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Generic_Rep_repOf(dictGeneric_0_loop *Constructor_Data_Generic_Rep_Generic, v_1_loop uint32) uint32 {
var dictGeneric_0 *Constructor_Data_Generic_Rep_Generic = dictGeneric_0_loop
_ = dictGeneric_0
var v_1 uint32 = v_1_loop
_ = v_1
return 513803634
}

func Call_Data_Generic_Rep_from(dict_0_loop *Constructor_Data_Generic_Rep_Generic) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Generic_Rep_Generic = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Generic_Rep_from__1498760952(dict_0_loop *Constructor_Data_Generic_Rep_Generic) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Generic_Rep_Generic = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Generic_Rep_to__1498760952(dict_0_loop *Constructor_Data_Generic_Rep_Generic) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Generic_Rep_Generic = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}


