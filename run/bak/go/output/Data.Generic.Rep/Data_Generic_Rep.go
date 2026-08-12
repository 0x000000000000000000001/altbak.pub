package Data_Generic_Rep

import (
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Symbol "gopurs/output/Data.Symbol"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Inl gopurs_runtime.Value
var once_Inl sync.Once
func Get_Inl() gopurs_runtime.Value {
	once_Inl.Do(func() {
		cache_Inl = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3478632216, UnsafePtr: unsafe.Pointer(&Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Inl
}

var cache_Inr gopurs_runtime.Value
var once_Inr sync.Once
func Get_Inr() gopurs_runtime.Value {
	once_Inr.Do(func() {
		cache_Inr = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 492034566, UnsafePtr: unsafe.Pointer(&Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Inr
}

var cache_Product gopurs_runtime.Value
var once_Product sync.Once
func Get_Product() gopurs_runtime.Value {
	once_Product.Do(func() {
		cache_Product = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1171963320, UnsafePtr: unsafe.Pointer(&Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Product
}

var cache_NoArguments gopurs_runtime.Value
var once_NoArguments sync.Once
func Get_NoArguments() gopurs_runtime.Value {
	once_NoArguments.Do(func() {
		cache_NoArguments = gopurs_runtime.Value{Type: 9, IntVal: int64(1454898258), UnsafePtr: nil}
	})
	return cache_NoArguments
}

var cache_Constructor gopurs_runtime.Value
var once_Constructor sync.Once
func Get_Constructor() gopurs_runtime.Value {
	once_Constructor.Do(func() {
		cache_Constructor = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Constructor(x_0_box)
})
	})
	return cache_Constructor
}

var cache_Argument gopurs_runtime.Value
var once_Argument sync.Once
func Get_Argument() gopurs_runtime.Value {
	once_Argument.Do(func() {
		cache_Argument = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Argument(x_0_box)
})
	})
	return cache_Argument
}

var cache_to gopurs_runtime.Value
var once_to sync.Once
func Get_to() gopurs_runtime.Value {
	once_to.Do(func() {
		cache_to = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_to(gopurs_runtime.CoerceToStruct[Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_to
}

var cache_showSum gopurs_runtime.Value
var once_showSum sync.Once
func Get_showSum() gopurs_runtime.Value {
	once_showSum.Do(func() {
		cache_showSum = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showSum(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showSum
}

var cache_showProduct gopurs_runtime.Value
var once_showProduct sync.Once
func Get_showProduct() gopurs_runtime.Value {
	once_showProduct.Do(func() {
		cache_showProduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showProduct(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showProduct
}

var cache_showNoArguments gopurs_runtime.Value
var once_showNoArguments sync.Once
func Get_showNoArguments() gopurs_runtime.Value {
	once_showNoArguments.Do(func() {
		cache_showNoArguments = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("NoArguments")
}))
	})
	return cache_showNoArguments
}

var cache_showConstructor gopurs_runtime.Value
var once_showConstructor sync.Once
func Get_showConstructor() gopurs_runtime.Value {
	once_showConstructor.Do(func() {
		cache_showConstructor = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictShow_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showConstructor(dictIsSymbol_0_box, dictShow_1_box)
})
	})
	return cache_showConstructor
}

var cache_showArgument gopurs_runtime.Value
var once_showArgument sync.Once
func Get_showArgument() gopurs_runtime.Value {
	once_showArgument.Do(func() {
		cache_showArgument = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showArgument(dictShow_0_box)
})
	})
	return cache_showArgument
}

var cache_repOf gopurs_runtime.Value
var once_repOf sync.Once
func Get_repOf() gopurs_runtime.Value {
	once_repOf.Do(func() {
		cache_repOf = gopurs_runtime.Func2(func(dictGeneric_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_repOf(gopurs_runtime.CoerceToStruct[Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dictGeneric_0_box), uint32(v_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_repOf
}

var cache_from gopurs_runtime.Value
var once_from sync.Once
func Get_from() gopurs_runtime.Value {
	once_from.Do(func() {
		cache_from = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_from(gopurs_runtime.CoerceToStruct[Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_from
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_show__3756561682 gopurs_runtime.Value
var once_show__3756561682 sync.Once
func Get_show__3756561682() gopurs_runtime.Value {
	once_show__3756561682.Do(func() {
		cache_show__3756561682 = gopurs_runtime.RecordGet(pkg_Data_Show.Get_showString(), "show")
	})
	return cache_show__3756561682
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_reflectSymbol__3416619207 gopurs_runtime.Value
var once_reflectSymbol__3416619207 sync.Once
func Get_reflectSymbol__3416619207() gopurs_runtime.Value {
	once_reflectSymbol__3416619207.Do(func() {
		cache_reflectSymbol__3416619207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__3416619207(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__3416619207
}

var cache_reflectSymbol__1166932993 gopurs_runtime.Value
var once_reflectSymbol__1166932993 sync.Once
func Get_reflectSymbol__1166932993() gopurs_runtime.Value {
	once_reflectSymbol__1166932993.Do(func() {
		cache_reflectSymbol__1166932993 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__1166932993(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__1166932993
}

type Constructor_Inl[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Inr[T_a any, T_b any] struct {
	Rc uint32
	V0 T_b
}


type Constructor_Product[T_a any, T_b any] struct {
	Rc uint32
	V0 T_a
	V1 T_b
}


type Constructor_NoArguments struct {
	Rc uint32
}


type Constructor_Generic[T_a any, T_rep any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1921946594] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "from": return c.V0
		case "to": return c.V1
		default: panic("Key not found in dictionary Constructor_Generic: " + key)
		}
	}
}


func Call_Constructor(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Argument(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_to(dict_0_loop *Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_showSum(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3478632216) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Inl "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Inl[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Str(")"))).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 492034566) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Inr "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Inr[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Str(")"))).StrVal())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
}

func Call_showProduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Product "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Product[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Str(")"))))).StrVal())
}))
}

func Call_showConstructor(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Constructor @"), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(Get_show__3756561682(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1, "show"), v_2), gopurs_runtime.Str(")"))))).StrVal())
}))
}

func Call_showArgument(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Argument "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_repOf(dictGeneric_0_loop *Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value], v_1_loop uint32) uint32 {
var dictGeneric_0 *Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dictGeneric_0_loop
_ = dictGeneric_0
var v_1 uint32 = v_1_loop
_ = v_1
return 513803634
}

func Call_from(dict_0_loop *Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Generic[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__3416619207(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__1166932993(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


