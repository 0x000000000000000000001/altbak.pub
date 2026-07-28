package Data_Function_Uncurried

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_runFn1 gopurs_runtime.Value
var once_runFn1 sync.Once
func Get_runFn1() gopurs_runtime.Value {
	once_runFn1.Do(func() {
		cache_runFn1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runFn1(f_0_box)
})
	})
	return cache_runFn1
}

var cache_mkFn1 gopurs_runtime.Value
var once_mkFn1 sync.Once
func Get_mkFn1() gopurs_runtime.Value {
	once_mkFn1.Do(func() {
		cache_mkFn1 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mkFn1(f_0_box)
})
	})
	return cache_mkFn1
}

func Call_runFn1(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return f_0
}

func Call_mkFn1(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return f_0
}

func Get_mkFn0() gopurs_runtime.Value {
	return _Gopurs_MkFn0
}

func Get_mkFn10() gopurs_runtime.Value {
	return _Gopurs_MkFn10
}

func Get_mkFn2() gopurs_runtime.Value {
	return _Gopurs_MkFn2
}

func Get_mkFn3() gopurs_runtime.Value {
	return _Gopurs_MkFn3
}

func Get_mkFn4() gopurs_runtime.Value {
	return _Gopurs_MkFn4
}

func Get_mkFn5() gopurs_runtime.Value {
	return _Gopurs_MkFn5
}

func Get_mkFn6() gopurs_runtime.Value {
	return _Gopurs_MkFn6
}

func Get_mkFn7() gopurs_runtime.Value {
	return _Gopurs_MkFn7
}

func Get_mkFn8() gopurs_runtime.Value {
	return _Gopurs_MkFn8
}

func Get_mkFn9() gopurs_runtime.Value {
	return _Gopurs_MkFn9
}

func Get_runFn0() gopurs_runtime.Value {
	return _Gopurs_RunFn0
}

func Get_runFn10() gopurs_runtime.Value {
	return _Gopurs_RunFn10
}

func Get_runFn2() gopurs_runtime.Value {
	return _Gopurs_RunFn2
}

func Get_runFn3() gopurs_runtime.Value {
	return _Gopurs_RunFn3
}

func Get_runFn4() gopurs_runtime.Value {
	return _Gopurs_RunFn4
}

func Get_runFn5() gopurs_runtime.Value {
	return _Gopurs_RunFn5
}

func Get_runFn6() gopurs_runtime.Value {
	return _Gopurs_RunFn6
}

func Get_runFn7() gopurs_runtime.Value {
	return _Gopurs_RunFn7
}

func Get_runFn8() gopurs_runtime.Value {
	return _Gopurs_RunFn8
}

func Get_runFn9() gopurs_runtime.Value {
	return _Gopurs_RunFn9
}
