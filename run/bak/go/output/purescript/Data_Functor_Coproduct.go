package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Coproduct_Coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_Coproduct sync.Once
func Get_Data_Functor_Coproduct_Coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_Coproduct.Do(func() {
		cache_Data_Functor_Coproduct_Coproduct = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_Coproduct(x_0_box)
})
	})
	return cache_Data_Functor_Coproduct_Coproduct
}

var cache_Data_Functor_Coproduct_showCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_showCoproduct sync.Once
func Get_Data_Functor_Coproduct_showCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_showCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_showCoproduct = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_showCoproduct(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_Data_Functor_Coproduct_showCoproduct
}

var cache_Data_Functor_Coproduct_right gopurs_runtime.Value
var once_Data_Functor_Coproduct_right sync.Once
func Get_Data_Functor_Coproduct_right() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right.Do(func() {
		cache_Data_Functor_Coproduct_right = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right
}

var cache_Data_Functor_Coproduct_newtypeCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_newtypeCoproduct sync.Once
func Get_Data_Functor_Coproduct_newtypeCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_newtypeCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_newtypeCoproduct = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Functor_Coproduct_newtypeCoproduct
}

var cache_Data_Functor_Coproduct_left gopurs_runtime.Value
var once_Data_Functor_Coproduct_left sync.Once
func Get_Data_Functor_Coproduct_left() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_left.Do(func() {
		cache_Data_Functor_Coproduct_left = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_left(fa_0_box)
})
	})
	return cache_Data_Functor_Coproduct_left
}

var cache_Data_Functor_Coproduct_functorCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_functorCoproduct sync.Once
func Get_Data_Functor_Coproduct_functorCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_functorCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_functorCoproduct = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, dictFunctor1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_functorCoproduct(dictFunctor_0_box, dictFunctor1_1_box)
})
	})
	return cache_Data_Functor_Coproduct_functorCoproduct
}

var cache_Data_Functor_Coproduct_eq1Coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_eq1Coproduct sync.Once
func Get_Data_Functor_Coproduct_eq1Coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_eq1Coproduct.Do(func() {
		cache_Data_Functor_Coproduct_eq1Coproduct = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_eq1Coproduct(dictEq1_0_box, dictEq11_1_box)
})
	})
	return cache_Data_Functor_Coproduct_eq1Coproduct
}

var cache_Data_Functor_Coproduct_eqCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_eqCoproduct sync.Once
func Get_Data_Functor_Coproduct_eqCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_eqCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_eqCoproduct = gopurs_runtime.Func3(func(dictEq1_0_box gopurs_runtime.Value, dictEq11_1_box gopurs_runtime.Value, dictEq_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_eqCoproduct(dictEq1_0_box, dictEq11_1_box, dictEq_2_box)
})
	})
	return cache_Data_Functor_Coproduct_eqCoproduct
}

var cache_Data_Functor_Coproduct_ord1Coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_ord1Coproduct sync.Once
func Get_Data_Functor_Coproduct_ord1Coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_ord1Coproduct.Do(func() {
		cache_Data_Functor_Coproduct_ord1Coproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_ord1Coproduct(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Coproduct_ord1Coproduct
}

var cache_Data_Functor_Coproduct_ordCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_ordCoproduct sync.Once
func Get_Data_Functor_Coproduct_ordCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_ordCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_ordCoproduct = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_ordCoproduct(dictOrd1_0_box)
})
	})
	return cache_Data_Functor_Coproduct_ordCoproduct
}

var cache_Data_Functor_Coproduct_coproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct sync.Once
func Get_Data_Functor_Coproduct_coproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct.Do(func() {
		cache_Data_Functor_Coproduct_coproduct = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_coproduct(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Functor_Coproduct_coproduct
}

var cache_Data_Functor_Coproduct_extendCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_extendCoproduct sync.Once
func Get_Data_Functor_Coproduct_extendCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_extendCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_extendCoproduct = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_extendCoproduct(dictExtend_0_box)
})
	})
	return cache_Data_Functor_Coproduct_extendCoproduct
}

var cache_Data_Functor_Coproduct_comonadCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_comonadCoproduct sync.Once
func Get_Data_Functor_Coproduct_comonadCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_comonadCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_comonadCoproduct = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_comonadCoproduct(dictComonad_0_box)
})
	})
	return cache_Data_Functor_Coproduct_comonadCoproduct
}

