with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs.orig', 'r') as f:
    orig_lines = f.read().split('\n')

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    cur_lines = f.read().split('\n')

# Find start and end of App block in orig
orig_app_start = -1
for i, line in enumerate(orig_lines):
    if "mbIntrinsic = case getVar (unwrapTcoExpr flatFn) of" in line:
        orig_app_start = i - 1
        break

orig_app_end = -1
for i in range(orig_app_start, len(orig_lines)):
    if "UncurriedApp fn args ->" in orig_lines[i]:
        orig_app_end = i - 1
        break

orig_app_block = orig_lines[orig_app_start:orig_app_end]

# Find start and end in cur
cur_app_start = -1
for i, line in enumerate(cur_lines):
    if "mbIntrinsic = case getVar (unwrapTcoExpr flatFn) of" in line:
        cur_app_start = i - 1
        break

cur_app_end = -1
for i in range(cur_app_start, len(cur_lines)):
    if "UncurriedApp fn args ->" in cur_lines[i]:
        cur_app_end = i - 1
        break

if cur_app_start != -1 and cur_app_end != -1:
    new_lines = cur_lines[:cur_app_start] + orig_app_block + cur_lines[cur_app_end+1:]
    with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
        f.write('\n'.join(new_lines))
    print("Restored App block")
else:
    print("Could not find App block in cur")
