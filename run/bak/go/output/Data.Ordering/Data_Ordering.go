package Data_Ordering

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_LT gopurs_runtime.Value
var once_LT sync.Once
func Get_LT() gopurs_runtime.Value {
	once_LT.Do(func() {
		cache_LT = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
	})
	return cache_LT
}

var cache_GT gopurs_runtime.Value
var once_GT sync.Once
func Get_GT() gopurs_runtime.Value {
	once_GT.Do(func() {
		cache_GT = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
	})
	return cache_GT
}

var cache_EQ gopurs_runtime.Value
var once_EQ sync.Once
func Get_EQ() gopurs_runtime.Value {
	once_EQ.Do(func() {
		cache_EQ = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
	})
	return cache_EQ
}

var cache_showOrdering gopurs_runtime.Value
var once_showOrdering sync.Once
func Get_showOrdering() gopurs_runtime.Value {
	once_showOrdering.Do(func() {
		cache_showOrdering = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_0.IntVal) == 1527465420) {
__t0 = gopurs_runtime.Str("LT")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
__t0 = gopurs_runtime.Str("GT")
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
__t0 = gopurs_runtime.Str("EQ")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
	})
	return cache_showOrdering
}

var cache_semigroupOrdering gopurs_runtime.Value
var once_semigroupOrdering sync.Once
func Get_semigroupOrdering() gopurs_runtime.Value {
	once_semigroupOrdering.Do(func() {
		cache_semigroupOrdering = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_0.IntVal) == 1527465420) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 902936544) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(v1_1.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_semigroupOrdering
}

var cache_invert gopurs_runtime.Value
var once_invert sync.Once
func Get_invert() gopurs_runtime.Value {
	once_invert.Do(func() {
		cache_invert = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_invert(uint32(v_0_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_invert
}

var cache_eqOrdering gopurs_runtime.Value
var once_eqOrdering sync.Once
func Get_eqOrdering() gopurs_runtime.Value {
	once_eqOrdering.Do(func() {
		cache_eqOrdering = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (uint32(v_0.IntVal) == 1527465420) {
var __t1 gopurs_runtime.Value
{
if (uint32(v1_1.IntVal) == 1527465420) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (uint32(v_0.IntVal) == 380165415) {
var __t2 gopurs_runtime.Value
{
if (uint32(v1_1.IntVal) == 380165415) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(false)
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if ((uint32(v_0.IntVal) == 902936544)) && ((uint32(v1_1.IntVal) == 902936544)) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
})
}))
	})
	return cache_eqOrdering
}

type Constructor_LT struct {
	Rc uint32
}


type Constructor_GT struct {
	Rc uint32
}


type Constructor_EQ struct {
	Rc uint32
}


func Call_invert(v_0_loop uint32) uint32 {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 380165415) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 == 902936544) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 == 1527465420) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(__t0.IntVal)
}


