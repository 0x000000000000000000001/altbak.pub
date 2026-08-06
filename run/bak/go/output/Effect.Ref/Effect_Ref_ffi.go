package Effect_Ref



import (
	"sync"
	"gopurs/output/gopurs_runtime"
)

type RefState struct {
	mu  sync.Mutex
	val gopurs_runtime.Value
}

func _New(val gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Any(&RefState{val: val})
	})
}

func NewWithSelf(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		ref := &RefState{}
		ref.val = gopurs_runtime.Apply(f, gopurs_runtime.Any(ref))
		return gopurs_runtime.Any(ref)
	})
}

func Read(ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		state := ref.AnyVal().(*RefState)
		state.mu.Lock()
		defer state.mu.Unlock()
		return state.val
	})
}

func ModifyImpl(f gopurs_runtime.Value, ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		state := ref.AnyVal().(*RefState)
		state.mu.Lock()
		defer state.mu.Unlock()
		
		res := gopurs_runtime.Apply(f, state.val)
		record := gopurs_runtime.UnboxObject(res)
		
		state.val = record["state"].(gopurs_runtime.Value)
		return record["value"].(gopurs_runtime.Value)
	})
}

func Write(val gopurs_runtime.Value, ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		state := ref.AnyVal().(*RefState)
		state.mu.Lock()
		state.val = val
		state.mu.Unlock()
		return gopurs_runtime.Any(nil)
	})
}


// --- Auto-generated FFI wrappers ---
var _Gopurs__New = // TAST: (Func [(TypeVar s)] (ADT ["Effect","Effect"] [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _New(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ModifyImpl = // TAST: (Func [(Func [(TypeVar s)] (Record [state: (TypeVar s), value: (TypeVar b)])), (ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(TypeVar b)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := ModifyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_NewWithSelf = // TAST: (Func [(Func [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (TypeVar s))] (ADT ["Effect","Effect"] [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := NewWithSelf(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Read = // TAST: (Func [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(TypeVar s)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Read(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Write = // TAST: (Func [(TypeVar s), (ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Write(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})