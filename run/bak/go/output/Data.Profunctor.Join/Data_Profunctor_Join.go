package Data_Profunctor_Join

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Profunctor "gopurs/output/Data.Profunctor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Join gopurs_runtime.Value
var once_Join sync.Once
func Get_Join() gopurs_runtime.Value {
	once_Join.Do(func() {
		cache_Join = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Join(x_0_box)
})
	})
	return cache_Join
}

var cache_showJoin gopurs_runtime.Value
var once_showJoin sync.Once
func Get_showJoin() gopurs_runtime.Value {
	once_showJoin.Do(func() {
		cache_showJoin = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showJoin(dictShow_0_box)
})
	})
	return cache_showJoin
}

var cache_semigroupJoin gopurs_runtime.Value
var once_semigroupJoin sync.Once
func Get_semigroupJoin() gopurs_runtime.Value {
	once_semigroupJoin.Do(func() {
		cache_semigroupJoin = gopurs_runtime.Func(func(dictSemigroupoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupJoin(dictSemigroupoid_0_box)
})
	})
	return cache_semigroupJoin
}

var cache_ordJoin gopurs_runtime.Value
var once_ordJoin sync.Once
func Get_ordJoin() gopurs_runtime.Value {
	once_ordJoin.Do(func() {
		cache_ordJoin = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordJoin(dictOrd_0_box)
})
	})
	return cache_ordJoin
}

var cache_newtypeJoin gopurs_runtime.Value
var once_newtypeJoin sync.Once
func Get_newtypeJoin() gopurs_runtime.Value {
	once_newtypeJoin.Do(func() {
		cache_newtypeJoin = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeJoin
}

var cache_monoidJoin gopurs_runtime.Value
var once_monoidJoin sync.Once
func Get_monoidJoin() gopurs_runtime.Value {
	once_monoidJoin.Do(func() {
		cache_monoidJoin = gopurs_runtime.Func(func(dictCategory_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidJoin(dictCategory_0_box)
})
	})
	return cache_monoidJoin
}

var cache_invariantJoin gopurs_runtime.Value
var once_invariantJoin sync.Once
func Get_invariantJoin() gopurs_runtime.Value {
	once_invariantJoin.Do(func() {
		cache_invariantJoin = gopurs_runtime.Func(func(dictProfunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_invariantJoin(dictProfunctor_0_box)
})
	})
	return cache_invariantJoin
}

var cache_eqJoin gopurs_runtime.Value
var once_eqJoin sync.Once
func Get_eqJoin() gopurs_runtime.Value {
	once_eqJoin.Do(func() {
		cache_eqJoin = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqJoin(dictEq_0_box)
})
	})
	return cache_eqJoin
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_dimap__1466332548 gopurs_runtime.Value
var once_dimap__1466332548 sync.Once
func Get_dimap__1466332548() gopurs_runtime.Value {
	once_dimap__1466332548.Do(func() {
		cache_dimap__1466332548 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dimap__1466332548(gopurs_runtime.CoerceToStruct[pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_dimap__1466332548
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

var cache_show__255526802 gopurs_runtime.Value
var once_show__255526802 sync.Once
func Get_show__255526802() gopurs_runtime.Value {
	once_show__255526802.Do(func() {
		cache_show__255526802 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__255526802(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__255526802
}

func Call_Join(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showJoin(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(Join "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semigroupJoin(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), v_1, v1_2)
})
}))
}

func Call_ordJoin(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_monoidJoin(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupJoin1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "compose"), v_2, v1_3)
})
}))
_ = semigroupJoin1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupJoin1_1_0
}), gopurs_runtime.RecordGet(dictCategory_0, "identity"))
}

func Call_invariantJoin(dictProfunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), g_2, f_1, v_3)
})
})
}))
}

func Call_eqJoin(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_dimap__1466332548(dict_0_loop *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Profunctor.Constructor_Profunctor[gopurs_runtime.Value] = dict_0_loop
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

func Call_show__255526802(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


