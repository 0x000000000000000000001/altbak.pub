package Control_Monad_ST_Class

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Effect "gopurs/output/Effect"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monadSTST gopurs_runtime.Value
var once_monadSTST sync.Once
func Get_monadSTST() gopurs_runtime.Value {
	once_monadSTST.Do(func() {
		cache_monadSTST = gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_monadST()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_monadSTST
}

var cache_monadSTEffect gopurs_runtime.Value
var once_monadSTEffect sync.Once
func Get_monadSTEffect() gopurs_runtime.Value {
	once_monadSTEffect.Do(func() {
		cache_monadSTEffect = gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}), pkg_Unsafe_Coerce.Get_unsafeCoerce())
	})
	return cache_monadSTEffect
}

var cache_liftST gopurs_runtime.Value
var once_liftST sync.Once
func Get_liftST() gopurs_runtime.Value {
	once_liftST.Do(func() {
		cache_liftST = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftST(gopurs_runtime.CoerceToStruct[Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftST
}

var cache_toEffect__4169273813 gopurs_runtime.Value
var once_toEffect__4169273813 sync.Once
func Get_toEffect__4169273813() gopurs_runtime.Value {
	once_toEffect__4169273813.Do(func() {
		cache_toEffect__4169273813 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_toEffect__4169273813
}

var cache_applicativeST__2868811880 gopurs_runtime.Value
var once_applicativeST__2868811880 sync.Once
func Get_applicativeST__2868811880() gopurs_runtime.Value {
	once_applicativeST__2868811880.Do(func() {
		cache_applicativeST__2868811880 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_applyST()
}), pkg_Control_Monad_ST_Internal.Get_pure_())
	})
	return cache_applicativeST__2868811880
}

var cache_applyST__2741064779 gopurs_runtime.Value
var once_applyST__2741064779 sync.Once
func Get_applyST__2741064779() gopurs_runtime.Value {
	once_applyST__2741064779.Do(func() {
		cache_applyST__2741064779 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_monadST(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_monadST(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_functorST()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyST__2741064779
}

var cache_bindST__4187656679 gopurs_runtime.Value
var once_bindST__4187656679 sync.Once
func Get_bindST__4187656679() gopurs_runtime.Value {
	once_bindST__4187656679.Do(func() {
		cache_bindST__4187656679 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_applyST()
}), pkg_Control_Monad_ST_Internal.Get_bind_())
	})
	return cache_bindST__4187656679
}

var cache_functorST__2441840241 gopurs_runtime.Value
var once_functorST__2441840241 sync.Once
func Get_functorST__2441840241() gopurs_runtime.Value {
	once_functorST__2441840241.Do(func() {
		cache_functorST__2441840241 = gopurs_runtime.RecordDict1("map", pkg_Control_Monad_ST_Internal.Get_map_())
	})
	return cache_functorST__2441840241
}

var cache_monadST__1413783571 gopurs_runtime.Value
var once_monadST__1413783571 sync.Once
func Get_monadST__1413783571() gopurs_runtime.Value {
	once_monadST__1413783571.Do(func() {
		cache_monadST__1413783571 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_applicativeST()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_bindST()
}))
	})
	return cache_monadST__1413783571
}

var cache_applicativeEffect__1969567048 gopurs_runtime.Value
var once_applicativeEffect__1969567048 sync.Once
func Get_applicativeEffect__1969567048() gopurs_runtime.Value {
	once_applicativeEffect__1969567048.Do(func() {
		cache_applicativeEffect__1969567048 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_pureE())
	})
	return cache_applicativeEffect__1969567048
}

var cache_applyEffect__2014400020 gopurs_runtime.Value
var once_applyEffect__2014400020 sync.Once
func Get_applyEffect__2014400020() gopurs_runtime.Value {
	once_applyEffect__2014400020.Do(func() {
		cache_applyEffect__2014400020 = func() gopurs_runtime.Value {
Bind1_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_0_0
Applicative0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_monadEffect(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_functorEffect()
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, f_2, gopurs_runtime.Func(func(f_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_0_0.V1, a_3, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_1_1.V1, gopurs_runtime.Apply(f_prime_4, a_prime_5))
}))
}))
})
}))
}()
	})
	return cache_applyEffect__2014400020
}

var cache_bindEffect__3856311079 gopurs_runtime.Value
var once_bindEffect__3856311079 sync.Once
func Get_bindEffect__3856311079() gopurs_runtime.Value {
	once_bindEffect__3856311079.Do(func() {
		cache_bindEffect__3856311079 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applyEffect()
}), pkg_Effect.Get_bindE())
	})
	return cache_bindEffect__3856311079
}

var cache_functorEffect__3107547953 gopurs_runtime.Value
var once_functorEffect__3107547953 sync.Once
func Get_functorEffect__3107547953() gopurs_runtime.Value {
	once_functorEffect__3107547953.Do(func() {
		cache_functorEffect__3107547953 = func() gopurs_runtime.Value {
Apply0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_0_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_0_0.V1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), f_1), a_2)
})
}))
}()
	})
	return cache_functorEffect__3107547953
}

var cache_monadEffect__3527935219 gopurs_runtime.Value
var once_monadEffect__3527935219 sync.Once
func Get_monadEffect__3527935219() gopurs_runtime.Value {
	once_monadEffect__3527935219.Do(func() {
		cache_monadEffect__3527935219 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_applicativeEffect()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_bindEffect()
}))
	})
	return cache_monadEffect__3527935219
}

type Constructor_MonadST[T_s any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2155655715] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "liftST": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadST: " + key)
		}
	}
}


func Call_liftST(dict_0_loop *Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


