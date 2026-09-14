from pathlib import Path
import hashlib
import json
import re

HERE = Path(__file__).resolve().parent
STAGES = ['naive', 'pushdown', 'precise', 'fusion', 'specialized_drop', 'reuse', 'retained_fields']
COUNTERS = ['new', 'dup', 'drop_last', 'drop_shared', 'unwrap_unique', 'unwrap_shared_fallback', 'reuse', 'scalar_update', 'reset']
SCENARIOS = ['unique', 'shared_root', 'shared_child', 'weak_root', 'mixed', 'weak_child']
result = {'counter_names': COUNTERS, 'scenarios': SCENARIOS, 'stages': {}}
for stage in STAGES:
    rows = {}
    for line in (HERE / f'build/{stage}.log').read_text().splitlines():
        if line.startswith('CASE '):
            _, name, mode, counters = line.split(' ', 3)
            rows.setdefault(name, {})[SCENARIOS[int(mode)]] = dict(zip(COUNTERS, json.loads(counters)))
        elif line.startswith('PEAK '):
            _, name, mode, peak = line.split()
            rows[name][SCENARIOS[int(mode)]]['peak_live_payloads_including_input'] = int(peak)
    result['stages'][stage] = rows
(HERE / 'summary.json').write_text(json.dumps(result, indent=2)+'\n')
files = [HERE/'prototype.py', HERE/'runtime.rs', HERE/'source-ir.json', HERE/'ir-checks.json']
files += [HERE/f'{stage}.ir.json' for stage in STAGES]
files += [HERE/f'build/{stage}.rs' for stage in STAGES]
files += [HERE/f'build/{stage}-native' for stage in STAGES]
metadata = {'files_sha256': {str(p.relative_to(HERE)):hashlib.sha256(p.read_bytes()).hexdigest() for p in files},
            'representation': 'native Rc; counted wrapper is one word and delegates ownership to native Rc; all 7 stages additionally compiled with --cfg untracked as direct std::rc::Rc',
            'validation_cases': 7 * 96 * 2, 'timing_executed_by_this_task': False}
(HERE/'metadata.json').write_text(json.dumps(metadata, indent=2)+'\n')
print('summarized 7 stages x 96 cases; recorded generated Rust and native binary hashes')
