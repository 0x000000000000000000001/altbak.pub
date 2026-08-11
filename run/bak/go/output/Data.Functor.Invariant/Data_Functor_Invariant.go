package Data_Functor_Invariant

import (
	pkg_Data_Functor "gopurs/output/Data.Functor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_invariantMultiplicative gopurs_runtime.Value
var once_invariantMultiplicative sync.Once
func Get_invariantMultiplicative() gopurs_runtime.Value {
	once_invariantMultiplicative.Do(func() {
		cache_invariantMultiplicative = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_invariantMultiplicative
}

var cache_invariantEndo gopurs_runtime.Value
var once_invariantEndo sync.Once
func Get_invariantEndo() gopurs_runtime.Value {
	once_invariantEndo.Do(func() {
		cache_invariantEndo = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(ab_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ba_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ab_0, gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(ba_1, x_3)))
})
})
})
}))
	})
	return cache_invariantEndo
}

var cache_invariantDual gopurs_runtime.Value
var once_invariantDual sync.Once
func Get_invariantDual() gopurs_runtime.Value {
	once_invariantDual.Do(func() {
		cache_invariantDual = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_invariantDual
}

var cache_invariantDisj gopurs_runtime.Value
var once_invariantDisj sync.Once
func Get_invariantDisj() gopurs_runtime.Value {
	once_invariantDisj.Do(func() {
		cache_invariantDisj = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_invariantDisj
}

var cache_invariantConj gopurs_runtime.Value
var once_invariantConj sync.Once
func Get_invariantConj() gopurs_runtime.Value {
	once_invariantConj.Do(func() {
		cache_invariantConj = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_invariantConj
}

var cache_invariantAdditive gopurs_runtime.Value
var once_invariantAdditive sync.Once
func Get_invariantAdditive() gopurs_runtime.Value {
	once_invariantAdditive.Do(func() {
		cache_invariantAdditive = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_invariantAdditive
}

var cache_imapF gopurs_runtime.Value
var once_imapF sync.Once
func Get_imapF() gopurs_runtime.Value {
	once_imapF.Do(func() {
		cache_imapF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_imapF(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return cache_imapF
}

var cache_invariantArray gopurs_runtime.Value
var once_invariantArray sync.Once
func Get_invariantArray() gopurs_runtime.Value {
	once_invariantArray.Do(func() {
		cache_invariantArray = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), f_0)
})
}))
	})
	return cache_invariantArray
}

var cache_invariantFn gopurs_runtime.Value
var once_invariantFn sync.Once
func Get_invariantFn() gopurs_runtime.Value {
	once_invariantFn.Do(func() {
		cache_invariantFn = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorFn(), "map"), f_0)
})
}))
	})
	return cache_invariantFn
}

var cache_imap gopurs_runtime.Value
var once_imap sync.Once
func Get_imap() gopurs_runtime.Value {
	once_imap.Do(func() {
		cache_imap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_imap(dict_0_box)
})
	})
	return cache_imap
}

var cache_imap__gopurs_runtime_Value_2950557085 gopurs_runtime.Value
var once_imap__gopurs_runtime_Value_2950557085 sync.Once
func Get_imap__gopurs_runtime_Value_2950557085() gopurs_runtime.Value {
	once_imap__gopurs_runtime_Value_2950557085.Do(func() {
		cache_imap__gopurs_runtime_Value_2950557085 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_imap__gopurs_runtime_Value_2950557085(dict_0_box)
})
	})
	return cache_imap__gopurs_runtime_Value_2950557085
}

var cache_invariantAlternate gopurs_runtime.Value
var once_invariantAlternate sync.Once
func Get_invariantAlternate() gopurs_runtime.Value {
	once_invariantAlternate.Do(func() {
		cache_invariantAlternate = gopurs_runtime.Func(func(dictInvariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_invariantAlternate(dictInvariant_0_box)
})
	})
	return cache_invariantAlternate
}

func Call_imapF(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
}

func Call_imap(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "imap")
}

func Call_imap__gopurs_runtime_Value_2950557085(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "imap")
}

func Call_invariantAlternate(dictInvariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
_ = dictInvariant_0
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictInvariant_0, "imap"), f_1, g_2, v_3)
})
})
}))
}


