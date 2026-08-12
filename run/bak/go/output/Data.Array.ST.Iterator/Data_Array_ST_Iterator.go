package Data_Array_ST_Iterator

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Control_Monad_ST_Uncurried "gopurs/output/Control.Monad.ST.Uncurried"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Array_ST "gopurs/output/Data.Array.ST"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void
}

var cache_void1 gopurs_runtime.Value
var once_void1 sync.Once
func Get_void1() gopurs_runtime.Value {
	once_void1.Do(func() {
		cache_void1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_void1
}

var cache_Iterator gopurs_runtime.Value
var once_Iterator sync.Once
func Get_Iterator() gopurs_runtime.Value {
	once_Iterator.Do(func() {
		cache_Iterator = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(&Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Iterator
}

var cache_peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		cache_peek = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek(gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_peek
}

var cache_next gopurs_runtime.Value
var once_next sync.Once
func Get_next() gopurs_runtime.Value {
	once_next.Do(func() {
		cache_next = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_next(gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_next
}

var cache_pushWhile gopurs_runtime.Value
var once_pushWhile sync.Once
func Get_pushWhile() gopurs_runtime.Value {
	once_pushWhile.Do(func() {
		cache_pushWhile = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pushWhile(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_1_box), array_2_box)
})
	})
	return cache_pushWhile
}

var cache_pushAll gopurs_runtime.Value
var once_pushAll sync.Once
func Get_pushAll() gopurs_runtime.Value {
	once_pushAll.Do(func() {
		cache_pushAll = gopurs_runtime.Apply(Get_pushWhile(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return cache_pushAll
}

var cache_iterator gopurs_runtime.Value
var once_iterator sync.Once
func Get_iterator() gopurs_runtime.Value {
	once_iterator.Do(func() {
		cache_iterator = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterator(f_0_box)
})
	})
	return cache_iterator
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func2(func(iter_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_iterate(gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_0_box), f_1_box)
})
	})
	return cache_iterate
}

var cache_exhausted gopurs_runtime.Value
var once_exhausted sync.Once
func Get_exhausted() gopurs_runtime.Value {
	once_exhausted.Do(func() {
		cache_exhausted = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), pkg_Data_Maybe.Get_isNothing())
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, Call_peek(gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](x_1)))
})
}()
	})
	return cache_exhausted
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

var cache_pure__3079134646 gopurs_runtime.Value
var once_pure__3079134646 sync.Once
func Get_pure__3079134646() gopurs_runtime.Value {
	once_pure__3079134646.Do(func() {
		cache_pure__3079134646 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure")
	})
	return cache_pure__3079134646
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

var cache_bind__3352508289 gopurs_runtime.Value
var once_bind__3352508289 sync.Once
func Get_bind__3352508289() gopurs_runtime.Value {
	once_bind__3352508289.Do(func() {
		cache_bind__3352508289 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3352508289
}

var cache_bind__3115293729 gopurs_runtime.Value
var once_bind__3115293729 sync.Once
func Get_bind__3115293729() gopurs_runtime.Value {
	once_bind__3115293729.Do(func() {
		cache_bind__3115293729 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3115293729
}

var cache_bind__3897039777 gopurs_runtime.Value
var once_bind__3897039777 sync.Once
func Get_bind__3897039777() gopurs_runtime.Value {
	once_bind__3897039777.Do(func() {
		cache_bind__3897039777 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3897039777
}

var cache_bind__3937147233 gopurs_runtime.Value
var once_bind__3937147233 sync.Once
func Get_bind__3937147233() gopurs_runtime.Value {
	once_bind__3937147233.Do(func() {
		cache_bind__3937147233 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind")
	})
	return cache_bind__3937147233
}

var cache_applicativeST__3091537981 gopurs_runtime.Value
var once_applicativeST__3091537981 sync.Once
func Get_applicativeST__3091537981() gopurs_runtime.Value {
	once_applicativeST__3091537981.Do(func() {
		cache_applicativeST__3091537981 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_applyST()
}), pkg_Control_Monad_ST_Internal.Get_pure_())
	})
	return cache_applicativeST__3091537981
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

var cache_bindST__2435660861 gopurs_runtime.Value
var once_bindST__2435660861 sync.Once
func Get_bindST__2435660861() gopurs_runtime.Value {
	once_bindST__2435660861.Do(func() {
		cache_bindST__2435660861 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_applyST()
}), pkg_Control_Monad_ST_Internal.Get_bind_())
	})
	return cache_bindST__2435660861
}

var cache_functorST__4062753802 gopurs_runtime.Value
var once_functorST__4062753802 sync.Once
func Get_functorST__4062753802() gopurs_runtime.Value {
	once_functorST__4062753802.Do(func() {
		cache_functorST__4062753802 = gopurs_runtime.RecordDict1("map", pkg_Control_Monad_ST_Internal.Get_map_())
	})
	return cache_functorST__4062753802
}

var cache_functorST__2441840241 gopurs_runtime.Value
var once_functorST__2441840241 sync.Once
func Get_functorST__2441840241() gopurs_runtime.Value {
	once_functorST__2441840241.Do(func() {
		cache_functorST__2441840241 = gopurs_runtime.RecordDict1("map", pkg_Control_Monad_ST_Internal.Get_map_())
	})
	return cache_functorST__2441840241
}

var cache_modify__3866314397 gopurs_runtime.Value
var once_modify__3866314397 sync.Once
func Get_modify__3866314397() gopurs_runtime.Value {
	once_modify__3866314397.Do(func() {
		cache_modify__3866314397 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__3866314397(f_0_box)
})
	})
	return cache_modify__3866314397
}

