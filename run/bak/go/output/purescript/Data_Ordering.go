package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Ordering_LT gopurs_runtime.Value
var once_Data_Ordering_LT sync.Once
func Get_Data_Ordering_LT() gopurs_runtime.Value {
	once_Data_Ordering_LT.Do(func() {
		cache_Data_Ordering_LT = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
	})
	return cache_Data_Ordering_LT
}

var cache_Data_Ordering_GT gopurs_runtime.Value
var once_Data_Ordering_GT sync.Once
func Get_Data_Ordering_GT() gopurs_runtime.Value {
	once_Data_Ordering_GT.Do(func() {
		cache_Data_Ordering_GT = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
	})
	return cache_Data_Ordering_GT
}

var cache_Data_Ordering_EQ gopurs_runtime.Value
var once_Data_Ordering_EQ sync.Once
func Get_Data_Ordering_EQ() gopurs_runtime.Value {
	once_Data_Ordering_EQ.Do(func() {
		cache_Data_Ordering_EQ = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
	})
	return cache_Data_Ordering_EQ
}

var cache_Data_Ordering_showOrdering gopurs_runtime.Value
var once_Data_Ordering_showOrdering sync.Once
func Get_Data_Ordering_showOrdering() gopurs_runtime.Value {
	once_Data_Ordering_showOrdering.Do(func() {
		cache_Data_Ordering_showOrdering = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 string
{
if (uint32(v_0.IntVal) == 1527465420) {
__t0 = "LT"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
__t0 = "GT"
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
__t0 = "EQ"
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_0:
return gopurs_runtime.Str(__t0)
}))
	})
	return cache_Data_Ordering_showOrdering
}

var cache_Data_Ordering_semigroupOrdering gopurs_runtime.Value
var once_Data_Ordering_semigroupOrdering sync.Once
func Get_Data_Ordering_semigroupOrdering() gopurs_runtime.Value {
	once_Data_Ordering_semigroupOrdering.Do(func() {
		cache_Data_Ordering_semigroupOrdering = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 uint32
{
if (uint32(v_0.IntVal) == 1527465420) {
__t0 = 1527465420
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
__t0 = 380165415
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
__t0 = uint32(v1_1.IntVal)
goto end_branch_0
} else {

}
}
{
__t0 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Ordering_semigroupOrdering
}

var cache_Data_Ordering_invert gopurs_runtime.Value
var once_Data_Ordering_invert sync.Once
func Get_Data_Ordering_invert() gopurs_runtime.Value {
	once_Data_Ordering_invert.Do(func() {
		cache_Data_Ordering_invert = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Ordering_invert(uint32(v_0_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_Data_Ordering_invert
}

var cache_Data_Ordering_eqOrdering gopurs_runtime.Value
var once_Data_Ordering_eqOrdering sync.Once
func Get_Data_Ordering_eqOrdering() gopurs_runtime.Value {
	once_Data_Ordering_eqOrdering.Do(func() {
		cache_Data_Ordering_eqOrdering = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 bool
{
if (uint32(v_0.IntVal) == 1527465420) {
var __t0 bool
{
if (uint32(v1_1.IntVal) == 1527465420) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t2 = __t0
goto end_branch_2
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
var __t1 bool
{
if (uint32(v1_1.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((uint32(v_0.IntVal) == 902936544)) && ((uint32(v1_1.IntVal) == 902936544)) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})
}))
	})
	return cache_Data_Ordering_eqOrdering
}

type Constructor_Data_Ordering_LT struct {
	Rc uint32
}


type Constructor_Data_Ordering_GT struct {
	Rc uint32
}


type Constructor_Data_Ordering_EQ struct {
	Rc uint32
}


func Call_Data_Ordering_invert(v_0_loop uint32) uint32 {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 uint32
{
if (v_0 == 380165415) {
__t0 = 1527465420
goto end_branch_0
} else {

}
}
{
if (v_0 == 902936544) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
if (v_0 == 1527465420) {
__t0 = 380165415
goto end_branch_0
} else {

}
}
{
__t0 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_0:
return __t0
}


