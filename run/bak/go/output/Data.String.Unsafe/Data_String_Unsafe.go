package Data_String_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_char gopurs_runtime.Value
var once_char sync.Once
func Get_char() gopurs_runtime.Value {
	once_char.Do(func() {
		cache_char = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Char(arg0.StrVal()))
})
	})
	return cache_char
}

var cache_charAt gopurs_runtime.Value
var once_charAt sync.Once
func Get_charAt() gopurs_runtime.Value {
	once_charAt.Do(func() {
		cache_charAt = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(CharAt(arg0.IntVal, arg1.StrVal()))
})
	})
	return cache_charAt
}


