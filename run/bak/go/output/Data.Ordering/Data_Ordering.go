package Data_Ordering

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_LT gopurs_runtime.Value
var once_LT sync.Once
func Get_LT() gopurs_runtime.Value {
	once_LT.Do(func() {
		cache_LT = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
	})
	return cache_LT
}

var cache_GT gopurs_runtime.Value
var once_GT sync.Once
func Get_GT() gopurs_runtime.Value {
	once_GT.Do(func() {
		cache_GT = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
	})
	return cache_GT
}

var cache_EQ gopurs_runtime.Value
var once_EQ sync.Once
func Get_EQ() gopurs_runtime.Value {
	once_EQ.Do(func() {
		cache_EQ = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
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
if (v_0.Type == 9 && v_0.IntVal == 1527465420) {
__t0 = gopurs_runtime.Str("LT")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t0 = gopurs_runtime.Str("GT")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 902936544) {
__t0 = gopurs_runtime.Str("EQ")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return cache_showOrdering
}

var cache_semigroupOrdering gopurs_runtime.Value
var once_semigroupOrdering sync.Once
func Get_semigroupOrdering() gopurs_runtime.Value {
	once_semigroupOrdering.Do(func() {
		cache_semigroupOrdering = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1527465420) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 902936544) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
	})
	return cache_semigroupOrdering
}

var cache_invert gopurs_runtime.Value
var once_invert sync.Once
func Get_invert() gopurs_runtime.Value {
	once_invert.Do(func() {
		cache_invert = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_invert(v_0_box.IntVal)
})
	})
	return cache_invert
}

var cache_eqOrdering gopurs_runtime.Value
var once_eqOrdering sync.Once
func Get_eqOrdering() gopurs_runtime.Value {
	once_eqOrdering.Do(func() {
		cache_eqOrdering = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1527465420) {
__t0 = gopurs_runtime.Bool((v1_1.Type == 9 && v1_1.IntVal == 1527465420))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 380165415) {
__t0 = gopurs_runtime.Bool((v1_1.Type == 9 && v1_1.IntVal == 380165415))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((v_0.Type == 9 && v_0.IntVal == 902936544)) && ((v1_1.Type == 9 && v1_1.IntVal == 902936544)))
}
end_branch_0:
return __t0
}))
	})
	return cache_eqOrdering
}

type Constructor_LT struct {
	
}


type Constructor_GT struct {
	
}


type Constructor_EQ struct {
	
}


func Call_invert(v_0_loop int64) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 380165415) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 == 902936544) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 == 1527465420) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}


