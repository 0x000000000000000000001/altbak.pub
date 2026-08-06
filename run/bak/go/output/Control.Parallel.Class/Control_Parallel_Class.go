package Control_Parallel_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Maybe_Trans "gopurs/output/Control.Monad.Maybe.Trans"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Control_Monad_Except_Trans "gopurs/output/Control.Monad.Except.Trans"
	pkg_Data_Functor_Costar "gopurs/output/Data.Functor.Costar"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	unsafe "unsafe"
)

var cache_ParCont gopurs_runtime.Value
var once_ParCont sync.Once
func Get_ParCont() gopurs_runtime.Value {
	once_ParCont.Do(func() {
		cache_ParCont = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ParCont(x_0_box)
})
	})
	return cache_ParCont
}

var cache_sequential gopurs_runtime.Value
var once_sequential sync.Once
func Get_sequential() gopurs_runtime.Value {
	once_sequential.Do(func() {
		cache_sequential = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential(dict_0_box)
})
	})
	return cache_sequential
}

var cache_sequential__gopurs_runtime_Value_3901640075 gopurs_runtime.Value
var once_sequential__gopurs_runtime_Value_3901640075 sync.Once
func Get_sequential__gopurs_runtime_Value_3901640075() gopurs_runtime.Value {
	once_sequential__gopurs_runtime_Value_3901640075.Do(func() {
		cache_sequential__gopurs_runtime_Value_3901640075 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequential__gopurs_runtime_Value_3901640075(dict_0_box)
})
	})
	return cache_sequential__gopurs_runtime_Value_3901640075
}

var cache_parallel gopurs_runtime.Value
var once_parallel sync.Once
func Get_parallel() gopurs_runtime.Value {
	once_parallel.Do(func() {
		cache_parallel = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel(dict_0_box)
})
	})
	return cache_parallel
}

var cache_parallel__gopurs_runtime_Value_470498798 gopurs_runtime.Value
var once_parallel__gopurs_runtime_Value_470498798 sync.Once
func Get_parallel__gopurs_runtime_Value_470498798() gopurs_runtime.Value {
	once_parallel__gopurs_runtime_Value_470498798.Do(func() {
		cache_parallel__gopurs_runtime_Value_470498798 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel__gopurs_runtime_Value_470498798(dict_0_box)
})
	})
	return cache_parallel__gopurs_runtime_Value_470498798
}

var cache_parallel__gopurs_runtime_Value_3901640075 gopurs_runtime.Value
var once_parallel__gopurs_runtime_Value_3901640075 sync.Once
func Get_parallel__gopurs_runtime_Value_3901640075() gopurs_runtime.Value {
	once_parallel__gopurs_runtime_Value_3901640075.Do(func() {
		cache_parallel__gopurs_runtime_Value_3901640075 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parallel__gopurs_runtime_Value_3901640075(dict_0_box)
})
	})
	return cache_parallel__gopurs_runtime_Value_3901640075
}

var cache_newtypeParCont gopurs_runtime.Value
var once_newtypeParCont sync.Once
func Get_newtypeParCont() gopurs_runtime.Value {
	once_newtypeParCont.Do(func() {
		cache_newtypeParCont = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeParCont
}

var cache_monadParWriterT gopurs_runtime.Value
var once_monadParWriterT sync.Once
func Get_monadParWriterT() gopurs_runtime.Value {
	once_monadParWriterT.Do(func() {
		cache_monadParWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParWriterT(dictMonoid_0_box)
})
	})
	return cache_monadParWriterT
}

var cache_monadParStar gopurs_runtime.Value
var once_monadParStar sync.Once
func Get_monadParStar() gopurs_runtime.Value {
	once_monadParStar.Do(func() {
		cache_monadParStar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParStar(dictParallel_0_box)
})
	})
	return cache_monadParStar
}

var cache_monadParReaderT gopurs_runtime.Value
var once_monadParReaderT sync.Once
func Get_monadParReaderT() gopurs_runtime.Value {
	once_monadParReaderT.Do(func() {
		cache_monadParReaderT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParReaderT(dictParallel_0_box)
})
	})
	return cache_monadParReaderT
}

var cache_monadParMaybeT gopurs_runtime.Value
var once_monadParMaybeT sync.Once
func Get_monadParMaybeT() gopurs_runtime.Value {
	once_monadParMaybeT.Do(func() {
		cache_monadParMaybeT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParMaybeT(dictParallel_0_box)
})
	})
	return cache_monadParMaybeT
}

var cache_monadParExceptT gopurs_runtime.Value
var once_monadParExceptT sync.Once
func Get_monadParExceptT() gopurs_runtime.Value {
	once_monadParExceptT.Do(func() {
		cache_monadParExceptT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParExceptT(dictParallel_0_box)
})
	})
	return cache_monadParExceptT
}

