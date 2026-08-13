package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Parallel_Class_Parallel_dollarDict gopurs_runtime.Value
var once_Control_Parallel_Class_Parallel_dollarDict sync.Once
func Get_Control_Parallel_Class_Parallel_dollarDict() gopurs_runtime.Value {
	once_Control_Parallel_Class_Parallel_dollarDict.Do(func() {
		cache_Control_Parallel_Class_Parallel_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_Parallel_dollarDict(x_0_box)
})
	})
	return cache_Control_Parallel_Class_Parallel_dollarDict
}

var cache_Control_Parallel_Class_ParCont gopurs_runtime.Value
var once_Control_Parallel_Class_ParCont sync.Once
func Get_Control_Parallel_Class_ParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_ParCont.Do(func() {
		cache_Control_Parallel_Class_ParCont = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_ParCont(x_0_box)
})
	})
	return cache_Control_Parallel_Class_ParCont
}

var cache_Control_Parallel_Class_sequential gopurs_runtime.Value
var once_Control_Parallel_Class_sequential sync.Once
func Get_Control_Parallel_Class_sequential() gopurs_runtime.Value {
	once_Control_Parallel_Class_sequential.Do(func() {
		cache_Control_Parallel_Class_sequential = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_sequential
}

var cache_Control_Parallel_Class_parallel gopurs_runtime.Value
var once_Control_Parallel_Class_parallel sync.Once
func Get_Control_Parallel_Class_parallel() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel.Do(func() {
		cache_Control_Parallel_Class_parallel = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_parallel
}

var cache_Control_Parallel_Class_newtypeParCont gopurs_runtime.Value
var once_Control_Parallel_Class_newtypeParCont sync.Once
func Get_Control_Parallel_Class_newtypeParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_newtypeParCont.Do(func() {
		cache_Control_Parallel_Class_newtypeParCont = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Parallel_Class_newtypeParCont
}

var cache_Control_Parallel_Class_monadParWriterT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParWriterT sync.Once
func Get_Control_Parallel_Class_monadParWriterT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParWriterT.Do(func() {
		cache_Control_Parallel_Class_monadParWriterT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_monadParWriterT(dictMonoid_0_box)
})
	})
	return cache_Control_Parallel_Class_monadParWriterT
}

var cache_Control_Parallel_Class_monadParStar gopurs_runtime.Value
var once_Control_Parallel_Class_monadParStar sync.Once
func Get_Control_Parallel_Class_monadParStar() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParStar.Do(func() {
		cache_Control_Parallel_Class_monadParStar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_monadParStar(dictParallel_0_box)
})
	})
	return cache_Control_Parallel_Class_monadParStar
}

var cache_Control_Parallel_Class_monadParReaderT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParReaderT sync.Once
func Get_Control_Parallel_Class_monadParReaderT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParReaderT.Do(func() {
		cache_Control_Parallel_Class_monadParReaderT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_monadParReaderT(dictParallel_0_box)
})
	})
	return cache_Control_Parallel_Class_monadParReaderT
}

var cache_Control_Parallel_Class_monadParMaybeT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParMaybeT sync.Once
func Get_Control_Parallel_Class_monadParMaybeT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParMaybeT.Do(func() {
		cache_Control_Parallel_Class_monadParMaybeT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_monadParMaybeT(dictParallel_0_box)
})
	})
	return cache_Control_Parallel_Class_monadParMaybeT
}

var cache_Control_Parallel_Class_monadParExceptT gopurs_runtime.Value
var once_Control_Parallel_Class_monadParExceptT sync.Once
func Get_Control_Parallel_Class_monadParExceptT() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParExceptT.Do(func() {
		cache_Control_Parallel_Class_monadParExceptT = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_monadParExceptT(dictParallel_0_box)
})
	})
	return cache_Control_Parallel_Class_monadParExceptT
}

