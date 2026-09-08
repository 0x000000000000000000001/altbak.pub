package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Parallel_Class_Parallel_dollar_Dict gopurs_runtime.Value
var once_Control_Parallel_Class_Parallel_dollar_Dict sync.Once
func Get_Control_Parallel_Class_Parallel_dollar_Dict() gopurs_runtime.Value {
	once_Control_Parallel_Class_Parallel_dollar_Dict.Do(func() {
		cache_Control_Parallel_Class_Parallel_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer(Call_Control_Parallel_Class_Parallel_dollar_Dict(func() struct{
	Apply0 gopurs_runtime.Value
	Apply1 gopurs_runtime.Value
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Apply0 gopurs_runtime.Value
	Apply1 gopurs_runtime.Value
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}{}
					clone.Apply0 = gopurs_runtime.RecordGet(orig, "Apply0")
					clone.Apply1 = gopurs_runtime.RecordGet(orig, "Apply1")
					clone.parallel = gopurs_runtime.RecordGet(orig, "parallel")
					clone.sequential = gopurs_runtime.RecordGet(orig, "sequential")
					return clone
				}()))}
})
	})
	return cache_Control_Parallel_Class_Parallel_dollar_Dict
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
return Call_Control_Parallel_Class_sequential(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_sequential
}

var cache_Control_Parallel_Class_parallel gopurs_runtime.Value
var once_Control_Parallel_Class_parallel sync.Once
func Get_Control_Parallel_Class_parallel() gopurs_runtime.Value {
	once_Control_Parallel_Class_parallel.Do(func() {
		cache_Control_Parallel_Class_parallel = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Parallel_Class_parallel(gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Parallel_Class_parallel
}

var cache_Control_Parallel_Class_newtypeParCont gopurs_runtime.Value
var once_Control_Parallel_Class_newtypeParCont sync.Once
func Get_Control_Parallel_Class_newtypeParCont() gopurs_runtime.Value {
	once_Control_Parallel_Class_newtypeParCont.Do(func() {
		cache_Control_Parallel_Class_newtypeParCont = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer((&Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})}))}
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

type Constructor_Control_Parallel_Class_Parallel[T_f any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[327692956] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Parallel_Class_Parallel[any, any])(ptr)
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


func Call_Control_Parallel_Class_Parallel_dollar_Dict(x_0_loop struct{
	Apply0 gopurs_runtime.Value
	Apply1 gopurs_runtime.Value
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
}) *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Apply0 gopurs_runtime.Value
	Apply1 gopurs_runtime.Value
	parallel gopurs_runtime.Value
	sequential gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Apply0", "Apply1", "parallel", "sequential"}, []gopurs_runtime.Value{orig.Apply0, orig.Apply1, orig.parallel, orig.sequential})
				}())
}

func Call_Control_Parallel_Class_ParCont(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Parallel_Class_sequential(dict_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Parallel_Class_parallel(dict_0_loop *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Parallel_Class_monadParWriterT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): applyWriterT__193435443_1_0 shape=App(Var) bindingType=Any
applyWriterT__193435443_1_0 := gopurs_runtime.Apply(Get_Control_Monad_Writer_Trans_applyWriterT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = applyWriterT__193435443_1_0
return gopurs_runtime.Func(func(dictParallel_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyWriterT1_3_1 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply0"), gopurs_runtime.Value{})))
_ = applyWriterT1_3_1
// TAST (Let): applyWriterT2_4_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar f) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar w)])])])
applyWriterT2_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(applyWriterT__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "Apply1"), gopurs_runtime.Value{})))
_ = applyWriterT2_4_2
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT1_3_1)}
}), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyWriterT2_4_2)}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "parallel"), v_5)
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_2, "sequential"), v_5)
})}))}
})
}

func Call_Control_Parallel_Class_monadParStar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorStar1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=Any
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, gopurs_runtime.Apply(v_4, x_6))
})
})})
_ = functorStar1_2_2
// TAST (Let): applyStar_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar m), (TypeVar a)])])
applyStar_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_2_2)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, a_5), gopurs_runtime.Apply(v1_4, a_5))
})})
_ = applyStar_1_0
// TAST (Let): __local_var_2_6 shape=App(Other) bindingType=Any
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
// TAST (Let): __local_var_3_8 shape=App(Other) bindingType=Any
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
// TAST (Let): functorStar1_3_7 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
functorStar1_3_7 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_9 shape=App(Other) bindingType=Any
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), f_4)
_ = __local_var_6_9
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_9, gopurs_runtime.Apply(v_5, x_7))
})
})})
_ = functorStar1_3_7
// TAST (Let): applyStar1_2_5 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar b)])) [(TypeVar f), (TypeVar a)])])
applyStar1_2_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStar1_3_7)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "apply"), gopurs_runtime.Apply(v_4, a_6), gopurs_runtime.Apply(v1_5, a_6))
})})
_ = applyStar1_2_5
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyStar1_2_5)}
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(v_3, x_4))
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(v_3, x_4))
})}))}
}

