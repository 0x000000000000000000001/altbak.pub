package Data_Generic_Rep

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var Inl gopurs_runtime.Value
var once_Inl sync.Once
func Get_Inl() gopurs_runtime.Value {
	once_Inl.Do(func() {
		Inl = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Inl", value0)
})
	})
	return Inl
}

var Inr gopurs_runtime.Value
var once_Inr sync.Once
func Get_Inr() gopurs_runtime.Value {
	once_Inr.Do(func() {
		Inr = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor1("Inr", value0)
})
	})
	return Inr
}

var Product gopurs_runtime.Value
var once_Product sync.Once
func Get_Product() gopurs_runtime.Value {
	once_Product.Do(func() {
		Product = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Product", value0, value1)
})
})
	})
	return Product
}

var NoArguments gopurs_runtime.Value
var once_NoArguments sync.Once
func Get_NoArguments() gopurs_runtime.Value {
	once_NoArguments.Do(func() {
		NoArguments = gopurs_runtime.Constructor0("NoArguments")
	})
	return NoArguments
}

var Constructor gopurs_runtime.Value
var once_Constructor sync.Once
func Get_Constructor() gopurs_runtime.Value {
	once_Constructor.Do(func() {
		Constructor = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Constructor
}

var Argument gopurs_runtime.Value
var once_Argument sync.Once
func Get_Argument() gopurs_runtime.Value {
	once_Argument.Do(func() {
		Argument = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return Argument
}

var to gopurs_runtime.Value
var once_to sync.Once
func Get_to() gopurs_runtime.Value {
	once_to.Do(func() {
		to = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "to")
}()
})
	})
	return to
}

var showSum gopurs_runtime.Value
var once_showSum sync.Once
func Get_showSum() gopurs_runtime.Value {
	once_showSum.Do(func() {
		showSum = gopurs_runtime.Func2(Call_showSum)
	})
	return showSum
}

var showProduct gopurs_runtime.Value
var once_showProduct sync.Once
func Get_showProduct() gopurs_runtime.Value {
	once_showProduct.Do(func() {
		showProduct = gopurs_runtime.Func2(Call_showProduct)
	})
	return showProduct
}

var showNoArguments gopurs_runtime.Value
var once_showNoArguments sync.Once
func Get_showNoArguments() gopurs_runtime.Value {
	once_showNoArguments.Do(func() {
		showNoArguments = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("NoArguments")
}))
	})
	return showNoArguments
}

var showConstructor gopurs_runtime.Value
var once_showConstructor sync.Once
func Get_showConstructor() gopurs_runtime.Value {
	once_showConstructor.Do(func() {
		showConstructor = gopurs_runtime.Func2(Call_showConstructor)
	})
	return showConstructor
}

var showArgument gopurs_runtime.Value
var once_showArgument sync.Once
func Get_showArgument() gopurs_runtime.Value {
	once_showArgument.Do(func() {
		showArgument = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Argument " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), v_1).StrVal + ")")
}))
}()
})
	})
	return showArgument
}

var repOf gopurs_runtime.Value
var once_repOf sync.Once
func Get_repOf() gopurs_runtime.Value {
	once_repOf.Do(func() {
		repOf = gopurs_runtime.Func2(Call_repOf)
	})
	return repOf
}

var from gopurs_runtime.Value
var once_from sync.Once
func Get_from() gopurs_runtime.Value {
	once_from.Do(func() {
		from = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "from")
}()
})
	})
	return from
}

func Call_showSum(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Inl").IntVal != 0 {
__t0 = gopurs_runtime.Str("(Inl " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + ")")
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v_2.StrVal == "Inr").IntVal != 0 {
__t0 = gopurs_runtime.Str("(Inr " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + ")")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}

func Call_showProduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Product " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0]).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1_loop, "show"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1]).StrVal + ")")
}))
}

func Call_showConstructor(dictIsSymbol_0_loop gopurs_runtime.Value, dictShow_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictShow_1 gopurs_runtime.Value = dictShow_1_loop
_ = dictShow_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(Constructor @" + gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0_loop, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))).StrVal + " " + gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_1_loop, "show"), v_2).StrVal + ")")
}))
}

func Call_repOf(dictGeneric_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Constructor0("Proxy")
}