var cache_Control_Parallel_Class_monadParCostar gopurs_runtime.Value
var once_Control_Parallel_Class_monadParCostar sync.Once
func Get_Control_Parallel_Class_monadParCostar() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParCostar.Do(func() {
		cache_Control_Parallel_Class_monadParCostar = gopurs_runtime.Func(func(dictParallel_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_monadParCostar(dictParallel_0_box)
})
	})
	return cache_Control_Parallel_Class_monadParCostar
}

var cache_Control_Parallel_Class_monadParParCont gopurs_runtime.Value
var once_Control_Parallel_Class_monadParParCont sync.Once
func Get_Control_Parallel_Class_monadParParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_monadParParCont.Do(func() {
		cache_Control_Parallel_Class_monadParParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0_box)
})
	})
	return cache_Control_Parallel_Class_monadParParCont
}

var cache_Control_Parallel_Class_functorParCont gopurs_runtime.Value
var once_Control_Parallel_Class_functorParCont sync.Once
func Get_Control_Parallel_Class_functorParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_functorParCont.Do(func() {
		cache_Control_Parallel_Class_functorParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0_box)
})
	})
	return cache_Control_Parallel_Class_functorParCont
}

var cache_Control_Parallel_Class_applyParCont gopurs_runtime.Value
var once_Control_Parallel_Class_applyParCont sync.Once
func Get_Control_Parallel_Class_applyParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_applyParCont.Do(func() {
		cache_Control_Parallel_Class_applyParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0_box)
})
	})
	return cache_Control_Parallel_Class_applyParCont
}

var cache_Control_Parallel_Class_applicativeParCont gopurs_runtime.Value
var once_Control_Parallel_Class_applicativeParCont sync.Once
func Get_Control_Parallel_Class_applicativeParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_applicativeParCont.Do(func() {
		cache_Control_Parallel_Class_applicativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_applicativeParCont(dictMonadEffect_0_box)
})
	})
	return cache_Control_Parallel_Class_applicativeParCont
}

var cache_Control_Parallel_Class_altParCont gopurs_runtime.Value
var once_Control_Parallel_Class_altParCont sync.Once
func Get_Control_Parallel_Class_altParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_altParCont.Do(func() {
		cache_Control_Parallel_Class_altParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_altParCont(dictMonadEffect_0_box)
})
	})
	return cache_Control_Parallel_Class_altParCont
}

var cache_Control_Parallel_Class_plusParCont gopurs_runtime.Value
var once_Control_Parallel_Class_plusParCont sync.Once
func Get_Control_Parallel_Class_plusParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_plusParCont.Do(func() {
		cache_Control_Parallel_Class_plusParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_plusParCont(dictMonadEffect_0_box)
})
	})
	return cache_Control_Parallel_Class_plusParCont
}

var cache_Control_Parallel_Class_alternativeParCont gopurs_runtime.Value
var once_Control_Parallel_Class_alternativeParCont sync.Once
func Get_Control_Parallel_Class_alternativeParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_alternativeParCont.Do(func() {
		cache_Control_Parallel_Class_alternativeParCont = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_alternativeParCont(dictMonadEffect_0_box)
})
	})
	return cache_Control_Parallel_Class_alternativeParCont
}

