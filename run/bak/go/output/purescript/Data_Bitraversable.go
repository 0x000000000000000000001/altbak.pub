package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bitraversable_identity gopurs_runtime.Value
var once_Data_Bitraversable_identity sync.Once
func Get_Data_Bitraversable_identity() gopurs_runtime.Value {
	once_Data_Bitraversable_identity.Do(func() {
		cache_Data_Bitraversable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_identity(x_0_box)
})
	})
	return cache_Data_Bitraversable_identity
}

var cache_Data_Bitraversable_identity1 gopurs_runtime.Value
var once_Data_Bitraversable_identity1 sync.Once
func Get_Data_Bitraversable_identity1() gopurs_runtime.Value {
	once_Data_Bitraversable_identity1.Do(func() {
		cache_Data_Bitraversable_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_identity1(x_0_box)
})
	})
	return cache_Data_Bitraversable_identity1
}

var cache_Data_Bitraversable_Bitraversable_dollarDict gopurs_runtime.Value
var once_Data_Bitraversable_Bitraversable_dollarDict sync.Once
func Get_Data_Bitraversable_Bitraversable_dollarDict() gopurs_runtime.Value {
	once_Data_Bitraversable_Bitraversable_dollarDict.Do(func() {
		cache_Data_Bitraversable_Bitraversable_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_Bitraversable_dollarDict(x_0_box)
})
	})
	return cache_Data_Bitraversable_Bitraversable_dollarDict
}

var cache_Data_Bitraversable_bitraverse gopurs_runtime.Value
var once_Data_Bitraversable_bitraverse sync.Once
func Get_Data_Bitraversable_bitraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverse.Do(func() {
		cache_Data_Bitraversable_bitraverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dict_0_box))
})
	})
	return cache_Data_Bitraversable_bitraverse
}

var cache_Data_Bitraversable_lfor gopurs_runtime.Value
var once_Data_Bitraversable_lfor sync.Once
func Get_Data_Bitraversable_lfor() gopurs_runtime.Value {
	once_Data_Bitraversable_lfor.Do(func() {
		cache_Data_Bitraversable_lfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_lfor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_lfor
}

var cache_Data_Bitraversable_ltraverse gopurs_runtime.Value
var once_Data_Bitraversable_ltraverse sync.Once
func Get_Data_Bitraversable_ltraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_ltraverse.Do(func() {
		cache_Data_Bitraversable_ltraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_ltraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_ltraverse
}

var cache_Data_Bitraversable_rfor gopurs_runtime.Value
var once_Data_Bitraversable_rfor sync.Once
func Get_Data_Bitraversable_rfor() gopurs_runtime.Value {
	once_Data_Bitraversable_rfor.Do(func() {
		cache_Data_Bitraversable_rfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_rfor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_rfor
}

var cache_Data_Bitraversable_rtraverse gopurs_runtime.Value
var once_Data_Bitraversable_rtraverse sync.Once
func Get_Data_Bitraversable_rtraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_rtraverse.Do(func() {
		cache_Data_Bitraversable_rtraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_rtraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_rtraverse
}

var cache_Data_Bitraversable_bitraversableTuple gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableTuple sync.Once
func Get_Data_Bitraversable_bitraversableTuple() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableTuple.Do(func() {
		cache_Data_Bitraversable_bitraversableTuple = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifoldable_bifoldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifunctor_bifunctorTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_2 -> *Constructor_Control_Apply_Apply
Apply0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_2
// TAST (Let): Functor0_2_3 -> *Constructor_Data_Functor_Functor
Functor0_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_3
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_3.V0), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)), gopurs_runtime.Apply(g_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1))
})
})
})
}))
	})
	return cache_Data_Bitraversable_bitraversableTuple
}

var cache_Data_Bitraversable_bitraversableJoker gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableJoker sync.Once
func Get_Data_Bitraversable_bitraversableJoker() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableJoker.Do(func() {
		cache_Data_Bitraversable_bitraversableJoker = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableJoker(dictTraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableJoker
}

var cache_Data_Bitraversable_bitraversableEither gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableEither sync.Once
func Get_Data_Bitraversable_bitraversableEither() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableEither.Do(func() {
		cache_Data_Bitraversable_bitraversableEither = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifoldable_bifoldableEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifunctor_bifunctorEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Left(), (*Constructor_Data_Either_Left)(v_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), (*Constructor_Data_Either_Right)(v_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_2 -> *Constructor_Data_Functor_Functor
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Either_Left(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Left)(v2_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Either_Right(), gopurs_runtime.Apply(v1_3, (*Constructor_Data_Either_Right)(v2_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})
})
}))
	})
	return cache_Data_Bitraversable_bitraversableEither
}

