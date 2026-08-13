package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Clown_Clown gopurs_runtime.Value
var once_Data_Functor_Clown_Clown sync.Once
func Get_Data_Functor_Clown_Clown() gopurs_runtime.Value {
	once_Data_Functor_Clown_Clown.Do(func() {
		cache_Data_Functor_Clown_Clown = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_Clown(x_0_box)
})
	})
	return cache_Data_Functor_Clown_Clown
}

var cache_Data_Functor_Clown_showClown gopurs_runtime.Value
var once_Data_Functor_Clown_showClown sync.Once
func Get_Data_Functor_Clown_showClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_showClown.Do(func() {
		cache_Data_Functor_Clown_showClown = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_showClown(dictShow_0_box)
})
	})
	return cache_Data_Functor_Clown_showClown
}

var cache_Data_Functor_Clown_profunctorClown gopurs_runtime.Value
var once_Data_Functor_Clown_profunctorClown sync.Once
func Get_Data_Functor_Clown_profunctorClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_profunctorClown.Do(func() {
		cache_Data_Functor_Clown_profunctorClown = gopurs_runtime.Func(func(dictContravariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_profunctorClown(dictContravariant_0_box)
})
	})
	return cache_Data_Functor_Clown_profunctorClown
}

var cache_Data_Functor_Clown_ordClown gopurs_runtime.Value
var once_Data_Functor_Clown_ordClown sync.Once
func Get_Data_Functor_Clown_ordClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_ordClown.Do(func() {
		cache_Data_Functor_Clown_ordClown = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_ordClown(dictOrd_0_box)
})
	})
	return cache_Data_Functor_Clown_ordClown
}

var cache_Data_Functor_Clown_newtypeClown gopurs_runtime.Value
var once_Data_Functor_Clown_newtypeClown sync.Once
func Get_Data_Functor_Clown_newtypeClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_newtypeClown.Do(func() {
		cache_Data_Functor_Clown_newtypeClown = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Data_Functor_Clown_newtypeClown
}

var cache_Data_Functor_Clown_hoistClown gopurs_runtime.Value
var once_Data_Functor_Clown_hoistClown sync.Once
func Get_Data_Functor_Clown_hoistClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_hoistClown.Do(func() {
		cache_Data_Functor_Clown_hoistClown = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_hoistClown(f_0_box, v_1_box)
})
	})
	return cache_Data_Functor_Clown_hoistClown
}

var cache_Data_Functor_Clown_functorClown gopurs_runtime.Value
var once_Data_Functor_Clown_functorClown sync.Once
func Get_Data_Functor_Clown_functorClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_functorClown.Do(func() {
		cache_Data_Functor_Clown_functorClown = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
})})}
	})
	return cache_Data_Functor_Clown_functorClown
}

var cache_Data_Functor_Clown_eqClown gopurs_runtime.Value
var once_Data_Functor_Clown_eqClown sync.Once
func Get_Data_Functor_Clown_eqClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_eqClown.Do(func() {
		cache_Data_Functor_Clown_eqClown = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_eqClown(dictEq_0_box)
})
	})
	return cache_Data_Functor_Clown_eqClown
}

var cache_Data_Functor_Clown_bifunctorClown gopurs_runtime.Value
var once_Data_Functor_Clown_bifunctorClown sync.Once
func Get_Data_Functor_Clown_bifunctorClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_bifunctorClown.Do(func() {
		cache_Data_Functor_Clown_bifunctorClown = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_bifunctorClown(dictFunctor_0_box)
})
	})
	return cache_Data_Functor_Clown_bifunctorClown
}

var cache_Data_Functor_Clown_biapplyClown gopurs_runtime.Value
var once_Data_Functor_Clown_biapplyClown sync.Once
func Get_Data_Functor_Clown_biapplyClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_biapplyClown.Do(func() {
		cache_Data_Functor_Clown_biapplyClown = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_biapplyClown(dictApply_0_box)
})
	})
	return cache_Data_Functor_Clown_biapplyClown
}

var cache_Data_Functor_Clown_biapplicativeClown gopurs_runtime.Value
var once_Data_Functor_Clown_biapplicativeClown sync.Once
func Get_Data_Functor_Clown_biapplicativeClown() gopurs_runtime.Value {
	once_Data_Functor_Clown_biapplicativeClown.Do(func() {
		cache_Data_Functor_Clown_biapplicativeClown = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Clown_biapplicativeClown(dictApplicative_0_box)
})
	})
	return cache_Data_Functor_Clown_biapplicativeClown
}

func Call_Data_Functor_Clown_Clown(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Clown_showClown(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Clown ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
})})}
}

func Call_Data_Functor_Clown_profunctorClown(dictContravariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictContravariant_0 gopurs_runtime.Value = dictContravariant_0_loop
_ = dictContravariant_0
return gopurs_runtime.Value{Type: 9, IntVal: 2367018778, UnsafePtr: unsafe.Pointer(&Constructor_Data_Profunctor_Profunctor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictContravariant_0, "cmap"), f_1, v1_3)
})
})
})})}
}

func Call_Data_Functor_Clown_ordClown(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0))}
}

func Call_Data_Functor_Clown_hoistClown(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Data_Functor_Clown_eqClown(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0))}
}

func Call_Data_Functor_Clown_bifunctorClown(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(&Constructor_Data_Bifunctor_Bifunctor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, v1_3)
})
})
})})}
}

func Call_Data_Functor_Clown_biapplyClown(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorClown1_1_0 -> *Constructor_Data_Bifunctor_Bifunctor
bifunctorClown1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, v1_4)
})
})
})))
_ = bifunctorClown1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(&Constructor_Control_Biapply_Biapply{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorClown1_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), v_2, v1_3)
})
})})}
}

func Call_Data_Functor_Clown_biapplicativeClown(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifunctorClown1_2_2 -> *Constructor_Data_Bifunctor_Bifunctor
bifunctorClown1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, v1_5)
})
})
})))
_ = bifunctorClown1_2_2
// TAST (Let): biapplyClown1_1_0 -> *Constructor_Control_Biapply_Biapply
biapplyClown1_1_0 := &Constructor_Control_Biapply_Biapply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorClown1_2_2)}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), v_3, v1_4)
})
})}
_ = biapplyClown1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3949191309, UnsafePtr: unsafe.Pointer(&Constructor_Control_Biapplicative_Biapplicative{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(biapplyClown1_1_0)}
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), a_2)
})
})})}
}


