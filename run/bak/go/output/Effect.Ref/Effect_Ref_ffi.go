package Effect_Ref

import "gopurs/output/gopurs_runtime"


func _New(val interface{}, _ interface{}) interface{} {
	return map[string]interface{}{"value": val}
}
func NewWithSelf(f func(interface{}) interface{}, _ interface{}) interface{} {
	ref := map[string]interface{}{}
	ref["value"] = f(ref)
	return ref
}
func Read(ref interface{}, _ interface{}) interface{} {
	return ref.(map[string]interface{})["value"]
}
func ModifyImpl(f func(interface{}) interface{}, ref interface{}, _ interface{}) interface{} {
	t := f(ref.(map[string]interface{})["value"]).(map[string]interface{})
	ref.(map[string]interface{})["value"] = t["state"]
	return t["value"]
}
func Write(val interface{}, ref interface{}, _ interface{}) interface{} {
	ref.(map[string]interface{})["value"] = val
	return nil
}


// --- Auto-generated FFI wrappers ---
var _Gopurs__New = // TAST: (Func [(TypeVar s)] (ADT ["Effect","Effect"] [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _New(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ModifyImpl = // TAST: (Func [(Func [(TypeVar s)] (Record [state: (TypeVar s), value: (TypeVar b)])), (ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(TypeVar b)]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := ModifyImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_NewWithSelf = // TAST: (Func [(Func [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (TypeVar s))] (ADT ["Effect","Effect"] [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_res := NewWithSelf(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Read = // TAST: (Func [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(TypeVar s)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Read(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Write = // TAST: (Func [(TypeVar s), (ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := Write(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})