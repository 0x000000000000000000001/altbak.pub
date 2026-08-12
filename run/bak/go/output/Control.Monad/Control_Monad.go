package Control_Monad

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_whenM gopurs_runtime.Value
var once_whenM sync.Once
func Get_whenM() gopurs_runtime.Value {
	once_whenM.Do(func() {
		cache_whenM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_whenM(gopurs_runtime.CoerceToStruct[Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_whenM
}

var cache_unlessM gopurs_runtime.Value
var once_unlessM sync.Once
func Get_unlessM() gopurs_runtime.Value {
	once_unlessM.Do(func() {
		cache_unlessM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unlessM(gopurs_runtime.CoerceToStruct[Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
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
		cache_liftM1 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftM1(gopurs_runtime.CoerceToStruct[Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_liftM1
}

var cache_ap gopurs_runtime.Value
var once_ap sync.Once
func Get_ap() gopurs_runtime.Value {
	once_ap.Do(func() {
		cache_ap = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ap(gopurs_runtime.CoerceToStruct[Constructor_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_ap
}

var cache_applicativeArray__1604836744 gopurs_runtime.Value
var once_applicativeArray__1604836744 sync.Once
func Get_applicativeArray__1604836744() gopurs_runtime.Value {
	once_applicativeArray__1604836744.Do(func() {
		cache_applicativeArray__1604836744 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}))
	})
	return cache_applicativeArray__1604836744
}

var cache_applicativeFn__3751223912 gopurs_runtime.Value
var once_applicativeFn__3751223912 sync.Once
func Get_applicativeFn__3751223912() gopurs_runtime.Value {
	once_applicativeFn__3751223912.Do(func() {
		cache_applicativeFn__3751223912 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
}))
	})
	return cache_applicativeFn__3751223912
}

var cache_applicativeProxy__1913125352 gopurs_runtime.Value
var once_applicativeProxy__1913125352 sync.Once
func Get_applicativeProxy__1913125352() gopurs_runtime.Value {
	once_applicativeProxy__1913125352.Do(func() {
		cache_applicativeProxy__1913125352 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
}))
	})
	return cache_applicativeProxy__1913125352
}

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_unless__1954875249 gopurs_runtime.Value
var once_unless__1954875249 sync.Once
func Get_unless__1954875249() gopurs_runtime.Value {
	once_unless__1954875249.Do(func() {
		cache_unless__1954875249 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unless__1954875249(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_unless__1954875249
}

var cache_when__1954875249 gopurs_runtime.Value
var once_when__1954875249 sync.Once
func Get_when__1954875249() gopurs_runtime.Value {
	once_when__1954875249.Do(func() {
		cache_when__1954875249 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_when__1954875249(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), (v_1_box.IntVal) != (0), v1_2_box)
})
	})
	return cache_when__1954875249
}

var cache_applyArray__2998472828 gopurs_runtime.Value
var once_applyArray__2998472828 sync.Once
func Get_applyArray__2998472828() gopurs_runtime.Value {
	once_applyArray__2998472828.Do(func() {
		cache_applyArray__2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Control_Apply.Get_arrayApply())
	})
	return cache_applyArray__2998472828
}

var cache_applyFn__4042184691 gopurs_runtime.Value
var once_applyFn__4042184691 sync.Once
func Get_applyFn__4042184691() gopurs_runtime.Value {
	once_applyFn__4042184691.Do(func() {
		cache_applyFn__4042184691 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorFn()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, x_2, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_applyFn__4042184691
}

var cache_applyProxy__2261709491 gopurs_runtime.Value
var once_applyProxy__2261709491 sync.Once
func Get_applyProxy__2261709491() gopurs_runtime.Value {
	once_applyProxy__2261709491.Do(func() {
		cache_applyProxy__2261709491 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_applyProxy__2261709491
}

var cache_bind__3225218311 gopurs_runtime.Value
var once_bind__3225218311 sync.Once
func Get_bind__3225218311() gopurs_runtime.Value {
	once_bind__3225218311.Do(func() {
		cache_bind__3225218311 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__3225218311(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__3225218311
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__2370822215 gopurs_runtime.Value
var once_bind__2370822215 sync.Once
func Get_bind__2370822215() gopurs_runtime.Value {
	once_bind__2370822215.Do(func() {
		cache_bind__2370822215 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2370822215(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2370822215
}

var cache_bindArray__1650562023 gopurs_runtime.Value
var once_bindArray__1650562023 sync.Once
func Get_bindArray__1650562023() gopurs_runtime.Value {
	once_bindArray__1650562023.Do(func() {
		cache_bindArray__1650562023 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), pkg_Control_Bind.Get_arrayBind())
	})
	return cache_bindArray__1650562023
}

var cache_bindFn__1648334822 gopurs_runtime.Value
var once_bindFn__1648334822 sync.Once
func Get_bindFn__1648334822() gopurs_runtime.Value {
	once_bindFn__1648334822.Do(func() {
		cache_bindFn__1648334822 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyFn()
}), gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(m_0, x_2), x_2)
})
})
}))
	})
	return cache_bindFn__1648334822
}

var cache_bindProxy__347077479 gopurs_runtime.Value
var once_bindProxy__347077479 sync.Once
func Get_bindProxy__347077479() gopurs_runtime.Value {
	once_bindProxy__347077479.Do(func() {
		cache_bindProxy__347077479 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_bindProxy__347077479
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

var cache_functorFn__20325936 gopurs_runtime.Value
var once_functorFn__20325936 sync.Once
func Get_functorFn__20325936() gopurs_runtime.Value {
	once_functorFn__20325936.Do(func() {
		cache_functorFn__20325936 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn__20325936
}

var cache_functorProxy__1157108209 gopurs_runtime.Value
var once_functorProxy__1157108209 sync.Once
func Get_functorProxy__1157108209() gopurs_runtime.Value {
	once_functorProxy__1157108209.Do(func() {
		cache_functorProxy__1157108209 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_functorProxy__1157108209
}

type Constructor_Monad[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[778916621] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Monad[gopurs_runtime.Value])(ptr)
		switch key {
		case "Applicative0": return c.V0
		case "Bind1": return c.V1
		default: panic("Key not found in dictionary Constructor_Monad: " + key)
		}
	}
}


func Call_whenM(dictMonad_0_loop *Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(mb_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, mb_3, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (b_5.IntVal) != (0) {
__t2 = m_4
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(Applicative0_2_1.V1, pkg_Data_Unit.Get_unit())
}
end_branch_2:
return __t2
}))
})
})
}

func Call_unlessM(dictMonad_0_loop *Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(mb_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, mb_3, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if ((b_5.IntVal) != (0)) != (true) {
__t2 = m_4
goto end_branch_2
} else {

}
}
{
if (b_5.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_2_1.V1, pkg_Data_Unit.Get_unit())
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
})
})
}

func Call_liftM1(dictMonad_0_loop *Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, a_4, gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Apply(f_3, a_prime_5))
}))
})
})
}

func Call_ap(dictMonad_0_loop *Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V1, gopurs_runtime.Value{}))
_ = Bind1_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_1_0.V1, a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Applicative0_2_1.V1, gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
})
})
}

func Call_pure__2935994064(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unless__1954875249(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 bool = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if (v_1) != (true) {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
if v_1 {
__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_when__1954875249(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], v_1_loop bool, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 bool = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t0 gopurs_runtime.Value
{
if v_1 {
__t0 = v1_2
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit())
}
end_branch_0:
return __t0
}

func Call_bind__3225218311(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2370822215(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


