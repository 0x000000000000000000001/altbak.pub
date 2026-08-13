package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Partial_crashWith gopurs_runtime.Value
var once_Partial_crashWith sync.Once
func Get_Partial_crashWith() gopurs_runtime.Value {
	once_Partial_crashWith.Do(func() {
		cache_Partial_crashWith = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Partial_crashWith(_dollar__unused_0_box)
})
	})
	return cache_Partial_crashWith
}

var cache_Partial_crash gopurs_runtime.Value
var once_Partial_crash sync.Once
func Get_Partial_crash() gopurs_runtime.Value {
	once_Partial_crash.Do(func() {
		cache_Partial_crash = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Partial_crash(_dollar__unused_0_box)
})
	})
	return cache_Partial_crash
}

var cache_Partial_crashWith__1894115486 gopurs_runtime.Value
var once_Partial_crashWith__1894115486 sync.Once
func Get_Partial_crashWith__1894115486() gopurs_runtime.Value {
	once_Partial_crashWith__1894115486.Do(func() {
		cache_Partial_crashWith__1894115486 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Partial_crashWith__1894115486(_dollar__unused_0_box)
})
	})
	return cache_Partial_crashWith__1894115486
}

var cache_Partial_crashWith__286259978 gopurs_runtime.Value
var once_Partial_crashWith__286259978 sync.Once
func Get_Partial_crashWith__286259978() gopurs_runtime.Value {
	once_Partial_crashWith__286259978.Do(func() {
		cache_Partial_crashWith__286259978 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Partial_crashWith__286259978(_dollar__unused_0_box)
})
	})
	return cache_Partial_crashWith__286259978
}

var cache_Partial_crashWith__3377663964 gopurs_runtime.Value
var once_Partial_crashWith__3377663964 sync.Once
func Get_Partial_crashWith__3377663964() gopurs_runtime.Value {
	once_Partial_crashWith__3377663964.Do(func() {
		cache_Partial_crashWith__3377663964 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Partial_crashWith__3377663964(_dollar__unused_0_box)
})
	})
	return cache_Partial_crashWith__3377663964
}

var cache_Partial_crashWith__2090867676 gopurs_runtime.Value
var once_Partial_crashWith__2090867676 sync.Once
func Get_Partial_crashWith__2090867676() gopurs_runtime.Value {
	once_Partial_crashWith__2090867676.Do(func() {
		cache_Partial_crashWith__2090867676 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Partial_crashWith__2090867676(_dollar__unused_0_box)
})
	})
	return cache_Partial_crashWith__2090867676
}

var cache_Partial_crashWith__572531106 gopurs_runtime.Value
var once_Partial_crashWith__572531106 sync.Once
func Get_Partial_crashWith__572531106() gopurs_runtime.Value {
	once_Partial_crashWith__572531106.Do(func() {
		cache_Partial_crashWith__572531106 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Partial_crashWith__572531106(_dollar__unused_0_box)
})
	})
	return cache_Partial_crashWith__572531106
}

func Call_Partial_crashWith(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get_Partial__crashWith()
}

func Call_Partial_crash(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str("Partial.crash: partial function"))
}

func Call_Partial_crashWith__1894115486(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get_Partial__crashWith()
}

func Call_Partial_crashWith__286259978(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get_Partial__crashWith()
}

func Call_Partial_crashWith__3377663964(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get_Partial__crashWith()
}

func Call_Partial_crashWith__2090867676(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get_Partial__crashWith()
}

func Call_Partial_crashWith__572531106(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get_Partial__crashWith()
}

func Get_Partial__crashWith() gopurs_runtime.Value {
	return _Gopurs_Partial__CrashWith
}
