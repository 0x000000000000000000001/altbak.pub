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
return gopurs_runtime.Bool(true)
}))
	})
	return genericEqNoConstructors
}

var genericEqNoArguments gopurs_runtime.Value
var once_genericEqNoArguments sync.Once
func Get_genericEqNoArguments() gopurs_runtime.Value {
	once_genericEqNoArguments.Do(func() {
		genericEqNoArguments = gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return genericEqNoArguments
}

var genericEqArgument gopurs_runtime.Value
var once_genericEqArgument sync.Once
func Get_genericEqArgument() gopurs_runtime.Value {
	once_genericEqArgument.Do(func() {
		genericEqArgument = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0_loop, "eq"), v_1, v1_2)
}))
}()
})
	})
	return genericEqArgument
}

var genericEq_prime gopurs_runtime.Value
var once_genericEq_prime sync.Once
func Get_genericEq_prime() gopurs_runtime.Value {
	once_genericEq_prime.Do(func() {
		genericEq_prime = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "genericEq'")
}()
})
	})
	return genericEq_prime
}

var genericEqConstructor gopurs_runtime.Value
var once_genericEqConstructor sync.Once
func Get_genericEqConstructor() gopurs_runtime.Value {
	once_genericEqConstructor.Do(func() {
		genericEqConstructor = gopurs_runtime.Func(func(dictGenericEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0_loop, "genericEq'"), v_1, v1_2)
}))
}()
})
	})
	return genericEqConstructor
}

var genericEqProduct gopurs_runtime.Value
var once_genericEqProduct sync.Once
func Get_genericEqProduct() gopurs_runtime.Value {
	once_genericEqProduct.Do(func() {
		genericEqProduct = gopurs_runtime.Func2(Call_genericEqProduct)
	})
	return genericEqProduct
}

var genericEqSum gopurs_runtime.Value
var once_genericEqSum sync.Once
func Get_genericEqSum() gopurs_runtime.Value {
	once_genericEqSum.Do(func() {
		genericEqSum = gopurs_runtime.Func2(Call_genericEqSum)
	})
	return genericEqSum
}

var genericEq gopurs_runtime.Value
var once_genericEq sync.Once
func Get_genericEq() gopurs_runtime.Value {
	once_genericEq.Do(func() {
		genericEq = gopurs_runtime.Func4(Call_genericEq)
	})
	return genericEq
}

func Call_genericEqProduct(dictGenericEq_0_loop gopurs_runtime.Value, dictGenericEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
var dictGenericEq1_1 gopurs_runtime.Value = dictGenericEq1_1_loop
_ = dictGenericEq1_1
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0_loop, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1_loop, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[1], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[1]).IntVal != 0)
}))
}

func Call_genericEqSum(dictGenericEq_0_loop gopurs_runtime.Value, dictGenericEq1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGenericEq_0 gopurs_runtime.Value = dictGenericEq_0_loop
_ = dictGenericEq_0
var dictGenericEq1_1 gopurs_runtime.Value = dictGenericEq1_1_loop
_ = dictGenericEq1_1
return gopurs_runtime.RecordDict1("genericEq'", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_2.StrVal == "Inl").IntVal != 0 {
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(v1_3.StrVal == "Inl").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_0_loop, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(gopurs_runtime.Bool(v_2.StrVal == "Inr").IntVal != 0 && gopurs_runtime.Bool(v1_3.StrVal == "Inr").IntVal != 0 && gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq1_1_loop, "genericEq'"), (*[1024]gopurs_runtime.Value)(v_2.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v1_3.UnsafePtr)[0]).IntVal != 0)
}
end_branch_0:
return __t0
}))
}

func Call_genericEq(dictGeneric_0_loop gopurs_runtime.Value, dictGenericEq_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictGeneric_0 gopurs_runtime.Value = dictGeneric_0_loop
_ = dictGeneric_0
var dictGenericEq_1 gopurs_runtime.Value = dictGenericEq_1_loop
_ = dictGenericEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictGenericEq_1_loop, "genericEq'"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), x_2_loop), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictGeneric_0_loop, "from"), y_3_loop))
}