var cache_modify__781734141 gopurs_runtime.Value
var once_modify__781734141 sync.Once
func Get_modify__781734141() gopurs_runtime.Value {
	once_modify__781734141.Do(func() {
		cache_modify__781734141 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__781734141(f_0_box)
})
	})
	return cache_modify__781734141
}

var cache_modify_prime__1497736571 gopurs_runtime.Value
var once_modify_prime__1497736571 sync.Once
func Get_modify_prime__1497736571() gopurs_runtime.Value {
	once_modify_prime__1497736571.Do(func() {
		cache_modify_prime__1497736571 = pkg_Control_Monad_ST_Internal.Get_modifyImpl()
	})
	return cache_modify_prime__1497736571
}

var cache_new__3579768924 gopurs_runtime.Value
var once_new__3579768924 sync.Once
func Get_new__3579768924() gopurs_runtime.Value {
	once_new__3579768924.Do(func() {
		cache_new__3579768924 = pkg_Control_Monad_ST_Internal.Get_newImpl()
	})
	return cache_new__3579768924
}

var cache_new__122671164 gopurs_runtime.Value
var once_new__122671164 sync.Once
func Get_new__122671164() gopurs_runtime.Value {
	once_new__122671164.Do(func() {
		cache_new__122671164 = pkg_Control_Monad_ST_Internal.Get_newImpl()
	})
	return cache_new__122671164
}

var cache_new__2010968700 gopurs_runtime.Value
var once_new__2010968700 sync.Once
func Get_new__2010968700() gopurs_runtime.Value {
	once_new__2010968700.Do(func() {
		cache_new__2010968700 = pkg_Control_Monad_ST_Internal.Get_newImpl()
	})
	return cache_new__2010968700
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

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_next__2731492779 gopurs_runtime.Value
var once_next__2731492779 sync.Once
func Get_next__2731492779() gopurs_runtime.Value {
	once_next__2731492779.Do(func() {
		cache_next__2731492779 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_next__2731492779(gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_next__2731492779
}

var cache_peek__201669949 gopurs_runtime.Value
var once_peek__201669949 sync.Once
func Get_peek__201669949() gopurs_runtime.Value {
	once_peek__201669949.Do(func() {
		cache_peek__201669949 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek__201669949(gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_peek__201669949
}

var cache_peek__2731492779 gopurs_runtime.Value
var once_peek__2731492779 sync.Once
func Get_peek__2731492779() gopurs_runtime.Value {
	once_peek__2731492779.Do(func() {
		cache_peek__2731492779 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek__2731492779(gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_peek__2731492779
}

var cache_pushWhile__2298419255 gopurs_runtime.Value
var once_pushWhile__2298419255 sync.Once
func Get_pushWhile__2298419255() gopurs_runtime.Value {
	once_pushWhile__2298419255.Do(func() {
		cache_pushWhile__2298419255 = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, iter_1_box gopurs_runtime.Value, array_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pushWhile__2298419255(p_0_box, gopurs_runtime.CoerceToStruct[Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_1_box), array_2_box)
})
	})
	return cache_pushWhile__2298419255
}

var cache_push__1557574173 gopurs_runtime.Value
var once_push__1557574173 sync.Once
func Get_push__1557574173() gopurs_runtime.Value {
	once_push__1557574173.Do(func() {
		cache_push__1557574173 = gopurs_runtime.Apply(pkg_Control_Monad_ST_Uncurried.Get_runSTFn2(), pkg_Data_Array_ST.Get_pushImpl())
	})
	return cache_push__1557574173
}

var cache_const__220790420 gopurs_runtime.Value
var once_const__220790420 sync.Once
func Get_const__220790420() gopurs_runtime.Value {
	once_const__220790420.Do(func() {
		cache_const__220790420 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__220790420(a_0_box, v_1_box)
})
	})
	return cache_const__220790420
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__2174973445 gopurs_runtime.Value
var once_map__2174973445 sync.Once
func Get_map__2174973445() gopurs_runtime.Value {
	once_map__2174973445.Do(func() {
		cache_map__2174973445 = gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map")
	})
	return cache_map__2174973445
}

var cache_isNothing__2591355336 gopurs_runtime.Value
var once_isNothing__2591355336 sync.Once
func Get_isNothing__2591355336() gopurs_runtime.Value {
	once_isNothing__2591355336.Do(func() {
		cache_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2591355336
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

type Constructor_Iterator[T_r any, T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func Call_peek(v_0_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Apply((*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, i_2))
}))
}

func Call_next(v_0_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_3_1
*(__local_var_1_0.PtrVal().(*interface{})) = gopurs_runtime.Apply2(Get_add__560788792(), __local_var_3_1, gopurs_runtime.Int(1))
return gopurs_runtime.Apply2(Get_add__560788792(), __local_var_3_1, gopurs_runtime.Int(1))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Apply((*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, i_2))
}))
}))
}

func Call_pushWhile(p_0_loop gopurs_runtime.Value, iter_1_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_peek(iter_1), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((mx_4.Type == 9 && mx_4.IntVal == 930809136 && mx_4.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0).IntVal) != (0)) {
__local_var_5_0 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0
_ = __local_var_5_0
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_5_0, array_2)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), Call_next(iter_1))
}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
}
end_branch_1:
return __t1
})))
}))
}