var cache_Data_Functor_Coproduct_bihoistCoproduct gopurs_runtime.Value
var once_Data_Functor_Coproduct_bihoistCoproduct sync.Once
func Get_Data_Functor_Coproduct_bihoistCoproduct() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_bihoistCoproduct.Do(func() {
		cache_Data_Functor_Coproduct_bihoistCoproduct = gopurs_runtime.Func3(func(natF_0_box gopurs_runtime.Value, natG_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_bihoistCoproduct(natF_0_box, natG_1_box, v_2_box)
})
	})
	return cache_Data_Functor_Coproduct_bihoistCoproduct
}

var cache_Data_Functor_Coproduct_coproduct__79520197 gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct__79520197 sync.Once
func Get_Data_Functor_Coproduct_coproduct__79520197() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct__79520197.Do(func() {
		cache_Data_Functor_Coproduct_coproduct__79520197 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_coproduct__79520197(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Functor_Coproduct_coproduct__79520197
}

var cache_Data_Functor_Coproduct_coproduct__413515331 gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct__413515331 sync.Once
func Get_Data_Functor_Coproduct_coproduct__413515331() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct__413515331.Do(func() {
		cache_Data_Functor_Coproduct_coproduct__413515331 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_coproduct__413515331(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Functor_Coproduct_coproduct__413515331
}

var cache_Data_Functor_Coproduct_coproduct__829064685 gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct__829064685 sync.Once
func Get_Data_Functor_Coproduct_coproduct__829064685() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct__829064685.Do(func() {
		cache_Data_Functor_Coproduct_coproduct__829064685 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_coproduct__829064685(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Functor_Coproduct_coproduct__829064685
}

var cache_Data_Functor_Coproduct_coproduct__2507487339 gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct__2507487339 sync.Once
func Get_Data_Functor_Coproduct_coproduct__2507487339() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct__2507487339.Do(func() {
		cache_Data_Functor_Coproduct_coproduct__2507487339 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Functor_Coproduct_coproduct__2507487339(v_0_box, v1_1_box, v2_2_box))}
})
	})
	return cache_Data_Functor_Coproduct_coproduct__2507487339
}

var cache_Data_Functor_Coproduct_coproduct__1706612365 gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct__1706612365 sync.Once
func Get_Data_Functor_Coproduct_coproduct__1706612365() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct__1706612365.Do(func() {
		cache_Data_Functor_Coproduct_coproduct__1706612365 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_coproduct__1706612365(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Functor_Coproduct_coproduct__1706612365
}

var cache_Data_Functor_Coproduct_coproduct__1642299426 gopurs_runtime.Value
var once_Data_Functor_Coproduct_coproduct__1642299426 sync.Once
func Get_Data_Functor_Coproduct_coproduct__1642299426() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_coproduct__1642299426.Do(func() {
		cache_Data_Functor_Coproduct_coproduct__1642299426 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_coproduct__1642299426(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Functor_Coproduct_coproduct__1642299426
}

var cache_Data_Functor_Coproduct_left__786465523 gopurs_runtime.Value
var once_Data_Functor_Coproduct_left__786465523 sync.Once
func Get_Data_Functor_Coproduct_left__786465523() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_left__786465523.Do(func() {
		cache_Data_Functor_Coproduct_left__786465523 = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_left__786465523(fa_0_box)
})
	})
	return cache_Data_Functor_Coproduct_left__786465523
}

var cache_Data_Functor_Coproduct_left__1785278009 gopurs_runtime.Value
var once_Data_Functor_Coproduct_left__1785278009 sync.Once
func Get_Data_Functor_Coproduct_left__1785278009() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_left__1785278009.Do(func() {
		cache_Data_Functor_Coproduct_left__1785278009 = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_left__1785278009(fa_0_box)
})
	})
	return cache_Data_Functor_Coproduct_left__1785278009
}

var cache_Data_Functor_Coproduct_left__4018477375 gopurs_runtime.Value
var once_Data_Functor_Coproduct_left__4018477375 sync.Once
func Get_Data_Functor_Coproduct_left__4018477375() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_left__4018477375.Do(func() {
		cache_Data_Functor_Coproduct_left__4018477375 = gopurs_runtime.Func(func(fa_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_left__4018477375(fa_0_box)
})
	})
	return cache_Data_Functor_Coproduct_left__4018477375
}

