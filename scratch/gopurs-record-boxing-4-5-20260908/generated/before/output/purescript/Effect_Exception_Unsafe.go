package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Exception_Unsafe_unsafeThrowException gopurs_runtime.Value
var once_Effect_Exception_Unsafe_unsafeThrowException sync.Once
func Get_Effect_Exception_Unsafe_unsafeThrowException() gopurs_runtime.Value {
	once_Effect_Exception_Unsafe_unsafeThrowException.Do(func() {
		cache_Effect_Exception_Unsafe_unsafeThrowException = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Exception_Unsafe_unsafeThrowException(x_0_box)
})
	})
	return cache_Effect_Exception_Unsafe_unsafeThrowException
}

var cache_Effect_Exception_Unsafe_unsafeThrow gopurs_runtime.Value
var once_Effect_Exception_Unsafe_unsafeThrow sync.Once
func Get_Effect_Exception_Unsafe_unsafeThrow() gopurs_runtime.Value {
	once_Effect_Exception_Unsafe_unsafeThrow.Do(func() {
		cache_Effect_Exception_Unsafe_unsafeThrow = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Exception_Unsafe_unsafeThrow(x_0_box.StrVal())
})
	})
	return cache_Effect_Exception_Unsafe_unsafeThrow
}

func Call_Effect_Exception_Unsafe_unsafeThrowException(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply(Get_Effect_Exception_throwException(), x_0))
}

func Call_Effect_Exception_Unsafe_unsafeThrow(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Effect_Unsafe_unsafePerformEffect(), gopurs_runtime.Apply(Get_Effect_Exception_throwException(), gopurs_runtime.Apply(Get_Effect_Exception_error(), gopurs_runtime.Str(x_0))))
}


