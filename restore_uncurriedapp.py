with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    lines = f.read().split('\n')

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs.orig', 'r') as f:
    orig_content = f.read()

# Extract UncurriedApp block from orig
start_marker = "      UncurriedApp fn args ->"
end_marker = "      UncurriedAbs args body -> liftIfNeeded \\_ ->"
orig_start = orig_content.find(start_marker)
orig_end = orig_content.find(end_marker)

uncurried_app_block = orig_content[orig_start:orig_end]

# Now let's find where to insert it in CodeGen.purs
# We just need to find "UncurriedAbs args body -> liftIfNeeded \_ ->" in CodeGen.purs and insert it BEFORE it.

insert_idx = -1
for i, line in enumerate(lines):
    if "UncurriedAbs args body -> liftIfNeeded \\_ ->" in line:
        insert_idx = i
        break

if insert_idx != -1:
    new_content = '\n'.join(lines[:insert_idx]) + '\n' + uncurried_app_block + '\n'.join(lines[insert_idx:])
    with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
        f.write(new_content)
    print("Restored UncurriedApp block")
else:
    print("Failed to find insert location")
