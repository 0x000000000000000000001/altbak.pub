package Control_Monad_List_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Lazy "gopurs/output/Data.Lazy"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_identity(gopurs_runtime.UnboxAny(x_0_box)))
})
	})
	return cache_identity
}

var cache_Yield gopurs_runtime.Value
var once_Yield sync.Once
func Get_Yield() gopurs_runtime.Value {
	once_Yield.Do(func() {
		cache_Yield = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(value0), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(value1, inner_arg0))
}})}
})
})
	})
	return cache_Yield
}

var cache_Skip gopurs_runtime.Value
var once_Skip sync.Once
func Get_Skip() gopurs_runtime.Value {
	once_Skip.Do(func() {
		cache_Skip = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(value0, inner_arg0))
}})}
})
	})
	return cache_Skip
}

var cache_Done gopurs_runtime.Value
var once_Done sync.Once
func Get_Done() gopurs_runtime.Value {
	once_Done.Do(func() {
		cache_Done = gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}
	})
	return cache_Done
}

var cache_ListT gopurs_runtime.Value
var once_ListT sync.Once
func Get_ListT() gopurs_runtime.Value {
	once_ListT.Do(func() {
		cache_ListT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ListT(x_0_box)
})
	})
	return cache_ListT
}

var cache_wrapLazy gopurs_runtime.Value
var once_wrapLazy sync.Once
func Get_wrapLazy() gopurs_runtime.Value {
	once_wrapLazy.Do(func() {
		cache_wrapLazy = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_wrapLazy(dictApplicative_0_box, func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_1_box, inner_arg0))
}))
})
	})
	return cache_wrapLazy
}

var cache_wrapEffect gopurs_runtime.Value
var once_wrapEffect sync.Once
func Get_wrapEffect() gopurs_runtime.Value {
	once_wrapEffect.Do(func() {
		cache_wrapEffect = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_wrapEffect(dictFunctor_0_box, gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_wrapEffect
}

var cache_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225 gopurs_runtime.Value
var once_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225 sync.Once
func Get_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225() gopurs_runtime.Value {
	once_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225.Do(func() {
		cache_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225(dictFunctor_0_box, gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225
}

var cache_unfold gopurs_runtime.Value
var once_unfold sync.Once
func Get_unfold() gopurs_runtime.Value {
	once_unfold.Do(func() {
		cache_unfold = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unfold(dictMonad_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(z_2_box)))
})
	})
	return cache_unfold
}

var cache_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807 gopurs_runtime.Value
var once_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807 sync.Once
func Get_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807() gopurs_runtime.Value {
	once_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807.Do(func() {
		cache_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807 = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807(dictMonad_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(z_2_box)))
})
	})
	return cache_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(dictMonad_0_box)
})
	})
	return cache_uncons
}

var cache_uncons__func_gopurs_runtime_Value__interface____interface___1472790614 gopurs_runtime.Value
var once_uncons__func_gopurs_runtime_Value__interface____interface___1472790614 sync.Once
func Get_uncons__func_gopurs_runtime_Value__interface____interface___1472790614() gopurs_runtime.Value {
	once_uncons__func_gopurs_runtime_Value__interface____interface___1472790614.Do(func() {
		cache_uncons__func_gopurs_runtime_Value__interface____interface___1472790614 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons__func_gopurs_runtime_Value__interface____interface___1472790614(dictMonad_0_box)
})
	})
	return cache_uncons__func_gopurs_runtime_Value__interface____interface___1472790614
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail(dictMonad_0_box)
})
	})
	return cache_tail
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(dictApplicative_0_box)
})
	})
	return cache_takeWhile
}

var cache_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412 gopurs_runtime.Value
var once_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412 sync.Once
func Get_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412() gopurs_runtime.Value {
	once_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412.Do(func() {
		cache_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412(dictApplicative_0_box)
})
	})
	return cache_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value, l_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_scanl(dictMonad_0_box, func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(b_2_box), gopurs_runtime.UnboxAny(l_3_box)))
})
	})
	return cache_scanl
}

var cache_prepend_prime gopurs_runtime.Value
var once_prepend_prime sync.Once
func Get_prepend_prime() gopurs_runtime.Value {
	once_prepend_prime.Do(func() {
		cache_prepend_prime = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_prepend_prime(dictApplicative_0_box, gopurs_runtime.UnboxAny(h_1_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(t_2_box, inner_arg0))
}))
})
	})
	return cache_prepend_prime
}

var cache_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542 gopurs_runtime.Value
var once_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542 sync.Once
func Get_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542() gopurs_runtime.Value {
	once_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542.Do(func() {
		cache_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542(dictApplicative_0_box, gopurs_runtime.UnboxAny(h_1_box), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(t_2_box, inner_arg0))
}))
})
	})
	return cache_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542
}

var cache_prepend gopurs_runtime.Value
var once_prepend sync.Once
func Get_prepend() gopurs_runtime.Value {
	once_prepend.Do(func() {
		cache_prepend = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_prepend(dictApplicative_0_box, gopurs_runtime.UnboxAny(h_1_box), gopurs_runtime.UnboxAny(t_2_box)))
})
	})
	return cache_prepend
}

var cache_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070 gopurs_runtime.Value
var once_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070 sync.Once
func Get_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070() gopurs_runtime.Value {
	once_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070.Do(func() {
		cache_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, h_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070(dictApplicative_0_box, gopurs_runtime.UnboxAny(h_1_box), gopurs_runtime.UnboxAny(t_2_box)))
})
	})
	return cache_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070
}

var cache_nil gopurs_runtime.Value
var once_nil sync.Once
func Get_nil() gopurs_runtime.Value {
	once_nil.Do(func() {
		cache_nil = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_nil(dictApplicative_0_box))
})
	})
	return cache_nil
}