var cache_Control_Parallel_Class_parallel__2242335472 gopurs_runtime.Value
var once_Control_Parallel_Class_parallel__2242335472 sync.Once
func Get_Control_Parallel_Class_parallel__2242335472() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel__2242335472.Do(func() {
		cache_Control_Parallel_Class_parallel__2242335472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_parallel__2242335472(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_parallel__2242335472
}

var cache_Control_Parallel_Class_parallel__412321264 gopurs_runtime.Value
var once_Control_Parallel_Class_parallel__412321264 sync.Once
func Get_Control_Parallel_Class_parallel__412321264() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel__412321264.Do(func() {
		cache_Control_Parallel_Class_parallel__412321264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_parallel__412321264(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_parallel__412321264
}

var cache_Control_Parallel_Class_parallel__1250658000 gopurs_runtime.Value
var once_Control_Parallel_Class_parallel__1250658000 sync.Once
func Get_Control_Parallel_Class_parallel__1250658000() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel__1250658000.Do(func() {
		cache_Control_Parallel_Class_parallel__1250658000 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_parallel__1250658000(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_parallel__1250658000
}

var cache_Control_Parallel_Class_parallel__4223476656 gopurs_runtime.Value
var once_Control_Parallel_Class_parallel__4223476656 sync.Once
func Get_Control_Parallel_Class_parallel__4223476656() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel__4223476656.Do(func() {
		cache_Control_Parallel_Class_parallel__4223476656 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_parallel__4223476656(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_parallel__4223476656
}

var cache_Control_Parallel_Class_parallel__1734610934 gopurs_runtime.Value
var once_Control_Parallel_Class_parallel__1734610934 sync.Once
func Get_Control_Parallel_Class_parallel__1734610934() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel__1734610934.Do(func() {
		cache_Control_Parallel_Class_parallel__1734610934 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_parallel__1734610934(__eta0_0_box)
})
	})
	return cache_Control_Parallel_Class_parallel__1734610934
}

var cache_Control_Parallel_Class_sequential__2242335472 gopurs_runtime.Value
var once_Control_Parallel_Class_sequential__2242335472 sync.Once
func Get_Control_Parallel_Class_sequential__2242335472() gopurs_runtime.Value {
	once_Control_Parallel_Class_sequential__2242335472.Do(func() {
		cache_Control_Parallel_Class_sequential__2242335472 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_sequential__2242335472(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_sequential__2242335472
}

var cache_Control_Parallel_Class_sequential__412321264 gopurs_runtime.Value
var once_Control_Parallel_Class_sequential__412321264 sync.Once
func Get_Control_Parallel_Class_sequential__412321264() gopurs_runtime.Value {
	once_Control_Parallel_Class_sequential__412321264.Do(func() {
		cache_Control_Parallel_Class_sequential__412321264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_sequential__412321264(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_sequential__412321264
}

var cache_Control_Parallel_Class_sequential__1250658000 gopurs_runtime.Value
var once_Control_Parallel_Class_sequential__1250658000 sync.Once
func Get_Control_Parallel_Class_sequential__1250658000() gopurs_runtime.Value {
	once_Control_Parallel_Class_sequential__1250658000.Do(func() {
		cache_Control_Parallel_Class_sequential__1250658000 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_sequential__1250658000(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_sequential__1250658000
}

type Constructor_Control_Parallel_Class_Parallel struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[327692956] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Parallel_Class_Parallel)(ptr)
		_ = c
		switch key {
		case "Apply0": return gopurs_runtime.Box(c.V0)
		case "Apply1": return gopurs_runtime.Box(c.V1)
		case "parallel": return gopurs_runtime.Box(c.V2)
		case "sequential": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Control_Parallel_Class_Parallel: " + key)
		}
	}
}


func Call_Control_Parallel_Class_Parallel_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Parallel_Class_ParCont(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Parallel_Class_sequential(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Parallel_Class_parallel(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Parallel_Class_monadParWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): applyWriterT_1_0 -> gopurs_runtime.Value
applyWriterT_1_0 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_applyWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyWriterT_1_0
return gopurs_runtime.Func(func(dictParallel_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT1_3_1 -> gopurs_runtime.Value
applyWriterT1_3_1 := gopurs_runtime.Apply(applyWriterT_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply0"), gopurs_runtime.Value{}))
_ = applyWriterT1_3_1
// TAST (Let): applyWriterT2_4_2 -> gopurs_runtime.Value
applyWriterT2_4_2 := gopurs_runtime.Apply(applyWriterT_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply1"), gopurs_runtime.Value{}))
_ = applyWriterT2_4_2
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT1_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "parallel"), v_5)
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "sequential"), v_5)
}))
})
}

func Call_Control_Parallel_Class_monadParStar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorStar1_2_2 -> gopurs_runtime.Value
functorStar1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})
}))
_ = functorStar1_2_2
// TAST (Let): applyStar_1_0 -> gopurs_runtime.Value
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
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
// TAST (Let): __local_var_3_8 -> gopurs_runtime.Value
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
// TAST (Let): functorStar1_3_7 -> gopurs_runtime.Value
functorStar1_3_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_9 -> gopurs_runtime.Value
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), f_4)
_ = __local_var_6_9
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_9, gopurs_runtime.Apply(v_5, x_7))
})
})
}))
_ = functorStar1_3_7
// TAST (Let): applyStar1_2_5 -> gopurs_runtime.Value
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

