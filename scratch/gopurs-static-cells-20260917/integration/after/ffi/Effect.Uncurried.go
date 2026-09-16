package Effect_Uncurried

import "gopurs/output/gopurs_runtime"

func MkEffectFn1(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), gopurs_runtime.Value{})
	})
}

func MkEffectFn2(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func2(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), gopurs_runtime.Value{})
	})
}

func MkEffectFn3(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func3(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), gopurs_runtime.Value{})
	})
}

func MkEffectFn4(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func4(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), gopurs_runtime.Value{})
	})
}

func MkEffectFn5(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func5(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), gopurs_runtime.Value{})
	})
}

func MkEffectFn6(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func6(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), gopurs_runtime.Value{})
	})
}

func MkEffectFn7(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func7(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), gopurs_runtime.Value{})
	})
}

func MkEffectFn8(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func8(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), a8), gopurs_runtime.Value{})
	})
}

func MkEffectFn9(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func9(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), a8), a9), gopurs_runtime.Value{})
	})
}

func MkEffectFn10(f gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func10(func(a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value, a10 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f, a1), a2), a3), a4), a5), a6), a7), a8), a9), a10), gopurs_runtime.Value{})
	})
}

func RunEffectFn1(f gopurs_runtime.Value, a1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(f, a1)
	})
}

func RunEffectFn2(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply2(f, a1, a2)
	})
}

func RunEffectFn3(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply3(f, a1, a2, a3)
	})
}

func RunEffectFn4(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply4(f, a1, a2, a3, a4)
	})
}

func RunEffectFn5(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply5(f, a1, a2, a3, a4, a5)
	})
}

func RunEffectFn6(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply6(f, a1, a2, a3, a4, a5, a6)
	})
}

func RunEffectFn7(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply7(f, a1, a2, a3, a4, a5, a6, a7)
	})
}

func RunEffectFn8(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply8(f, a1, a2, a3, a4, a5, a6, a7, a8)
	})
}

func RunEffectFn9(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply9(f, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	})
}

func RunEffectFn10(f gopurs_runtime.Value, a1 gopurs_runtime.Value, a2 gopurs_runtime.Value, a3 gopurs_runtime.Value, a4 gopurs_runtime.Value, a5 gopurs_runtime.Value, a6 gopurs_runtime.Value, a7 gopurs_runtime.Value, a8 gopurs_runtime.Value, a9 gopurs_runtime.Value, a10 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply10(f, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	})
}

