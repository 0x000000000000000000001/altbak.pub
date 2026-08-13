package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Divisible_Divisible_dollarDict gopurs_runtime.Value
var once_Data_Divisible_Divisible_dollarDict sync.Once
func Get_Data_Divisible_Divisible_dollarDict() gopurs_runtime.Value {
	once_Data_Divisible_Divisible_dollarDict.Do(func() {
		cache_Data_Divisible_Divisible_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divisible_Divisible_dollarDict(x_0_box)
})
	})
	return cache_Data_Divisible_Divisible_dollarDict
}

var cache_Data_Divisible_divisiblePredicate gopurs_runtime.Value
var once_Data_Divisible_divisiblePredicate sync.Once
func Get_Data_Divisible_divisiblePredicate() gopurs_runtime.Value {
	once_Data_Divisible_divisiblePredicate.Do(func() {
		cache_Data_Divisible_divisiblePredicate = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divide_dividePredicate()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_Data_Divisible_divisiblePredicate
}

var cache_Data_Divisible_divisibleOp gopurs_runtime.Value
var once_Data_Divisible_divisibleOp sync.Once
func Get_Data_Divisible_divisibleOp() gopurs_runtime.Value {
	once_Data_Divisible_divisibleOp.Do(func() {
		cache_Data_Divisible_divisibleOp = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divisible_divisibleOp(dictMonoid_0_box)
})
	})
	return cache_Data_Divisible_divisibleOp
}

var cache_Data_Divisible_divisibleEquivalence gopurs_runtime.Value
var once_Data_Divisible_divisibleEquivalence sync.Once
func Get_Data_Divisible_divisibleEquivalence() gopurs_runtime.Value {
	once_Data_Divisible_divisibleEquivalence.Do(func() {
		cache_Data_Divisible_divisibleEquivalence = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divide_divideEquivalence()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Divisible_divisibleEquivalence
}

var cache_Data_Divisible_divisibleComparison gopurs_runtime.Value
var once_Data_Divisible_divisibleComparison sync.Once
func Get_Data_Divisible_divisibleComparison() gopurs_runtime.Value {
	once_Data_Divisible_divisibleComparison.Do(func() {
		cache_Data_Divisible_divisibleComparison = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divide_divideComparison()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Divisible_divisibleComparison
}

var cache_Data_Divisible_conquer gopurs_runtime.Value
var once_Data_Divisible_conquer sync.Once
func Get_Data_Divisible_conquer() gopurs_runtime.Value {
	once_Data_Divisible_conquer.Do(func() {
		cache_Data_Divisible_conquer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Divisible_conquer(dict_0_box)
})
	})
	return cache_Data_Divisible_conquer
}

var cache_Data_Divisible_divisibleComparison__661164760 gopurs_runtime.Value
var once_Data_Divisible_divisibleComparison__661164760 sync.Once
func Get_Data_Divisible_divisibleComparison__661164760() gopurs_runtime.Value {
	once_Data_Divisible_divisibleComparison__661164760.Do(func() {
		cache_Data_Divisible_divisibleComparison__661164760 = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divide_divideComparison()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Divisible_divisibleComparison__661164760
}

var cache_Data_Divisible_divisibleEquivalence__4236776696 gopurs_runtime.Value
var once_Data_Divisible_divisibleEquivalence__4236776696 sync.Once
func Get_Data_Divisible_divisibleEquivalence__4236776696() gopurs_runtime.Value {
	once_Data_Divisible_divisibleEquivalence__4236776696.Do(func() {
		cache_Data_Divisible_divisibleEquivalence__4236776696 = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divide_divideEquivalence()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Divisible_divisibleEquivalence__4236776696
}

var cache_Data_Divisible_divisiblePredicate__1930744184 gopurs_runtime.Value
var once_Data_Divisible_divisiblePredicate__1930744184 sync.Once
func Get_Data_Divisible_divisiblePredicate__1930744184() gopurs_runtime.Value {
	once_Data_Divisible_divisiblePredicate__1930744184.Do(func() {
		cache_Data_Divisible_divisiblePredicate__1930744184 = gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Divide_dividePredicate()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_Data_Divisible_divisiblePredicate__1930744184
}

type Constructor_Data_Divisible_Divisible[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2265116602] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Divisible_Divisible[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Divide0": return gopurs_runtime.Box(c.V0)
		case "conquer": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Divisible_Divisible: " + key)
		}
	}
}


func Call_Data_Divisible_Divisible_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Divisible_divisibleOp(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): divideOp_1_0 -> gopurs_runtime.Value
divideOp_1_0 := gopurs_runtime.Apply(Get_Data_Divide_divideOp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = divideOp_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("Divide0", "conquer", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return divideOp_1_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}))
}

func Call_Data_Divisible_conquer(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "conquer")
}


