import re

with open('src/Test/BenchCheck.purs', 'r') as file:
    content = file.read()

content = content.replace('import Bench (benchNow)', 'import Bench (benchNow_)')
content = content.replace('benchNow', 'benchNow_ unit')
content = content.replace('benchNow_ unit_', 'benchNow_') # fix double replacement

with open('src/Test/BenchCheck.purs', 'w') as file:
    file.write(content)
