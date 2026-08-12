package TestLoop

import (
	pkg_Data_Ring "gopurs/output/Data.Ring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_loop gopurs_runtime.Value
var once_loop sync.Once
func Get_loop() gopurs_runtime.Value {
	once_loop.Do(func() {
		cache_loop = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_loop(v_0_box.IntVal))
})
	})
	return cache_loop
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

func Call_loop(v_0_loop int64) int64 {
loop:
for {
if false { continue loop }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 int64
{
if (v_0) == (0) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
v_0_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal
continue loop
__t0 = gopurs_runtime.Value{}.IntVal
}
end_branch_0:
return __t0
}
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}


