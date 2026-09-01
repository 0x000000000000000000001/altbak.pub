package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Implicit_foo gopurs_runtime.Value
var once_Test_Implicit_foo sync.Once

func Get_Test_Implicit_foo() gopurs_runtime.Value {
	once_Test_Implicit_foo.Do(func() {
		cache_Test_Implicit_foo = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Implicit_foo(x_0_box)
		})
	})
	return cache_Test_Implicit_foo
}

var cache_Test_Implicit_foo__3883318852 gopurs_runtime.Value
var once_Test_Implicit_foo__3883318852 sync.Once

func Get_Test_Implicit_foo__3883318852() gopurs_runtime.Value {
	once_Test_Implicit_foo__3883318852.Do(func() {
		cache_Test_Implicit_foo__3883318852 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_Implicit_foo__3883318852(x_0_box.IntVal))
		})
	})
	return cache_Test_Implicit_foo__3883318852
}

var cache_Test_Implicit_bar gopurs_runtime.Value
var once_Test_Implicit_bar sync.Once

func Get_Test_Implicit_bar() gopurs_runtime.Value {
	once_Test_Implicit_bar.Do(func() {
		cache_Test_Implicit_bar = gopurs_runtime.Int(42)
	})
	return cache_Test_Implicit_bar
}

func Call_Test_Implicit_foo(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Test_Implicit_foo__3883318852(x_0_loop int64) int64 {
foo__3883318852:
	for {
		if false {
			continue foo__3883318852
		}
		var x_0 int64 = x_0_loop
		_ = x_0
		return x_0
	}
}