func Call_Control_Parallel_Class_monadParReaderT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 -> gopurs_runtime.Value
functorReaderT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})
}))
_ = functorReaderT1_2_2
// TAST (Let): applyReaderT_1_0 -> gopurs_runtime.Value
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
// TAST (Let): __local_var_2_6 -> gopurs_runtime.Value
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
// TAST (Let): __local_var_3_8 -> gopurs_runtime.Value
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
// TAST (Let): functorReaderT1_3_7 -> gopurs_runtime.Value
functorReaderT1_3_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 -> gopurs_runtime.Value
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_9, gopurs_runtime.Apply(v_6, x_7))
})
})
}))
_ = functorReaderT1_3_7
// TAST (Let): applyReaderT1_2_5 -> gopurs_runtime.Value
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

func Call_Control_Parallel_Class_monadParMaybeT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorCompose2_4_4 -> gopurs_runtime.Value
functorCompose2_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), v_5)
})
}))
_ = functorCompose2_4_4
// TAST (Let): applyCompose_1_0 -> gopurs_runtime.Value
applyCompose_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_4_4
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.RecordGet(Get_Data_Maybe_applyMaybe(), "apply"), v_5), v1_6)
})
}))
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyMaybeT_3_6 -> gopurs_runtime.Value
applyMaybeT_3_6 := gopurs_runtime.Apply(Get_Control_Monad_Maybe_Trans_applyMaybeT(), dictMonad_2)
_ = applyMaybeT_3_6
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyMaybeT_3_6
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_1_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
}))
})
}

func Call_Control_Parallel_Class_monadParExceptT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorCompose2_4_4 -> gopurs_runtime.Value
functorCompose2_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_6.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Either_Right)(m_6.UnsafePtr).V0)})}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}), v_5)
})
}))
_ = functorCompose2_4_4
// TAST (Let): applyCompose_1_0 -> gopurs_runtime.Value
applyCompose_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose2_4_4
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.RecordGet(Get_Data_Either_applyEither(), "apply"), v_5), v1_6)
})
}))
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyExceptT_3_6 -> gopurs_runtime.Value
applyExceptT_3_6 := gopurs_runtime.Apply(Get_Control_Monad_Except_Trans_applyExceptT(), dictMonad_2)
_ = applyExceptT_3_6
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyExceptT_3_6
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyCompose_1_0
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
}))
})
}

func Call_Control_Parallel_Class_monadParCostar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
return gopurs_runtime.RecordDict4("Apply0", "Apply1", "parallel", "sequential", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_Costar_applyCostar()
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_Costar_applyCostar()
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

func Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): functorContT1_1_1 -> gopurs_runtime.Value
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
// TAST (Let): applyContT_1_0 -> gopurs_runtime.Value
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
return Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0)
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return v_2
}))
}

func Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): functorContT_1_0 -> *Constructor_Data_Functor_Functor
functorContT_1_0 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
})}
_ = functorContT_1_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_1_0.V0), f_2)
_ = __local_var_3_2
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, x_4)
})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4)
})
}))
}

func Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorContT_3_1 -> *Constructor_Data_Functor_Functor
functorContT_3_1 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_3, a_6))
}))
})
})
})}
_ = functorContT_3_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_3_1.V0), f_4)
_ = __local_var_5_3
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
})
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, x_6)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), rb_6)), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_7})}, ra_5))
goto end_branch_4
} else {

}
}
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr != nil) {
__t4 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*Constructor_Data_Maybe_Just)(mb_8.UnsafePtr).V0))
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
})), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), ra_5)), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, b_8})}, rb_6))
goto end_branch_5
} else {

}
}
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just)(ma_9.UnsafePtr).V0, b_8))
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

func Call_Control_Parallel_Class_applicativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): applyParCont1_1_0 -> gopurs_runtime.Value
applyParCont1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorContT_3_2 -> *Constructor_Data_Functor_Functor
functorContT_3_2 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_3, a_6))
}))
})
})
})}
_ = functorContT_3_2
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_3_2.V0), f_4)
_ = __local_var_5_4
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, x_6)
})
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), rb_6)), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_7})}, ra_5))
goto end_branch_5
} else {

}
}
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*Constructor_Data_Maybe_Just)(mb_8.UnsafePtr).V0))
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
})), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), ra_5)), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, b_8})}, rb_6))
goto end_branch_6
} else {

}
}
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just)(ma_9.UnsafePtr).V0, b_8))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
}))
}))
}))
}))
})
})
}))
_ = applyParCont1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyParCont1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
})
}))
}

func Call_Control_Parallel_Class_altParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
// TAST (Let): functorContT_4_4 -> *Constructor_Data_Functor_Functor
functorContT_4_4 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_6, gopurs_runtime.Apply(f_4, a_7))
}))
})
})
})}
_ = functorContT_4_4
// TAST (Let): functorParCont1_4_3 -> gopurs_runtime.Value
functorParCont1_4_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_4_4.V0), f_5)
_ = __local_var_6_6
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, x_7)
})
_ = __local_var_6_5
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, x_7)
})
}))
_ = functorParCont1_4_3
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_4_3
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(done_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_8)), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (b_10.IntVal) != (0) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), Get_Data_Unit_unit())
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_8)), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_9)
}))
}
end_branch_7:
return __t7
}))
})), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_8)), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (b_11.IntVal) != (0) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), Get_Data_Unit_unit())
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_8)), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_10)
}))
}
end_branch_8:
return __t8
}))
}))
}))
}))
})
})
}))
}

func Call_Control_Parallel_Class_plusParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Monad0_2_2 -> gopurs_runtime.Value
Monad0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_2
// TAST (Let): Bind1_3_3 -> *Constructor_Control_Bind_Bind
Bind1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_3
// TAST (Let): Applicative0_4_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_4
// TAST (Let): functorContT_5_6 -> *Constructor_Data_Functor_Functor
functorContT_5_6 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, gopurs_runtime.Apply(f_5, a_8))
}))
})
})
})}
_ = functorContT_5_6
// TAST (Let): functorParCont1_5_5 -> gopurs_runtime.Value
functorParCont1_5_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_5_6.V0), f_6)
_ = __local_var_7_8
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_8, x_8)
})
_ = __local_var_7_7
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_7, x_8)
})
}))
_ = functorParCont1_5_5
// TAST (Let): altParCont1_2_1 -> gopurs_runtime.Value
altParCont1_2_1 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_5_5
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(done_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_9)), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (b_11.IntVal) != (0) {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_4.V1), Get_Data_Unit_unit())
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_9)), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_10)
}))
}
end_branch_9:
return __t9
}))
})), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_9)), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (b_12.IntVal) != (0) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_4.V1), Get_Data_Unit_unit())
goto end_branch_10
} else {

}
}
{
__t10 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_9)), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_11)
}))
}
end_branch_10:
return __t10
}))
}))
}))
}))
})
})
}))
_ = altParCont1_2_1
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altParCont1_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
}))
}

