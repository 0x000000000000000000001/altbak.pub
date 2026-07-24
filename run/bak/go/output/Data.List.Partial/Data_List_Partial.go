package Data_List_Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		tail = gopurs_runtime.Func2(Call_tail)
	})
	return tail
}

var last gopurs_runtime.Value
var once_last sync.Once
func Get_last() gopurs_runtime.Value {
	once_last.Do(func() {
		last = gopurs_runtime.Func2(Call_last)
	})
	return last
}

var init_ gopurs_runtime.Value
var once_init_ sync.Once
func Get_init_() gopurs_runtime.Value {
	once_init_.Do(func() {
		init_ = gopurs_runtime.Func2(Call_init_)
	})
	return init_
}

var head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		head = gopurs_runtime.Func2(Call_head)
	})
	return head
}

func Call_tail(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1_loop.StrVal == "Cons").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1]
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

func Call_last(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
last:
for {
if false { continue last }
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1_loop.StrVal == "Cons").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1].StrVal == "Nil").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0]
goto end_branch_1
} else {

}
}
{
__t1 = Call_last(gopurs_runtime.Value{}, (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1])
}
end_branch_1:
__t0 = __t1
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
}

func Call_init_(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
init_:
for {
if false { continue init_ }
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1_loop.StrVal == "Cons").IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool((*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1].StrVal == "Nil").IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nil")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor2("Cons", (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0], gopurs_runtime.Apply2(Get_init_(), gopurs_runtime.Value{}, (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[1]))
}
end_branch_1:
__t0 = __t1
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
}

func Call_head(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_1_loop.StrVal == "Cons").IntVal != 0 {
__t0 = (*[1024]gopurs_runtime.Value)(v_1_loop.UnsafePtr)[0]
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


