#!/usr/bin/env python3
"""Measure three archived variants in three rotated process rounds."""
from pathlib import Path
import datetime, hashlib, importlib.util, json, os, statistics, subprocess, sys

sys.dont_write_bytecode = True
base = Path(__file__).resolve().parent
logs = base / 'suite-pair-measurements'
logs.mkdir()
validator_path = Path('/Users/0x1/Documents/htdocs/altbak.pub/bin/benchmark/validate.py')
spec = importlib.util.spec_from_file_location('validator', validator_path)
validator = importlib.util.module_from_spec(spec)
spec.loader.exec_module(validator)
env = dict(os.environ, GOGC='800', GOWORK='off')
env.pop('PPROF', None)
sha = lambda p: hashlib.sha256(p.read_bytes()).hexdigest()
plan = [('before', 'pure'), ('integrated', 'pure')]
runs = []
for repetition in range(1, 6):
    offset = (repetition - 1) % 2
    for key, mode in plan[offset:] + plan[:offset]:
        binary = base / key / 'benchmark'
        digest = sha(binary)
        started = datetime.datetime.now(datetime.timezone.utc).isoformat()
        result = subprocess.run([str(binary)], cwd=binary.parent, env=env,
                                capture_output=True, text=True, timeout=60)
        (logs / f'{key}-{repetition}.stdout.log').write_text(result.stdout)
        (logs / f'{key}-{repetition}.stderr.log').write_text(result.stderr)
        assert result.returncode == 0, result.stderr
        assert sha(binary) == digest
        validated = validator.validate_output(result.stdout, mode)
        record = dict(key=key, repetition=repetition, started_utc=started,
                      binary_sha256=digest, validated=validated)
        runs.append(record)
        (logs / f'{key}-{repetition}.json').write_text(json.dumps(record, indent=2) + '\n')
        print(key, repetition, validated['sum_displayed_lines_ms'], flush=True)
summary = {}
for key, mode in plan:
    records = [r['validated'] for r in runs if r['key'] == key]
    medians = [statistics.median(c) for c in zip(*(r['times_us'] for r in records))]
    summary[key] = dict(times_us=medians, sum_row_medians_ms=sum(medians)/1000,
                        church_us=medians[6], rbtree_us=medians[8],
                        process_totals_ms=[r['sum_displayed_lines_ms'] for r in records])
report = dict(protocol='5 alternating pairs; 3 global warmups, 3/test, best of 10; row medians; 14 numeric results/process validated; PGO off',
              environment={k: env.get(k) for k in ['GOGC', 'GOWORK', 'GOMAXPROCS', 'GOMEMLIMIT']},
              validator_sha256=sha(validator_path), runs=runs, summary=summary)
(base / 'suite-pair-results.json').write_text(json.dumps(report, indent=2) + '\n')
print(json.dumps(summary, indent=2))