func Call_Control_Parallel_Class_monadParReaderT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorReaderT1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_2_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "map"), x_3)
_ = __local_var_4_4
return gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(v_5, x_6))
})
})})
_ = functorReaderT1_2_2
// TAST (Let): applyReaderT_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar e), (TypeVar m)])])
applyReaderT_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_2_2)}
}), gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply(v_3, r_5), gopurs_runtime.Apply(v1_4, r_5))
})})
_ = applyReaderT_1_0
// TAST (Let): __local_var_2_6 shape=App(Other) bindingType=Any
__local_var_2_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_2_6
// TAST (Let): __local_var_3_8 shape=App(Other) bindingType=Any
__local_var_3_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_6, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_8
// TAST (Let): functorReaderT1_3_7 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar r), (TypeVar m)])])
functorReaderT1_3_7 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_9 shape=App(Other) bindingType=(ForAll [a, b] (Func [(Func [(TypeVar a)] (TypeVar b)), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar f) [(TypeVar b)])))
__local_var_5_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_8, "map"), x_4)
_ = __local_var_5_9
return gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_9, gopurs_runtime.Apply(v_6, x_7))
})
})})
_ = functorReaderT1_3_7
// TAST (Let): applyReaderT1_2_5 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(TypeVar r)] (TypeApp (TypeVar m) [(TypeVar a)])) [(TypeVar e), (TypeVar f)])])
applyReaderT1_2_5 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorReaderT1_3_7)}
}), gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_6, "apply"), gopurs_runtime.Apply(v_4, r_6), gopurs_runtime.Apply(v1_5, r_6))
})})
_ = applyReaderT1_2_5
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyReaderT1_2_5)}
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), gopurs_runtime.Apply(v_3, x_4))
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), gopurs_runtime.Apply(v_3, x_4))
})}))}
}

func Call_Control_Parallel_Class_monadParMaybeT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorCompose2_3_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_3_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_6)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_6.UnsafePtr).V0), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), v_5)
})})
_ = functorCompose2_3_3
// TAST (Let): applyCompose_1_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (ADT ["Data","Maybe","Maybe"] [])])])
applyCompose_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_3_3)}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_applyMaybe()).V1), v_4), v1_5)
})})
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyMaybeT_3_7 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])])])
applyMaybeT_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Monad_Maybe_Trans_applyMaybeT(), dictMonad_2))
_ = applyMaybeT_3_7
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyMaybeT_3_7)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyCompose_1_0)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
})}))}
})
}

func Call_Control_Parallel_Class_monadParExceptT(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "Apply1"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_4 shape=App(Other) bindingType=Any
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorCompose2_3_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose2_3_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(f_4, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_6.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
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
})})
_ = functorCompose2_3_3
// TAST (Let): applyCompose_1_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (ADT ["Data","Either","Either"] [(TypeVar e)])])])
applyCompose_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose2_3_3)}
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.RecordGet(Get_Data_Either_applyEither(), "apply"), v_4), v1_5)
})})
_ = applyCompose_1_0
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): applyExceptT_3_6 shape=App(Var) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (TypeVar m) [(ADT ["Data","Either","Either"] [(TypeVar e), (TypeVar a)])])])
applyExceptT_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Monad_Except_Trans_applyExceptT(), dictMonad_2))
_ = applyExceptT_3_6
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyExceptT_3_6)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyCompose_1_0)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), v_4)
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), v_4)
})}))}
})
}

func Call_Control_Parallel_Class_monadParCostar(dictParallel_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictParallel_0 gopurs_runtime.Value = dictParallel_0_loop
_ = dictParallel_0
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Functor_Costar_applyCostar()))}
}), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Functor_Costar_applyCostar()))}
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "sequential"), x_2))
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictParallel_0, "parallel"), x_2))
})}))}
}

func Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): functorContT1_1_1 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [(TypeVar r), (TypeVar m)])])
functorContT1_1_1 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT1_1_1
// TAST (Let): applyContT_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m)])])
applyContT_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorContT1_1_1)}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})})
_ = applyContT_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 327692956, UnsafePtr: unsafe.Pointer((&Constructor_Control_Parallel_Class_Parallel[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyContT_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0)))}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return v_2
})}))}
}

func Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): functorContT_1_0 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [Any])
functorContT_1_0 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})})
_ = functorContT_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 shape=App(Other) bindingType=Any
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_1_0.V0), f_2)
_ = __local_var_3_2
// TAST (Let): __local_var_4_3 shape=Other bindingType=(Func [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])] (TypeApp Any [(TypeVar a)]))
__local_var_4_3 := gopurs_runtime.RecordGet(Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0), "sequential")
_ = __local_var_4_3
// TAST (Let): __local_var_3_1 shape=Let(Let(Abs(App(Other)))) bindingType=(Func [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])] (TypeApp Any [(TypeVar b)]))
__local_var_3_1 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(__local_var_4_3, x_5))
})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Call_Control_Parallel_Class_monadParParCont(dictMonadEffect_0), "parallel"), gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})}))}
}

func Call_Control_Parallel_Class_applyParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Bind1_1_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Call_Control_Parallel_Class_functorParCont(dictMonadEffect_0)))}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), rb_6)), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mb_8)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{a_7, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, ra_5))
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mb_8)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(mb_8.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
})), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), ra_5)), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ma_9)
if (__t_tag_4 == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{b_8, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, rb_6))
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ma_9)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ma_9.UnsafePtr).V0, b_8))
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
})}))}
}

func Call_Control_Parallel_Class_applicativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Bind1_1_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): applyParCont1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
applyParCont1_1_0 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorContT_3_2 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [Any])
functorContT_3_2 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_3, a_6))
}))
})})
_ = functorContT_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=App(Other) bindingType=Any
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_3_2.V0), f_4)
_ = __local_var_5_4
// TAST (Let): __local_var_5_3 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])] (TypeApp Any [(TypeVar b)]))
__local_var_5_3 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, x_6)
})
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
})
})}))}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), rb_6)), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mb_8)
if (__t_tag_5 == nil) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{a_7, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, ra_5))
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mb_8)
if (__t_tag_6 != nil) {
__t7 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(mb_8.UnsafePtr).V0))
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
})), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), ra_5)), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ma_9)
if (__t_tag_8 == nil) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{b_8, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, rb_6))
goto end_branch_10
} else {

}
}
{
var __t_tag_9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ma_9)
if (__t_tag_9 != nil) {
__t10 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ma_9.UnsafePtr).V0, b_8))
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}))
}))
}))
}))
}))
})})
_ = applyParCont1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyParCont1_1_0)}
}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
})}))}
}

func Call_Control_Parallel_Class_altParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): Applicative0_3_2 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_2
// TAST (Let): functorContT_4_4 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [Any])
functorContT_4_4 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_6, gopurs_runtime.Apply(f_4, a_7))
}))
})})
_ = functorContT_4_4
// TAST (Let): functorParCont1_4_3 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
functorParCont1_4_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 shape=App(Other) bindingType=Any
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_4_4.V0), f_5)
_ = __local_var_6_6
// TAST (Let): __local_var_6_5 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])] (TypeApp Any [(TypeVar b)]))
__local_var_6_5 := gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, x_7)
})
_ = __local_var_6_5
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, x_7)
})
})})
_ = functorParCont1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorParCont1_4_3)}
}), gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_8)), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_9)
}))
}
end_branch_7:
return __t7
}))
})), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_8)), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, a_10)
}))
}
end_branch_8:
return __t8
}))
}))
}))
}))
})}))}
}

func Call_Control_Parallel_Class_plusParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Monad0_2_2 shape=App(Other) bindingType=Any
Monad0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_2
// TAST (Let): Bind1_3_3 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_3
// TAST (Let): Applicative0_4_4 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_4
// TAST (Let): functorContT_5_6 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [Any])
functorContT_5_6 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_6, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_7, gopurs_runtime.Apply(f_5, a_8))
}))
})})
_ = functorContT_5_6
// TAST (Let): functorParCont1_5_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
functorParCont1_5_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 shape=App(Other) bindingType=Any
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_5_6.V0), f_6)
_ = __local_var_7_8
// TAST (Let): __local_var_7_7 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])] (TypeApp Any [(TypeVar b)]))
__local_var_7_7 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_8, x_8)
})
_ = __local_var_7_7
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_7, x_8)
})
})})
_ = functorParCont1_5_5
// TAST (Let): altParCont1_2_1 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
altParCont1_2_1 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorParCont1_5_5)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_9)), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_10)
}))
}
end_branch_9:
return __t9
}))
})), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t10 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_9)), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, a_11)
}))
}
end_branch_10:
return __t10
}))
}))
}))
}))
})})
_ = altParCont1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer((&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altParCont1_2_1)}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), Get_Data_Unit_unit())
})}))}
}

