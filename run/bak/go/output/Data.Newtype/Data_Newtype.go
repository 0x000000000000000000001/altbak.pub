package Data_Newtype

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var cache_wrap gopurs_runtime.Value
var once_wrap sync.Once
func Get_wrap() gopurs_runtime.Value {
	once_wrap.Do(func() {
		cache_wrap = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrap(_dollar__unused_0_box)
})
	})
	return cache_wrap
}

var cache_wrap__gopurs_runtime_Value_3267718003 gopurs_runtime.Value
var once_wrap__gopurs_runtime_Value_3267718003 sync.Once
func Get_wrap__gopurs_runtime_Value_3267718003() gopurs_runtime.Value {
	once_wrap__gopurs_runtime_Value_3267718003.Do(func() {
		cache_wrap__gopurs_runtime_Value_3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_wrap__gopurs_runtime_Value_3267718003(_dollar__unused_0_box)
})
	})
	return cache_wrap__gopurs_runtime_Value_3267718003
}

var cache_unwrap gopurs_runtime.Value
var once_unwrap sync.Once
func Get_unwrap() gopurs_runtime.Value {
	once_unwrap.Do(func() {
		cache_unwrap = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap(_dollar__unused_0_box)
})
	})
	return cache_unwrap
}

var cache_unwrap__gopurs_runtime_Value_3267718003 gopurs_runtime.Value
var once_unwrap__gopurs_runtime_Value_3267718003 sync.Once
func Get_unwrap__gopurs_runtime_Value_3267718003() gopurs_runtime.Value {
	once_unwrap__gopurs_runtime_Value_3267718003.Do(func() {
		cache_unwrap__gopurs_runtime_Value_3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__gopurs_runtime_Value_3267718003(_dollar__unused_0_box)
})
	})
	return cache_unwrap__gopurs_runtime_Value_3267718003
}