var cache_monadParCostar gopurs_runtime.Value
var once_monadParCostar sync.Once
func Get_monadParCostar() gopurs_runtime.Value {
	once_monadParCostar.Do(func() {
		cache_monadParCostar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParCostar(dictParallel_0_box)
})
	})
	return cache_monadParCostar
}

var cache_monadParParCont gopurs_runtime.Value
var once_monadParParCont sync.Once
func Get_monadParParCont() gopurs_runtime.Value {
	once_monadParParCont.Do(func() {
		cache_monadParParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadParParCont(dictMonadEffect_0_box)
})
	})
	return cache_monadParParCont
}

var cache_functorParCont gopurs_runtime.Value
var once_functorParCont sync.Once
func Get_functorParCont() gopurs_runtime.Value {
	once_functorParCont.Do(func() {
		cache_functorParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorParCont(dictMonadEffect_0_box)
})
	})
	return cache_functorParCont
}

var cache_applyParCont gopurs_runtime.Value
var once_applyParCont sync.Once
func Get_applyParCont() gopurs_runtime.Value {
	once_applyParCont.Do(func() {
		cache_applyParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyParCont(dictMonadEffect_0_box)
})
	})
	return cache_applyParCont
}

var cache_applicativeParCont gopurs_runtime.Value
var once_applicativeParCont sync.Once
func Get_applicativeParCont() gopurs_runtime.Value {
	once_applicativeParCont.Do(func() {
		cache_applicativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeParCont(dictMonadEffect_0_box)
})
	})
	return cache_applicativeParCont
}

var cache_altParCont gopurs_runtime.Value
var once_altParCont sync.Once
func Get_altParCont() gopurs_runtime.Value {
	once_altParCont.Do(func() {
		cache_altParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altParCont(dictMonadEffect_0_box)
})
	})
	return cache_altParCont
}

var cache_plusParCont gopurs_runtime.Value
var once_plusParCont sync.Once
func Get_plusParCont() gopurs_runtime.Value {
	once_plusParCont.Do(func() {
		cache_plusParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusParCont(dictMonadEffect_0_box)
})
	})
	return cache_plusParCont
}

var cache_alternativeParCont gopurs_runtime.Value
var once_alternativeParCont sync.Once
func Get_alternativeParCont() gopurs_runtime.Value {
	once_alternativeParCont.Do(func() {
		cache_alternativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeParCont(dictMonadEffect_0_box)
})
	})
	return cache_alternativeParCont
}

func Call_ParCont(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_sequential(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "sequential")
}

func Call_sequential__gopurs_runtime_Value_3901640075(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "sequential")
}

func Call_parallel(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "parallel")
}

func Call_parallel__gopurs_runtime_Value_470498798(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "parallel")
}

func Call_parallel__gopurs_runtime_Value_3901640075(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "parallel")
}

func Call_monadParWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
applyWriterT_2_1 := gopurs_runtime.Func(func(dictApply_2 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_2, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_3_2
functorWriterT1_4_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Functor0_3_2, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1})}
}))
_ = __local_var_5_4
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, v_6)
})
}))
_ = functorWriterT1_4_3
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_4_3
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_2, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_3_2, "map"), gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v3_7.UnsafePtr).V1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v4_8.UnsafePtr).V1)})}
})
}), v_5), v1_6)
})
}))
})
_ = applyWriterT_2_1
return gopurs_runtime.Func(func(dictParallel_3 gopurs_runtime.Value) gopurs_runtime.Value {
applyWriterT1_4_5 := gopurs_runtime.UncurriedApp(applyWriterT_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "Apply0"), gopurs_runtime.Value{}))
_ = applyWriterT1_4_5
applyWriterT2_5_6 := gopurs_runtime.UncurriedApp(applyWriterT_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "Apply1"), gopurs_runtime.Value{}))
_ = applyWriterT2_5_6
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT1_4_5
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_5_6
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "parallel"), v_6)
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_3, "sequential"), v_6)
}))
})
}

func Call_monadParStar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorStar1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
}))
_ = functorStar1_2_2
applyStar_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
})
})
}))
_ = applyStar_1_0
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
functorStar1_3_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), f_4)
_ = __local_var_6_9
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_9, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_7
applyStar1_2_5 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStar1_3_7
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})
})
}))
_ = applyStar1_2_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyStar1_2_5
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(v_3, x_4))
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(v_3, x_4))
})
}))
}

func Call_monadParReaderT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorReaderT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
}))
_ = functorReaderT1_2_2
applyReaderT_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})
})
}))
_ = applyReaderT_1_0
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
functorReaderT1_3_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_9, gopurs_runtime.Apply(v_6, x_7))
})
})
}))
_ = functorReaderT1_3_7
applyReaderT1_2_5 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorReaderT1_3_7
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})
})
}))
_ = applyReaderT1_2_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyReaderT1_2_5
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(v_3, x_4))
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(v_3, x_4))
})
}))
}

