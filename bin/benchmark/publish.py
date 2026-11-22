#!/usr/bin/env python3
"""Publish complete validated campaigns, preserving the user's column headers."""
import argparse
import hashlib
from html import escape
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
    'Go': ['go-pure', 'go-dev', 'psgo-pure', 'go-ffi', 'go-fficc'],
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
WIP_TOTAL_LABEL = '**Excluded WIP Total**'
GRAND_TOTAL_LABEL = '**Total incl. Excluded WIP**'
WIP_CELL = re.compile(r'(?:< )?([\d.]+) μs')
TOTAL_CELL = re.compile(r'(?:< )?([\d.]+) ms')


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


NEW_LAYOUT_GROUPS = tuple(title for title in GROUPS if title != 'Extended Results')
CELL_RATIO = re.compile(r'([\d.]+)x')
CELL_MS = re.compile(r'([\d.]+) ms')
CELL_US = re.compile(r'([\d.]+) μs')


def table_cells(line):
    return [cell.strip() for cell in line.split('|')[1:-1]]


def find_benchmark_row(lines, label):
    prefix = '| [' + label + ']('
    for i, line in enumerate(lines):
        if line.startswith(prefix):
            return i
    raise ValueError('Missing benchmark row: ' + label)


def find_table_row(lines, label):
    prefix = '| ' + label + ' |'
    return next(i for i, line in enumerate(lines) if line.startswith(prefix))


def metric_row(lines, start, label):
    for i in range(start + 1, len(lines)):
        if lines[i].startswith('| ' + label + ' |'):
            return i
        if lines[i].startswith('| **') or lines[i].startswith('| ['):
            break
    raise ValueError('Missing ' + label + ' row')


def cell_ratio(cell):
    match = CELL_RATIO.search(cell)
    return None if match is None else float(match.group(1))


def cell_ms(cell):
    match = CELL_MS.search(cell)
    return None if match is None else float(match.group(1))


def cell_us(cell):
    match = CELL_US.search(cell)
    return None if match is None else float(match.group(1))


def new_layout(lines):
    return any(line.startswith('| **Stable** |') for line in lines)


def section_lines(text, title):
    pattern = r'(#### ' + re.escape(title) + r'\n)(.*?)(?=\n#### |\n### |\n## |\Z)'
    match = re.search(pattern, text, re.S)
    if match is None:
        return None
    return match, match[2].splitlines()


def bold_value(value):
    value = value.strip()
    if not value or value == '—':
        return value
    value = re.sub(r'\*\*((?:~ |< )?[\d.]+(?: (?:μs|ms|s)|x))\*\*', r'\1', value)
    if re.fullmatch(r'(?:~ |< )?[\d.]+ (?:μs|ms|s)', value) or re.fullmatch(r'[\d.]+x', value):
        return '<strong>' + escape(value) + '</strong>'
    value = re.sub(r'^((?:~ |< )?[\d.]+ (?:μs|ms|s))',
                   lambda match: '<strong>' + escape(match[1]) + '</strong>', value)
    return re.sub(r'(/(?:C|JS) = )([\d.]+x)', r'\1<strong>\2</strong>', value)


def bold_total_line(line):
    parts = line.split('|')
    start = 2 if line.startswith('|') else 1
    end = len(parts) - 1 if line.rstrip().endswith('|') else len(parts)
    for i in range(start, end):
        raw = parts[i]
        value = raw.strip()
        if not value or value == '—':
            continue
        bolded = bold_value(value)
        if bolded != value:
            parts[i] = raw.replace(value, bolded, 1)
    return '|'.join(parts)


def bold_totals(text):
    """Bold the numeric cells of total rows and their metric rows."""
    result = []
    in_total = False
    for line in text.splitlines():
        if line.startswith('| **Total'):
            result.append(bold_total_line(line))
            in_total = True
        elif line.startswith('**Total'):
            result.append(bold_total_line(line))
            in_total = False
        elif in_total and line.startswith('| ↳ '):
            result.append(bold_total_line(line))
        else:
            in_total = False
            result.append(line)
    return '\n'.join(result) + '\n'


