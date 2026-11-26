#!/usr/bin/env python3
"""Publish complete validated campaigns, preserving the user's column headers."""
import argparse
from decimal import Decimal
import hashlib
from html import escape, unescape
import json
import math
import os
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
    """Use two decimals at magnitude >= 1, otherwise two significant digits."""
    if value == 0 or abs(value) >= 1:
        return f'{value:.2f}'
    rounded = Decimal(format(value, '#.2g'))
    # Rounding a value just below one may cross the ordinary-decimal boundary.
    return format(rounded, '.2f' if abs(rounded) >= 1 else 'f')


def format_table_numbers(text):
    """Normalize existing timing/ratio displays without turning bounds into data."""
    number = re.compile(r'(?<![\w./=])(-?\d+(?:\.\d+)?)(\s*(?:μs|µs|ns|ms|s|x|×))(?!\w)')
    lines = text.splitlines(keepends=True)
    for index, line in enumerate(lines):
        if '|' not in line or line.lstrip().startswith('>'):
            continue
        def replace(match):
            if re.search(r'<\s*$', line[:match.start()]):
                return match.group(0)
            return format_number(float(match.group(1))) + match.group(2)
        lines[index] = number.sub(replace, line)
    return ''.join(lines)


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
            in_total = True
        elif in_total and line.startswith(('| ↳ ', '↳ ')):
            result.append(bold_total_line(line))
        else:
            in_total = False
            result.append(line)
    return '\n'.join(result) + '\n'


def timing_measurement(cell, key, label, measurements, catalogue):
    """Read a time, retaining published bounds and its precision provenance."""
    plain = unescape(re.sub(r'</?strong>', '', cell)).strip()
    if plain == '—':
        return ('missing', None, False)
    if plain.startswith('<'):
        return ('bounded', None, False)
    match = re.fullmatch(r'~?\s*([\d.]+) μs', plain)
    if match is None:
        raise ValueError('Invalid timing cell for ' + key + ': ' + label)
    published = float(match[1])
    if published <= 0:
        return ('bounded', None, False)
    incoming = measurements.get(key, {})
    stored = catalogue.get('columns', {}).get(key, {}).get(label, {})
    if label in incoming:
        value, approximate = incoming[label], False
    elif stored.get('display') == cell:
        value, approximate = stored['time_us'], False
    else:
        value, approximate = published, True
    if not math.isfinite(value) or value <= 0:
        raise ValueError('Invalid statistical timing for ' + key + ': ' + label)
    return ('measured', value, approximate)


def summarize_ratios(candidates, references):
    """Summarize matched positive times; a missing measurement invalidates a block."""
    if len(candidates) != len(references):
        raise ValueError('Mismatched ratio populations')
    pairs = list(zip(candidates, references))
    if not pairs or any('missing' in (value[0], reference[0]) for value, reference in pairs):
        return None
    measured = [(value, reference) for value, reference in pairs
                if value[0] == reference[0] == 'measured']
    if not measured:
        return None
    ratios = [value[1] / reference[1] for value, reference in measured]
    return {
        'μ': statistics.fmean(ratios),
        'Q₂': statistics.median(ratios),
        'G': math.exp(statistics.fmean(math.log(value) for value in ratios)),
        'count': len(ratios), 'population': len(pairs),
        'approximate': any(value[2] or reference[2] for value, reference in measured),
    }


def statistics_cell(summary, statistic):
    if summary is None:
        return '—'
    value = ('≈ ' if summary['approximate'] else '') + format_number(summary[statistic]) + 'x'
    cell = '<strong>' + escape(value) + '</strong>'
    if summary['count'] != summary['population']:
        cell += ' (n=' + str(summary['count']) + '/' + str(summary['population']) + ')'
    return cell


def campaign_times(results, wip=False):
    return {key: dict(zip(result['library']['labels'], result['library']['times_us']))
            if wip else dict(zip(EXTENDED_ROWS if key.endswith('-x') else ROWS, result['times_us']))
            for key, result in results.items() if not wip or not key.endswith('-x')}


def updated_ratio_catalogue(catalogue, measurements, origins):
    """Retain raw times from validated publications without dropping older blocks."""
    updated = {**catalogue, 'sources': dict(catalogue.get('sources', {})),
               'columns': {key: dict(values) for key, values in catalogue.get('columns', {}).items()}}
    for key, values in measurements.items():
        source = origins[key]['path']
        updated['sources'][source] = origins[key]['sha256']
        column = updated['columns'].setdefault(key, {})
        for label, value in values.items():
            column[label] = {'time_us': value, 'display': format_time(value), 'source': source}
    if 'missing_columns' in catalogue:
        updated['missing_columns'] = []
        for item in catalogue['missing_columns']:
            remaining = [label for label in item['missing']
                         if label not in measurements.get(item['key'], {})]
            if remaining:
                updated['missing_columns'].append({**item, 'missing': remaining})
    return updated


