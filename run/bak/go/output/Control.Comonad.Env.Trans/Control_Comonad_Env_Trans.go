package Control_Comonad_Env_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var EnvT gopurs_runtime.Value
var once_EnvT sync.Once
func Get_EnvT() gopurs_runtime.Value {
	once_EnvT.Do(func() {
		EnvT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return EnvT
}

var withEnvT gopurs_runtime.Value
var once_withEnvT sync.Once
func Get_withEnvT() gopurs_runtime.Value {
	once_withEnvT.Do(func() {
		withEnvT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
})
	})
	return withEnvT
}

var runEnvT gopurs_runtime.Value
var once_runEnvT sync.Once
func Get_runEnvT() gopurs_runtime.Value {
	once_runEnvT.Do(func() {
		runEnvT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runEnvT
}

var newtypeEnvT gopurs_runtime.Value
var once_newtypeEnvT sync.Once
func Get_newtypeEnvT() gopurs_runtime.Value {
	once_newtypeEnvT.Do(func() {
		newtypeEnvT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeEnvT
}

var mapEnvT gopurs_runtime.Value
var once_mapEnvT sync.Once
func Get_mapEnvT() gopurs_runtime.Value {
	once_mapEnvT.Do(func() {
		mapEnvT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
	})
	return mapEnvT
}

var functorEnvT gopurs_runtime.Value
var once_functorEnvT sync.Once
func Get_functorEnvT() gopurs_runtime.Value {
	once_functorEnvT.Do(func() {
		functorEnvT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_1), v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
})
	})
	return functorEnvT
}

var functorWithIndexEnvT gopurs_runtime.Value
var once_functorWithIndexEnvT sync.Once
func Get_functorWithIndexEnvT() gopurs_runtime.Value {
	once_functorWithIndexEnvT.Do(func() {
		functorWithIndexEnvT = gopurs_runtime.Func(func(dictFunctorWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorEnvT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctorWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], f_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
})})
})
	})
	return functorWithIndexEnvT
}

var foldableEnvT gopurs_runtime.Value
var once_foldableEnvT sync.Once
func Get_foldableEnvT() gopurs_runtime.Value {
	once_foldableEnvT.Do(func() {
		foldableEnvT = gopurs_runtime.Func(func(dictFoldable_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldl": gopurs_runtime.Func(func(fn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], fn_1), a_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(fn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], fn_1), a_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(dictFoldable_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_1)
return gopurs_runtime.Func(func(fn_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_2_0, fn_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
})})
})
	})
	return foldableEnvT
}

var foldableWithIndexEnvT gopurs_runtime.Value
var once_foldableWithIndexEnvT sync.Once
func Get_foldableWithIndexEnvT() gopurs_runtime.Value {
	once_foldableWithIndexEnvT.Do(func() {
		foldableWithIndexEnvT = gopurs_runtime.Func(func(dictFoldableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable0"], gopurs_runtime.Value{})
foldableEnvT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldl": gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], fn_2), a_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], fn_2), a_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_3_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_2)
return gopurs_runtime.Func(func(fn_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_3_2, fn_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldlWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldlWithIndex"], f_3), a_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldrWithIndex": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldrWithIndex"], f_3), a_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMapWithIndex": gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_4_3 := gopurs_runtime.Apply(dictFoldableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["foldMapWithIndex"], dictMonoid_3)
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMapWithIndex1_4_3, f_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "Foldable0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_2_1
})})
})
	})
	return foldableWithIndexEnvT
}

var traversableEnvT gopurs_runtime.Value
var once_traversableEnvT sync.Once
func Get_traversableEnvT() gopurs_runtime.Value {
	once_traversableEnvT.Do(func() {
		traversableEnvT = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorEnvT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
__local_var_3_2 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["Foldable1"], gopurs_runtime.Value{})
foldableEnvT1_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"foldl": gopurs_runtime.Func(func(fn_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldl"], fn_4), a_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldr": gopurs_runtime.Func(func(fn_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldr"], fn_4), a_5), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
}), "foldMap": gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["foldMap"], dictMonoid_4)
return gopurs_runtime.Func(func(fn_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(foldMap1_5_4, fn_6), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"sequence": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_5 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["sequence"], dictApplicative_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(sequence1_6_5, v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
}), "traverse": gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_6 := gopurs_runtime.Apply(dictTraversable_0.PtrVal.(map[string]gopurs_runtime.Value)["traverse"], dictApplicative_5)
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_5.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(traverse1_6_6, f_7), v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
}), "Foldable1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_4_3
})})
})
	})
	return traversableEnvT
}

var traversableWithIndexEnvT gopurs_runtime.Value
var once_traversableWithIndexEnvT sync.Once
func Get_traversableWithIndexEnvT() gopurs_runtime.Value {
	once_traversableWithIndexEnvT.Do(func() {
		traversableWithIndexEnvT = gopurs_runtime.Func(func(dictTraversableWithIndex_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FunctorWithIndex0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorEnvT1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
functorWithIndexEnvT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mapWithIndex": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["mapWithIndex"], f_4), v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_3_3
})})
foldableWithIndexEnvT1_4_4 := gopurs_runtime.Apply(Get_foldableWithIndexEnvT(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["FoldableWithIndex1"], gopurs_runtime.Value{}))
traversableEnvT1_5_5 := gopurs_runtime.Apply(Get_traversableEnvT(), gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["Traversable2"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"traverseWithIndex": gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex1_7_6 := gopurs_runtime.Apply(dictTraversableWithIndex_0.PtrVal.(map[string]gopurs_runtime.Value)["traverseWithIndex"], dictApplicative_6)
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictApplicative_6.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Apply(gopurs_runtime.Apply(traverseWithIndex1_7_6, f_8), v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))
})
})
}), "FunctorWithIndex0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexEnvT1_3_2
}), "FoldableWithIndex1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexEnvT1_4_4
}), "Traversable2": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableEnvT1_5_5
})})
})
	})
	return traversableWithIndexEnvT
}

var extendEnvT gopurs_runtime.Value
var once_extendEnvT sync.Once
func Get_extendEnvT() gopurs_runtime.Value {
	once_extendEnvT.Do(func() {
		extendEnvT = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorEnvT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], f_3), gopurs_runtime.Apply(gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]))})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
})})
})
	})
	return extendEnvT
}

var comonadTransEnvT gopurs_runtime.Value
var once_comonadTransEnvT sync.Once
func Get_comonadTransEnvT() gopurs_runtime.Value {
	once_comonadTransEnvT.Do(func() {
		comonadTransEnvT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lower": gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
})})
	})
	return comonadTransEnvT
}

var comonadEnvT gopurs_runtime.Value
var once_comonadEnvT sync.Once
func Get_comonadEnvT() gopurs_runtime.Value {
	once_comonadEnvT.Do(func() {
		comonadEnvT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
extendEnvT1_1_0 := gopurs_runtime.Apply(Get_extendEnvT(), gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}), "Extend0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return extendEnvT1_1_0
})})
})
	})
	return comonadEnvT
}


