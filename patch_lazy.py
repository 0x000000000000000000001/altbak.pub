with open('run/bak/go/output/purescript/Data_Lazy_FFI.go', 'r') as f:
    content = f.read()

content = content.replace('"sync"\n', '')

patch_lazy = """func Data_Lazy_Defer(thunk gopurs_runtime.Value) gopurs_runtime.Value {
	var evaluated bool
	var result gopurs_runtime.Value

	return gopurs_runtime.Func(func(_dollar__unused gopurs_runtime.Value) gopurs_runtime.Value {
		if !evaluated {
			result = gopurs_runtime.Apply(thunk, gopurs_runtime.Value{})
			evaluated = true
		}
		return result
	})
}"""

import re
content = re.sub(r'func Data_Lazy_Defer.*?return result\n\t}\)\n}', patch_lazy, content, flags=re.DOTALL)

with open('run/bak/go/output/purescript/Data_Lazy_FFI.go', 'w') as f:
    f.write(content)