def backend_status_layout(text):
    """Recognize per-column statuses without assigning them to reference tables."""
    return any(line.startswith('|') and table_cells(line)
               and table_cells(line)[0] == ''
               and any(re.match(r'\*\*(Stable|WIP|Backlog)\*\*', cell)
                       for cell in table_cells(line)[1:])
               for line in text.splitlines())


def backend_status_rows(lines, columns):
    header = next(line for line in lines if 'Benchmark' in line and '|' in line)
    headers = [cell.strip() for cell in header.strip().strip('|').split('|')][1:]
    if len(headers) != columns:
        raise ValueError('Unexpected backend status column count')
    controlled = [bool(re.search(r'\[(?:gopurs|purust|phpurs|sharpurs|javapurs)\]\(', cell))
                  for cell in headers]
    states = [None] * columns
    rows = []
    for index, line in enumerate(lines):
        if not line.startswith('|'):
            continue
        cells = table_cells(line)
        if len(cells) != columns + 1:
            continue
        if not cells[0]:
            for column, cell in enumerate(cells[1:]):
                status = re.match(r'\*\*(Stable|WIP|Backlog)\*\*', cell)
                if status:
                    if not controlled[column]:
                        raise ValueError('A reference or external backend cannot have a status')
                    states[column] = status[1]
        elif cells[0].startswith('['):
            rows.append((row_label(cells[0]), index, cells, list(states)))
    return controlled, rows


def validate_backend_status_rows(text, groups, labels):
    for title in groups:
        found = section_lines(text, title)
        if found is None:
            raise ValueError('Missing section: ' + title)
        for label in labels:
            find_benchmark_row(found[1], label)


def backend_status_times(text, measurements, catalogue):
    sections, times = {}, {}
    for title in NEW_LAYOUT_GROUPS:
        found = section_lines(text, title)
        if found is None:
            continue
        _, lines = found
        keys = GROUPS[title]
        controlled, rows = backend_status_rows(lines, len(keys))
        sections[title] = (controlled, rows)
        for column, key in enumerate(keys):
            times[key] = {}
            for label, _, cells, states in rows:
                times[key][label] = (('missing', None, False) if states[column] == 'Backlog'
                                    else timing_measurement(cells[column + 1], key, label,
                                                            measurements, catalogue))
    return sections, times


def matched_status_summary(candidate, reference):
    """Use only matching measured cases; expose incomplete reference coverage."""
    pairs = [(value, reference[label]) for label, value in candidate.items()
             if reference.get(label, ('missing', None, False))[0] == 'measured']
    if not pairs:
        return None
    ratios = [value[1] / ref[1] for value, ref in pairs]
    return {'/': sum(value[1] for value, _ in pairs) / sum(ref[1] for _, ref in pairs),
            'μ': statistics.fmean(ratios), 'Q₂': statistics.median(ratios),
            'G': math.exp(statistics.fmean(math.log(value) for value in ratios)),
            'count': len(pairs), 'population': len(candidate),
            'approximate': any(value[2] or ref[2] for value, ref in pairs)}


