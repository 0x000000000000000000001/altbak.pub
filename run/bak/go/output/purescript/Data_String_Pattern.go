package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_String_Pattern_Replacement gopurs_runtime.Value
var once_Data_String_Pattern_Replacement sync.Once
func Get_Data_String_Pattern_Replacement() gopurs_runtime.Value {
	once_Data_String_Pattern_Replacement.Do(func() {
		cache_Data_String_Pattern_Replacement = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Pattern_Replacement(x_0_box)
})
	})
	return cache_Data_String_Pattern_Replacement
}

var cache_Data_String_Pattern_Pattern gopurs_runtime.Value
var once_Data_String_Pattern_Pattern sync.Once
func Get_Data_String_Pattern_Pattern() gopurs_runtime.Value {
	once_Data_String_Pattern_Pattern.Do(func() {
		cache_Data_String_Pattern_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Pattern_Pattern(x_0_box)
})
	})
	return cache_Data_String_Pattern_Pattern
}

var cache_Data_String_Pattern_showReplacement gopurs_runtime.Value
var once_Data_String_Pattern_showReplacement sync.Once
func Get_Data_String_Pattern_showReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_showReplacement.Do(func() {
		cache_Data_String_Pattern_showReplacement = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Replacement ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_String_Pattern_showReplacement
}

var cache_Data_String_Pattern_showPattern gopurs_runtime.Value
var once_Data_String_Pattern_showPattern sync.Once
func Get_Data_String_Pattern_showPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_showPattern.Do(func() {
		cache_Data_String_Pattern_showPattern = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_String_Pattern_showPattern
}

var cache_Data_String_Pattern_newtypeReplacement gopurs_runtime.Value
var once_Data_String_Pattern_newtypeReplacement sync.Once
func Get_Data_String_Pattern_newtypeReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_newtypeReplacement.Do(func() {
		cache_Data_String_Pattern_newtypeReplacement = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_String_Pattern_newtypeReplacement
}

var cache_Data_String_Pattern_newtypePattern gopurs_runtime.Value
var once_Data_String_Pattern_newtypePattern sync.Once
func Get_Data_String_Pattern_newtypePattern() gopurs_runtime.Value {
	once_Data_String_Pattern_newtypePattern.Do(func() {
		cache_Data_String_Pattern_newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_String_Pattern_newtypePattern
}

var cache_Data_String_Pattern_eqReplacement gopurs_runtime.Value
var once_Data_String_Pattern_eqReplacement sync.Once
func Get_Data_String_Pattern_eqReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_eqReplacement.Do(func() {
		cache_Data_String_Pattern_eqReplacement = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
})
}))
	})
	return cache_Data_String_Pattern_eqReplacement
}

var cache_Data_String_Pattern_ordReplacement gopurs_runtime.Value
var once_Data_String_Pattern_ordReplacement sync.Once
func Get_Data_String_Pattern_ordReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_ordReplacement.Do(func() {
		cache_Data_String_Pattern_ordReplacement = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_1.StrVal()) == (y_2.StrVal()))
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_Data_String_Pattern_ordReplacement
}

var cache_Data_String_Pattern_eqPattern gopurs_runtime.Value
var once_Data_String_Pattern_eqPattern sync.Once
func Get_Data_String_Pattern_eqPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_eqPattern.Do(func() {
		cache_Data_String_Pattern_eqPattern = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
})
}))
	})
	return cache_Data_String_Pattern_eqPattern
}

var cache_Data_String_Pattern_ordPattern gopurs_runtime.Value
var once_Data_String_Pattern_ordPattern sync.Once
func Get_Data_String_Pattern_ordPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_ordPattern.Do(func() {
		cache_Data_String_Pattern_ordPattern = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_1.StrVal()) == (y_2.StrVal()))
})
}))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_Data_String_Pattern_ordPattern
}

func Call_Data_String_Pattern_Replacement(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_Pattern_Pattern(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}


