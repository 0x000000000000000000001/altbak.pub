#!/usr/bin/env python3
"""Prepare validation against exactly the requested compiled output."""
from pathlib import Path
import hashlib, re
HERE = Path(__file__).resolve().parent
SOURCE = Path('/Users/0x1/Documents/htdocs/altbak.pub-purust/output/purust_output/Purs_Test_RBTree/src/lib.rs')
source = SOURCE.read_text()
assert hashlib.sha256(source.encode()).hexdigest() == 'f5479b1b43a8a967665e4aec683d62a5c2c12d1db4d57c65f2a7c732a128964a'
# Retain dependency-free declarations/functions verbatim, with only module path adaptation.
start = source.index('#[derive(Clone, Copy)]')
source = source[start:]
functions = []
prefix = source[:source.index('fn Test_RBTree___purust_rebuild_E')]
for match in re.finditer(r'(?m)^(?:pub )?fn (Test_RBTree_\w+)\(', source):
    name = match.group(1)
    if name in {'Test_RBTree_R','Test_RBTree_B','Test_RBTree_E','Test_RBTree_T','Test_RBTree_describe','Test_RBTree_act'}: continue
    begin = match.start()
    brace = source.index('{', match.end())
    level = 1
    end = brace + 1
    while level:
        if source[end] == '{': level += 1
        elif source[end] == '}': level -= 1
        end += 1
    # Preserve inline attribute immediately preceding this function.
    if source[max(0,begin-10):begin] == '#[inline]\n': begin -= 10
    functions.append(source[begin:end])
generated = prefix + '\n'.join(functions)
generated = generated.replace('crate::', 'crate::generated::')
(HERE/'generated.rs').write_text(generated)
print(f'source_sha256={hashlib.sha256(SOURCE.read_bytes()).hexdigest()}')
print(f'functions={len(functions)}')