def refresh_backend_status_totals(text, measurements=None, catalogue=None):
    measurements = {} if measurements is None else measurements
    catalogue = {} if catalogue is None else catalogue
    sections, times = backend_status_times(text, measurements, catalogue)
    for title, (controlled, rows) in sections.items():
        match, lines = section_lines(text, title)
        keys = GROUPS[title]
        for index in range(len(lines) - 1, -1, -1):
            cells = table_cells(lines[index]) if lines[index].startswith('|') else []
            if not cells or cells[0] not in ['**Total Stable**', '**Total WIP**', '**Total**']:
                continue
            total_label = cells[0]
            status = {'**Total Stable**': 'Stable', '**Total WIP**': 'WIP'}.get(total_label)
            values, summaries = [], {'C': [], 'JS': []}
            for column, key in enumerate(keys):
                active = controlled[column] if status else not controlled[column]
                if not active:
                    values.append('')
                    for reference in summaries:
                        summaries[reference].append('')
                    continue
                candidate = {label: times[key][label] for label, _, _, states in rows
                             if (status is None or states[column] == status)
                             and times[key][label][0] == 'measured'}
                time_ms = sum(value[1] for value in candidate.values()) / 1000
                time = ('≈ ' + format_number(time_ms) + ' ms' if any(value[2] for value in candidate.values())
                        else format_time(time_ms, 'ms'))
                total = '<strong>' + time + '</strong>' if candidate or status else '—'
                if status is None and candidate and len(candidate) != len(rows):
                    total += ' (n=' + str(len(candidate)) + '/' + str(len(rows)) + ')'
                values.append(total)
                for reference, reference_key in [('C', 'c'), ('JS', 'js-pure')]:
                    summary = (None if reference == 'JS' and reference_key == key else
                               matched_status_summary(candidate, times.get(reference_key, {})))
                    if summary is not None and reference_key == key:
                        summary['approximate'] = False  # A reference compared with itself is exactly one.
                    summaries[reference].append(summary)
            end = index + 1
            while end < len(lines) and lines[end].startswith('| ↳ '):
                end += 1
            metrics = []
            for reference in ['C', 'JS']:
                for statistic in ['/', 'μ', 'Q₂', 'G']:
                    label = '↳ ' + ('' if statistic == '/' else statistic) + '/' + reference
                    metrics.append('| ' + label + ' | ' + ' | '.join(
                        '' if summary == '' else statistics_cell(summary, statistic)
                        for summary in summaries[reference]) + ' |')
            lines[index:end] = ['| ' + total_label + ' | ' + ' | '.join(values) + ' |'] + metrics
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return text


def render_backend_status(original, measurements):
    """Update measured cells without publishing or reclassifying Backlog cases."""
    text = original
    references = {}
    for title, reference in [('JavaScript', 'JS'), ('C (reference)', 'C')]:
        found = section_lines(text, title)
        if found is None:
            references[reference] = {}
            continue
        _, lines = found
        key = GROUPS[title][0]
        _, rows = backend_status_rows(lines, len(GROUPS[title]))
        references[reference] = {}
        for label, _, cells, _ in rows:
            cell = format_time(measurements[key][label]) if label in measurements.get(key, {}) else cells[1]
            references[reference][label] = timing_measurement(cell, key, label, measurements, {})
    for title in NEW_LAYOUT_GROUPS:
        found = section_lines(text, title)
        if found is None:
            continue
        match, lines = found
        keys = GROUPS[title]
        _, rows = backend_status_rows(lines, len(keys))
        for label, index, cells, states in reversed(rows):
            end = index + 1
            while end < len(lines) and lines[end].startswith('| ↳ '):
                end += 1
            old_metrics = {table_cells(line)[0]: table_cells(line)[1:]
                           for line in lines[index + 1:end]}
            changed = any(label in values for values in measurements.values())
            for column, key in enumerate(keys):
                if states[column] == 'Backlog':
                    cells[column + 1] = '—'
                elif label in measurements.get(key, {}):
                    cells[column + 1] = format_time(measurements[key][label])
            replacement = ['| ' + ' | '.join(cells) + ' |']
            if any(cell_us(cell) is not None for cell in cells[1:]):
                for reference in ['C', 'JS']:
                    metric = '↳ /' + reference
                    values = list(old_metrics.get(metric, ['—'] * len(keys)))
                    for column, key in enumerate(keys):
                        if states[column] == 'Backlog' or (reference == 'JS' and key == 'js-pure'):
                            values[column] = '—'
                        elif changed:
                            value = timing_measurement(cells[column + 1], key, label, measurements, {})
                            ref = references[reference].get(label, ('missing', None, False))
                            values[column] = (('≈ ' if value[2] or ref[2] else '') + format_number(value[1] / ref[1]) + 'x'
                                              if value[0] == ref[0] == 'measured' else '—')
                    replacement.append('| ' + metric + ' | ' + ' | '.join(values) + ' |')
            lines[index:end] = replacement
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return refresh_backend_status_totals(text, measurements)


