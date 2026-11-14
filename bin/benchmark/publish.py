#!/usr/bin/env python3
"""Publish complete validated campaigns, preserving the user's column headers."""
import argparse
import hashlib
import json
import math
from pathlib import Path
import re
import statistics
import sys

sys.dont_write_bytecode = True
from validate import CASES, EXTENDED_CASES, expected_cases
from measure import fingerprints

ROOT = Path(__file__).resolve().parents[2]
GROUPS = {
    'JavaScript': ['js-pure', 'es-pure', 'js-ffi', 'js-fficc'],
    'Go': ['go-pure', 'psgo-pure', 'go-ffi', 'go-fficc'],
    'Scheme': ['scm-pure', 'scm-ffi', 'scm-fficc'],
    'Erlang': ['erl-pure', 'erl-ffi', 'erl-fficc'],
    'PHP': ['php-pure', 'php-ffi', 'php-fficc'],
    'Rust': ['purust-numeric', 'fable-sharpurs', 'fable-native', 'rust-ffi', 'rust-fficc'],
    'C++': ['cpp-pure', 'cpp-ffi', 'cpp-fficc'],
    'F#/C#': ['sharp-pure', 'sharp-ffi', 'sharp-fficc'],
    'Java': ['java-pure', 'java-ffi', 'java-fficc'],
    'Koka': ['koka'], 'Haskell': ['haskell'], 'OCaml': ['ocaml'], 'C (reference)': ['c'],
    'Extended Results': ['js-x', 'es-x', 'go-x', 'rust-x'],
}
ROWS = ['AST Evaluation', 'Fibonacci', 'List Processing', 'Tail Call Optimization',
        'Deep Record Updates', 'Ackermann', 'Church Numerals', 'Prime Sieve',
        'Red-Black Tree', 'Polymorphism', 'State Monad', 'Lazy Evaluation',
        'Array Processing', 'RowToList']
EXTENDED_ROWS = ['File I/O', 'STArray Operations', 'String Operations', 'Aff Operations', 'Parallelism']
WIP_GROUPS = {
    'JavaScript': ['js-pure', 'es-pure', 'js-ffi', 'js-fficc'],
    'Go': ['go-pure', 'go-dev', 'psgo-pure', 'go-ffi', 'go-fficc'],
}


def row_label(cell):
    """Read the benchmark name while keeping source links in the table."""
    return re.sub(r'^\[([^\]]+)\]\([^)]+\)', r'\1', cell.strip())


def format_number(value):
    """Limit display precision without presenting a positive value as zero."""
    return '< 0.01' if 0 < value < 0.005 else f'{value:.2f}'


def format_time(value, unit='μs'):
    number = format_number(value)
    return ('' if number.startswith('<') else '~ ') + number + ' ' + unit


def format_wip_cell(value, c_reference, js_reference):
    def ratio(reference):
        return '—' if reference is None or reference <= 0 else format_number(value / reference) + 'x'
    cell = format_time(value) + '<br>(/C = ' + ratio(c_reference) + ')'
    js_ratio = ratio(js_reference)
    return cell if js_ratio == '1.00x' else cell + '<br>(/JS = ' + js_ratio + ')'


def reference_times(text, title):
    """Use the same-row first column of the published reference table."""
    match = re.search(r'#### ' + re.escape(title) + r'\n(.*?)(?=\n#### |\n### |\n## |\Z)', text, re.S)
    result = {}
    if match:
        for line in match[1].splitlines():
            cells = line.split('|')
            timing = re.match(r'\s*~\s+([\d.]+) μs', cells[1]) if len(cells) > 1 else None
            if timing:
                result[row_label(cells[0])] = float(timing[1])
    return result


def verify(key, result):
    extended = key.endswith('-x')
    wip = result['mode'] == 'wip'
    expected = ([case['value'] for case in expected_cases('wip')] if wip else
                [case['value'] for case in EXTENDED_CASES] if extended else list(CASES.values()))
    expected = [str(v) for v in expected]
    if not result.get('values_validated') or result['values'] != expected:
        raise ValueError('Missing or incorrect oracles: ' + key)
    samples = result.get('process_runs', [])
    if len(samples) != 3:
        raise ValueError('Expected three processes: ' + key)
    if any(not s.get('values_validated') or s['values'] != expected for s in samples):
        raise ValueError('Invalid per-process outputs: ' + key)
    if any(s.get('labels') != result.get('labels') for s in samples):
        raise ValueError('Inconsistent per-process labels: ' + key)
    if result.get('labels') != [case['label'] for case in expected_cases(result['mode'])]:
        raise ValueError('Incorrect benchmark labels: ' + key)
    if len(result['times_us']) != len(expected):
        raise ValueError('Incomplete column: ' + key)
    for i, value in enumerate(result['times_us']):
        if not math.isfinite(value) or value <= 0:
            raise ValueError('Invalid timing: ' + key)
        median = statistics.median(s['times_us'][i] for s in samples)
        if not math.isclose(value, median, rel_tol=1e-12, abs_tol=1e-12):
            raise ValueError('Cell is not the median of recorded processes: ' + key)
    if wip:
        if result['total_ms'] is not None:
            raise ValueError('WIP cases do not have an aggregate score: ' + key)
    elif not math.isclose(result['total_ms'], sum(result['times_us']) / 1000, rel_tol=1e-12, abs_tol=1e-12):
        raise ValueError('Inconsistent total: ' + key)