def refresh_new_totals(text):
    """Recompute the Stable/Incl WIP aggregate rows from displayed cells."""
    js_match, js_lines = section_lines(text, 'JavaScript')
    if js_match is None or not new_layout(js_lines):
        return text
    js_stable = cell_ms(table_cells(js_lines[find_table_row(js_lines, '**Total Stable**')])[1])
    js_incl_rows = []
    js_block = None
    for js_index, js_line in enumerate(js_lines):
        if js_line.startswith('| **Incl WIP**'):
            js_block = 'incl'
        elif js_line.startswith('| **Excl WIP**'):
            js_block = 'excl'
        elif js_block == 'incl' and js_line.startswith('| ['):
            js_incl_rows.append(js_index)
    if len(js_incl_rows) != 23:
        raise ValueError('Unexpected Incl WIP composition in JavaScript')
    js_incl = sum(cell_us(table_cells(js_lines[i])[1]) for i in js_incl_rows) / 1000
    js_combined = js_stable + js_incl
    for title in NEW_LAYOUT_GROUPS:
        found = section_lines(text, title)
        if found is None:
            continue
        match, lines = found
        if not new_layout(lines):
            continue
        header = next(i for i, line in enumerate(lines) if line.strip())
        columns = len(table_cells(lines[header + 1]))
        block, benchmarks = None, []
        for i, line in enumerate(lines):
            if line.startswith('| **Stable**'):
                block = 'stable'
            elif line.startswith('| **Incl WIP**'):
                block = 'incl'
            elif line.startswith('| **Excl WIP**'):
                block = 'excl'
            elif line.startswith('| ['):
                label = line[len('| ['):line.index('](')]
                benchmarks.append((block, label, i))
        stable = [row for row in benchmarks if row[0] == 'stable']
        incl = [row for row in benchmarks if row[0] == 'incl']
        if len(stable) != 14 or len(incl) != 23:
            raise ValueError('Unexpected Stable/Incl WIP composition in ' + title)

        def ratios(rows, metric):
            result = []
            for _, _, index in rows:
                result.append([cell_ratio(cell)
                               for cell in table_cells(lines[metric_row(lines, index, metric)])[1:]])
            return result

        stable_c = ratios(stable, '↳ /C')
        stable_js = ratios(stable, '↳ /JS')
        incl_js = ratios(incl, '↳ /JS')
        all_js = stable_js + incl_js
        total_row = find_table_row(lines, '**Total Stable**')
        stable_ms = [cell_ms(cell) for cell in table_cells(lines[total_row])[1:]]

        def incl_total(i):
            times = [cell_us(table_cells(lines[index])[i + 1]) for _, _, index in incl]
            if any(value is None for value in times):
                return None
            return sum(times) / 1000

        incl_ms = [incl_total(i) for i in range(columns)]
        combined_ms = [None if incl_ms[i] is None else stable_ms[i] + incl_ms[i]
                       for i in range(columns)]

        def mean_cells(rows, available=None):
            cells = []
            for i in range(columns):
                if available is not None and not available[i]:
                    cells.append('—')
                    continue
                values = [row[i] for row in rows if row[i] is not None]
                cells.append('—' if not values else
                             format_number(sum(values) / len(values)) + 'x')
            return cells

        for label, rows in [('↳ μ/C', stable_c), ('↳ μ/JS', stable_js)]:
            lines[metric_row(lines, total_row, label)] = (
                '| ' + label + ' | ' + ' | '.join(mean_cells(rows)) + ' |')

        incl_row = find_table_row(lines, '**Total Incl WIP**')
        lines[incl_row] = ('| **Total Incl WIP** | '
                           + ' | '.join('—' if value is None else format_time(value, 'ms')
                                        for value in incl_ms) + ' |')
        lines[metric_row(lines, incl_row, '↳ /C')] = (
            '| ↳ /C | ' + ' | '.join('—' for _ in range(columns)) + ' |')
        lines[metric_row(lines, incl_row, '↳ /JS')] = (
            '| ↳ /JS | ' + ' | '.join(
                '—' if incl_ms[i] is None or (title == 'JavaScript' and i == 0)
                else format_number(incl_ms[i] / js_incl) + 'x'
                for i in range(columns)) + ' |')
        lines[metric_row(lines, incl_row, '↳ μ/JS')] = (
            '| ↳ μ/JS | ' + ' | '.join(mean_cells(incl_js, [v is not None for v in incl_ms])) + ' |')

        combined_row = find_table_row(lines, '**Total Stable + Incl WIP**')
        lines[combined_row] = ('| **Total Stable + Incl WIP** | '
                               + ' | '.join('—' if value is None else format_time(value, 'ms')
                                            for value in combined_ms) + ' |')
        lines[metric_row(lines, combined_row, '↳ /C')] = (
            '| ↳ /C | ' + ' | '.join('—' for _ in range(columns)) + ' |')
        lines[metric_row(lines, combined_row, '↳ /JS')] = (
            '| ↳ /JS | ' + ' | '.join(
                '—' if combined_ms[i] is None or (title == 'JavaScript' and i == 0)
                else format_number(combined_ms[i] / js_combined) + 'x'
                for i in range(columns)) + ' |')
        lines[metric_row(lines, combined_row, '↳ μ/JS')] = (
            '| ↳ μ/JS | ' + ' | '.join(mean_cells(all_js, [v is not None for v in combined_ms])) + ' |')

        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return bold_totals(text)


