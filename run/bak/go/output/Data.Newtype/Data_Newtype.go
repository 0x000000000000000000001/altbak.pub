package Data_Newtype

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var wrap gopurs_runtime.Value
var once_wrap sync.Once
func Get_wrap() gopurs_runtime.Value {
	once_wrap.Do(func() {
		wrap = gopurs_runtime.Func(func(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}()
})
	})
	return wrap
}

var unwrap gopurs_runtime.Value
var once_unwrap sync.Once
func Get_unwrap() gopurs_runtime.Value {
	once_unwrap.Do(func() {
		unwrap = gopurs_runtime.Func(func(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}()
})
	})
	return unwrap
}

var underF2 gopurs_runtime.Value
var once_underF2 sync.Once
func Get_underF2() gopurs_runtime.Value {
	once_underF2.Do(func() {
		underF2 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_underF2(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return underF2
}

var underF gopurs_runtime.Value
var once_underF sync.Once
func Get_underF() gopurs_runtime.Value {
	once_underF.Do(func() {
		underF = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_underF(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return underF
}

var under2 gopurs_runtime.Value
var once_under2 sync.Once
func Get_under2() gopurs_runtime.Value {
	once_under2.Do(func() {
		under2 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_under2(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return under2
}

var under gopurs_runtime.Value
var once_under sync.Once
func Get_under() gopurs_runtime.Value {
	once_under.Do(func() {
		under = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_under(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return under
}

var un gopurs_runtime.Value
var once_un sync.Once
func Get_un() gopurs_runtime.Value {
	once_un.Do(func() {
		un = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_un(_dollar__unused_0_box, v_1_box)
})
	})
	return un
}

var traverse gopurs_runtime.Value
var once_traverse sync.Once
func Get_traverse() gopurs_runtime.Value {
	once_traverse.Do(func() {
		traverse = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return traverse
}

var overF2 gopurs_runtime.Value
var once_overF2 sync.Once
func Get_overF2() gopurs_runtime.Value {
	once_overF2.Do(func() {
		overF2 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_overF2(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return overF2
}

var overF gopurs_runtime.Value
var once_overF sync.Once
func Get_overF() gopurs_runtime.Value {
	once_overF.Do(func() {
		overF = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_overF(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return overF
}

var over2 gopurs_runtime.Value
var once_over2 sync.Once
func Get_over2() gopurs_runtime.Value {
	once_over2.Do(func() {
		over2 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over2(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return over2
}

var over gopurs_runtime.Value
var once_over sync.Once
func Get_over() gopurs_runtime.Value {
	once_over.Do(func() {
		over = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_over(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return over
}

var newtypeMultiplicative gopurs_runtime.Value
var once_newtypeMultiplicative sync.Once
func Get_newtypeMultiplicative() gopurs_runtime.Value {
	once_newtypeMultiplicative.Do(func() {
		newtypeMultiplicative = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeMultiplicative
}

var newtypeLast gopurs_runtime.Value
var once_newtypeLast sync.Once
func Get_newtypeLast() gopurs_runtime.Value {
	once_newtypeLast.Do(func() {
		newtypeLast = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeLast
}

var newtypeFirst gopurs_runtime.Value
var once_newtypeFirst sync.Once
func Get_newtypeFirst() gopurs_runtime.Value {
	once_newtypeFirst.Do(func() {
		newtypeFirst = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeFirst
}

var newtypeEndo gopurs_runtime.Value
var once_newtypeEndo sync.Once
func Get_newtypeEndo() gopurs_runtime.Value {
	once_newtypeEndo.Do(func() {
		newtypeEndo = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeEndo
}

var newtypeDual gopurs_runtime.Value
var once_newtypeDual sync.Once
func Get_newtypeDual() gopurs_runtime.Value {
	once_newtypeDual.Do(func() {
		newtypeDual = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeDual
}

var newtypeDisj gopurs_runtime.Value
var once_newtypeDisj sync.Once
func Get_newtypeDisj() gopurs_runtime.Value {
	once_newtypeDisj.Do(func() {
		newtypeDisj = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeDisj
}

var newtypeConj gopurs_runtime.Value
var once_newtypeConj sync.Once
func Get_newtypeConj() gopurs_runtime.Value {
	once_newtypeConj.Do(func() {
		newtypeConj = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeConj
}

var newtypeAdditive gopurs_runtime.Value
var once_newtypeAdditive sync.Once
func Get_newtypeAdditive() gopurs_runtime.Value {
	once_newtypeAdditive.Do(func() {
		newtypeAdditive = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeAdditive
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, fn_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(_dollar__unused_0_box, fn_1_box, t_2_box)
})
	})
	return modify
}

var collect gopurs_runtime.Value
var once_collect sync.Once
func Get_collect() gopurs_runtime.Value {
	once_collect.Do(func() {
		collect = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_collect(_dollar__unused_0_box, _dollar__unused_1_box, v_2_box)
})
	})
	return collect
}

var alaF gopurs_runtime.Value
var once_alaF sync.Once
func Get_alaF() gopurs_runtime.Value {
	once_alaF.Do(func() {
		alaF = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, _dollar__unused_3_box, v_4_box)
})
	})
	return alaF
}

var ala gopurs_runtime.Value
var once_ala sync.Once
func Get_ala() gopurs_runtime.Value {
	once_ala.Do(func() {
		ala = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ala(_dollar__unused_0_box, _dollar__unused_1_box, _dollar__unused_2_box, v_3_box, f_4_box)
})
	})
	return ala
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