var cache_nil__func_gopurs_runtime_Value__interface___30547265 gopurs_runtime.Value
var once_nil__func_gopurs_runtime_Value__interface___30547265 sync.Once
func Get_nil__func_gopurs_runtime_Value__interface___30547265() gopurs_runtime.Value {
	once_nil__func_gopurs_runtime_Value__interface___30547265.Do(func() {
		cache_nil__func_gopurs_runtime_Value__interface___30547265 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_nil__func_gopurs_runtime_Value__interface___30547265(dictApplicative_0_box))
})
	})
	return cache_nil__func_gopurs_runtime_Value__interface___30547265
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(dictApplicative_0_box)
})
	})
	return cache_singleton
}

var cache_singleton__func_gopurs_runtime_Value__interface____interface___3816438021 gopurs_runtime.Value
var once_singleton__func_gopurs_runtime_Value__interface____interface___3816438021 sync.Once
func Get_singleton__func_gopurs_runtime_Value__interface____interface___3816438021() gopurs_runtime.Value {
	once_singleton__func_gopurs_runtime_Value__interface____interface___3816438021.Do(func() {
		cache_singleton__func_gopurs_runtime_Value__interface____interface___3816438021 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton__func_gopurs_runtime_Value__interface____interface___3816438021(dictApplicative_0_box)
})
	})
	return cache_singleton__func_gopurs_runtime_Value__interface____interface___3816438021
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take(dictApplicative_0_box)
})
	})
	return cache_take
}

var cache_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118 gopurs_runtime.Value
var once_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118 sync.Once
func Get_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118() gopurs_runtime.Value {
	once_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118.Do(func() {
		cache_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118(dictApplicative_0_box)
})
	})
	return cache_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118
}

var cache_zipWith_prime gopurs_runtime.Value
var once_zipWith_prime sync.Once
func Get_zipWith_prime() gopurs_runtime.Value {
	once_zipWith_prime.Do(func() {
		cache_zipWith_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith_prime(dictMonad_0_box)
})
	})
	return cache_zipWith_prime
}

var cache_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796 gopurs_runtime.Value
var once_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796 sync.Once
func Get_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796() gopurs_runtime.Value {
	once_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796.Do(func() {
		cache_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796(dictMonad_0_box)
})
	})
	return cache_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796
}

var cache_zipWith gopurs_runtime.Value
var once_zipWith sync.Once
func Get_zipWith() gopurs_runtime.Value {
	once_zipWith.Do(func() {
		cache_zipWith = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zipWith(dictMonad_0_box)
})
	})
	return cache_zipWith
}

var cache_newtypeListT gopurs_runtime.Value
var once_newtypeListT sync.Once
func Get_newtypeListT() gopurs_runtime.Value {
	once_newtypeListT.Do(func() {
		cache_newtypeListT = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))))
	})
	return cache_newtypeListT
}

var cache_mapMaybe gopurs_runtime.Value
var once_mapMaybe sync.Once
func Get_mapMaybe() gopurs_runtime.Value {
	once_mapMaybe.Do(func() {
		cache_mapMaybe = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mapMaybe(dictFunctor_0_box, func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, gopurs_runtime.UnboxAny(v_2_box)))
})
	})
	return cache_mapMaybe
}

var cache_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890 gopurs_runtime.Value
var once_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890 sync.Once
func Get_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890() gopurs_runtime.Value {
	once_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890.Do(func() {
		cache_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890(dictFunctor_0_box, func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, gopurs_runtime.UnboxAny(v_2_box)))
})
	})
	return cache_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890
}

var cache_iterate gopurs_runtime.Value
var once_iterate sync.Once
func Get_iterate() gopurs_runtime.Value {
	once_iterate.Do(func() {
		cache_iterate = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_iterate(dictMonad_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_iterate
}

var cache_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728 gopurs_runtime.Value
var once_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728 sync.Once
func Get_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728() gopurs_runtime.Value {
	once_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728.Do(func() {
		cache_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728 = gopurs_runtime.Func3(func(dictMonad_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728(dictMonad_0_box, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(a_2_box)))
})
	})
	return cache_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728
}

var cache_repeat gopurs_runtime.Value
var once_repeat sync.Once
func Get_repeat() gopurs_runtime.Value {
	once_repeat.Do(func() {
		cache_repeat = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_repeat(dictMonad_0_box)
})
	})
	return cache_repeat
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(dictMonad_0_box)
})
	})
	return cache_head
}

var cache_functorListT gopurs_runtime.Value
var once_functorListT sync.Once
func Get_functorListT() gopurs_runtime.Value {
	once_functorListT.Do(func() {
		cache_functorListT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_functorListT(dictFunctor_0_box))
})
	})
	return cache_functorListT
}

var cache_fromEffect gopurs_runtime.Value
var once_fromEffect sync.Once
func Get_fromEffect() gopurs_runtime.Value {
	once_fromEffect.Do(func() {
		cache_fromEffect = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEffect(dictApplicative_0_box)
})
	})
	return cache_fromEffect
}

var cache_monadTransListT gopurs_runtime.Value
var once_monadTransListT sync.Once
func Get_monadTransListT() gopurs_runtime.Value {
	once_monadTransListT.Do(func() {
		cache_monadTransListT = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_2_1
return gopurs_runtime.Func(func(fa_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_4), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
}), inner_arg0))
}})})
}), fa_3)
})
}))))
	})
	return cache_monadTransListT
}

var cache_foldlRec_prime gopurs_runtime.Value
var once_foldlRec_prime sync.Once
func Get_foldlRec_prime() gopurs_runtime.Value {
	once_foldlRec_prime.Do(func() {
		cache_foldlRec_prime = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec_prime(dictMonadRec_0_box)
})
	})
	return cache_foldlRec_prime
}

