package purescript



import (
	"sync"
	"gopurs/output/gopurs_runtime"
)

type RefState struct {
	mu  sync.Mutex
	val gopurs_runtime.Value
}

func Effect_Ref__New(val gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Any(&RefState{val: val})
	})
}

func Effect_Ref_NewWithSelf(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		ref := &RefState{}
		ref.val = gopurs_runtime.Apply(f, gopurs_runtime.Any(ref))
		return gopurs_runtime.Any(ref)
	})
}

func Effect_Ref_Read(ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		state := ref.AnyVal().(*RefState)
		state.mu.Lock()
		defer state.mu.Unlock()
		return state.val
	})
}

func Effect_Ref_ModifyImpl(f gopurs_runtime.Value, ref gopurs_runtime.Value) gopurs_runtime.Value {
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

func Effect_Ref_Modify_(f gopurs_runtime.Value, ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		state := ref.AnyVal().(*RefState)
		state.mu.Lock()
		defer state.mu.Unlock()
		
		state.val = gopurs_runtime.Apply(f, state.val)
		return gopurs_runtime.Any(nil)
	})
}

func Effect_Ref_Write(val gopurs_runtime.Value, ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		state := ref.AnyVal().(*RefState)
		state.mu.Lock()
		state.val = val
		state.mu.Unlock()
		return gopurs_runtime.Any(nil)
	})
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Effect_Ref__New = // TAST: (ForAll [s] (Func [(TypeVar s)] (ADT ["Effect","Effect"] [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Ref__New(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Ref_ModifyImpl = // TAST: (ForAll [s, b] (Func [(Func [(TypeVar s)] (Record (Row [state: (TypeVar s), value: (TypeVar b)] Empty))), (ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(TypeVar b)])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Effect_Ref_ModifyImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Ref_Modify_ = // TAST: (ForAll [s] (Func [(Func [(TypeVar s)] (TypeVar s)), (ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [Unit])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Effect_Ref_Modify_(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Ref_NewWithSelf = // TAST: (ForAll [s] (Func [(Func [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (TypeVar s))] (ADT ["Effect","Effect"] [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Ref_NewWithSelf(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Ref_Read = // TAST: (ForAll [s] (Func [(ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [(TypeVar s)])))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Effect_Ref_Read(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Effect_Ref_Write = // TAST: (ForAll [s] (Func [(TypeVar s), (ADT ["Effect","Ref","Ref"] [(TypeVar s)])] (ADT ["Effect","Effect"] [Unit])))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := Effect_Ref_Write(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})