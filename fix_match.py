import re

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Replace resE.exprType /= TypeValue with a check that also excludes TypeGenericParam
pattern = r'resE\.exprType \/=\s*TypeValue'
replacement = r'(case resE.exprType of\n                                    TypeValue -> false\n                                    TypeGenericParam _ -> false\n                                    _ -> true)'

content = re.sub(pattern, replacement, content)

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