var cache_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185 gopurs_runtime.Value
var once_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185 sync.Once
func Get_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185() gopurs_runtime.Value {
	once_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185.Do(func() {
		cache_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185(dictMonadRec_0_box)
})
	})
	return cache_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185
}

var cache_runListTRec gopurs_runtime.Value
var once_runListTRec sync.Once
func Get_runListTRec() gopurs_runtime.Value {
	once_runListTRec.Do(func() {
		cache_runListTRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListTRec(dictMonadRec_0_box)
})
	})
	return cache_runListTRec
}

var cache_foldlRec gopurs_runtime.Value
var once_foldlRec sync.Once
func Get_foldlRec() gopurs_runtime.Value {
	once_foldlRec.Do(func() {
		cache_foldlRec = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlRec(dictMonadRec_0_box)
})
	})
	return cache_foldlRec
}

var cache_foldl_prime gopurs_runtime.Value
var once_foldl_prime sync.Once
func Get_foldl_prime() gopurs_runtime.Value {
	once_foldl_prime.Do(func() {
		cache_foldl_prime = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_prime(dictMonad_0_box)
})
	})
	return cache_foldl_prime
}

var cache_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616 gopurs_runtime.Value
var once_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616 sync.Once
func Get_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616() gopurs_runtime.Value {
	once_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616.Do(func() {
		cache_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616(dictMonad_0_box)
})
	})
	return cache_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616
}

var cache_runListT gopurs_runtime.Value
var once_runListT sync.Once
func Get_runListT() gopurs_runtime.Value {
	once_runListT.Do(func() {
		cache_runListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runListT(dictMonad_0_box)
})
	})
	return cache_runListT
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(dictMonad_0_box)
})
	})
	return cache_foldl
}

var cache_filter gopurs_runtime.Value
var once_filter sync.Once
func Get_filter() gopurs_runtime.Value {
	once_filter.Do(func() {
		cache_filter = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_filter(dictFunctor_0_box, func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, gopurs_runtime.UnboxAny(v_2_box)))
})
	})
	return cache_filter
}

var cache_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519 gopurs_runtime.Value
var once_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519 sync.Once
func Get_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519() gopurs_runtime.Value {
	once_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519.Do(func() {
		cache_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519(dictFunctor_0_box, func(inner_arg0 interface{}) bool {
return (gopurs_runtime.Apply(f_1_box, gopurs_runtime.Any(inner_arg0)).IntVal) != (0)
}, gopurs_runtime.UnboxAny(v_2_box)))
})
	})
	return cache_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(dictApplicative_0_box)
})
	})
	return cache_dropWhile
}

var cache_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412 gopurs_runtime.Value
var once_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412 sync.Once
func Get_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412() gopurs_runtime.Value {
	once_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412.Do(func() {
		cache_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412(dictApplicative_0_box)
})
	})
	return cache_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop(dictApplicative_0_box)
})
	})
	return cache_drop
}

var cache_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118 gopurs_runtime.Value
var once_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118 sync.Once
func Get_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118() gopurs_runtime.Value {
	once_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118.Do(func() {
		cache_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118(dictApplicative_0_box)
})
	})
	return cache_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_cons(dictApplicative_0_box, func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(lh_1_box, inner_arg0))
}, func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(t_2_box, inner_arg0))
}))
})
	})
	return cache_cons
}

var cache_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302 gopurs_runtime.Value
var once_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302 sync.Once
func Get_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302() gopurs_runtime.Value {
	once_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302.Do(func() {
		cache_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302 = gopurs_runtime.Func3(func(dictApplicative_0_box gopurs_runtime.Value, lh_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302(dictApplicative_0_box, func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(lh_1_box, inner_arg0))
}, func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(t_2_box, inner_arg0))
}))
})
	})
	return cache_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302
}

var cache_unfoldable1ListT gopurs_runtime.Value
var once_unfoldable1ListT sync.Once
func Get_unfoldable1ListT() gopurs_runtime.Value {
	once_unfoldable1ListT.Do(func() {
		cache_unfoldable1ListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unfoldable1ListT(dictMonad_0_box))
})
	})
	return cache_unfoldable1ListT
}

var cache_unfoldableListT gopurs_runtime.Value
var once_unfoldableListT sync.Once
func Get_unfoldableListT() gopurs_runtime.Value {
	once_unfoldableListT.Do(func() {
		cache_unfoldableListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unfoldableListT(dictMonad_0_box))
})
	})
	return cache_unfoldableListT
}

var cache_semigroupListT gopurs_runtime.Value
var once_semigroupListT sync.Once
func Get_semigroupListT() gopurs_runtime.Value {
	once_semigroupListT.Do(func() {
		cache_semigroupListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_semigroupListT(dictApplicative_0_box))
})
	})
	return cache_semigroupListT
}

var cache_concat gopurs_runtime.Value
var once_concat sync.Once
func Get_concat() gopurs_runtime.Value {
	once_concat.Do(func() {
		cache_concat = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_concat(dictApplicative_0_box)
})
	})
	return cache_concat
}

var cache_monoidListT gopurs_runtime.Value
var once_monoidListT sync.Once
func Get_monoidListT() gopurs_runtime.Value {
	once_monoidListT.Do(func() {
		cache_monoidListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monoidListT(dictApplicative_0_box))
})
	})
	return cache_monoidListT
}

var cache_catMaybes gopurs_runtime.Value
var once_catMaybes sync.Once
func Get_catMaybes() gopurs_runtime.Value {
	once_catMaybes.Do(func() {
		cache_catMaybes = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_catMaybes(dictFunctor_0_box)
})
	})
	return cache_catMaybes
}

var cache_monadListT gopurs_runtime.Value
var once_monadListT sync.Once
func Get_monadListT() gopurs_runtime.Value {
	once_monadListT.Do(func() {
		cache_monadListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadListT(dictMonad_0_box))
})
	})
	return cache_monadListT
}

