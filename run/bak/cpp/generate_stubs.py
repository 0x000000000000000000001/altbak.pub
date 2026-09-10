import sys
import re
import collections

if len(sys.argv) < 3:
    print("Usage: generate_stubs.py <make.log> <output.cc>")
    sys.exit(1)

log_file = sys.argv[1]
out_file = sys.argv[2]

with open(log_file, 'r') as f:
    text = f.read()

# Pattern for undefined reference (Apple Clang)
# e.g.,   "Effect_Console::log(purescript::boxed const&)", referenced from:
pattern = r'^\s*"([a-zA-Z0-9_]+)::([a-zA-Z0-9_]+)\((.*?)\)", referenced from:'
matches = re.findall(pattern, text, re.MULTILINE)

namespaces = collections.defaultdict(list)
for ns, func, args in set(matches):
    if ns == "purescript" or ns == "std":
        continue
    namespaces[ns].append((func, args))

with open(out_file, 'w') as out:
    out.write('#include "purescript.h"\n\n')
    for ns, funcs in namespaces.items():
        out.write(f"namespace {ns} {{\n")
        for func, args in funcs:
            # We don't name the arg so it avoids unused warning, just signature
            if args:
                out.write(f"    auto {func}({args}) -> purescript::boxed {{\n")
            else:
                out.write(f"    auto {func}() -> purescript::boxed {{\n")
            out.write(f'        return purescript::boxed();\n')
            out.write(f"    }}\n")
        out.write(f"}}\n\n")

print(f"Generated stubs for {len(namespaces)} namespaces.")
