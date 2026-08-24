import re

with open('../phpurs/phpurs/src/Phpurs/CodeGen.purs', 'r') as f:
    content = f.read()

content = content.replace('false false a.nextId expr', 'false a.nextId expr')
content = content.replace('false false a.nextId argExpr', 'false a.nextId argExpr')
content = content.replace('false false acc.nextId flatFn', 'false acc.nextId flatFn')
content = content.replace('false false resFn.nextId innerBody', 'false resFn.nextId innerBody')
content = content.replace('false false nextId e', 'false nextId e')
content = content.replace('false false nextId e1', 'false nextId e1')
content = content.replace('false false res1.nextId e2', 'false res1.nextId e2')
content = content.replace('false false a.nextId', 'false a.nextId')
content = content.replace('false false acc.nextId', 'false acc.nextId')
content = content.replace('false false resFn.nextId', 'false resFn.nextId')
content = content.replace('false false nextId', 'false nextId')
content = content.replace('false false res1.nextId', 'false res1.nextId')
content = content.replace('false false resCond.nextId', 'false resCond.nextId')
content = content.replace('false false resVal.nextId', 'false resVal.nextId')
content = content.replace('true false 0 fn.body', 'true 0 fn.body')
content = content.replace('false false 0 expr', 'false 0 expr')
content = content.replace('false false 0 fn.body', 'false 0 fn.body')

with open('../phpurs/phpurs/src/Phpurs/CodeGen.purs', 'w') as f:
    f.write(content)