var cache_Data_Functor_Coproduct_right__1193983057 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__1193983057 sync.Once
func Get_Data_Functor_Coproduct_right__1193983057() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__1193983057.Do(func() {
		cache_Data_Functor_Coproduct_right__1193983057 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__1193983057(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__1193983057
}

var cache_Data_Functor_Coproduct_right__1623275833 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__1623275833 sync.Once
func Get_Data_Functor_Coproduct_right__1623275833() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__1623275833.Do(func() {
		cache_Data_Functor_Coproduct_right__1623275833 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__1623275833(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__1623275833
}

var cache_Data_Functor_Coproduct_right__856638295 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__856638295 sync.Once
func Get_Data_Functor_Coproduct_right__856638295() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__856638295.Do(func() {
		cache_Data_Functor_Coproduct_right__856638295 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__856638295(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__856638295
}

var cache_Data_Functor_Coproduct_right__1175236921 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__1175236921 sync.Once
func Get_Data_Functor_Coproduct_right__1175236921() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__1175236921.Do(func() {
		cache_Data_Functor_Coproduct_right__1175236921 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__1175236921(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__1175236921
}

var cache_Data_Functor_Coproduct_right__3235459543 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__3235459543 sync.Once
func Get_Data_Functor_Coproduct_right__3235459543() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__3235459543.Do(func() {
		cache_Data_Functor_Coproduct_right__3235459543 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__3235459543(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__3235459543
}

var cache_Data_Functor_Coproduct_right__2015984185 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__2015984185 sync.Once
func Get_Data_Functor_Coproduct_right__2015984185() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__2015984185.Do(func() {
		cache_Data_Functor_Coproduct_right__2015984185 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__2015984185(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__2015984185
}

var cache_Data_Functor_Coproduct_right__176310615 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__176310615 sync.Once
func Get_Data_Functor_Coproduct_right__176310615() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__176310615.Do(func() {
		cache_Data_Functor_Coproduct_right__176310615 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__176310615(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__176310615
}

var cache_Data_Functor_Coproduct_right__1052251193 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__1052251193 sync.Once
func Get_Data_Functor_Coproduct_right__1052251193() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__1052251193.Do(func() {
		cache_Data_Functor_Coproduct_right__1052251193 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__1052251193(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__1052251193
}

var cache_Data_Functor_Coproduct_right__2641088471 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__2641088471 sync.Once
func Get_Data_Functor_Coproduct_right__2641088471() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__2641088471.Do(func() {
		cache_Data_Functor_Coproduct_right__2641088471 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__2641088471(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__2641088471
}

var cache_Data_Functor_Coproduct_right__3220060985 gopurs_runtime.Value
var once_Data_Functor_Coproduct_right__3220060985 sync.Once
func Get_Data_Functor_Coproduct_right__3220060985() gopurs_runtime.Value {
	once_Data_Functor_Coproduct_right__3220060985.Do(func() {
		cache_Data_Functor_Coproduct_right__3220060985 = gopurs_runtime.Func(func(ga_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Coproduct_right__3220060985(ga_0_box)
})
	})
	return cache_Data_Functor_Coproduct_right__3220060985
}

func Call_Data_Functor_Coproduct_Coproduct(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Coproduct_showCoproduct(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Str((("(left ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")"))
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Str((("(right ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal())) + (")"))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
}

func Call_Data_Functor_Coproduct_right(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_left(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_Data_Functor_Coproduct_functorCoproduct(dictFunctor_0_loop gopurs_runtime.Value, dictFunctor1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var dictFunctor1_1 gopurs_runtime.Value = dictFunctor1_1_loop
_ = dictFunctor1_1
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Bifunctor_bifunctorEither(), "bimap"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_1, "map"), f_2), v_3)
})
}))
}

func Call_Data_Functor_Coproduct_eq1Coproduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (v_3.Type == 9 && v_3.IntVal == 3711209382) {
var __t0 bool
{
if (v1_4.Type == 9 && v1_4.IntVal == 3711209382) {
__t0 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 2465973597)) && ((v1_4.Type == 9 && v1_4.IntVal == 2465973597)) {
__t1 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
})
}))
}

func Call_Data_Functor_Coproduct_eqCoproduct(dictEq1_0_loop gopurs_runtime.Value, dictEq11_1_loop gopurs_runtime.Value, dictEq_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq11_1 gopurs_runtime.Value = dictEq11_1_loop
_ = dictEq11_1
var dictEq_2 gopurs_runtime.Value = dictEq_2_loop
_ = dictEq_2
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 bool
{
if (v_3.Type == 9 && v_3.IntVal == 3711209382) {
var __t0 bool
{
if (v1_4.Type == 9 && v1_4.IntVal == 3711209382) {
__t0 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((v_3.Type == 9 && v_3.IntVal == 2465973597)) && ((v1_4.Type == 9 && v1_4.IntVal == 2465973597)) {
__t1 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq11_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0).IntVal) != (0)
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return gopurs_runtime.Bool(__t1)
})
}))
}

func Call_Data_Functor_Coproduct_ord1Coproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictOrd11_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_2, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): eq1Coproduct2_3_1 -> gopurs_runtime.Value
eq1Coproduct2_3_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
var __t3 bool
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t3 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_4))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && ((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) {
__t4 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_4))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal) != (0)
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
return gopurs_runtime.Bool(__t4)
})
})
}))
_ = eq1Coproduct2_3_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1Coproduct2_3_1
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
var __t5 uint32
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t5 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 3711209382) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if ((v_5.Type == 9 && v_5.IntVal == 2465973597)) && ((v1_6.Type == 9 && v1_6.IntVal == 2465973597)) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd11_2, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_4))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_6.UnsafePtr).V0).IntVal)), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t6.IntVal)), UnsafePtr: nil}
})
})
}))
})
}

