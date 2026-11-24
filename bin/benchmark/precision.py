"""Refresh numeric tables from measured times without inventing old bounds.

Call ``refresh(text, catalogue, measurements, origins)`` after validating new
campaigns. The result is a README string and a catalogue with matching displays;
the function does not write either file. ``measurements`` and ``origins`` use
the same shapes as publish.updated_ratio_catalogue.
"""
import math

import publish as pub


def replace_cell(line, column, value, leading=True):
    parts = line.split('|')
    parts[column + (2 if leading else 1)] = ' ' + value + ' '
    return '|'.join(parts)


def refresh(text, catalogue, measurements=None, origins=None):
    measurements = {} if measurements is None else measurements
    origins = {} if origins is None else origins
    for values in measurements.values():
        if any(not math.isfinite(value) or value <= 0 for value in values.values()):
            raise ValueError('New measurements must be finite and positive')
    updated = pub.updated_ratio_catalogue(catalogue, measurements, origins)

    def timing(cell, key, label):
        if label in measurements.get(key, {}):
            value = measurements[key][label]
        else:
            stored = catalogue.get('columns', {}).get(key, {}).get(label, {})
            if stored.get('display') != cell:
                return cell
            value = stored.get('time_us')
        if not isinstance(value, (float, int)) or not math.isfinite(value) or value <= 0:
            raise ValueError('Invalid archived measurement: ' + key + ': ' + label)
        display = pub.format_time(value)
        entry = updated['columns'][key][label]
        updated['columns'][key][label] = {**entry, 'display': display}
        return display

    # A matching archived measurement can replace an old formatting bound.
    # An unmatched bound, zero or missing value remains exactly as published.
    for title, keys in pub.GROUPS.items():
        found = pub.section_lines(text, title)
        if found is None:
            continue
        match, lines = found
        leading = title != 'Extended Results'
        for index, line in enumerate(lines):
            if not line.startswith('| [' if leading else '['):
                continue
            cells = pub.table_cells(line) if leading else [part.strip() for part in line.split('|')[:-1]]
            if len(cells) != len(keys) + 1:
                raise ValueError('Unexpected timing columns: ' + title)
            label = pub.row_label(cells[0])
            for column, key in enumerate(keys):
                lines[index] = replace_cell(lines[index], column,
                                            timing(cells[column + 1], key, label), leading)
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]

    sections = {}
    for title in pub.NEW_LAYOUT_GROUPS:
        found = pub.section_lines(text, title)
        if found is None:
            continue
        _, lines = found
        section = {'rows': {}, 'stable': [], 'incl': [], 'totals': {}}
        block = None
        for line in lines:
            if line.startswith('| **Stable** |'):
                block = 'stable'
            elif line.startswith('| **Incl WIP** |'):
                block = 'incl'
            elif line.startswith('| **Excl WIP** |'):
                block = None
            elif line.startswith('| ['):
                cells = pub.table_cells(line)
                label = pub.row_label(cells[0])
                section['rows'][label] = [pub.timing_measurement(cell, key, label, {}, updated)
                                          for key, cell in zip(pub.GROUPS[title], cells[1:])]
                if block is not None:
                    section[block].append(label)
        section['combined'] = section['stable'] + section['incl']
        for block, total in [('stable', '**Total Stable**'), ('incl', '**Total Incl WIP**'),
                             ('combined', '**Total Stable + Incl WIP**')]:
            totals = []
            for column in range(len(pub.GROUPS[title])):
                values = [section['rows'][label][column] for label in section[block]]
                value = (sum(entry[1] for entry in values) / 1000
                         if values and all(entry[0] == 'measured' for entry in values) else None)
                # A published Stable total remains usable even when a tiny
                # component was historically rounded to a bound or zero.
                if value is None and block == 'stable':
                    old = pub.table_cells(lines[pub.find_table_row(lines, total)])[column + 1]
                    value = pub.cell_ms(old)
                if value is None and block == 'combined':
                    stable, incl = section['totals']['stable'][column], section['totals']['incl'][column]
                    if stable is not None and incl is not None:
                        value = stable + incl
                totals.append(value)
            section['totals'][block] = totals
        sections[title] = section

    def ratio(candidate, reference):
        if candidate is None or reference is None or candidate <= 0 or reference <= 0:
            return '—'
        return pub.format_number(candidate / reference) + 'x'

    for title, section in sections.items():
        match, lines = pub.section_lines(text, title)
        for label, values in section['rows'].items():
            row = pub.find_benchmark_row(lines, label)
            for suffix, reference_title in [('C', 'C (reference)'), ('JS', 'JavaScript')]:
                reference = sections.get(reference_title, {}).get('rows', {}).get(label, [('missing', None, False)])[0]
                index = pub.metric_row(lines, row, '↳ /' + suffix)
                for column, key in enumerate(pub.GROUPS[title]):
                    candidate = values[column]
                    cell = ('—' if (key == 'js-pure' and suffix == 'JS')
                            or candidate[0] != 'measured' or reference[0] != 'measured'
                            else ratio(candidate[1], reference[1]))
                    lines[index] = replace_cell(lines[index], column, cell)
        for block, total in [('stable', '**Total Stable**'), ('incl', '**Total Incl WIP**'),
                             ('combined', '**Total Stable + Incl WIP**')]:
            row = pub.find_table_row(lines, total)
            for column, value in enumerate(section['totals'][block]):
                cell = '—' if value is None else pub.bold_value(pub.format_time(value, 'ms'))
                lines[row] = replace_cell(lines[row], column, cell)
            for suffix, reference_title in [('C', 'C (reference)'), ('JS', 'JavaScript')]:
                reference = sections.get(reference_title, {}).get('totals', {}).get(block, [None])[0]
                index = pub.metric_row(lines, row, '↳ /' + suffix)
                for column, key in enumerate(pub.GROUPS[title]):
                    cell = ('—' if key == 'js-pure' and suffix == 'JS'
                            else ratio(section['totals'][block][column], reference))
                    lines[index] = replace_cell(lines[index], column, pub.bold_value(cell))
        text = text[:match.start(2)] + '\n'.join(lines) + '\n' + text[match.end(2):]
    return pub.format_table_numbers(pub.refresh_ratio_statistics(text, catalogue=updated)), updated