var cache_Data_Bitraversable_bitraversableConst gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableConst sync.Once
func Get_Data_Bitraversable_bitraversableConst() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableConst.Do(func() {
		cache_Data_Bitraversable_bitraversableConst = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifoldable_bifoldableConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Bifunctor_bifunctorConst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Const_Const(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Const_Const(), gopurs_runtime.Apply(f_2, v1_4))
})
})
})
}))
	})
	return cache_Data_Bitraversable_bitraversableConst
}

var cache_Data_Bitraversable_bitraversableClown gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableClown sync.Once
func Get_Data_Bitraversable_bitraversableClown() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableClown.Do(func() {
		cache_Data_Bitraversable_bitraversableClown = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableClown(dictTraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableClown
}

var cache_Data_Bitraversable_bisequenceDefault gopurs_runtime.Value
var once_Data_Bitraversable_bisequenceDefault sync.Once
func Get_Data_Bitraversable_bisequenceDefault() gopurs_runtime.Value {
	once_Data_Bitraversable_bisequenceDefault.Do(func() {
		cache_Data_Bitraversable_bisequenceDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bisequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_bisequenceDefault
}

var cache_Data_Bitraversable_bisequence gopurs_runtime.Value
var once_Data_Bitraversable_bisequence sync.Once
func Get_Data_Bitraversable_bisequence() gopurs_runtime.Value {
	once_Data_Bitraversable_bisequence.Do(func() {
		cache_Data_Bitraversable_bisequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bisequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dict_0_box))
})
	})
	return cache_Data_Bitraversable_bisequence
}

var cache_Data_Bitraversable_bitraversableFlip gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableFlip sync.Once
func Get_Data_Bitraversable_bitraversableFlip() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableFlip.Do(func() {
		cache_Data_Bitraversable_bitraversableFlip = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableFlip(dictBitraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableFlip
}

var cache_Data_Bitraversable_bitraversableProduct2 gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableProduct2 sync.Once
func Get_Data_Bitraversable_bitraversableProduct2() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableProduct2.Do(func() {
		cache_Data_Bitraversable_bitraversableProduct2 = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableProduct2(dictBitraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableProduct2
}

var cache_Data_Bitraversable_bitraverseDefault gopurs_runtime.Value
var once_Data_Bitraversable_bitraverseDefault sync.Once
func Get_Data_Bitraversable_bitraverseDefault() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverseDefault.Do(func() {
		cache_Data_Bitraversable_bitraverseDefault = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dictBitraversable_0_box))
})
	})
	return cache_Data_Bitraversable_bitraverseDefault
}

var cache_Data_Bitraversable_bifor gopurs_runtime.Value
var once_Data_Bitraversable_bifor sync.Once
func Get_Data_Bitraversable_bifor() gopurs_runtime.Value {
	once_Data_Bitraversable_bifor.Do(func() {
		cache_Data_Bitraversable_bifor = gopurs_runtime.Func5(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value, g_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bifor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box), t_2_box, f_3_box, g_4_box)
})
	})
	return cache_Data_Bitraversable_bifor
}

var cache_Data_Bitraversable_bisequence__3517827200 gopurs_runtime.Value
var once_Data_Bitraversable_bisequence__3517827200 sync.Once
func Get_Data_Bitraversable_bisequence__3517827200() gopurs_runtime.Value {
	once_Data_Bitraversable_bisequence__3517827200.Do(func() {
		cache_Data_Bitraversable_bisequence__3517827200 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bisequence__3517827200(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dict_0_box))
})
	})
	return cache_Data_Bitraversable_bisequence__3517827200
}

