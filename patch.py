import re

with open('run/bak/go/output/purescript/Data_Map_Internal.go', 'r') as f:
    content = f.read()

# Replace gopurs_runtime.Apply(go__go_3_0_17, arg) with go__go_3_0_17_native(arg)
content = re.sub(r'gopurs_runtime\.Apply\((go__go_3_0_\d+), ([^\)]+)\)', r'\1_native(\2)', content)

# Define native func
# go__go_3_0_17 = gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
def repl_func(m):
    go_name = m.group(1)
    arg_name = m.group(2)
    return f"""var {go_name}_native func({arg_name} gopurs_runtime.Value) gopurs_runtime.Value
{go_name}_native = func({arg_name} gopurs_runtime.Value) gopurs_runtime.Value {{
"""

content = re.sub(r'(go__go_3_0_\d+) = gopurs_runtime\.Func\(func\((v1_\d+)_val gopurs_runtime\.Value\) gopurs_runtime\.Value \{', repl_func, content)
content = re.sub(r'(go__go_3_0_\d+) = gopurs_runtime\.Func\(func\((v1_\d+) gopurs_runtime\.Value\) gopurs_runtime\.Value \{', repl_func, content)

# But wait, we also need to wrap it back for the top level assignment if needed?
# Actually insert doesn't return the closure, it calls it: gopurs_runtime.Apply(go_name, arg)
# Wait, insert returns the result of the closure, not the closure itself.
# Let's check what insert returns.

with open('run/bak/go/output/purescript/Data_Map_Internal.go', 'w') as f:
    f.write(content)

