package Effect_Exception_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Effect_Unsafe "gopurs/output/Effect.Unsafe"
	pkg_Effect_Exception "gopurs/output/Effect.Exception"
)

var cache_unsafeThrowException gopurs_runtime.Value
var once_unsafeThrowException sync.Once
func Get_unsafeThrowException() gopurs_runtime.Value {
	once_unsafeThrowException.Do(func() {
		cache_unsafeThrowException = gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), pkg_Effect_Unsafe.Get_unsafePerformEffect(), pkg_Effect_Exception.Get_throwException())
	})
	return cache_unsafeThrowException
}

var cache_unsafeThrow gopurs_runtime.Value
var once_unsafeThrow sync.Once
func Get_unsafeThrow() gopurs_runtime.Value {
	once_unsafeThrow.Do(func() {
		cache_unsafeThrow = gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_unsafeThrowException(), pkg_Effect_Exception.Get_error())
	})
	return cache_unsafeThrow
}