func Call_Control_Parallel_Class_alternativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Bind1_1_2 -> *Constructor_Control_Bind_Bind
Bind1_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_2
// TAST (Let): applyParCont1_1_1 -> gopurs_runtime.Value
applyParCont1_1_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorContT_3_3 -> *Constructor_Data_Functor_Functor
functorContT_3_3 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_3, a_6))
}))
})
})
})}
_ = functorContT_3_3
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_3_3.V0), f_4)
_ = __local_var_5_5
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, x_6)
})
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, x_6)
})
}))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), rb_6)), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_7})}, ra_5))
goto end_branch_6
} else {

}
}
{
if (mb_8.Type == 9 && mb_8.IntVal == 930809136 && mb_8.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*Constructor_Data_Maybe_Just)(mb_8.UnsafePtr).V0))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
})), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), ra_5)), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr == nil) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, b_8})}, rb_6))
goto end_branch_7
} else {

}
}
{
if (ma_9.Type == 9 && ma_9.IntVal == 930809136 && ma_9.UnsafePtr != nil) {
__t7 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just)(ma_9.UnsafePtr).V0, b_8))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}))
}))
}))
}))
}))
})
})
}))
_ = applyParCont1_1_1
// TAST (Let): applicativeParCont1_1_0 -> gopurs_runtime.Value
applicativeParCont1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyParCont1_1_1
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
})
}))
_ = applicativeParCont1_1_0
// TAST (Let): Applicative0_2_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_9
// TAST (Let): Monad0_3_11 -> gopurs_runtime.Value
Monad0_3_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_11
// TAST (Let): Bind1_4_12 -> *Constructor_Control_Bind_Bind
Bind1_4_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_12
// TAST (Let): Applicative0_5_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_13
// TAST (Let): functorContT_6_15 -> *Constructor_Data_Functor_Functor
functorContT_6_15 := &Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_7, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, gopurs_runtime.Apply(f_6, a_9))
}))
})
})
})}
_ = functorContT_6_15
// TAST (Let): functorParCont1_6_14 -> gopurs_runtime.Value
functorParCont1_6_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_17 -> gopurs_runtime.Value
__local_var_8_17 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_6_15.V0), f_7)
_ = __local_var_8_17
// TAST (Let): __local_var_8_16 -> gopurs_runtime.Value
__local_var_8_16 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_17, x_9)
})
_ = __local_var_8_16
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_16, x_9)
})
}))
_ = functorParCont1_6_14
// TAST (Let): altParCont1_3_10 -> gopurs_runtime.Value
altParCont1_3_10 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorParCont1_6_14
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(done_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), gopurs_runtime.Apply(v_7, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_10)), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (b_12.IntVal) != (0) {
__t18 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_13.V1), Get_Data_Unit_unit())
goto end_branch_18
} else {

}
}
{
__t18 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_10)), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_9, a_11)
}))
}
end_branch_18:
return __t18
}))
})), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_8, gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_10)), gopurs_runtime.Func(func(b_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (b_13.IntVal) != (0) {
__t19 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_13.V1), Get_Data_Unit_unit())
goto end_branch_19
} else {

}
}
{
__t19 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_12.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_10)), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_9, a_12)
}))
}
end_branch_19:
return __t19
}))
}))
}))
}))
})
})
}))
_ = altParCont1_3_10
// TAST (Let): plusParCont1_2_8 -> gopurs_runtime.Value
plusParCont1_2_8 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altParCont1_3_10
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_9.V1), Get_Data_Unit_unit())
}))
_ = plusParCont1_2_8
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeParCont1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusParCont1_2_8
}))
}

func Call_Control_Parallel_Class_parallel__2242335472(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Parallel_Class_parallel__412321264(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Parallel_Class_parallel__1250658000(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Parallel_Class_parallel__4223476656(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Parallel_Class_parallel__1734610934(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_Control_Parallel_Class_sequential__2242335472(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Parallel_Class_sequential__412321264(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Parallel_Class_sequential__1250658000(dict_0_loop *Constructor_Control_Parallel_Class_Parallel) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}