func Call_Control_Parallel_Class_alternativeParCont(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Bind1_1_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_2
// TAST (Let): applyParCont1_1_1 shape=Let(LitRecord) bindingType=(ADT ["Control","Apply","Apply"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
applyParCont1_1_1 := (&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorContT_3_3 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [Any])
functorContT_3_3 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_3, a_6))
}))
})})
_ = functorContT_3_3
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 shape=App(Other) bindingType=Any
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_3_3.V0), f_4)
_ = __local_var_5_5
// TAST (Let): __local_var_5_4 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])] (TypeApp Any [(TypeVar b)]))
__local_var_5_4 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, x_6)
})
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, x_6)
})
})}))}
}), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})), gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), rb_6)), gopurs_runtime.Func(func(mb_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mb_8)
if (__t_tag_6 == nil) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{a_7, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, ra_5))
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](mb_8)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(a_7, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(mb_8.UnsafePtr).V0))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
})), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), ra_5)), gopurs_runtime.Func(func(ma_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
var __t_tag_9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ma_9)
if (__t_tag_9 == nil) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{b_8, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, rb_6))
goto end_branch_11
} else {

}
}
{
var __t_tag_10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](ma_9)
if (__t_tag_10 != nil) {
__t11 = gopurs_runtime.Apply(k_4, gopurs_runtime.Apply((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(ma_9.UnsafePtr).V0, b_8))
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}))
}))
}))
}))
}))
})})
_ = applyParCont1_1_1
// TAST (Let): applicativeParCont1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
applicativeParCont1_1_0 := (&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(applyParCont1_1_1)}
}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, x_2)
})})
_ = applicativeParCont1_1_0
// TAST (Let): Applicative0_2_13 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_13
// TAST (Let): Monad0_3_15 shape=App(Other) bindingType=Any
Monad0_3_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_15
// TAST (Let): Bind1_4_16 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_4_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_16
// TAST (Let): Applicative0_5_17 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_5_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_17
// TAST (Let): functorContT_6_19 shape=LitRecord bindingType=(ADT ["Data","Functor","Functor"] [Any])
functorContT_6_19 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_7, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_8, gopurs_runtime.Apply(f_6, a_9))
}))
})})
_ = functorContT_6_19
// TAST (Let): functorParCont1_6_18 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
functorParCont1_6_18 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_21 shape=App(Other) bindingType=Any
__local_var_8_21 := gopurs_runtime.Apply(gopurs_runtime.Box(functorContT_6_19.V0), f_7)
_ = __local_var_8_21
// TAST (Let): __local_var_8_20 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])] (TypeApp Any [(TypeVar b)]))
__local_var_8_20 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_21, x_9)
})
_ = __local_var_8_20
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_20, x_9)
})
})})
_ = functorParCont1_6_18
// TAST (Let): altParCont1_3_14 shape=Let(Let(Let(Let(LitRecord)))) bindingType=(ADT ["Control","Alt","Alt"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
altParCont1_3_14 := (&Constructor_Control_Alt_Alt[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorParCont1_6_18)}
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref__new(), gopurs_runtime.Bool(false))), gopurs_runtime.Func(func(done_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(v_7, gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_10)), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 gopurs_runtime.Value
{
if (b_12.IntVal) != (0) {
__t22 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_17.V1), Get_Data_Unit_unit())
goto end_branch_22
} else {

}
}
{
__t22 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_10)), gopurs_runtime.Func(func(_dollar___unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_9, a_11)
}))
}
end_branch_22:
return __t22
}))
})), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_8, gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply(Get_Effect_Ref_read(), done_10)), gopurs_runtime.Func(func(b_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (b_13.IntVal) != (0) {
__t23 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_17.V1), Get_Data_Unit_unit())
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_16.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), gopurs_runtime.Apply2(Get_Effect_Ref_write(), gopurs_runtime.Bool(true), done_10)), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_9, a_12)
}))
}
end_branch_23:
return __t23
}))
}))
}))
}))
})})
_ = altParCont1_3_14
// TAST (Let): plusParCont1_2_12 shape=Let(Let(LitRecord)) bindingType=(ADT ["Control","Plus","Plus"] [(TypeApp (Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar r)]))] (TypeApp (TypeVar m) [(TypeVar r)])) [Unit, (TypeVar m), (TypeVar a)])])
plusParCont1_2_12 := (&Constructor_Control_Plus_Plus[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altParCont1_3_14)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_13.V1), Get_Data_Unit_unit())
})})
_ = plusParCont1_2_12
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer((&Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeParCont1_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusParCont1_2_12)}
})}))}
}


