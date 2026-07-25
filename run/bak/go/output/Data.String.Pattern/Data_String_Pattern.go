package Data_String_Pattern

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var cache_Replacement gopurs_runtime.Value
var once_Replacement sync.Once
func Get_Replacement() gopurs_runtime.Value {
	once_Replacement.Do(func() {
		cache_Replacement = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Replacement
}

var cache_Pattern gopurs_runtime.Value
var once_Pattern sync.Once
func Get_Pattern() gopurs_runtime.Value {
	once_Pattern.Do(func() {
		cache_Pattern = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_Pattern
}

var cache_showReplacement gopurs_runtime.Value
var once_showReplacement sync.Once
func Get_showReplacement() gopurs_runtime.Value {
	once_showReplacement.Do(func() {
		cache_showReplacement = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Replacement ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showReplacement
}

var cache_showPattern gopurs_runtime.Value
var once_showPattern sync.Once
func Get_showPattern() gopurs_runtime.Value {
	once_showPattern.Do(func() {
		cache_showPattern = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showPattern
}

var cache_newtypeReplacement gopurs_runtime.Value
var once_newtypeReplacement sync.Once
func Get_newtypeReplacement() gopurs_runtime.Value {
	once_newtypeReplacement.Do(func() {
		cache_newtypeReplacement = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeReplacement
}

var cache_newtypePattern gopurs_runtime.Value
var once_newtypePattern sync.Once
func Get_newtypePattern() gopurs_runtime.Value {
	once_newtypePattern.Do(func() {
		cache_newtypePattern = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypePattern
}

var cache_eqReplacement gopurs_runtime.Value
var once_eqReplacement sync.Once
func Get_eqReplacement() gopurs_runtime.Value {
	once_eqReplacement.Do(func() {
		cache_eqReplacement = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
}))
	})
	return cache_eqReplacement
}

var cache_ordReplacement gopurs_runtime.Value
var once_ordReplacement sync.Once
func Get_ordReplacement() gopurs_runtime.Value {
	once_ordReplacement.Do(func() {
		cache_ordReplacement = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordString(), "compare"), x_0, y_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqReplacement()
}))
	})
	return cache_ordReplacement
}

var cache_eqPattern gopurs_runtime.Value
var once_eqPattern sync.Once
func Get_eqPattern() gopurs_runtime.Value {
	once_eqPattern.Do(func() {
		cache_eqPattern = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
}))
	})
	return cache_eqPattern
}

var cache_ordPattern gopurs_runtime.Value
var once_ordPattern sync.Once
func Get_ordPattern() gopurs_runtime.Value {
	once_ordPattern.Do(func() {
		cache_ordPattern = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordString(), "compare"), x_0, y_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqPattern()
}))
	})
	return cache_ordPattern
}




