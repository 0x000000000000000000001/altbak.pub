import re

with open('run/bak/go/output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# First replace the apply calls for 73
content = content.replace("gopurs_runtime.Apply(go__go_3_0_73, ", "go__go_3_0_73_native(")
# And for 17
content = content.replace("gopurs_runtime.Apply(go__go_3_0_17, ", "go__go_3_0_17_native(")

# Now replace the function declarations
def repl_func(m):
    go_name = m.group(1)
    arg_name = m.group(2)
    return f"""var {go_name}_native func({arg_name} gopurs_runtime.Value) gopurs_runtime.Value
{go_name}_native = func({arg_name} gopurs_runtime.Value) gopurs_runtime.Value {{
"""

content = re.sub(r'var (go__go_3_0_\d+) gopurs_runtime\.Value\n_ = \1\n\1 = gopurs_runtime\.Func\(func\((v1_\d+) gopurs_runtime\.Value\) gopurs_runtime\.Value \{', r'var \1 gopurs_runtime.Value\n_ = \1\nvar \1_native func(\2 gopurs_runtime.Value) gopurs_runtime.Value\n\1_native = func(\2 gopurs_runtime.Value) gopurs_runtime.Value {', content)

# And we need to close the native func and wrap it back in a gopurs_runtime.Func!
# Wait, if we just define it as native, the original code returns it at the end of Call_...
# Let's find "return go__go_3_0_73" and inject the wrapper before it!
content = content.replace("return go__go_3_0_73\n}", "go__go_3_0_73 = gopurs_runtime.Func(go__go_3_0_73_native)\nreturn go__go_3_0_73\n}")
content = content.replace("return go__go_3_0_17\n}", "go__go_3_0_17 = gopurs_runtime.Func(go__go_3_0_17_native)\nreturn go__go_3_0_17\n}")

with open('run/bak/go/output/purescript/Data_Map_Internal.go', 'w') as f:
    f.write(content)