def refresh_extended_statistics(text, measurements, catalogue):
    """Use the extended table's own JS reference and preserve its ratio rows."""
    found = section_lines(text, 'Extended Results')
    if found is None:
        return text
    match, lines = found
    keys = GROUPS['Extended Results']
    times = {key: {} for key in keys}
    total_index = None
    for index, line in enumerate(lines):
        cells = [cell.strip() for cell in line.split('|')]
        label = row_label(cells[0])
        if label == '**Total Execution Time**':
            total_index = index
        if label not in EXTENDED_ROWS:
            continue
        if len(cells[1:-1]) != len(keys):
            raise ValueError('Unexpected statistical column count in Extended Results')
        for key, cell in zip(keys, cells[1:-1]):
            times[key][label] = timing_measurement(cell, key, label,
                                                   measurements, catalogue)
    if total_index is None or any(list(values) != EXTENDED_ROWS for values in times.values()):
        raise ValueError('Unexpected statistical population in Extended Results')
    summaries = [None if key == 'js-x' else summarize_ratios(
        list(times[key].values()), list(times['js-x'].values())) for key in keys]
    end = total_index + 1
    while end < len(lines) and lines[end].startswith('↳ '):
        end += 1
    old_metrics = {line.split('|', 1)[0].strip(): line
                   for line in lines[total_index + 1:end]}
    rows = []
    for reference in ['C', 'JS']:
        rows.append(old_metrics['↳ /' + reference])
        for statistic in ['μ', 'Q₂', 'G']:
            cells = ['—'] * len(keys) if reference == 'C' else [
                statistics_cell(summary, statistic) for summary in summaries]
            rows.append('↳ ' + statistic + '/' + reference + ' | ' + ' | '.join(cells) + ' |')
    lines[total_index + 1:end] = rows
    return text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]


def refresh_ratio_statistics(text, measurements=None, catalogue=None):
    """Refresh only μ, Q₂ and G rows, using times rather than rounded ratios.

    ``measurements`` maps backend keys to ``{label: time_us}`` and overrides
    archived times, including when both round to the same displayed cell.
    Archived times are usable only while their published cell still matches.
    """
    measurements = {} if measurements is None else measurements
    if catalogue is None:
        source = ROOT / 'bin/benchmark/ratio-measurements.json'
        catalogue = json.loads(source.read_text()) if source.exists() else {}
    if backend_status_layout(text):
        return refresh_extended_statistics(
            refresh_backend_status_totals(text, measurements, catalogue), measurements, catalogue)
    sections = {}
    for title in NEW_LAYOUT_GROUPS:
        found = section_lines(text, title)
        if found is None or not new_layout(found[1]):
            continue
        _, lines = found
        blocks, block = {'stable': [], 'incl': []}, None
        for line in lines:
            if line.startswith('| **Stable** |'):
                block = 'stable'
            elif line.startswith('| **Incl WIP** |'):
                block = 'incl'
            elif line.startswith('| **Excl WIP** |'):
                block = None
            elif block is not None and line.startswith('| ['):
                cells = table_cells(line)
                label = row_label(cells[0])
                if len(cells) != len(GROUPS[title]) + 1:
                    raise ValueError('Unexpected statistical column count in ' + title)
                blocks[block].append((label, cells[1:]))
        if len(blocks['stable']) != 14 or len(blocks['incl']) != 23:
            raise ValueError('Unexpected statistical population in ' + title)
        sections[title] = blocks

    def times(title, block, column):
        key = GROUPS[title][column]
        rows = sections[title][block] if block != 'combined' else (
            sections[title]['stable'] + sections[title]['incl'])
        return {label: timing_measurement(cells[column], key, label, measurements, catalogue)
                for label, cells in rows}

    for title in sections:
        match, lines = section_lines(text, title)
        for block, total in [('stable', '**Total Stable**'),
                             ('incl', '**Total Incl WIP**'),
                             ('combined', '**Total Stable + Incl WIP**')]:
            total_index = find_table_row(lines, total)
            end = total_index + 1
            while end < len(lines) and lines[end].startswith('| ↳ '):
                end += 1
            old_metrics = {table_cells(line)[0]: line for line in lines[total_index + 1:end]}
            metrics = []
            for reference_name, reference_title in [('C', 'C (reference)'), ('JS', 'JavaScript')]:
                metrics.append(old_metrics['↳ /' + reference_name])
                reference = (times(reference_title, block, 0) if reference_title in sections else {})
                summaries = []
                for column, key in enumerate(GROUPS[title]):
                    if key == 'js-pure' and reference_name == 'JS':
                        summaries.append(None)
                        continue
                    candidate = times(title, block, column)
                    summaries.append(summarize_ratios(
                        list(candidate.values()),
                        [reference.get(label, ('missing', None, False)) for label in candidate]))
                for statistic in ['μ', 'Q₂', 'G']:
                    metrics.append('| ↳ ' + statistic + '/' + reference_name + ' | '
                                   + ' | '.join(statistics_cell(value, statistic) for value in summaries) + ' |')
            lines[total_index + 1:end] = metrics
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return refresh_extended_statistics(text, measurements, catalogue)