var cache_bindListT gopurs_runtime.Value
var once_bindListT sync.Once
func Get_bindListT() gopurs_runtime.Value {
	once_bindListT.Do(func() {
		cache_bindListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindListT(dictMonad_0_box))
})
	})
	return cache_bindListT
}

var cache_applyListT gopurs_runtime.Value
var once_applyListT sync.Once
func Get_applyListT() gopurs_runtime.Value {
	once_applyListT.Do(func() {
		cache_applyListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applyListT(dictMonad_0_box))
})
	})
	return cache_applyListT
}

var cache_applicativeListT gopurs_runtime.Value
var once_applicativeListT sync.Once
func Get_applicativeListT() gopurs_runtime.Value {
	once_applicativeListT.Do(func() {
		cache_applicativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applicativeListT(dictMonad_0_box))
})
	})
	return cache_applicativeListT
}

var cache_monadEffectListT gopurs_runtime.Value
var once_monadEffectListT sync.Once
func Get_monadEffectListT() gopurs_runtime.Value {
	once_monadEffectListT.Do(func() {
		cache_monadEffectListT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadEffectListT(dictMonadEffect_0_box))
})
	})
	return cache_monadEffectListT
}

var cache_monadSTListT gopurs_runtime.Value
var once_monadSTListT sync.Once
func Get_monadSTListT() gopurs_runtime.Value {
	once_monadSTListT.Do(func() {
		cache_monadSTListT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadSTListT(dictMonadST_0_box))
})
	})
	return cache_monadSTListT
}

var cache_altListT gopurs_runtime.Value
var once_altListT sync.Once
func Get_altListT() gopurs_runtime.Value {
	once_altListT.Do(func() {
		cache_altListT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_altListT(dictApplicative_0_box))
})
	})
	return cache_altListT
}

var cache_plusListT gopurs_runtime.Value
var once_plusListT sync.Once
func Get_plusListT() gopurs_runtime.Value {
	once_plusListT.Do(func() {
		cache_plusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_plusListT(dictMonad_0_box))
})
	})
	return cache_plusListT
}

var cache_alternativeListT gopurs_runtime.Value
var once_alternativeListT sync.Once
func Get_alternativeListT() gopurs_runtime.Value {
	once_alternativeListT.Do(func() {
		cache_alternativeListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_alternativeListT(dictMonad_0_box))
})
	})
	return cache_alternativeListT
}

var cache_monadPlusListT gopurs_runtime.Value
var once_monadPlusListT sync.Once
func Get_monadPlusListT() gopurs_runtime.Value {
	once_monadPlusListT.Do(func() {
		cache_monadPlusListT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monadPlusListT(dictMonad_0_box))
})
	})
	return cache_monadPlusListT
}

type Constructor_Yield[T_a any, T_s any] struct {
	V0 T_a
	V1 func(gopurs_runtime.Value) interface{}
}


type Constructor_Skip[T_a any, T_s any] struct {
	V0 func(gopurs_runtime.Value) interface{}
}


type Constructor_Done[T_a any, T_s any] struct {
	
}


func Call_identity(x_0_loop interface{}) interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
return x_0
}

func Call_ListT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_wrapLazy(dictApplicative_0_loop gopurs_runtime.Value, v_1_loop func(gopurs_runtime.Value) interface{}) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var v_1 func(gopurs_runtime.Value) interface{} = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{v_1})})))
}

func Call_wrapEffect(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), inner_arg0))
}})})
}), gopurs_runtime.Any(v_1)))
}

func Call_wrapEffect__func_gopurs_runtime_Value__interface____interface___2088361225(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), inner_arg0))
}})})
}), gopurs_runtime.Any(v_1)))
}

func Call_unfold(dictMonad_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, z_2_loop interface{}) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var z_2 interface{} = z_2_loop
_ = z_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil) {
__local_var_4_1 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0).UnsafePtr).V0)
_ = __local_var_4_1
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0).UnsafePtr).V1)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unfold(dictMonad_0, f_1, gopurs_runtime.UnboxAny(__local_var_4_1)))
}), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Any(f_1(z_2))))
}

func Call_unfold__func_gopurs_runtime_Value__func_interface____interface____interface____interface___917382807(dictMonad_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, z_2_loop interface{}) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var z_2 interface{} = z_2_loop
_ = z_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil) {
__local_var_4_1 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0).UnsafePtr).V0)
_ = __local_var_4_1
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0).UnsafePtr).V1)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unfold(dictMonad_0, f_1, gopurs_runtime.UnboxAny(__local_var_4_1)))
}), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Any(f_1(z_2))))
}

func Call_uncons(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), v_2, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1320412129) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1(arg0))
}), pkg_Data_Unit.Get_unit()))})}})}))
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 813447293) {
__t1 = gopurs_runtime.Apply(Call_uncons(dictMonad_0), gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0(arg0))
}), pkg_Data_Unit.Get_unit()))
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 489128924) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
})
}

func Call_uncons__func_gopurs_runtime_Value__interface____interface___1472790614(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), v_2, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1320412129) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)), gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1(arg0))
}), pkg_Data_Unit.Get_unit()))})}})}))
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 813447293) {
__t1 = gopurs_runtime.Apply(Call_uncons(dictMonad_0), gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0(arg0))
}), pkg_Data_Unit.Get_unit()))
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 489128924) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
})
}

func Call_tail(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := Call_uncons(dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_snd()), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
}

func Call_takeWhile(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)).IntVal) != (0) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}

func Call_takeWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)).IntVal) != (0) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_takeWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}

