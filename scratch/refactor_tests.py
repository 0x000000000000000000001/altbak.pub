import os
import glob
import re

files = glob.glob('src/Test/*.purs')

for f in files:
    with open(f, 'r') as file:
        content = file.read()
    
    # describe
    content = re.sub(r'describe :: Effect Unit', r'describe :: Unit -> Effect Unit', content)
    content = re.sub(r'^describe =', r'describe _ =', content, flags=re.MULTILINE)
    
    # act
    content = re.sub(r'act :: Effect Unit', r'act :: Unit -> Effect Unit', content)
    content = re.sub(r'^act =', r'act _ =', content, flags=re.MULTILINE)
    
    with open(f, 'w') as file:
        file.write(content)
