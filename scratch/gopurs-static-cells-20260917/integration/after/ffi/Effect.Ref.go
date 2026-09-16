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

func Modify_(f gopurs_runtime.Value, ref gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		state := ref.AnyVal().(*RefState)
		state.mu.Lock()
		defer state.mu.Unlock()
		
		state.val = gopurs_runtime.Apply(f, state.val)
		return gopurs_runtime.Any(nil)
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
