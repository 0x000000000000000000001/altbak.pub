package Data_Eq_Generic

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var genericEqNoConstructors gopurs_runtime.Value
var once_genericEqNoConstructors sync.Once
func Get_genericEqNoConstructors() gopurs_runtime.Value {
	once_genericEqNoConstructors.Do(func() {
		genericEqNoConstructors = gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(true)
}))
	})
	return genericEqNoConstructors
}

var genericEqNoArguments gopurs_runtime.Value
var once_genericEqNoArguments sync.Once
func Get_genericEqNoArguments() gopurs_runtime.Value {
	once_genericEqNoArguments.Do(func() {
		genericEqNoArguments = gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(true)
}))
	})
	return genericEqNoArguments
}

var genericEqArgument gopurs_runtime.Value
var once_genericEqArgument sync.Once
func Get_genericEqArgument() gopurs_runtime.Value {
	once_genericEqArgument.Do(func() {
		genericEqArgument = gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), v_1, v1_2)
}))
})
	})
	return genericEqArgument
}

var genericEq_prime gopurs_runtime.Value
var once_genericEq_prime sync.Once
func Get_genericEq_prime() gopurs_runtime.Value {
	once_genericEq_prime.Do(func() {
		genericEq_prime = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "genericEq'")
})
	})
	return genericEq_prime
}

var genericEqConstructor gopurs_runtime.Value
var once_genericEqConstructor sync.Once
func Get_genericEqConstructor() gopurs_runtime.Value {
	once_genericEqConstructor.Do(func() {
		genericEqConstructor = gopurs_runtime.Func(func(dictGenericEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), v_1, v1_2)
}))
})
	})
	return genericEqConstructor
}

var genericEqProduct gopurs_runtime.Value
var once_genericEqProduct sync.Once
func Get_genericEqProduct() gopurs_runtime.Value {
	once_genericEqProduct.Do(func() {
		genericEqProduct = gopurs_runtime.Func2(func(dictGenericEq_0 gopurs_runtime.Value, dictGenericEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]).IntVal != 0)
}))
})
	})
	return genericEqProduct
}

var genericEqSum gopurs_runtime.Value
var once_genericEqSum sync.Once
func Get_genericEqSum() gopurs_runtime.Value {
	once_genericEqSum.Do(func() {
		genericEqSum = gopurs_runtime.Func2(func(dictGenericEq_0 gopurs_runtime.Value, dictGenericEq1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Inl").IntVal != 0 {
__t0 = gopurs_runtime.Boolean(gopurs_runtime.Bool(v1_3.StrVal == "Inl").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Boolean(gopurs_runtime.Bool(v_2.StrVal == "Inr").IntVal != 0 && gopurs_runtime.Bool(v1_3.StrVal == "Inr").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0)
}
end_branch_0:
return __t0
}))
})
	})
	return genericEqSum
}

var genericEq gopurs_runtime.Value
var once_genericEq sync.Once
func Get_genericEq() gopurs_runtime.Value {
	once_genericEq.Do(func() {
		genericEq = gopurs_runtime.Func4(func(dictGeneric_0 gopurs_runtime.Value, dictGenericEq_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_1, "genericEq'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), x_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0, "from"), y_3))
})
	})
	return genericEq
}