func Call_monadParMaybeT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_2
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
functorCompose2_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "map"), f_4), v_5)
})
}))
_ = functorCompose2_3_3
applyCompose_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_3
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_2, "map"), gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_applyMaybe(), "apply"), v_4), v1_5)
})
}))
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyMaybeT_3_5 := gopurs_runtime.Apply(pkg_Control_Monad_Maybe_Trans.Get_applyMaybeT(), dictMonad_2)
_ = applyMaybeT_3_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMaybeT_3_5
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_1_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
}))
})
}

func Call_monadParExceptT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_2
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_applyEither(), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
functorCompose2_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "map"), f_4), v_5)
})
}))
_ = functorCompose2_3_3
applyCompose_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_3_3
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_2, "map"), gopurs_runtime.RecordGet(pkg_Data_Either.Get_applyEither(), "apply"), v_4), v1_5)
})
}))
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
applyExceptT_3_5 := gopurs_runtime.Apply(pkg_Control_Monad_Except_Trans.Get_applyExceptT(), dictMonad_2)
_ = applyExceptT_3_5
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyExceptT_3_5
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_1_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
}))
})
}

func Call_monadParCostar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor_Costar.Get_applyCostar()
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), x_2))
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), x_2))
})
}))
}

func Call_monadParParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
functorContT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_1
applyContT_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
_ = applyContT_1_0
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT_1_0
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyParCont(dictMonadEffect_0)
}), Get_ParCont(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return v_2
}))
}

func Call_functorParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_monadParParCont(dictMonadEffect_0), "sequential"), x_2)
_ = __local_var_3_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_monadParParCont(dictMonadEffect_0), "parallel"), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_1, a_5))
}))
}))
})
}))
}

func Call_applyParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
discard1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Bind1_1_0)
_ = discard1_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorParCont(dictMonadEffect_0)
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_2 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
_ = __local_ref_2
var __local_iface_2 interface{} = __local_ref_2
return gopurs_runtime.Any(&__local_iface_2)
})), gopurs_runtime.Func(func(ra_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_3 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
_ = __local_ref_3
var __local_iface_3 interface{} = __local_ref_3
return gopurs_runtime.Any(&__local_iface_3)
})), gopurs_runtime.Func(func(rb_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(discard1_2_1, gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(rb_7.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Func(func(mb_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (mb_9.Type == 9 && mb_9.IntVal == 930809136 && mb_9.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(ra_6.PtrVal().(*interface{})) = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_8})}
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_8})}
}))
goto end_branch_4
} else {

}
}
{
if (mb_9.Type == 9 && mb_9.IntVal == 930809136 && mb_9.UnsafePtr != nil) {
__t4 = gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(a_8, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(mb_9.UnsafePtr).V0))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}))
})), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(ra_6.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Func(func(ma_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (ma_10.Type == 9 && ma_10.IntVal == 930809136 && ma_10.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(rb_7.PtrVal().(*interface{})) = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, b_9})}
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, b_9})}
}))
goto end_branch_5
} else {

}
}
{
if (ma_10.Type == 9 && ma_10.IntVal == 930809136 && ma_10.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(k_5, gopurs_runtime.Apply((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ma_10.UnsafePtr).V0, b_9))
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
}))
}))
}))
}))
})
})
}))
}

func Call_applicativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applyParCont1_1_0 := Call_applyParCont(dictMonadEffect_0)
_ = applyParCont1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyParCont1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_monadParParCont(dictMonadEffect_0), "parallel"), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
}))
}))
}

func Call_altParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
discard1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), Bind1_2_1)
_ = discard1_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_3
functorParCont1_5_4 := Call_functorParCont(dictMonadEffect_0)
_ = functorParCont1_5_4
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_5_4
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_ref_5 := gopurs_runtime.Bool(false)
_ = __local_ref_5
var __local_iface_5 interface{} = __local_ref_5
return gopurs_runtime.Any(&__local_iface_5)
})), gopurs_runtime.Func(func(done_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(discard1_3_2, gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(done_9.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (b_11.IntVal) != (0) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply2(discard1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(done_9.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
})), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_10)
}))
}
end_branch_6:
return __t6
}))
})), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return (*(done_9.PtrVal().(*interface{}))).(gopurs_runtime.Value)
})), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (b_12.IntVal) != (0) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), pkg_Data_Unit.Get_unit())
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply2(discard1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
*(done_9.PtrVal().(*interface{})) = gopurs_runtime.Bool(true)
return gopurs_runtime.Bool(true)
})), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_11)
}))
}
end_branch_7:
return __t7
}))
}))
}))
}))
})
})
}))
}

func Call_plusParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
altParCont1_1_0 := Call_altParCont(dictMonadEffect_0)
_ = altParCont1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altParCont1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}))
}

func Call_alternativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
applicativeParCont1_1_0 := Call_applicativeParCont(dictMonadEffect_0)
_ = applicativeParCont1_1_0
plusParCont1_2_1 := Call_plusParCont(dictMonadEffect_0)
_ = plusParCont1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeParCont1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusParCont1_2_1
}))
}


