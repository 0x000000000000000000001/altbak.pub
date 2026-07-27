package Record_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_unsafeDelete gopurs_runtime.Value
var once_unsafeDelete sync.Once
func Get_unsafeDelete() gopurs_runtime.Value {
	once_unsafeDelete.Do(func() {
		cache_unsafeDelete = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(UnsafeDelete(arg0.StrVal(), gopurs_runtime.UnboxAny(arg1)))
})
	})
	return cache_unsafeDelete
}

var cache_unsafeGet gopurs_runtime.Value
var once_unsafeGet sync.Once
func Get_unsafeGet() gopurs_runtime.Value {
	once_unsafeGet.Do(func() {
		cache_unsafeGet = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(UnsafeGet(arg0.StrVal(), gopurs_runtime.UnboxAny(arg1)))
})
	})
	return cache_unsafeGet
}

var cache_unsafeHas gopurs_runtime.Value
var once_unsafeHas sync.Once
func Get_unsafeHas() gopurs_runtime.Value {
	once_unsafeHas.Do(func() {
		cache_unsafeHas = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(UnsafeHas(arg0.StrVal(), gopurs_runtime.UnboxAny(arg1)))
})
	})
	return cache_unsafeHas
}

var cache_unsafeSet gopurs_runtime.Value
var once_unsafeSet sync.Once
func Get_unsafeSet() gopurs_runtime.Value {
	once_unsafeSet.Do(func() {
		cache_unsafeSet = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(UnsafeSet(arg0.StrVal(), gopurs_runtime.UnboxAny(arg1), gopurs_runtime.UnboxAny(arg2)))
})
	})
	return cache_unsafeSet
}


