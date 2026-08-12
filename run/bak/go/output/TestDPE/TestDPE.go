package TestDPE

import (
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_mempty_ gopurs_runtime.Value
var once_mempty_ sync.Once
func Get_mempty_() gopurs_runtime.Value {
	once_mempty_.Do(func() {
		cache_mempty_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty_(dict_0_box)
})
	})
	return cache_mempty_
}

var cache_mappend_ gopurs_runtime.Value
var once_mappend_ sync.Once
func Get_mappend_() gopurs_runtime.Value {
	once_mappend_.Do(func() {
		cache_mappend_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend_(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mappend_
}

var cache_polyLoop gopurs_runtime.Value
var once_polyLoop sync.Once
func Get_polyLoop() gopurs_runtime.Value {
	once_polyLoop.Do(func() {
		cache_polyLoop = gopurs_runtime.Func(func(dictMonoidish_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dictMonoidish_0_box))
})
	})
	return cache_polyLoop
}

var cache_intMonoidish gopurs_runtime.Value
var once_intMonoidish sync.Once
func Get_intMonoidish() gopurs_runtime.Value {
	once_intMonoidish.Do(func() {
		cache_intMonoidish = gopurs_runtime.RecordDict2("mappend_", "mempty_", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(Get_add__560788792(), a_0, b_1).IntVal)
})
}), gopurs_runtime.Int(0))
	})
	return cache_intMonoidish
}

var cache_test gopurs_runtime.Value
var once_test sync.Once
func Get_test() gopurs_runtime.Value {
	once_test.Do(func() {
		cache_test = gopurs_runtime.Int(Call_polyLoop__1285086188(10000000, gopurs_runtime.Int(0)).IntVal)
	})
	return cache_test
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

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mappend___3425151628 gopurs_runtime.Value
var once_mappend___3425151628 sync.Once
func Get_mappend___3425151628() gopurs_runtime.Value {
	once_mappend___3425151628.Do(func() {
		cache_mappend___3425151628 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mappend___3425151628(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mappend___3425151628
}

var cache_polyLoop__1285086188 gopurs_runtime.Value
var once_polyLoop__1285086188 sync.Once
func Get_polyLoop__1285086188() gopurs_runtime.Value {
	once_polyLoop__1285086188.Do(func() {
		cache_polyLoop__1285086188 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop__1285086188(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_polyLoop__1285086188
}

var cache_polyLoop__11377801 gopurs_runtime.Value
var once_polyLoop__11377801 sync.Once
func Get_polyLoop__11377801() gopurs_runtime.Value {
	once_polyLoop__11377801.Do(func() {
		cache_polyLoop__11377801 = gopurs_runtime.Func(func(dictMonoidish_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_polyLoop__11377801(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](dictMonoidish_0_box))
})
	})
	return cache_polyLoop__11377801
}

type Constructor_Monoidish[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}


func init() {
	gopurs_runtime.StructGetters[2768669742] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Monoidish[gopurs_runtime.Value])(ptr)
		switch key {
		case "mappend_": return c.V0
		case "mempty_": return c.V1
		default: panic("Key not found in dictionary Constructor_Monoidish: " + key)
		}
	}
}


func Call_mempty_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty_")
}

func Call_mappend_(dict_0_loop *Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_polyLoop(dictMonoidish_0_loop *Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
polyLoop:
for {
if false { continue polyLoop }
var dictMonoidish_0 *Constructor_Monoidish[gopurs_runtime.Value] = dictMonoidish_0_loop
_ = dictMonoidish_0
mempty_1_1_0 := dictMonoidish_0.V1
_ = mempty_1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_polyLoop(dictMonoidish_0), gopurs_runtime.Apply2(Get_sub__1043827704(), v_2, gopurs_runtime.Int(1)), gopurs_runtime.Apply2(dictMonoidish_0.V0, v1_3, mempty_1_1_0))
}
end_branch_1:
return __t1
})
})
}
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mappend___3425151628(dict_0_loop *Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Monoidish[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_polyLoop__1285086188(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Call_polyLoop(gopurs_runtime.CoerceToStruct[Constructor_Monoidish[gopurs_runtime.Value]](Get_intMonoidish())), gopurs_runtime.Int((v_0) - (1)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_intMonoidish(), "mappend_"), v1_1, gopurs_runtime.RecordGet(Get_intMonoidish(), "mempty_")))
}
end_branch_0:
return __t0
}

func Call_polyLoop__11377801(dictMonoidish_0_loop *Constructor_Monoidish[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoidish_0 *Constructor_Monoidish[gopurs_runtime.Value] = dictMonoidish_0_loop
_ = dictMonoidish_0
mempty_1_1_0 := dictMonoidish_0.V1
_ = mempty_1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) == (0) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(Call_polyLoop(dictMonoidish_0), gopurs_runtime.Int((v_2.IntVal) - (1)), gopurs_runtime.Apply2(dictMonoidish_0.V0, v1_3, mempty_1_1_0))
}
end_branch_1:
return __t1
})
})
}


