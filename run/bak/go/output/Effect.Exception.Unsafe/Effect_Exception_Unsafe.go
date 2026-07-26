package Effect_Exception_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Unsafe "gopurs/output/Effect.Unsafe"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
)

var cache_unsafeThrowException gopurs_runtime.Value
var once_unsafeThrowException sync.Once
func Get_unsafeThrowException() gopurs_runtime.Value {
	once_unsafeThrowException.Do(func() {
		cache_unsafeThrowException = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeThrowException(x_0_box)
})
	})
	return cache_unsafeThrowException
}

var cache_unsafeThrow gopurs_runtime.Value
var once_unsafeThrow sync.Once
func Get_unsafeThrow() gopurs_runtime.Value {
	once_unsafeThrow.Do(func() {
		cache_unsafeThrow = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeThrow(x_0_box.StrVal())
})
	})
	return cache_unsafeThrow
}

func Call_unsafeThrowException(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_throwException(), x_0))
}

func Call_unsafeThrow(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Effect_Unsafe.Get_unsafePerformEffect(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_throwException(), gopurs_runtime.Apply(pkg_Effect_Exception.Get_error(), gopurs_runtime.Str(x_0))))
}


