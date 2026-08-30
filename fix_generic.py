import re

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

pattern = r'exprTypeToGoType \(\s*unsafePerformEffect\s*\(\s*Ref\.read\s+helpersRef\s*\)\s*\)\.pointerAdtPaths \(\s*unsafePerformEffect\s*\(\s*Ref\.read\s+helpersRef\s*\)\s*\)\.enumAdts \(\s*unsafePerformEffect\s*\(\s*Ref\.read\s+helpersRef\s*\)\s*\)\.elidedCtors modNameStr'
replacement = r'exprTypeToGenericGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors (unsafePerformEffect (Ref.read helpersRef)).globalTypeVars modNameStr'

content = re.sub(pattern, replacement, content)

# Also fix `h.pointerAdtPaths`
pattern2 = r'exprTypeToGoType h\.pointerAdtPaths h\.enumAdts h\.elidedCtors modNameStr'
replacement2 = r'exprTypeToGenericGoType h.pointerAdtPaths h.enumAdts h.elidedCtors h.globalTypeVars modNameStr'
content = re.sub(pattern2, replacement2, content)

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