func Call_scanl(dictMonad_0_loop gopurs_runtime.Value, f_1_loop func(interface{}, interface{}) interface{}, b_2_loop interface{}, l_3_loop interface{}) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 func(interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var b_2 interface{} = b_2_loop
_ = b_2
var l_3 interface{} = l_3_loop
_ = l_3
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(Call_unfold(dictMonad_0, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = __local_var_5_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 1320412129) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(__local_var_5_0), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0))))), gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V1(arg0))
}), pkg_Data_Unit.Get_unit()))})}, gopurs_runtime.UnboxAny(__local_var_5_0)})}})})
goto end_branch_1
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(__local_var_5_0), gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0(arg0))
}), pkg_Data_Unit.Get_unit()))})}, gopurs_runtime.UnboxAny(__local_var_5_0)})}})})
goto end_branch_1
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{b_2, l_3})})))
}

func Call_prepend_prime(dictApplicative_0_loop gopurs_runtime.Value, h_1_loop interface{}, t_2_loop func(gopurs_runtime.Value) interface{}) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 interface{} = h_1_loop
_ = h_1
var t_2 func(gopurs_runtime.Value) interface{} = t_2_loop
_ = t_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{h_1, t_2})})))
}

func Call_prepend_prime__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___584341542(dictApplicative_0_loop gopurs_runtime.Value, h_1_loop interface{}, t_2_loop func(gopurs_runtime.Value) interface{}) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 interface{} = h_1_loop
_ = h_1
var t_2 func(gopurs_runtime.Value) interface{} = t_2_loop
_ = t_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{h_1, t_2})})))
}

func Call_prepend(dictApplicative_0_loop gopurs_runtime.Value, h_1_loop interface{}, t_2_loop interface{}) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 interface{} = h_1_loop
_ = h_1
var t_2 interface{} = t_2_loop
_ = t_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{h_1, func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(t_2)
}), inner_arg0))
}})})))
}

func Call_prepend__func_gopurs_runtime_Value__interface____interface____interface___1575160070(dictApplicative_0_loop gopurs_runtime.Value, h_1_loop interface{}, t_2_loop interface{}) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var h_1 interface{} = h_1_loop
_ = h_1
var t_2 interface{} = t_2_loop
_ = t_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{h_1, func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(t_2)
}), inner_arg0))
}})})))
}

func Call_nil(dictApplicative_0_loop gopurs_runtime.Value) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})))
}

func Call_nil__func_gopurs_runtime_Value__interface___30547265(dictApplicative_0_loop gopurs_runtime.Value) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})))
}

func Call_singleton(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_2), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}), inner_arg0))
}})}))
})
}

func Call_singleton__func_gopurs_runtime_Value__interface____interface___3816438021(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_2), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}), inner_arg0))
}})}))
})
}

func Call_take(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.IntVal) == (0) {
__t3 = nil1_1_0
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), gopurs_runtime.Int((v_3.IntVal) - (1))), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), v_3), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 489128924) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), v1_4)
}
end_branch_3:
return __t3
})
}

func Call_take__func_gopurs_runtime_Value__int64__interface____interface___3143403118(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_3.IntVal) == (0) {
__t3 = nil1_1_0
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v2_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_5.Type == 9 && v2_5.IntVal == 1320412129) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), gopurs_runtime.Int((v_3.IntVal) - (1))), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_take(dictApplicative_0), v_3), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_5.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
if (v2_5.Type == 9 && v2_5.IntVal == 489128924) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), v1_4)
}
end_branch_3:
return __t3
})
}

func Call_zipWith_prime(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_2_1
Bind1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_3_2
Functor0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_4_3
uncons1_5_4 := Call_uncons(dictMonad_0)
_ = uncons1_5_4
return gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, fa_7 gopurs_runtime.Value, fb_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
}), inner_arg0))
}})})
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fa_7), gopurs_runtime.Func(func(ua_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fb_8), gopurs_runtime.Func(func(ub_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (ub_10.Type == 9 && ub_10.IntVal == 930809136 && ub_10.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if (ua_9.Type == 9 && ua_9.IntVal == 930809136 && ua_9.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if ((ua_9.Type == 9 && ua_9.IntVal == 930809136 && ua_9.UnsafePtr != nil)) && ((ub_10.Type == 9 && ub_10.IntVal == 930809136 && ub_10.UnsafePtr != nil)) {
__local_var_11_6 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_9.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_11_6
__local_var_12_7 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_10.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_12_7
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_13), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_zipWith_prime(), dictMonad_0, f_6, __local_var_11_6, __local_var_12_7)
}), inner_arg0))
}})}))
}), gopurs_runtime.Apply2(f_6, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_9.UnsafePtr).V0).UnsafePtr).V0), gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_10.UnsafePtr).V0).UnsafePtr).V0)))
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
})))
})
}

func Call_zipWith_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___2966163796(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_2_1
Bind1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_3_2
Functor0_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_3_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_4_3
uncons1_5_4 := Call_uncons(dictMonad_0)
_ = uncons1_5_4
return gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, fa_7 gopurs_runtime.Value, fb_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
}), inner_arg0))
}})})
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fa_7), gopurs_runtime.Func(func(ua_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_3_2, "bind"), gopurs_runtime.Apply(uncons1_5_4, fb_8), gopurs_runtime.Func(func(ub_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (ub_10.Type == 9 && ub_10.IntVal == 930809136 && ub_10.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if (ua_9.Type == 9 && ua_9.IntVal == 930809136 && ua_9.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), nil1_2_1)
goto end_branch_5
} else {

}
}
{
if ((ua_9.Type == 9 && ua_9.IntVal == 930809136 && ua_9.UnsafePtr != nil)) && ((ub_10.Type == 9 && ub_10.IntVal == 930809136 && ub_10.UnsafePtr != nil)) {
__local_var_11_6 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_9.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_11_6
__local_var_12_7 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_10.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_12_7
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_4_3, "map"), gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_13), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_zipWith_prime(), dictMonad_0, f_6, __local_var_11_6, __local_var_12_7)
}), inner_arg0))
}})}))
}), gopurs_runtime.Apply2(f_6, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ua_9.UnsafePtr).V0).UnsafePtr).V0), gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(ub_10.UnsafePtr).V0).UnsafePtr).V0)))
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
})))
})
}

