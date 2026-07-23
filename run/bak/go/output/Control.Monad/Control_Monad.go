package Control_Monad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
)

var whenM gopurs_runtime.Value
var once_whenM sync.Once
func Get_whenM() gopurs_runtime.Value {
	once_whenM.Do(func() {
		whenM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
when_1_0 := gopurs_runtime.Apply(pkg_Control_Applicative.Get_when(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = when_1_0
return gopurs_runtime.Func2(func(mb_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), mb_2, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(when_1_0, b_4, m_3)
}))
})
})
	})
	return whenM
}

var unlessM gopurs_runtime.Value
var once_unlessM sync.Once
func Get_unlessM() gopurs_runtime.Value {
	once_unlessM.Do(func() {
		unlessM = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
unless_1_0 := gopurs_runtime.Apply(pkg_Control_Applicative.Get_unless(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = unless_1_0
return gopurs_runtime.Func2(func(mb_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), mb_2, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(unless_1_0, b_4, m_3)
}))
})
})
	})
	return unlessM
}

var monadProxy gopurs_runtime.Value
var once_monadProxy sync.Once
func Get_monadProxy() gopurs_runtime.Value {
	once_monadProxy.Do(func() {
		monadProxy = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeProxy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindProxy()
}))
	})
	return monadProxy
}

var monadFn gopurs_runtime.Value
var once_monadFn sync.Once
func Get_monadFn() gopurs_runtime.Value {
	once_monadFn.Do(func() {
		monadFn = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeFn()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindFn()
}))
	})
	return monadFn
}

var monadArray gopurs_runtime.Value
var once_monadArray sync.Once
func Get_monadArray() gopurs_runtime.Value {
	once_monadArray.Do(func() {
		monadArray = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindArray()
}))
	})
	return monadArray
}

var liftM1 gopurs_runtime.Value
var once_liftM1 sync.Once
func Get_liftM1() gopurs_runtime.Value {
	once_liftM1.Do(func() {
		liftM1 = gopurs_runtime.Func3(func(dictMonad_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), a_2, gopurs_runtime.Func(func(a_prime_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_1, a_prime_3))
}))
})
	})
	return liftM1
}

var ap gopurs_runtime.Value
var once_ap sync.Once
func Get_ap() gopurs_runtime.Value {
	once_ap.Do(func() {
		ap = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
})
	})
	return ap
}


