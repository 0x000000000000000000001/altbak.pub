package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Join_Join gopurs_runtime.Value
var once_Data_Profunctor_Join_Join sync.Once
func Get_Data_Profunctor_Join_Join() gopurs_runtime.Value {
	once_Data_Profunctor_Join_Join.Do(func() {
		cache_Data_Profunctor_Join_Join = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Join_Join(x_0_box)
})
	})
	return cache_Data_Profunctor_Join_Join
}

var cache_Data_Profunctor_Join_showJoin gopurs_runtime.Value
var once_Data_Profunctor_Join_showJoin sync.Once
func Get_Data_Profunctor_Join_showJoin() gopurs_runtime.Value {
	once_Data_Profunctor_Join_showJoin.Do(func() {
		cache_Data_Profunctor_Join_showJoin = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Join_showJoin(dictShow_0_box)
})
	})
	return cache_Data_Profunctor_Join_showJoin
}

var cache_Data_Profunctor_Join_semigroupJoin gopurs_runtime.Value
var once_Data_Profunctor_Join_semigroupJoin sync.Once
func Get_Data_Profunctor_Join_semigroupJoin() gopurs_runtime.Value {
	once_Data_Profunctor_Join_semigroupJoin.Do(func() {
		cache_Data_Profunctor_Join_semigroupJoin = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Join_semigroupJoin(dictSemigroupoid_0_box)
})
	})
	return cache_Data_Profunctor_Join_semigroupJoin
}

var cache_Data_Profunctor_Join_ordJoin gopurs_runtime.Value
var once_Data_Profunctor_Join_ordJoin sync.Once
func Get_Data_Profunctor_Join_ordJoin() gopurs_runtime.Value {
	once_Data_Profunctor_Join_ordJoin.Do(func() {
		cache_Data_Profunctor_Join_ordJoin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Join_ordJoin(dictOrd_0_box)
})
	})
	return cache_Data_Profunctor_Join_ordJoin
}

var cache_Data_Profunctor_Join_newtypeJoin gopurs_runtime.Value
var once_Data_Profunctor_Join_newtypeJoin sync.Once
func Get_Data_Profunctor_Join_newtypeJoin() gopurs_runtime.Value {
	once_Data_Profunctor_Join_newtypeJoin.Do(func() {
		cache_Data_Profunctor_Join_newtypeJoin = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Profunctor_Join_newtypeJoin
}

var cache_Data_Profunctor_Join_monoidJoin gopurs_runtime.Value
var once_Data_Profunctor_Join_monoidJoin sync.Once
func Get_Data_Profunctor_Join_monoidJoin() gopurs_runtime.Value {
	once_Data_Profunctor_Join_monoidJoin.Do(func() {
		cache_Data_Profunctor_Join_monoidJoin = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Join_monoidJoin(dictCategory_0_box)
})
	})
	return cache_Data_Profunctor_Join_monoidJoin
}

var cache_Data_Profunctor_Join_invariantJoin gopurs_runtime.Value
var once_Data_Profunctor_Join_invariantJoin sync.Once
func Get_Data_Profunctor_Join_invariantJoin() gopurs_runtime.Value {
	once_Data_Profunctor_Join_invariantJoin.Do(func() {
		cache_Data_Profunctor_Join_invariantJoin = gopurs_runtime.Func(func(dictProfunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Join_invariantJoin(dictProfunctor_0_box)
})
	})
	return cache_Data_Profunctor_Join_invariantJoin
}

var cache_Data_Profunctor_Join_eqJoin gopurs_runtime.Value
var once_Data_Profunctor_Join_eqJoin sync.Once
func Get_Data_Profunctor_Join_eqJoin() gopurs_runtime.Value {
	once_Data_Profunctor_Join_eqJoin.Do(func() {
		cache_Data_Profunctor_Join_eqJoin = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Join_eqJoin(dictEq_0_box)
})
	})
	return cache_Data_Profunctor_Join_eqJoin
}

func Call_Data_Profunctor_Join_Join(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Join_showJoin(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Join ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})})}
}

func Call_Data_Profunctor_Join_semigroupJoin(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), v_1, v1_2)
})
})})}
}

func Call_Data_Profunctor_Join_ordJoin(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0))}
}

func Call_Data_Profunctor_Join_monoidJoin(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupJoin1_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupJoin1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compose"), v_2, v1_3)
})
})))
_ = semigroupJoin1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoin1_1_0)}
}), gopurs_runtime.RecordGet(dictCategory_0, "identity")})}
}

func Call_Data_Profunctor_Join_invariantJoin(dictProfunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Invariant_Invariant{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), g_2, f_1, v_3)
})
})
})})}
}

func Call_Data_Profunctor_Join_eqJoin(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0))}
}