func Call_zipWith(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
zipWith_prime1_1_0 := gopurs_runtime.Apply(Get_zipWith_prime(), dictMonad_0)
_ = zipWith_prime1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(zipWith_prime1_1_0, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Apply2(f_2, a_3, b_4))
}))
})
}

func Call_mapMaybe(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}], v_2_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] = f_1_loop
_ = f_1
var v_2 interface{} = v_2_loop
_ = v_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))))})
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136 && __local_var_4_1.UnsafePtr == nil) {
__t2 = Get_Skip()
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136 && __local_var_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Apply(__t2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_1(gopurs_runtime.UnboxAny(arg0)))}
})), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1(arg0))
})))
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_1(gopurs_runtime.UnboxAny(arg0)))}
})), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Any(v_2)))
}

func Call_mapMaybe__func_gopurs_runtime_Value__func_interface____ptrData_Maybe_Constructor_Just[interface__]__interface____interface___1725066890(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}], v_2_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 func(interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] = f_1_loop
_ = f_1
var v_2 interface{} = v_2_loop
_ = v_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__local_var_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), Get_Yield(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))))})
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136 && __local_var_4_1.UnsafePtr == nil) {
__t2 = Get_Skip()
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1.Type == 9 && __local_var_4_1.IntVal == 930809136 && __local_var_4_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t0 = gopurs_runtime.Apply(__t2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_1(gopurs_runtime.UnboxAny(arg0)))}
})), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1(arg0))
})))
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(f_1(gopurs_runtime.UnboxAny(arg0)))}
})), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Any(v_2)))
}

func Call_iterate(dictMonad_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, a_2_loop interface{}) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(Call_unfold(dictMonad_0, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(x_3)))), gopurs_runtime.UnboxAny(x_3)})}})}))
}), gopurs_runtime.Any(inner_arg0)))
}, a_2)))
}

func Call_iterate__func_gopurs_runtime_Value__func_interface____interface____interface____interface___528763728(dictMonad_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) interface{}, a_2_loop interface{}) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var f_1 func(interface{}) interface{} = f_1_loop
_ = f_1
var a_2 interface{} = a_2_loop
_ = a_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(Call_unfold(dictMonad_0, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_1(gopurs_runtime.UnboxAny(x_3)))), gopurs_runtime.UnboxAny(x_3)})}})}))
}), gopurs_runtime.Any(inner_arg0)))
}, a_2)))
}

func Call_repeat(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply2(Get_iterate(), dictMonad_0, Get_identity())
}

func Call_head(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := Call_uncons(dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), pkg_Data_Tuple.Get_fst()), gopurs_runtime.Apply(uncons1_1_0, l_2))
})
}

func Call_functorListT(dictFunctor_0_loop gopurs_runtime.Value) interface{} {
functorListT:
for {
if false { continue functorListT }
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_1, gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0))), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Any(Call_functorListT(dictFunctor_0)), "map"), f_1), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Any(Call_functorListT(dictFunctor_0)), "map"), f_1), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), v_2)
})))
}
}

func Call_fromEffect(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
nil1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_1_0
return gopurs_runtime.Func(func(fa_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_3), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_1_0
}), inner_arg0))
}})})
}), fa_2)
})
}

func Call_foldlRec_prime(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
uncons1_4_3 := Call_uncons(Monad0_1_0)
_ = uncons1_4_3
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(uncons1_4_3, gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(__local_var_9_4)})}))
goto end_branch_5
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
__local_var_11_6 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_11_6
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(f_5, __local_var_9_4, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0).UnsafePtr).V0)), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_6))})}))
}))
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
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
}

func Call_foldlRec_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___1441292185(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
uncons1_4_3 := Call_uncons(Monad0_1_0)
_ = uncons1_4_3
return gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := gopurs_runtime.RecordGet(o_8, "a")
_ = __local_var_9_4
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(uncons1_4_3, gopurs_runtime.RecordGet(o_8, "b")), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(__local_var_9_4)})}))
goto end_branch_5
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 930809136 && v_10.UnsafePtr != nil) {
__local_var_11_6 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_11_6
__t5 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply2(f_5, __local_var_9_4, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_10.UnsafePtr).V0).UnsafePtr).V0)), gopurs_runtime.Func(func(b_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("a", "b", b_prime_12, __local_var_11_6))})}))
}))
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
}), gopurs_runtime.RecordDict2("a", "b", a_6, b_7))
})
}

func Call_runListTRec(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
return gopurs_runtime.Apply3(Get_foldlRec_prime(), dictMonadRec_0, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
}

func Call_foldlRec(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
uncons1_3_2 := Call_uncons(Monad0_1_0)
_ = uncons1_3_2
return gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.RecordGet(o_7, "a")
_ = __local_var_8_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_3_2, gopurs_runtime.RecordGet(o_7, "b")), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_9.Type == 9 && v_9.IntVal == 930809136 && v_9.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(__local_var_8_3)})}))
goto end_branch_4
} else {

}
}
{
if (v_9.Type == 9 && v_9.IntVal == 930809136 && v_9.UnsafePtr != nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("a", "b", gopurs_runtime.Apply2(f_4, __local_var_8_3, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_9.UnsafePtr).V0).UnsafePtr).V0)), gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_9.UnsafePtr).V0).UnsafePtr).V1)))})}))
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
}), gopurs_runtime.RecordDict2("a", "b", a_5, b_6))
})
}

func Call_foldl_prime(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_0
uncons1_2_1 := Call_uncons(dictMonad_0)
_ = uncons1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2 gopurs_runtime.Value
_ = loop_4_2
loop_4_2 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply(uncons1_2_1, l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_5)
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__local_var_8_4 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_8_4
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply2(f_3, b_5, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0).UnsafePtr).V0)), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Any(loop_4_2), a_9, __local_var_8_4)
}))
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
})
return gopurs_runtime.Any(loop_4_2)
})
}