var cache_Data_Bitraversable_bitraverse__3884078439 gopurs_runtime.Value
var once_Data_Bitraversable_bitraverse__3884078439 sync.Once
func Get_Data_Bitraversable_bitraverse__3884078439() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverse__3884078439.Do(func() {
		cache_Data_Bitraversable_bitraverse__3884078439 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraverse__3884078439(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dict_0_box))
})
	})
	return cache_Data_Bitraversable_bitraverse__3884078439
}

var cache_Data_Bitraversable_bitraverse__1091604167 gopurs_runtime.Value
var once_Data_Bitraversable_bitraverse__1091604167 sync.Once
func Get_Data_Bitraversable_bitraverse__1091604167() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverse__1091604167.Do(func() {
		cache_Data_Bitraversable_bitraverse__1091604167 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraverse__1091604167(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable](dict_0_box))
})
	})
	return cache_Data_Bitraversable_bitraverse__1091604167
}

var cache_Data_Bitraversable_bitraverse__4064111983 gopurs_runtime.Value
var once_Data_Bitraversable_bitraverse__4064111983 sync.Once
func Get_Data_Bitraversable_bitraverse__4064111983() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverse__4064111983.Do(func() {
		cache_Data_Bitraversable_bitraverse__4064111983 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraverse__4064111983(dictApplicative_0_box)
})
	})
	return cache_Data_Bitraversable_bitraverse__4064111983
}

type Constructor_Data_Bitraversable_Bitraversable struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3704227322] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bitraversable_Bitraversable)(ptr)
		_ = c
		switch key {
		case "Bifoldable1": return gopurs_runtime.Box(c.V0)
		case "Bifunctor0": return gopurs_runtime.Box(c.V1)
		case "bisequence": return gopurs_runtime.Box(c.V2)
		case "bitraverse": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Bitraversable_Bitraversable: " + key)
		}
	}
}


func Call_Data_Bitraversable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bitraversable_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bitraversable_Bitraversable_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bitraversable_bitraverse(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bitraversable_Bitraversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Bitraversable_lfor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): pure_2_0 -> gopurs_runtime.Value
pure_2_0 := gopurs_runtime.Box(dictApplicative_1.V1)
_ = pure_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_4, pure_2_0, t_3)
})
})
}

func Call_Data_Bitraversable_ltraverse(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): pure_2_0 -> gopurs_runtime.Value
pure_2_0 := gopurs_runtime.Box(dictApplicative_1.V1)
_ = pure_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_3, pure_2_0)
})
}

func Call_Data_Bitraversable_rfor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): pure_2_0 -> gopurs_runtime.Value
pure_2_0 := gopurs_runtime.Box(dictApplicative_1.V1)
_ = pure_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, pure_2_0, f_4, t_3)
})
})
}

func Call_Data_Bitraversable_rtraverse(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Box(dictApplicative_1.V1))
}

func Call_Data_Bitraversable_bitraversableJoker(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorJoker_1_0 -> gopurs_runtime.Value
bifunctorJoker_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), g_3, v1_4)
})
})
}))
_ = bifunctorJoker_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifoldableJoker_2_2 -> gopurs_runtime.Value
bifoldableJoker_2_2 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, r_5, v1_6)
})
})
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), r_4, u_5, v1_6)
})
})
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), r_4, u_5, v1_6)
})
})
})
}))
_ = bifoldableJoker_2_2
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableJoker_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_Joker_Joker(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 -> *Constructor_Data_Functor_Functor
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), Get_Data_Functor_Joker_Joker(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, r_6, v1_7))
})
})
})
}))
}

func Call_Data_Bitraversable_bitraversableClown(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorClown_1_0 -> gopurs_runtime.Value
bifunctorClown_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, v1_4)
})
})
}))
_ = bifunctorClown_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifoldableClown_2_2 -> gopurs_runtime.Value
bifoldableClown_2_2 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, l_4, v1_6)
})
})
})
}), gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), l_3, u_5, v1_6)
})
})
})
}), gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), l_3, u_5, v1_6)
})
})
})
}))
_ = bifoldableClown_2_2
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableClown_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_Clown_Clown(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 -> *Constructor_Data_Functor_Functor
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), Get_Data_Functor_Clown_Clown(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, l_5, v1_7))
})
})
})
}))
}

func Call_Data_Bitraversable_bisequenceDefault(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, Get_Data_Bitraversable_identity(), Get_Data_Bitraversable_identity1())
}

