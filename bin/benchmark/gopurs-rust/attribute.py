"""Attribute sampled collection/copy costs to their nearest application caller."""
import collections
import json
from pathlib import Path
import re
import sys

source = Path(sys.argv[1])
stack = []
families = {
    'collections': r'Purs_Data_Map_Internal::|Purs_Data_Set::',
    'qualified comparison': r'CoreFn_compareModuleNames|CoreFn_compareQualifiedIdent',
    'Value copy': r'purust_core..Value.*Clone.*clone',
    'String copy': r'5alloc6string.*5clone',
}
callers = {name: collections.Counter() for name in families}
operations = collections.Counter()

def finish():
    _, count, name, children = stack.pop()
    count -= children
    if count <= 0 or re.search(r'__ulock_wait|__psynch_cvwait|kevent|mach_msg.*trap|thread_switch', name):
        return
    frames = [frame[2] for frame in stack] + [name]
    for label, pattern in families.items():
        first = next((i for i, frame in enumerate(frames) if re.search(pattern, frame)), None)
        if first is None:
            continue
        caller = next((frame for frame in reversed(frames[:first])
                       if re.search(r'Purs_Gopurs_|Purs_PureScript_Backend_Optimizer_', frame)), '<runtime>')
        caller = re.sub(r'::h[a-f0-9]+$', '', caller)
        caller = caller.split('::_$u7b$$u7b$closure')[0]
        callers[label][caller] += count
        if label == 'collections':
            operation = re.sub(r'::h[a-f0-9]+$', '', frames[first]).split('::_$u7b$$u7b$closure')[0]
            operations[caller + ' -> ' + operation] += count

for line in source.read_text().splitlines():
    if line.startswith('Total number in stack'):
        break
    match = re.match(r'^([ +!:|]*)(\d+) (.+)$', line)
    if not match:
        continue
    prefix, count, name = match.groups()
    depth, count, name = len(prefix), int(count), name.split('  (in ')[0]
    while stack and stack[-1][0] >= depth:
        finish()
    if stack:
        stack[-1][3] += count
    stack.append([depth, count, name, 0])
while stack:
    finish()
print(json.dumps({**{name: values.most_common(30) for name, values in callers.items()},
                  'collection operations': operations.most_common(50)}, indent=2))
