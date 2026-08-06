package Data_Divisible

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Divide "gopurs/output/Data.Divide"
)

var cache_divisiblePredicate gopurs_runtime.Value
var once_divisiblePredicate sync.Once
func Get_divisiblePredicate() gopurs_runtime.Value {
	once_divisiblePredicate.Do(func() {
		cache_divisiblePredicate = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_dividePredicate()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_divisiblePredicate
}

var cache_divisibleOp gopurs_runtime.Value
var once_divisibleOp sync.Once
func Get_divisibleOp() gopurs_runtime.Value {
	once_divisibleOp.Do(func() {
		cache_divisibleOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_divisibleOp(dictMonoid_0_box)
})
	})
	return cache_divisibleOp
}

var cache_divisibleEquivalence gopurs_runtime.Value
var once_divisibleEquivalence sync.Once
func Get_divisibleEquivalence() gopurs_runtime.Value {
	once_divisibleEquivalence.Do(func() {
		cache_divisibleEquivalence = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideEquivalence()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_divisibleEquivalence
}

var cache_divisibleComparison gopurs_runtime.Value
var once_divisibleComparison sync.Once
func Get_divisibleComparison() gopurs_runtime.Value {
	once_divisibleComparison.Do(func() {
		cache_divisibleComparison = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Divide.Get_divideComparison()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
})
}))
	})
	return cache_divisibleComparison
}

var cache_conquer gopurs_runtime.Value
var once_conquer sync.Once
func Get_conquer() gopurs_runtime.Value {
	once_conquer.Do(func() {
		cache_conquer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conquer(dict_0_box)
})
	})
	return cache_conquer
}

func Call_divisibleOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
divideOp_1_0 := gopurs_runtime.Apply(pkg_Data_Divide.Get_divideOp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = divideOp_1_0
__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return divideOp_1_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
}

func Call_conquer(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "conquer")
}