var cache_underF2 gopurs_runtime.Value
var once_underF2 sync.Once
func Get_underF2() gopurs_runtime.Value {
	once_underF2.Do(func() {
		cache_underF2 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_underF2(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return cache_underF2
}

var cache_underF gopurs_runtime.Value
var once_underF sync.Once
func Get_underF() gopurs_runtime.Value {
	once_underF.Do(func() {
		cache_underF = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_underF(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return cache_underF
}

var cache_under2 gopurs_runtime.Value
var once_under2 sync.Once
func Get_under2() gopurs_runtime.Value {
	once_under2.Do(func() {
		cache_under2 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_under2(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_under2
}

var cache_under gopurs_runtime.Value
var once_under sync.Once
func Get_under() gopurs_runtime.Value {
	once_under.Do(func() {
		cache_under = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_under(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_under
}

var cache_un gopurs_runtime.Value
var once_un sync.Once
func Get_un() gopurs_runtime.Value {
	once_un.Do(func() {
		cache_un = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_un(_dollar__unused_0_box, v_1_box)
})
	})
	return cache_un
}

var cache_un__gopurs_runtime_Value_3470773042 gopurs_runtime.Value
var once_un__gopurs_runtime_Value_3470773042 sync.Once
func Get_un__gopurs_runtime_Value_3470773042() gopurs_runtime.Value {
	once_un__gopurs_runtime_Value_3470773042.Do(func() {
		cache_un__gopurs_runtime_Value_3470773042 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_un__gopurs_runtime_Value_3470773042(_dollar__unused_0_box, v_1_box)
})
	})
	return cache_un__gopurs_runtime_Value_3470773042
}

var cache_traverse gopurs_runtime.Value
var once_traverse sync.Once
func Get_traverse() gopurs_runtime.Value {
	once_traverse.Do(func() {
		cache_traverse = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_traverse
}

var cache_overF2 gopurs_runtime.Value
var once_overF2 sync.Once
func Get_overF2() gopurs_runtime.Value {
	once_overF2.Do(func() {
		cache_overF2 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_overF2(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return cache_overF2
}

var cache_overF gopurs_runtime.Value
var once_overF sync.Once
func Get_overF() gopurs_runtime.Value {
	once_overF.Do(func() {
		cache_overF = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_overF(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return cache_overF
}

var cache_over2 gopurs_runtime.Value
var once_over2 sync.Once
func Get_over2() gopurs_runtime.Value {
	once_over2.Do(func() {
		cache_over2 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over2(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_over2
}

var cache_over gopurs_runtime.Value
var once_over sync.Once
func Get_over() gopurs_runtime.Value {
	once_over.Do(func() {
		cache_over = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_over
}

var cache_over__gopurs_runtime_Value_1462660013 gopurs_runtime.Value
var once_over__gopurs_runtime_Value_1462660013 sync.Once
func Get_over__gopurs_runtime_Value_1462660013() gopurs_runtime.Value {
	once_over__gopurs_runtime_Value_1462660013.Do(func() {
		cache_over__gopurs_runtime_Value_1462660013 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over__gopurs_runtime_Value_1462660013(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_over__gopurs_runtime_Value_1462660013
}

var cache_newtypeMultiplicative gopurs_runtime.Value
var once_newtypeMultiplicative sync.Once
func Get_newtypeMultiplicative() gopurs_runtime.Value {
	once_newtypeMultiplicative.Do(func() {
		cache_newtypeMultiplicative = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeMultiplicative
}

var cache_newtypeLast gopurs_runtime.Value
var once_newtypeLast sync.Once
func Get_newtypeLast() gopurs_runtime.Value {
	once_newtypeLast.Do(func() {
		cache_newtypeLast = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeLast
}

var cache_newtypeFirst gopurs_runtime.Value
var once_newtypeFirst sync.Once
func Get_newtypeFirst() gopurs_runtime.Value {
	once_newtypeFirst.Do(func() {
		cache_newtypeFirst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeFirst
}

var cache_newtypeEndo gopurs_runtime.Value
var once_newtypeEndo sync.Once
func Get_newtypeEndo() gopurs_runtime.Value {
	once_newtypeEndo.Do(func() {
		cache_newtypeEndo = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeEndo
}

var cache_newtypeDual gopurs_runtime.Value
var once_newtypeDual sync.Once
func Get_newtypeDual() gopurs_runtime.Value {
	once_newtypeDual.Do(func() {
		cache_newtypeDual = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeDual
}

var cache_newtypeDisj gopurs_runtime.Value
var once_newtypeDisj sync.Once
func Get_newtypeDisj() gopurs_runtime.Value {
	once_newtypeDisj.Do(func() {
		cache_newtypeDisj = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeDisj
}

var cache_newtypeConj gopurs_runtime.Value
var once_newtypeConj sync.Once
func Get_newtypeConj() gopurs_runtime.Value {
	once_newtypeConj.Do(func() {
		cache_newtypeConj = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeConj
}

var cache_newtypeAdditive gopurs_runtime.Value
var once_newtypeAdditive sync.Once
func Get_newtypeAdditive() gopurs_runtime.Value {
	once_newtypeAdditive.Do(func() {
		cache_newtypeAdditive = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeAdditive
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, fn_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(_dollar__unused_0_box, fn_1_box, t_2_box)
})
	})
	return cache_modify
}

var cache_collect gopurs_runtime.Value
var once_collect sync.Once
func Get_collect() gopurs_runtime.Value {
	once_collect.Do(func() {
		cache_collect = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collect(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return cache_collect
}

var cache_alaF gopurs_runtime.Value
var once_alaF sync.Once
func Get_alaF() gopurs_runtime.Value {
	once_alaF.Do(func() {
		cache_alaF = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return cache_alaF
}

var cache_alaF__gopurs_runtime_Value_4085337484 gopurs_runtime.Value
var once_alaF__gopurs_runtime_Value_4085337484 sync.Once
func Get_alaF__gopurs_runtime_Value_4085337484() gopurs_runtime.Value {
	once_alaF__gopurs_runtime_Value_4085337484.Do(func() {
		cache_alaF__gopurs_runtime_Value_4085337484 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__gopurs_runtime_Value_4085337484(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return cache_alaF__gopurs_runtime_Value_4085337484
}

var cache_ala gopurs_runtime.Value
var once_ala sync.Once
func Get_ala() gopurs_runtime.Value {
	once_ala.Do(func() {
		cache_ala = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ala(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, v_3_box, f_4_box)
})
	})
	return cache_ala
}

var cache_ala__gopurs_runtime_Value_2430830187 gopurs_runtime.Value
var once_ala__gopurs_runtime_Value_2430830187 sync.Once
func Get_ala__gopurs_runtime_Value_2430830187() gopurs_runtime.Value {
	once_ala__gopurs_runtime_Value_2430830187.Do(func() {
		cache_ala__gopurs_runtime_Value_2430830187 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ala__gopurs_runtime_Value_2430830187(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, v_3_box, f_4_box)
})
	})
	return cache_ala__gopurs_runtime_Value_2430830187
}

func Call_wrap(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_wrap__gopurs_runtime_Value_3267718003(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__gopurs_runtime_Value_3267718003(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_underF2(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, _dollar__unused_3_loop gopurs_runtime.Value, v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 gopurs_runtime.Value = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_underF(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, _dollar__unused_3_loop gopurs_runtime.Value, v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 gopurs_runtime.Value = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_under2(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_under(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_un(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_un__gopurs_runtime_Value_3470773042(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_traverse(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_overF2(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, _dollar__unused_3_loop gopurs_runtime.Value, v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 gopurs_runtime.Value = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_overF(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, _dollar__unused_3_loop gopurs_runtime.Value, v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 gopurs_runtime.Value = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over2(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_over__gopurs_runtime_Value_1462660013(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_modify(_dollar__unused_0_loop gopurs_runtime.Value, fn_1_loop gopurs_runtime.Value, t_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var fn_1 gopurs_runtime.Value = fn_1_loop
_ = fn_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
return gopurs_runtime.Apply(fn_1, t_2)
}

func Call_collect(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_alaF(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, _dollar__unused_3_loop gopurs_runtime.Value, v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 gopurs_runtime.Value = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_alaF__gopurs_runtime_Value_4085337484(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, _dollar__unused_3_loop gopurs_runtime.Value, v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 gopurs_runtime.Value = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_ala(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
return gopurs_runtime.Apply(f_4, pkg_Unsafe_Coerce.Get_unsafeCoerce())
}

func Call_ala__gopurs_runtime_Value_2430830187(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop gopurs_runtime.Value, v_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 gopurs_runtime.Value = _dollar__unused_2_loop
_ = _dollar__unused_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
return gopurs_runtime.Apply(f_4, pkg_Unsafe_Coerce.Get_unsafeCoerce())
}