def render(original, results):
    text = original
    for title, keys in GROUPS.items():
        pattern = r'(#### ' + re.escape(title) + r'\n)(.*?)(?=\n#### |\n### |\n## |\n> |\Z)'
        matches = list(re.finditer(pattern, text, re.S))
        if len(matches) != 1:
            raise ValueError('Missing/ambiguous section: ' + title)
        match = matches[0]
        section = match[2]
        expected_rows = EXTENDED_ROWS if title == 'Extended Results' else ROWS
        lines = section.splitlines()
        seen = []
        for i, line in enumerate(lines):
            if '|' not in line:
                continue
            name_cell = line.split('|', 1)[0].strip()
            label = row_label(name_cell)
            if label not in expected_rows and label != '**Total Execution Time**':
                continue
            total = label == '**Total Execution Time**'
            values = []
            for key in keys:
                result = results[key]
                value = result['total_ms'] if total else result['times_us'][expected_rows.index(label)]
                values.append(format_time(value, 'ms' if total else 'μs'))
            lines[i] = name_cell + ' | ' + ' | '.join(values) + ' |'
            seen.append(label)
        if seen != expected_rows + ['**Total Execution Time**']:
            raise ValueError('Unexpected table row order: ' + title)
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return text


def render_wip(original, results):
    """Update only the 23 WIP rows in the nine JS/Go columns."""
    text = original
    labels = [case['label'] for case in expected_cases('wip')]
    c_reference = reference_times(original, 'C (reference)')
    js_reference = dict(zip(results['js-pure']['library']['labels'],
                            results['js-pure']['library']['times_us']))
    for title, keys in WIP_GROUPS.items():
        pattern = r'(#### ' + re.escape(title) + r'\n)(.*?)(?=\n#### |\n### |\n## |\Z)'
        matches = list(re.finditer(pattern, text, re.S))
        if len(matches) != 1:
            raise ValueError('Missing/ambiguous section: ' + title)
        match = matches[0]
        lines = match[2].splitlines()
        seen = []
        for index, line in enumerate(lines):
            name_cell = line.split('|', 1)[0].strip()
            label = row_label(name_cell)
            if label not in labels:
                continue
            cells = line.split('|')
            if len(cells) != len(keys) + 2 or cells[-1].strip():
                raise ValueError('Unexpected column count for ' + title + ': ' + label)
            values = [results[key]['library']['times_us'][labels.index(label)] for key in keys]
            lines[index] = name_cell + ' | ' + ' | '.join(
                format_wip_cell(value, c_reference.get(label), js_reference[label])
                for value in values) + ' |'
            seen.append(label)
        if seen != labels:
            raise ValueError('Unexpected WIP row order: ' + title)
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return text


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--campaign', type=Path, action='append', required=True)
    parser.add_argument('--readme', type=Path, default=ROOT / 'README.md')
    parser.add_argument('--wip-only', action='store_true', help='Publish the 23 new rows in the nine JS/Go columns')
    args = parser.parse_args()
    results = {}
    source_snapshot = None
    protocol = None
    for folder in args.campaign:
        manifest = json.loads((folder / 'manifest.json').read_text())
        if not manifest.get('completed_at_utc'):
            raise ValueError('Incomplete campaign: ' + str(folder))
        if source_snapshot is None:
            source_snapshot, protocol = manifest['sources'], manifest['protocol']
        elif manifest['sources'] != source_snapshot or manifest['protocol'] != protocol:
            raise ValueError('Campaigns have different sources or timing protocols')
        result_file = folder / 'results.json'
        if hashlib.sha256(result_file.read_bytes()).hexdigest() != manifest.get('results_sha256'):
            raise ValueError('Results differ from completed campaign manifest: ' + str(folder))
        incoming = json.loads(result_file.read_text())
        if results.keys() & incoming.keys():
            raise ValueError('Duplicate column in campaigns')
        results.update(incoming)
    if fingerprints() != source_snapshot:
        raise ValueError('Current sources differ from the completed campaigns')
    groups = WIP_GROUPS if args.wip_only else GROUPS
    for keys in groups.values():
        for key in keys:
            verify(key, results[key])
            if args.wip_only:
                verify(key, results[key]['library'])
    original = args.readme.read_text()
    updated = render_wip(original, results) if args.wip_only else render(original, results)
    if args.readme.read_text() != original:
        raise RuntimeError('README changed while preparing publication')
    args.readme.write_text(updated)
    print('Published 23 WIP rows across nine JS/Go columns; core cells preserved.' if args.wip_only
          else 'Published 39 complete validated columns; existing headers preserved.')


if __name__ == '__main__':
    main()
