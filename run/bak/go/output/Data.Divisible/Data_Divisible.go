package Data_Divisible

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Divide "gopurs/output/Data.Divide"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	unsafe "unsafe"
)

var divisiblePredicate gopurs_runtime.Value
var once_divisiblePredicate sync.Once
func Get_divisiblePredicate() gopurs_runtime.Value {
	once_divisiblePredicate.Do(func() {
		divisiblePredicate = gopurs_runtime.RecordDict2("conquer", "Divide0", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_dividePredicate()
}))
	})
	return divisiblePredicate
}

var divisibleOp gopurs_runtime.Value
var once_divisibleOp sync.Once
func Get_divisibleOp() gopurs_runtime.Value {
	once_divisibleOp.Do(func() {
		divisibleOp = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
divideOp_1_0 := gopurs_runtime.Apply(pkg_Data_Divide.Get_divideOp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = divideOp_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("conquer", "Divide0", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return divideOp_1_0
}))
}()
})
	})
	return divisibleOp
}

var divisibleEquivalence gopurs_runtime.Value
var once_divisibleEquivalence sync.Once
func Get_divisibleEquivalence() gopurs_runtime.Value {
	once_divisibleEquivalence.Do(func() {
		divisibleEquivalence = gopurs_runtime.RecordDict2("conquer", "Divide0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideEquivalence()
}))
	})
	return divisibleEquivalence
}

var divisibleComparison gopurs_runtime.Value
var once_divisibleComparison sync.Once
func Get_divisibleComparison() gopurs_runtime.Value {
	once_divisibleComparison.Do(func() {
		divisibleComparison = gopurs_runtime.RecordDict2("conquer", "Divide0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1111389260, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideComparison()
}))
	})
	return divisibleComparison
}

var conquer gopurs_runtime.Value
var once_conquer sync.Once
func Get_conquer() gopurs_runtime.Value {
	once_conquer.Do(func() {
		conquer = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "conquer")
}()
})
	})
	return conquer
}