func Call_foldl_prime__func_gopurs_runtime_Value__func_interface____interface____interface____interface____interface____interface___331401616(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_0
uncons1_2_1 := Call_uncons(dictMonad_0)
_ = uncons1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_4_2 gopurs_runtime.Value
_ = loop_4_2
loop_4_2 = gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply(uncons1_2_1, l_6), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_5)
goto end_branch_3
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__local_var_8_4 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_8_4
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply2(f_3, b_5, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0).UnsafePtr).V0)), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Any(loop_4_2), a_9, __local_var_8_4)
}))
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
})
return gopurs_runtime.Any(loop_4_2)
})
}

func Call_runListT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Apply3(Get_foldl_prime(), dictMonad_0, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), pkg_Data_Unit.Get_unit())
}), pkg_Data_Unit.Get_unit())
}

func Call_foldl(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
uncons1_1_0 := Call_uncons(dictMonad_0)
_ = uncons1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
var loop_3_1 gopurs_runtime.Value
_ = loop_3_1
loop_3_1 = gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(uncons1_1_0, l_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 930809136 && v_6.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), b_4)
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 930809136 && v_6.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(gopurs_runtime.Any(loop_3_1), gopurs_runtime.Apply2(f_2, b_4, gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0).UnsafePtr).V0)), gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_6.UnsafePtr).V0).UnsafePtr).V1))
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
return gopurs_runtime.Any(loop_3_1)
})
}

func Call_filter(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) bool, v_2_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 func(interface{}) bool = f_1_loop
_ = f_1
var v_2 interface{} = v_2_loop
_ = v_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
s_prime_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), func() gopurs_runtime.Value {
arr_val_filter5 := gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_1(gopurs_runtime.UnboxAny(arg0)))
})
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if gopurs_runtime.Apply(dictFunctor_0, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1(arg0))
}))
_ = s_prime_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)))).IntVal) != (0) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(s_prime_4_1, inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(s_prime_4_1, inner_arg0))
}})})
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), func() gopurs_runtime.Value {
arr_val_filter5 := gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_1(gopurs_runtime.UnboxAny(arg0)))
})
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if gopurs_runtime.Apply(dictFunctor_0, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Any(v_2)))
}

func Call_filter__func_gopurs_runtime_Value__func_interface____bool__interface____interface___2708958519(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop func(interface{}) bool, v_2_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 func(interface{}) bool = f_1_loop
_ = f_1
var v_2 interface{} = v_2_loop
_ = v_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1320412129) {
s_prime_4_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), func() gopurs_runtime.Value {
arr_val_filter5 := gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_1(gopurs_runtime.UnboxAny(arg0)))
})
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if gopurs_runtime.Apply(dictFunctor_0, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1(arg0))
}))
_ = s_prime_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(f_1(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)))).IntVal) != (0) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(s_prime_4_1, inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(s_prime_4_1, inner_arg0))
}})})
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 813447293) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), func() gopurs_runtime.Value {
arr_val_filter5 := gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(f_1(gopurs_runtime.UnboxAny(arg0)))
})
_ = arr_val_filter5
arr_go_filter5 := (*[]gopurs_runtime.Value)(arr_val_filter5.UnsafePtr)
_ = arr_go_filter5
res_go_filter5 := make([]gopurs_runtime.Value, 0)
_ = res_go_filter5
for _, v_filter5 := range *arr_go_filter5 {
if gopurs_runtime.Apply(dictFunctor_0, v_filter5).BoolVal() {
res_go_filter5 = append(res_go_filter5, v_filter5)
} else {

}
}
return gopurs_runtime.Array(res_go_filter5)
}(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_0
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 489128924) {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Any(v_2)))
}

func Call_dropWhile(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)).IntVal) != (0) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1(arg0))
}), inner_arg0))
}})})
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}

func Call_dropWhile__func_gopurs_runtime_Value__func_interface____bool__interface____interface___1576292412(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(f_2, gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)).IntVal) != (0) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1(arg0))
}), inner_arg0))
}})})
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_dropWhile(dictApplicative_0), f_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v_3)
})
}

func Call_drop(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), gopurs_runtime.Int((v_2.IntVal) - (1))), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), v_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v1_3)
}
end_branch_2:
return __t2
})
}

func Call_drop__func_gopurs_runtime_Value__int64__interface____interface___3143403118(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t2 = v1_3
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), gopurs_runtime.Int((v_2.IntVal) - (1))), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Apply(Call_drop(dictApplicative_0), v_2), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), v1_3)
}
end_branch_2:
return __t2
})
}

func Call_cons(dictApplicative_0_loop gopurs_runtime.Value, lh_1_loop func(gopurs_runtime.Value) interface{}, t_2_loop func(gopurs_runtime.Value) interface{}) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 func(gopurs_runtime.Value) interface{} = lh_1_loop
_ = lh_1
var t_2 func(gopurs_runtime.Value) interface{} = t_2_loop
_ = t_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(lh_1(pkg_Data_Unit.Get_unit()))), t_2})})))
}

func Call_cons__func_gopurs_runtime_Value__func_gopurs_runtime_Value__interface____func_gopurs_runtime_Value__interface____interface___3975735302(dictApplicative_0_loop gopurs_runtime.Value, lh_1_loop func(gopurs_runtime.Value) interface{}, t_2_loop func(gopurs_runtime.Value) interface{}) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var lh_1 func(gopurs_runtime.Value) interface{} = lh_1_loop
_ = lh_1
var t_2 func(gopurs_runtime.Value) interface{} = t_2_loop
_ = t_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any(lh_1(pkg_Data_Unit.Get_unit()))), t_2})})))
}