func Call_Data_Bitraversable_bisequence(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bitraversable_Bitraversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bitraversable_bitraversableFlip(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorFlip_1_0 -> gopurs_runtime.Value
bifunctorFlip_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), g_3, f_2, v_4)
})
})
}))
_ = bifunctorFlip_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifoldableFlip_2_2 -> gopurs_runtime.Value
bifoldableFlip_2_2 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, l_5, r_4, v_6)
})
})
})
}), gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldl"), l_4, r_3, u_5, v_6)
})
})
})
}), gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldr"), l_4, r_3, u_5, v_6)
})
})
})
}))
_ = bifoldableFlip_2_2
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableFlip_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_Flip_Flip(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 -> *Constructor_Data_Functor_Functor
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), Get_Data_Functor_Flip_Flip(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, l_6, r_5, v_7))
})
})
})
}))
}

func Call_Data_Bitraversable_bitraversableProduct2(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): bifoldableProduct2_2_1 -> gopurs_runtime.Value
bifoldableProduct2_2_1 := gopurs_runtime.Apply(Get_Data_Bifoldable_bifoldableProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct2_2_1
return gopurs_runtime.Func(func(dictBitraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): bifunctorProduct21_4_2 -> gopurs_runtime.Value
bifunctorProduct21_4_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Product2_Product2{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_5, g_6, (*Constructor_Data_Functor_Product2_Product2)(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_3, "bimap"), f_5, g_6, (*Constructor_Data_Functor_Product2_Product2)(v_7.UnsafePtr).V1)})}
})
})
}))
_ = bifunctorProduct21_4_2
// TAST (Let): bifoldableProduct21_5_4 -> gopurs_runtime.Value
bifoldableProduct21_5_4 := gopurs_runtime.Apply(bifoldableProduct2_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct21_5_4
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableProduct21_5_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct21_4_2
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_5 -> *Constructor_Control_Apply_Apply
Apply0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_5
// TAST (Let): Functor0_8_6 -> *Constructor_Data_Functor_Functor
Functor0_8_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_6
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_5.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_6.V0), Get_Data_Functor_Product2_Product2(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, (*Constructor_Data_Functor_Product2_Product2)(v_9.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable1_3, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, (*Constructor_Data_Functor_Product2_Product2)(v_9.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_7 -> *Constructor_Control_Apply_Apply
Apply0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_7
// TAST (Let): Functor0_8_8 -> *Constructor_Data_Functor_Functor
Functor0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_8
return gopurs_runtime.Func(func(l_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_7.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_8.V0), Get_Data_Functor_Product2_Product2(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, l_9, r_10, (*Constructor_Data_Functor_Product2_Product2)(v_11.UnsafePtr).V0)), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable1_3, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, l_9, r_10, (*Constructor_Data_Functor_Product2_Product2)(v_11.UnsafePtr).V1))
})
})
})
}))
})
}

func Call_Data_Bitraversable_bitraverseDefault(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable = dictBitraversable_0_loop
_ = dictBitraversable_0
// TAST (Let): Bifunctor0_1_0 -> *Constructor_Data_Bifunctor_Bifunctor
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor](gopurs_runtime.Apply(gopurs_runtime.Box(dictBitraversable_0.V1), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBitraversable_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_2))}, gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), f_3, g_4, t_5))
})
})
})
})
}

func Call_Data_Bitraversable_bifor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative, t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value, g_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
_ = dictApplicative_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
var g_4 gopurs_runtime.Value = g_4_loop
_ = g_4
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_3, g_4, t_2)
}

func Call_Data_Bitraversable_bisequence__3517827200(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bitraversable_Bitraversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bitraversable_bitraverse__3884078439(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bitraversable_Bitraversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Bitraversable_bitraverse__1091604167(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bitraversable_Bitraversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Bitraversable_bitraverse__4064111983(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 237113226) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 1992629780) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_DurationEnd)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2020675835) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartDuration(), gopurs_runtime.Apply(v1_4, (*Constructor_Data_Interval_StartDuration)(v2_5.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration)(v2_5.UnsafePtr).V1))
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 2281256335) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_DurationOnly(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationOnly)(v2_5.UnsafePtr).V0))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
})
})
}


