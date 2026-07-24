package Data_String_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		null = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(s_0.StrVal == "")
})
	})
	return null
}

var localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		localeCompare = gopurs_runtime.Apply3(Get__localeCompare(), gopurs_runtime.Constructor0("LT"), gopurs_runtime.Constructor0("EQ"), gopurs_runtime.Constructor0("GT"))
	})
	return localeCompare
}



func Get__localeCompare() gopurs_runtime.Value {
	return _Gopurs__LocaleCompare
}

func Get_joinWith() gopurs_runtime.Value {
	return _Gopurs_JoinWith
}

func Get_replace() gopurs_runtime.Value {
	return _Gopurs_Replace
}

func Get_replaceAll() gopurs_runtime.Value {
	return _Gopurs_ReplaceAll
}

func Get_split() gopurs_runtime.Value {
	return _Gopurs_Split
}

func Get_toLower() gopurs_runtime.Value {
	return _Gopurs_ToLower
}

func Get_toUpper() gopurs_runtime.Value {
	return _Gopurs_ToUpper
}

func Get_trim() gopurs_runtime.Value {
	return _Gopurs_Trim
}
