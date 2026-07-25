package Control_Monad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
)

var cache_whenM gopurs_runtime.Value
var once_whenM sync.Once
func Get_whenM() gopurs_runtime.Value {
	once_whenM.Do(func() {
		cache_whenM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
when_1_0 := gopurs_runtime.Apply(pkg_Control_Applicative.Get_when(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = when_1_0
return gopurs_runtime.Func2(func(mb_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), mb_2, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(when_1_0, b_4, m_3)
}))
})
}()
})
	})
	return cache_whenM
}

var cache_unlessM gopurs_runtime.Value
var once_unlessM sync.Once
func Get_unlessM() gopurs_runtime.Value {
	once_unlessM.Do(func() {
		cache_unlessM = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
unless_1_0 := gopurs_runtime.Apply(pkg_Control_Applicative.Get_unless(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = unless_1_0
return gopurs_runtime.Func2(func(mb_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), mb_2, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(unless_1_0, b_4, m_3)
}))
})
}()
})
	})
	return cache_unlessM
}

var cache_monadProxy gopurs_runtime.Value
var once_monadProxy sync.Once
func Get_monadProxy() gopurs_runtime.Value {
	once_monadProxy.Do(func() {
		cache_monadProxy = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeProxy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindProxy()
}))
	})
	return cache_monadProxy
}

var cache_monadFn gopurs_runtime.Value
var once_monadFn sync.Once
func Get_monadFn() gopurs_runtime.Value {
	once_monadFn.Do(func() {
		cache_monadFn = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeFn()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindFn()
}))
	})
	return cache_monadFn
}

var cache_monadArray gopurs_runtime.Value
var once_monadArray sync.Once
func Get_monadArray() gopurs_runtime.Value {
	once_monadArray.Do(func() {
		cache_monadArray = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindArray()
}))
	})
	return cache_monadArray
}

var cache_liftM1 gopurs_runtime.Value
var once_liftM1 sync.Once
func Get_liftM1() gopurs_runtime.Value {
	once_liftM1.Do(func() {
		cache_liftM1 = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftM1(dictMonad_0_box, f_1_box, a_2_box)
})
	})
	return cache_liftM1
}

var cache_ap gopurs_runtime.Value
var once_ap sync.Once
func Get_ap() gopurs_runtime.Value {
	once_ap.Do(func() {
		cache_ap = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}()
})
	})
	return cache_ap
}

func Call_liftM1(dictMonad_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), a_2, gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_1, a_prime_3))
}))
}