func Call_Data_Functor_Coproduct_ordCoproduct(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
// TAST (Let): ord1Coproduct1_1_0 -> gopurs_runtime.Value
ord1Coproduct1_1_0 := Call_Data_Functor_Coproduct_ord1Coproduct(dictOrd1_0)
_ = ord1Coproduct1_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictOrd11_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd11_3, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_4_2
return gopurs_runtime.Func(func(dictOrd_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_5, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_6_4
// TAST (Let): eqCoproduct3_6_3 -> gopurs_runtime.Value
eqCoproduct3_6_3 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 bool
{
if (v_7.Type == 9 && v_7.IntVal == 3711209382) {
var __t5 bool
{
if (v1_8.Type == 9 && v1_8.IntVal == 3711209382) {
__t5 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](__local_var_6_4))}, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0).IntVal) != (0)
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if ((v_7.Type == 9 && v_7.IntVal == 2465973597)) && ((v1_8.Type == 9 && v1_8.IntVal == 2465973597)) {
__t6 = (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_2, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](__local_var_6_4))}, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_8.UnsafePtr).V0).IntVal) != (0)
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
return gopurs_runtime.Bool(__t6)
})
}))
_ = eqCoproduct3_6_3
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return eqCoproduct3_6_3
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(ord1Coproduct1_1_0, dictOrd11_3), "compare1"), dictOrd_5))
})
})
}

func Call_Data_Functor_Coproduct_coproduct(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Functor_Coproduct_extendCoproduct(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
// TAST (Let): functorCoproduct1_1_0 -> gopurs_runtime.Value
functorCoproduct1_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct1_1_0
return gopurs_runtime.Func(func(dictExtend1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCoproduct2_3_1 -> gopurs_runtime.Value
functorCoproduct2_3_1 := gopurs_runtime.Apply(functorCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct2_3_1
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_5})})
}))
_ = __local_var_5_4
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_5_4, x_6)})}
})
_ = __local_var_5_3
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend1_2, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_6})})
}))
_ = __local_var_6_6
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_6, x_7)})}
})
_ = __local_var_6_5
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(__local_var_5_3, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(__local_var_6_5, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, x_6)
})
}))
})
}

func Call_Data_Functor_Coproduct_comonadCoproduct(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): extendCoproduct1_1_0 -> gopurs_runtime.Value
extendCoproduct1_1_0 := Call_Data_Functor_Coproduct_extendCoproduct(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct1_1_0
return gopurs_runtime.Func(func(dictComonad1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): extendCoproduct2_3_1 -> gopurs_runtime.Value
extendCoproduct2_3_1 := gopurs_runtime.Apply(extendCoproduct1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "Extend0"), gopurs_runtime.Value{}))
_ = extendCoproduct2_3_1
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return extendCoproduct2_3_1
}), gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad1_2, "extract"), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0)
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
}

func Call_Data_Functor_Coproduct_bihoistCoproduct(natF_0_loop gopurs_runtime.Value, natG_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var natF_0 gopurs_runtime.Value = natF_0_loop
_ = natF_0
var natG_1 gopurs_runtime.Value = natG_1_loop
_ = natG_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Bifunctor_bifunctorEither(), "bimap"), natF_0, natG_1, v_2)
}

func Call_Data_Functor_Coproduct_coproduct__79520197(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Functor_Coproduct_coproduct__413515331(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Functor_Coproduct_coproduct__829064685(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Functor_Coproduct_coproduct__2507487339(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0)
}

func Call_Data_Functor_Coproduct_coproduct__1706612365(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Functor_Coproduct_coproduct__1642299426(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_Data_Functor_Coproduct_left__786465523(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_Data_Functor_Coproduct_left__1785278009(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_Data_Functor_Coproduct_left__4018477375(fa_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0})}
}

func Call_Data_Functor_Coproduct_right__1193983057(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__1623275833(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__856638295(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__1175236921(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__3235459543(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__2015984185(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__176310615(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__1052251193(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__2641088471(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}

func Call_Data_Functor_Coproduct_right__3220060985(ga_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var ga_0 gopurs_runtime.Value = ga_0_loop
_ = ga_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, ga_0})}
}


