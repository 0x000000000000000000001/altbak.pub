package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Functor_Invariant_Invariant_dollarDict gopurs_runtime.Value
var once_Data_Functor_Invariant_Invariant_dollarDict sync.Once
func Get_Data_Functor_Invariant_Invariant_dollarDict() gopurs_runtime.Value {
	once_Data_Functor_Invariant_Invariant_dollarDict.Do(func() {
		cache_Data_Functor_Invariant_Invariant_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Invariant_Invariant_dollarDict(x_0_box)
})
	})
	return cache_Data_Functor_Invariant_Invariant_dollarDict
}

var cache_Data_Functor_Invariant_invariantMultiplicative gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantMultiplicative sync.Once
func Get_Data_Functor_Invariant_invariantMultiplicative() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantMultiplicative.Do(func() {
		cache_Data_Functor_Invariant_invariantMultiplicative = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_Data_Functor_Invariant_invariantMultiplicative
}

var cache_Data_Functor_Invariant_invariantEndo gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantEndo sync.Once
func Get_Data_Functor_Invariant_invariantEndo() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantEndo.Do(func() {
		cache_Data_Functor_Invariant_invariantEndo = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(ab_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ba_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(ab_0, gopurs_runtime.Apply(v_2, gopurs_runtime.Apply(ba_1, x_3)))
})
})
})
}))
	})
	return cache_Data_Functor_Invariant_invariantEndo
}

var cache_Data_Functor_Invariant_invariantDual gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantDual sync.Once
func Get_Data_Functor_Invariant_invariantDual() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantDual.Do(func() {
		cache_Data_Functor_Invariant_invariantDual = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_Data_Functor_Invariant_invariantDual
}

var cache_Data_Functor_Invariant_invariantDisj gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantDisj sync.Once
func Get_Data_Functor_Invariant_invariantDisj() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantDisj.Do(func() {
		cache_Data_Functor_Invariant_invariantDisj = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_Data_Functor_Invariant_invariantDisj
}

var cache_Data_Functor_Invariant_invariantConj gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantConj sync.Once
func Get_Data_Functor_Invariant_invariantConj() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantConj.Do(func() {
		cache_Data_Functor_Invariant_invariantConj = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_Data_Functor_Invariant_invariantConj
}

var cache_Data_Functor_Invariant_invariantAdditive gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantAdditive sync.Once
func Get_Data_Functor_Invariant_invariantAdditive() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantAdditive.Do(func() {
		cache_Data_Functor_Invariant_invariantAdditive = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_Data_Functor_Invariant_invariantAdditive
}

var cache_Data_Functor_Invariant_imapF gopurs_runtime.Value
var once_Data_Functor_Invariant_imapF sync.Once
func Get_Data_Functor_Invariant_imapF() gopurs_runtime.Value {
	once_Data_Functor_Invariant_imapF.Do(func() {
		cache_Data_Functor_Invariant_imapF = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Invariant_imapF(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Data_Functor_Invariant_imapF
}

var cache_Data_Functor_Invariant_invariantArray gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantArray sync.Once
func Get_Data_Functor_Invariant_invariantArray() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantArray.Do(func() {
		cache_Data_Functor_Invariant_invariantArray = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), f_0)
})
}))
	})
	return cache_Data_Functor_Invariant_invariantArray
}

var cache_Data_Functor_Invariant_invariantFn gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantFn sync.Once
func Get_Data_Functor_Invariant_invariantFn() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantFn.Do(func() {
		cache_Data_Functor_Invariant_invariantFn = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), f_0)
})
}))
	})
	return cache_Data_Functor_Invariant_invariantFn
}

var cache_Data_Functor_Invariant_imap gopurs_runtime.Value
var once_Data_Functor_Invariant_imap sync.Once
func Get_Data_Functor_Invariant_imap() gopurs_runtime.Value {
	once_Data_Functor_Invariant_imap.Do(func() {
		cache_Data_Functor_Invariant_imap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Invariant_imap(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant](dict_0_box))
})
	})
	return cache_Data_Functor_Invariant_imap
}

var cache_Data_Functor_Invariant_invariantAlternate gopurs_runtime.Value
var once_Data_Functor_Invariant_invariantAlternate sync.Once
func Get_Data_Functor_Invariant_invariantAlternate() gopurs_runtime.Value {
	once_Data_Functor_Invariant_invariantAlternate.Do(func() {
		cache_Data_Functor_Invariant_invariantAlternate = gopurs_runtime.Func(func(dictInvariant_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Invariant_invariantAlternate(dictInvariant_0_box)
})
	})
	return cache_Data_Functor_Invariant_invariantAlternate
}

var cache_Data_Functor_Invariant_imap__2950557085 gopurs_runtime.Value
var once_Data_Functor_Invariant_imap__2950557085 sync.Once
func Get_Data_Functor_Invariant_imap__2950557085() gopurs_runtime.Value {
	once_Data_Functor_Invariant_imap__2950557085.Do(func() {
		cache_Data_Functor_Invariant_imap__2950557085 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Invariant_imap__2950557085(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Invariant_Invariant](dict_0_box))
})
	})
	return cache_Data_Functor_Invariant_imap__2950557085
}

type Constructor_Data_Functor_Invariant_Invariant struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2396985522] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Functor_Invariant_Invariant)(ptr)
		_ = c
		switch key {
		case "imap": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Functor_Invariant_Invariant: " + key)
		}
	}
}


func Call_Data_Functor_Invariant_Invariant_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Invariant_imapF(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_0.V0), f_1)
}

func Call_Data_Functor_Invariant_imap(dict_0_loop *Constructor_Data_Functor_Invariant_Invariant) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Invariant_Invariant = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Functor_Invariant_invariantAlternate(dictInvariant_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Functor_Invariant_imap__2950557085(dict_0_loop *Constructor_Data_Functor_Invariant_Invariant) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Functor_Invariant_Invariant = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}