func Call_unfoldable1ListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
var go__5_2 gopurs_runtime.Value
_ = go__5_2
go__5_2 = gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 930809136 && __t_tag_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
}), inner_arg0))
}})}))
goto end_branch_3
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 930809136 && __t_tag_5.UnsafePtr != nil) {
__local_var_7_6 := gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1).UnsafePtr).V0)
_ = __local_var_7_6
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), pkg_Data_Unit.Get_unit())), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Any(go__5_2), gopurs_runtime.Apply(f_3, __local_var_7_6))
}), inner_arg0))
}})}))
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
return gopurs_runtime.Apply(gopurs_runtime.Any(go__5_2), gopurs_runtime.Apply(f_3, b_4))
})))
}

func Call_unfoldableListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_2_1
unfoldable1ListT1_3_2 := gopurs_runtime.Any(Call_unfoldable1ListT(dictMonad_0))
_ = unfoldable1ListT1_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return unfoldable1ListT1_3_2
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__6_3 gopurs_runtime.Value
_ = go__6_3
go__6_3 = gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr == nil) {
__t4 = nil1_2_1
goto end_branch_4
} else {

}
}
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__local_var_8_5 := gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0).UnsafePtr).V1)
_ = __local_var_8_5
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_applicativeLazy(), "pure"), gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Any((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_7.UnsafePtr).V0).UnsafePtr).V0), pkg_Data_Unit.Get_unit())), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Any(go__6_3), gopurs_runtime.Apply(f_4, __local_var_8_5))
}), inner_arg0))
}})}))
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
return gopurs_runtime.Apply(gopurs_runtime.Any(go__6_3), gopurs_runtime.Apply(f_4, b_5))
})))
}

func Call_semigroupListT(dictApplicative_0_loop gopurs_runtime.Value) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("append", Call_concat(dictApplicative_0)))
}

func Call_concat(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 1320412129) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_concat(dictApplicative_0), v1_5, y_3)
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 813447293) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_concat(dictApplicative_0), v1_5, y_3)
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 489128924) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return y_3
}), inner_arg0))
}})})
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), x_2)
})
}

func Call_monoidListT(dictApplicative_0_loop gopurs_runtime.Value) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
semigroupListT1_1_0 := gopurs_runtime.RecordDict1("append", Call_concat(dictApplicative_0))
_ = semigroupListT1_1_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupListT1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))))
}

func Call_catMaybes(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply2(Get_mapMaybe(), dictFunctor_0, Get_identity())
}

func Call_monadListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applicativeListT(dictMonad_0))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindListT(dictMonad_0))
})))
}

func Call_bindListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
append_1_0 := Call_concat(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = append_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applyListT(dictMonad_0))
}), gopurs_runtime.Func2(func(fa_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1320412129) {
__local_var_6_3 := gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
_ = __local_var_6_3
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(append_1_0, gopurs_runtime.Apply(f_4, __local_var_6_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Any(Call_bindListT(dictMonad_0)), "bind"), s_prime_7, f_4))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 813447293) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 813447293, UnsafePtr: unsafe.Pointer(&Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Lazy.Get_functorLazy(), "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Any(Call_bindListT(dictMonad_0)), "bind"), v1_6, f_4)
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_Skip[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0(arg0))
})), inner_arg0))
}})})
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 489128924) {
__t2 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), fa_3)
})))
}

func Call_applyListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
functorListT1_1_0 := gopurs_runtime.Any(Call_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
_ = functorListT1_1_0
__local_var_2_1 := gopurs_runtime.Any(Call_bindListT(dictMonad_0))
_ = __local_var_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), f_3, gopurs_runtime.Func(func(f_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "bind"), a_4, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Any(Call_applicativeListT(dictMonad_0)), "pure"), gopurs_runtime.Apply(f_prime_5, a_prime_6))
}))
}))
})))
}

func Call_applicativeListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_0
nil1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))
_ = nil1_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applyListT(dictMonad_0))
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1320412129, UnsafePtr: unsafe.Pointer(&Constructor_Yield[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(a_3), func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return nil1_2_1
}), inner_arg0))
}})}))
})))
}

func Call_monadEffectListT(dictMonadEffect_0_loop gopurs_runtime.Value) interface{} {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applicativeListT(Monad0_1_0))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindListT(Monad0_1_0))
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
})))
}

func Call_monadSTListT(dictMonadST_0_loop gopurs_runtime.Value) interface{} {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadListT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applicativeListT(Monad0_1_0))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindListT(Monad0_1_0))
}))
_ = monadListT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransListT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
})))
}

func Call_altListT(dictApplicative_0_loop gopurs_runtime.Value) interface{} {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
functorListT1_1_0 := gopurs_runtime.Any(Call_functorListT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})))
_ = functorListT1_1_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorListT1_1_0
}), Call_concat(dictApplicative_0)))
}

func Call_plusListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Applicative0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = Applicative0_1_0
altListT1_2_1 := gopurs_runtime.Any(Call_altListT(Applicative0_1_0))
_ = altListT1_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altListT1_2_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Applicative0_1_0, "pure"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 489128924, UnsafePtr: nil}))))
}

func Call_alternativeListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeListT1_1_0 := gopurs_runtime.Any(Call_applicativeListT(dictMonad_0))
_ = applicativeListT1_1_0
plusListT1_2_1 := gopurs_runtime.Any(Call_plusListT(dictMonad_0))
_ = plusListT1_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeListT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusListT1_2_1
})))
}

func Call_monadPlusListT(dictMonad_0_loop gopurs_runtime.Value) interface{} {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadListT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_applicativeListT(dictMonad_0))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindListT(dictMonad_0))
}))
_ = monadListT1_1_0
alternativeListT1_2_1 := gopurs_runtime.Any(Call_alternativeListT(dictMonad_0))
_ = alternativeListT1_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeListT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadListT1_1_0
})))
}
