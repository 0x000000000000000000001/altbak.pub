package Data_Functor_Invariant

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor "gopurs/output/Data.Functor"
)

var invariantMultiplicative gopurs_runtime.Value
var once_invariantMultiplicative sync.Once
func Get_invariantMultiplicative() gopurs_runtime.Value {
	once_invariantMultiplicative.Do(func() {
		invariantMultiplicative = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
}))
	})
	return invariantMultiplicative
}

var invariantEndo gopurs_runtime.Value
var once_invariantEndo sync.Once
func Get_invariantEndo() gopurs_runtime.Value {
	once_invariantEndo.Do(func() {
		invariantEndo = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func4(func(ab_0 gopurs_runtime.Value, ba_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ab_0, gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(ba_1, x_3)))
}))
	})
	return invariantEndo
}

var invariantDual gopurs_runtime.Value
var once_invariantDual sync.Once
func Get_invariantDual() gopurs_runtime.Value {
	once_invariantDual.Do(func() {
		invariantDual = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
}))
	})
	return invariantDual
}

var invariantDisj gopurs_runtime.Value
var once_invariantDisj sync.Once
func Get_invariantDisj() gopurs_runtime.Value {
	once_invariantDisj.Do(func() {
		invariantDisj = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
}))
	})
	return invariantDisj
}

var invariantConj gopurs_runtime.Value
var once_invariantConj sync.Once
func Get_invariantConj() gopurs_runtime.Value {
	once_invariantConj.Do(func() {
		invariantConj = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
}))
	})
	return invariantConj
}

var invariantAdditive gopurs_runtime.Value
var once_invariantAdditive sync.Once
func Get_invariantAdditive() gopurs_runtime.Value {
	once_invariantAdditive.Do(func() {
		invariantAdditive = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
}))
	})
	return invariantAdditive
}

var imapF gopurs_runtime.Value
var once_imapF sync.Once
func Get_imapF() gopurs_runtime.Value {
	once_imapF.Do(func() {
		imapF = gopurs_runtime.Func3(func(dictFunctor_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1)
})
	})
	return imapF
}

var invariantArray gopurs_runtime.Value
var once_invariantArray sync.Once
func Get_invariantArray() gopurs_runtime.Value {
	once_invariantArray.Do(func() {
		invariantArray = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Functor.Get_arrayMap(), f_0)
}))
	})
	return invariantArray
}

var invariantFn gopurs_runtime.Value
var once_invariantFn sync.Once
func Get_invariantFn() gopurs_runtime.Value {
	once_invariantFn.Do(func() {
		invariantFn = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_2, x_3))
}))
	})
	return invariantFn
}

var imap gopurs_runtime.Value
var once_imap sync.Once
func Get_imap() gopurs_runtime.Value {
	once_imap.Do(func() {
		imap = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "imap")
})
	})
	return imap
}

var invariantAlternate gopurs_runtime.Value
var once_invariantAlternate sync.Once
func Get_invariantAlternate() gopurs_runtime.Value {
	once_invariantAlternate.Do(func() {
		invariantAlternate = gopurs_runtime.Func(func(dictInvariant_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictInvariant_0, "imap"), f_1, g_2, v_3)
}))
})
	})
	return invariantAlternate
}