def render(original, results):
    text = original
    c_reference = dict(zip(ROWS, results['c']['times_us']))
    js_reference = dict(zip(ROWS, results['js-pure']['times_us']))
    for title, keys in GROUPS.items():
        pattern = r'(#### ' + re.escape(title) + r'\n)(.*?)(?=\n#### |\n### |\n## |\n> |\Z)'
        matches = list(re.finditer(pattern, text, re.S))
        if len(matches) != 1:
            raise ValueError('Missing/ambiguous section: ' + title)
        match = matches[0]
        lines = match[2].splitlines()
        if title in NEW_LAYOUT_GROUPS:
            for label in ROWS:
                index = find_benchmark_row(lines, label)
                if len(table_cells(lines[index])) != len(keys) + 1:
                    raise ValueError('Unexpected column count for ' + title + ': ' + label)
                values = [results[key]['times_us'][ROWS.index(label)] for key in keys]
                lines[index] = ('| ' + table_cells(lines[index])[0] + ' | '
                                + ' | '.join(format_time(value, 'μs') for value in values) + ' |')
                c_row = metric_row(lines, index, '↳ /C')
                c_ref = c_reference.get(label)
                c_values = (['—' if c_ref is None else format_number(value / c_ref) + 'x'
                             for value in values])
                lines[c_row] = '| ↳ /C | ' + ' | '.join(c_values) + ' |'
                js_row = metric_row(lines, index, '↳ /JS')
                js_values = []
                for column, value in enumerate(values):
                    if title == 'JavaScript' and column == 0:
                        js_values.append('—')
                    else:
                        js_values.append(format_number(value / js_reference[label]) + 'x')
                lines[js_row] = '| ↳ /JS | ' + ' | '.join(js_values) + ' |'
            total_row = find_table_row(lines, '**Total Stable**')
            if len(table_cells(lines[total_row])) != len(keys) + 1:
                raise ValueError('Unexpected column count for ' + title + ': Total Stable')
            lines[total_row] = ('| **Total Stable** | ' + ' | '.join(
                format_time(results[key]['total_ms'], 'ms') for key in keys) + ' |')
            c_row = metric_row(lines, total_row, '↳ /C')
            c_total = results['c']['total_ms']
            lines[c_row] = '| ↳ /C | ' + ' | '.join(
                format_number(results[key]['total_ms'] / c_total) + 'x' for key in keys) + ' |'
            js_row = metric_row(lines, total_row, '↳ /JS')
            js_values = []
            for column, key in enumerate(keys):
                if title == 'JavaScript' and column == 0:
                    js_values.append('—')
                else:
                    js_values.append(format_number(results[key]['total_ms']
                                                   / results['js-pure']['total_ms']) + 'x')
            lines[js_row] = '| ↳ /JS | ' + ' | '.join(js_values) + ' |'
        else:
            expected_rows = EXTENDED_ROWS if title == 'Extended Results' else ROWS
            seen = []
            for i, line in enumerate(lines):
                if '|' not in line:
                    continue
                name_cell = line.split('|', 1)[0].strip()
                label = row_label(name_cell)
                if label not in expected_rows and label != '**Total Execution Time**':
                    continue
                total = label == '**Total Execution Time**'
                old_cells = line.split('|')[1:-1]
                new_cells = []
                for cell, key in zip(old_cells, keys):
                    result = results[key]
                    value = result['total_ms'] if total else result['times_us'][expected_rows.index(label)]
                    suffix = cell.split('<br>', 1)[1] if '<br>' in cell else ''
                    new_cells.append(format_time(value, 'ms' if total else 'μs')
                                     + ('<br>' + suffix if suffix else ''))
                lines[i] = name_cell + ' | ' + ' | '.join(new_cells) + ' |'
                seen.append(label)
            if seen != expected_rows + ['**Total Execution Time**']:
                raise ValueError('Unexpected table row order: ' + title)
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return refresh_new_totals(text)


def render_wip(original, results):
    """Update the 23 WIP rows in the nine JS/Go columns and their totals."""
    text = original
    labels = [case['label'] for case in expected_cases('wip')]
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
        for label in labels:
            index = find_benchmark_row(lines, label)
            if len(table_cells(lines[index])) != len(keys) + 1:
                raise ValueError('Unexpected column count for ' + title + ': ' + label)
            values = [results[key]['library']['times_us'][labels.index(label)] for key in keys]
            lines[index] = ('| ' + table_cells(lines[index])[0] + ' | '
                            + ' | '.join(format_time(value, 'μs') for value in values) + ' |')
            lines[metric_row(lines, index, '↳ /C')] = (
                '| ↳ /C | ' + ' | '.join('—' for _ in keys) + ' |')
            js_values = []
            for column, value in enumerate(values):
                if title == 'JavaScript' and column == 0:
                    js_values.append('—')
                else:
                    js_values.append(format_number(value / js_reference[label]) + 'x')
            lines[metric_row(lines, index, '↳ /JS')] = (
                '| ↳ /JS | ' + ' | '.join(js_values) + ' |')
            seen.append(label)
        if seen != labels:
            raise ValueError('Unexpected WIP row order: ' + title)
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return refresh_new_totals(text)


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
    print('Published 23 WIP rows and their totals across nine JS/Go columns; core cells preserved.' if args.wip_only
          else 'Published 40 complete validated columns; existing headers preserved.')


if __name__ == '__main__':
    main()