func Call_iterator(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.Apply(Get_Iterator(), f_0), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Int(0)))
}

func Call_iterate(iter_0_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var iter_0 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_0_loop
_ = iter_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_2.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_next(iter_0), gopurs_runtime.Func(func(mx_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (mx_3.Type == 9 && mx_3.IntVal == 930809136 && mx_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(Get_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_2.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})))
}))
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_modify__3866314397(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_modify__781734141(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_modifyImpl(), gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_2_0 := gopurs_runtime.Apply(f_0, s_1)
_ = s_prime_2_0
return gopurs_runtime.RecordDict2("state", "value", s_prime_2_0, s_prime_2_0)
}))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_next__2731492779(v_0_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
_ = __local_var_3_1
*(__local_var_1_0.PtrVal().(*interface{})) = gopurs_runtime.Apply2(Get_add__560788792(), __local_var_3_1, gopurs_runtime.Int(1))
return gopurs_runtime.Apply2(Get_add__560788792(), __local_var_3_1, gopurs_runtime.Int(1))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Apply((*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, i_2))
}))
}))
}

func Call_peek__201669949(v_0_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Apply((*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, i_2))
}))
}

func Call_peek__2731492779(v_0_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
__local_var_1_0 := (*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(__local_var_1_0.PtrVal().(*interface{}))).(gopurs_runtime.Value)
}), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_applicativeST(), "pure"), gopurs_runtime.Apply((*Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, i_2))
}))
}

func Call_pushWhile__2298419255(p_0_loop gopurs_runtime.Value, iter_1_loop *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value], array_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var iter_1 *Constructor_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] = iter_1_loop
_ = iter_1
var array_2 gopurs_runtime.Value = array_2_loop
_ = array_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Apply(pkg_Control_Monad_ST_Internal.Get_newImpl(), gopurs_runtime.Bool(false)), gopurs_runtime.Func(func(go__break_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Monad_ST_Internal.Get_while(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_functorST(), "map"), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(go__break_3.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), Call_peek(iter_1), gopurs_runtime.Func(func(mx_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if ((mx_4.Type == 9 && mx_4.IntVal == 930809136 && mx_4.UnsafePtr != nil)) && ((gopurs_runtime.Apply(p_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0).IntVal) != (0)) {
__local_var_5_0 := (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mx_4.UnsafePtr).V0
_ = __local_var_5_0
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Control_Monad_ST_Internal.Get_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(pkg_Data_Array_ST.Get_pushImpl(), __local_var_5_0, array_2)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_void(), Call_next(iter_1))
}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply(Get_void1(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(go__break_3.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
}))
}
end_branch_1:
return __t1
})))
}))
}

func Call_const__220790420(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_isNothing__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}