def refresh_new_totals(text, measurements=None, catalogue=None):
    """Recompute the Stable/Incl WIP aggregate rows from displayed cells."""
    if backend_status_layout(text):
        return refresh_ratio_statistics(text, measurements, catalogue)
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
        columns = len(GROUPS[title])
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
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return bold_totals(refresh_ratio_statistics(text, measurements, catalogue))


def render(original, results):
    per_backend = backend_status_layout(original)
    if per_backend:
        validate_backend_status_rows(original, NEW_LAYOUT_GROUPS, ROWS)
    text = render_backend_status(original, campaign_times(results)) if per_backend else original
    c_reference = dict(zip(ROWS, results['c']['times_us']))
    js_reference = dict(zip(ROWS, results['js-pure']['times_us']))
    for title, keys in GROUPS.items():
        if per_backend and title in NEW_LAYOUT_GROUPS:
            continue
        pattern = r'(#### ' + re.escape(title) + r'\n)(.*?)(?=\n#### |\n### |\n## |\Z)'
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
                    new_cells.append(format_time(value, 'ms' if total else 'μs'))
                lines[i] = name_cell + ' | ' + ' | '.join(new_cells) + ' |'
                reference = results['js-x']
                js_value = reference['total_ms'] if total else reference['times_us'][expected_rows.index(label)]
                ratios = ['—' if key == 'js-x' else format_number(
                    (results[key]['total_ms'] if total else results[key]['times_us'][expected_rows.index(label)])
                    / js_value) + 'x' for key in keys]
                for metric, values in [('C', ['—'] * len(keys)), ('JS', ratios)]:
                    ratio_index = next(j for j in range(i + 1, len(lines))
                                       if lines[j].startswith('↳ /' + metric + ' |'))
                    lines[ratio_index] = '↳ /' + metric + ' | ' + ' | '.join(values) + ' |'
                seen.append(label)
            if seen != expected_rows + ['**Total Execution Time**']:
                raise ValueError('Unexpected table row order: ' + title)
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    if per_backend:
        return bold_totals(refresh_extended_statistics(text, campaign_times(results), {}))
    return refresh_new_totals(text, campaign_times(results))


def render_wip(original, results):
    """Update the 23 WIP rows in the nine JS/Go columns and their totals."""
    if backend_status_layout(original):
        validate_backend_status_rows(original, WIP_GROUPS,
                                     [case['label'] for case in expected_cases('wip')])
        return render_backend_status(original, campaign_times(results, wip=True))
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
    return refresh_new_totals(text, campaign_times(results, wip=True))


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--campaign', type=Path, action='append', required=True)
    parser.add_argument('--readme', type=Path, default=ROOT / 'README.md')
    parser.add_argument('--wip-only', action='store_true', help='Publish the 23 new rows in the nine JS/Go columns')
    args = parser.parse_args()
    results = {}
    origins = {}
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
        origins.update({key: {'path': os.path.relpath(result_file.resolve(), ROOT),
                              'sha256': manifest['results_sha256']} for key in incoming})
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
    updated = format_table_numbers(updated)
    catalogue_path = ROOT / 'bin/benchmark/ratio-measurements.json'
    catalogue_text = None
    if args.readme.resolve() == (ROOT / 'README.md').resolve() and catalogue_path.exists():
        catalogue_original = catalogue_path.read_text() if catalogue_path.exists() else None
        catalogue = json.loads(catalogue_original) if catalogue_original is not None else {}
        catalogue_text = json.dumps(updated_ratio_catalogue(
            catalogue, campaign_times(results, wip=args.wip_only), origins),
            ensure_ascii=False, indent=2) + '\n'
    if args.readme.read_text() != original:
        raise RuntimeError('README changed while preparing publication')
    if catalogue_text is not None:
        latest = catalogue_path.read_text() if catalogue_path.exists() else None
        if latest != catalogue_original:
            raise RuntimeError('Ratio catalogue changed while preparing publication')
    args.readme.write_text(updated)
    if catalogue_text is not None:
        catalogue_path.write_text(catalogue_text)
    print('Published 23 WIP rows and their totals across nine JS/Go columns; core cells preserved.' if args.wip_only
          else 'Published 40 complete validated columns; existing headers preserved.')


if __name__ == '__main__':
    main()
