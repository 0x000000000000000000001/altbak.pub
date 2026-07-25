package Data_Profunctor_Join

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Join gopurs_runtime.Value
var once_Join sync.Once
func Get_Join() gopurs_runtime.Value {
	once_Join.Do(func() {
		cache_Join = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Join
}

var cache_showJoin gopurs_runtime.Value
var once_showJoin sync.Once
func Get_showJoin() gopurs_runtime.Value {
	once_showJoin.Do(func() {
		cache_showJoin = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Join ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showJoin
}

var cache_semigroupJoin gopurs_runtime.Value
var once_semigroupJoin sync.Once
func Get_semigroupJoin() gopurs_runtime.Value {
	once_semigroupJoin.Do(func() {
		cache_semigroupJoin = gopurs_runtime.Func(func(dictSemigroupoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), v_1, v1_2)
}))
}()
})
	})
	return cache_semigroupJoin
}

var cache_ordJoin gopurs_runtime.Value
var once_ordJoin sync.Once
func Get_ordJoin() gopurs_runtime.Value {
	once_ordJoin.Do(func() {
		cache_ordJoin = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}()
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
		cache_monoidJoin = gopurs_runtime.Func(func(dictCategory_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictCategory_0 gopurs_runtime.Value = dictCategory_0_loop
_ = dictCategory_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCategory_0, "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_1_0
semigroupJoin1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "compose"), v_2, v1_3)
}))
_ = semigroupJoin1_2_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(dictCategory_0, "identity"), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupJoin1_2_1
}))
}()
})
	})
	return cache_monoidJoin
}

var cache_invariantJoin gopurs_runtime.Value
var once_invariantJoin sync.Once
func Get_invariantJoin() gopurs_runtime.Value {
	once_invariantJoin.Do(func() {
		cache_invariantJoin = gopurs_runtime.Func(func(dictProfunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictProfunctor_0 gopurs_runtime.Value = dictProfunctor_0_loop
_ = dictProfunctor_0
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictProfunctor_0, "dimap"), g_2, f_1, v_3)
}))
}()
})
	})
	return cache_invariantJoin
}

var cache_eqJoin gopurs_runtime.Value
var once_eqJoin sync.Once
func Get_eqJoin() gopurs_runtime.Value {
	once_eqJoin.Do(func() {
		cache_eqJoin = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}()
})
	})
	return cache_eqJoin
}




