#!/usr/bin/env python3
"""Publish the additional C references without changing any backend timing."""
import difflib
import hashlib
import json
from pathlib import Path
import re
import sys

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / 'bin/benchmark'))
from publish import (GROUPS, backend_status_rows, format_number, format_time,
                     matched_status_summary, section_lines, statistics_cell, table_cells)

ARCHIVE = ROOT / 'docs/benchmark-results/2026-11-26-c-library.json'
WORK = ROOT / 'var/benchmark/c-library-20261126'
METRICS = ('/', 'μ', 'Q₂', 'G')


def row(cells):
    return '| ' + ' | '.join(cells) + ' |'


def measurement(value, approximate=False):
    return ('measured', value, approximate)


def main():
    archive = json.loads(ARCHIVE.read_text())
    refs_path = WORK / 'references.json'
    refs = json.loads(refs_path.read_text()) if refs_path.exists() else archive['publication']['references']
    # Extractor verifies each historical value against its current published cell.
    columns = refs['columns']
    values = {key: {name: entry['time_us'] for name, entry in column.items()}
              for key, column in columns.items()}
    new_c = archive['medians_us']
    c_values = values['c'] | new_c
    references = {name: measurement(value, columns['c'].get(name, {}).get('approximate', False))
                  for name, value in c_values.items()}
    original = (ROOT / 'README.md').read_text()
    assert hashlib.sha256(original.encode()).hexdigest() == archive['readme_before_sha256']
    text = original
    audit = {}
    for title in ('JavaScript', 'Go', 'C (reference)'):
        found, lines = section_lines(text, title)
        keys = GROUPS[title]
        controlled, rows = backend_status_rows(lines, len(keys))
        positions = {label: (index, cells, statuses) for label, index, cells, statuses in rows}
        for name, c_time in new_c.items():
            index, cells, statuses = positions[name]
            if title == 'C (reference)':
                # Add ratio rows inside the existing table, in reverse order below.
                continue
            updated = ['↳ /C']
            for column, key in enumerate(keys):
                if cells[column + 1] == '—' or statuses[column] == 'Backlog':
                    updated.append('—')
                else:
                    value = values[key][name]
                    assert format_time(value) == cells[column + 1]
                    updated.append(format_number(value / c_time) + 'x')
            assert table_cells(lines[index + 1])[0] == '↳ /C'
            lines[index + 1] = row(updated)

        # Recompute only aggregates whose matched C population has grown.
        for index, line in enumerate(lines):
            cells = table_cells(line) if line.startswith('|') else []
            if not cells or cells[0] not in ('**Total WIP**', '**Total**'):
                continue
            status = 'WIP' if cells[0] == '**Total WIP**' else None
            summaries = []
            for column, key in enumerate(keys):
                active = controlled[column] if status else not controlled[column]
                if not active:
                    summaries.append('')
                    continue
                names = [name for name, _, row_cells, states in rows
                         if (status is None or states[column] == status)
                         and row_cells[column + 1] != '—'
                         and not row_cells[column + 1].startswith('<')]
                if title == 'C (reference)':
                    names += list(new_c)
                    candidates = {name: references[name] for name in names}
                    target = {name: measurement(value, columns['js-pure'][name]['approximate'])
                              for name, value in values['js-pure'].items()}
                    total = sum(c_values[name] for name in names) / 1000
                    approximate = any(entry[2] for entry in candidates.values())
                    cells[column + 1] = '<strong>' + ('≈ ' + format_number(total) + ' ms' if approximate
                                                     else format_time(total, 'ms')) + '</strong>'
                    lines[index] = row(cells)
                else:
                    candidates = {name: measurement(values[key][name], columns[key][name]['approximate']) for name in names}
                    target = references
                summary = matched_status_summary(candidates, target)
                assert summary['count'] == summary['population'] == len(names)
                summaries.append(summary)
                audit[key] = summary
            reference_name = 'JS' if title == 'C (reference)' else 'C'
            for metric in METRICS:
                label = '↳ ' + ('' if metric == '/' else metric) + '/' + reference_name
                matches = [j for j in range(index + 1, min(index + 9, len(lines)))
                           if table_cells(lines[j])[0] == label]
                assert len(matches) == 1
                lines[matches[0]] = row([label] + ['' if summary == '' else statistics_cell(summary, metric)
                                                for summary in summaries])
        if title == 'C (reference)':
            for name in reversed(list(new_c)):
                index, cells, _ = positions[name]
                assert cells[1] == '—'
                cells[1] = format_time(new_c[name])
                lines[index:index + 1] = [row(cells), row(['↳ /C', '—']),
                                         row(['↳ /JS', format_number(new_c[name] / values['js-pure'][name]) + 'x'])]
        text = text[:found.start(2)] + '\n'.join(lines) + '\n' + text[found.end(2):]

    text = text.replace('for example, Go/C covers 14 of the 37 measured Go cases, while Go/JS covers all 37.',
                        'Go/C and Go/JS both cover all 37 measured Go cases.')
    text = text.replace('#### C (reference)\n',
                        '#### C (reference)\n\nThe 23 [library references](tmp/c-library) complement the existing core reference. '
                        '[Measurements and validation](docs/benchmark-results/2026-11-26-c-library.md).\n')

    # Preserve every non-C timing, every status and every existing Go/JS ratio.
    for title in GROUPS:
        before, after = section_lines(original, title), section_lines(text, title)
        if before is None:
            continue
        if title not in ('JavaScript', 'Go', 'C (reference)'):
            assert before[1] == after[1], title
            continue
        if title != 'C (reference)':
            old_rows = [line for line in before[1] if line.startswith('| [') or line.startswith('|  |')
                        or re.match(r'\| ↳ (?:μ|Q₂|G)?/JS ', line) or line.startswith('| **')]
            new_rows = [line for line in after[1] if line.startswith('| [') or line.startswith('|  |')
                        or re.match(r'\| ↳ (?:μ|Q₂|G)?/JS ', line) or line.startswith('| **')]
            assert old_rows == new_rows, title
    separators = re.compile(r'^\|?\s*:?-{3,}:?\s*(?:\|\s*:?-{3,}:?\s*)+\|?\s*$', re.M)
    assert len(separators.findall(text)) == len(separators.findall(original)) == 17
    assert audit['go-pure']['count'] == 37 and audit['c']['count'] == 40
    (WORK / 'README.after.md').write_text(text)
    (WORK / 'README.diff').write_text(''.join(difflib.unified_diff(original.splitlines(True), text.splitlines(True),
                                                               fromfile='README.before.md', tofile='README.md')))
    archive['publication'] = {
        'references': refs, 'aggregates': audit, 'historical_c_times_unchanged': True,
        'backend_times_statuses_js_ratios_unchanged': True, 'table_count': 17,
        'readme_after_sha256': hashlib.sha256(text.encode()).hexdigest(),
        'script': {'path': str(Path(__file__).relative_to(ROOT)), 'text': Path(__file__).read_text()},
    }
    if '--write' in sys.argv:
        assert (ROOT / 'README.md').read_text() == original
        (ROOT / 'README.md').write_text(text)
        ARCHIVE.write_text(json.dumps(archive, indent=2, ensure_ascii=False, allow_nan=False) + '\n')
    print(json.dumps(audit['go-pure'], indent=2))


if __name__ == '__main__':
    main()
