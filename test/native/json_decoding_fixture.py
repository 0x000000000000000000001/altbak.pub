#!/usr/bin/env python3
"""Independent JSON decoding inputs and full-result oracle; no compiler output.

Run without arguments to check the frozen corpus; --write regenerates it.
Successful outcomes are built alongside their input from explicit constructors.
Errors are specified from Argonaut's public decoding contract, not decoded here.
"""
import argparse
import copy
import hashlib
import json
from pathlib import Path
import unittest


FIXTURES = Path(__file__).resolve().parents[1] / 'fixtures' / 'json-decoding'
ABSENT = object()


def compact(value):
    return json.dumps(value, ensure_ascii=False, separators=(',', ':'), allow_nan=False)


def canonical(value):
    if isinstance(value, list):
        return [canonical(item) for item in value]
    if isinstance(value, dict):
        return {key: canonical(value[key]) for key in sorted(value)}
    if isinstance(value, float) and value.is_integer():
        return int(value)
    return value


def fingerprint(value):
    text = compact(canonical(value))
    for char in '<>&\u2028\u2029':
        text = text.replace(char, '\\u%04x' % ord(char))
    return hashlib.sha256(text.encode('utf-8')).hexdigest()


def profile(city, scores, note=ABSENT):
    source = {'city': city, 'scores': scores}
    if note is not ABSENT:
        source['note'] = note
    return source, {'city': city, 'scores': scores,
                    'note': None if note is ABSENT else note}


def user(ident, name, active, tags, details=ABSENT):
    source = {'id': ident, 'name': name, 'active': active, 'tags': tags}
    expected = dict(source, profile=None)
    if details is None:
        source['profile'] = None
    elif details is not ABSENT:
        source['profile'], expected['profile'] = details
    return source, expected


def view(path, duration=ABSENT):
    source = {'tag': 'view', 'path': path}
    if duration is not ABSENT:
        source['duration'] = duration
    return source, {'tag': 'view', 'path': path,
                    'duration': None if duration is ABSENT else duration}


def purchase(order_id, items):
    source = {'tag': 'purchase', 'orderId': order_id, 'items': items}
    return source, copy.deepcopy(source)


def payload(users=(), events=(), next_page=ABSENT, version=1):
    source = {'version': version, 'users': [pair[0] for pair in users],
              'events': [pair[0] for pair in events]}
    if next_page is not ABSENT:
        source['next'] = next_page
    expected = {'version': version, 'users': [pair[1] for pair in users],
                'events': [pair[1] for pair in events],
                'next': None if next_page is ABSENT else next_page}
    return source, expected


def error(*path, expected=None):
    lines = ['An error occurred while decoding a JSON value:']
    for kind, value in path:
        if kind == 'key':
            lines.append("  At object key '%s':" % value)
        elif kind == 'index':
            lines.append('  At array index %d:' % value)
        elif kind == 'named':
            lines.append("  Under '%s':" % value)
        else:
            raise ValueError(kind)
    lines.append('  No value was found.' if expected is None
                 else "  Expected value of type '%s'." % expected)
    return {'error': '\n'.join(lines)}


def make_cases():
    cases = []

    def success(name, pair, benchmark=False):
        cases.append((name, pair[0], {'value': pair[1]}, benchmark))

    def failure(name, value, outcome):
        cases.append((name, value, outcome, False))

    success('empty-arrays', payload(), True)
    success('flat-record-arrays', payload(
        users=[user(i, 'user-%04d' % i, i % 2 == 0, ['member', 'flat'], None)
               for i in range(1000)],
        events=[view('/page/%d' % i, i % 500) for i in range(1000)],
        next_page='page-2'), True)
    success('nested-records-and-variants', payload(
        users=[user(i, 'nested-%04d' % i, True, ['member', 'group-%d' % (i % 9)],
                    profile('city-%d' % (i % 31),
                            [i + j / 8 for j in range(16)], 'profile-%d' % i))
               for i in range(500)],
        events=[purchase(i, [{'sku': 'sku-%d-%d' % (i, j),
                              'quantity': j + 1, 'price': i + j / 4}
                             for j in range(4)]) for i in range(500)]), True)
    success('null-and-missing-optionals', payload(
        users=[user(i, 'optional-%04d' % i, False, [],
                    (ABSENT if i % 3 == 0 else None if i % 3 == 1
                     else profile('Paris', [0, 1.5], None if i % 2 else ABSENT)))
               for i in range(750)],
        events=[view('/optional/%d' % i, None if i % 2 else ABSENT)
                for i in range(750)], next_page=None), True)
    text = 'Éléonore 東京 😀 𝄞 "quoted" \\ slash\nline\ttab <>&\u2028\u2029'
    success('unicode-escaping-and-numbers', payload(
        users=[user(i - 225, text + str(i), i % 2 == 0, [text, ''],
                    profile('Montréal', [-10.5, 0, 0.125, 1000000.25], text))
               for i in range(450)],
        events=[purchase(i, [{'sku': text, 'quantity': 1, 'price': 0.125},
                              {'sku': '', 'quantity': -2, 'price': 1000000.25}])
                for i in range(225)] + [view(text, 2147483647), view('', -2147483648)],
        next_page=text, version=2147483647), True)

    failure('wrong-root-type', [], error(expected='Object'))
    missing = payload()[0]
    del missing['version']
    failure('missing-required-version', missing, error(('key', 'version')))
    fractional = payload()[0]
    fractional['version'] = 1.5
    failure('fractional-int', fractional, error(('key', 'version'), expected='Integer'))
    wrong_boolean = payload(users=[user(1, 'a', 0, [])])[0]
    failure('wrong-boolean', wrong_boolean,
            error(('key', 'users'), ('named', 'Array'), ('index', 0),
                  ('key', 'active'), expected='Boolean'))
    wrong_profile = payload(users=[user(1, 'a', True, [])])[0]
    wrong_profile['users'][0]['profile'] = 42
    failure('invalid-optional-record', wrong_profile,
            error(('key', 'users'), ('named', 'Array'), ('index', 0),
                  ('key', 'profile'), expected='Object'))
    wrong_note = payload(users=[user(1, 'a', True, [], profile('Paris', [], 42))])[0]
    failure('invalid-nested-optional-string', wrong_note,
            error(('key', 'users'), ('named', 'Array'), ('index', 0),
                  ('key', 'profile'), ('key', 'note'), expected='String'))
    unknown_tag = payload()[0]
    unknown_tag['events'] = [{'tag': 'delete'}]
    failure('unknown-variant-tag', unknown_tag,
            error(('key', 'events'), ('named', 'Array'), ('index', 0),
                  expected='Event tag'))
    wrong_quantity = payload(events=[purchase(1, [
        {'sku': 'ok', 'quantity': 1, 'price': 1.25},
        {'sku': 'bad', 'quantity': '2', 'price': 2.5}])])[0]
    failure('nested-array-element-type', wrong_quantity,
            error(('key', 'events'), ('named', 'Array'), ('index', 0),
                  ('key', 'items'), ('named', 'Array'), ('index', 1),
                  ('key', 'quantity'), expected='Number'))
    wrong_array = payload()[0]
    wrong_array['users'] = {}
    failure('wrong-array-container', wrong_array,
            error(('key', 'users'), expected='Array'))
    failure('first-error-priority',
            {'version': 'bad', 'users': False, 'next': 12,
             'events': [{'tag': 'unknown'}, {'tag': 1}]},
            error(('key', 'events'), ('named', 'Array'), ('index', 0),
                  expected='Event tag'))
    success('optional-fields-missing', payload(
        users=[user(0, 'missing', True, [], profile('Paris', [])),
               user(1, 'missing-profile', False, [])], events=[view('/')]))
    success('optional-fields-null', payload(
        users=[user(0, 'null', True, [], profile('Paris', [], None)),
               user(1, 'null-profile', False, [], None)],
        events=[view('/', None)], next_page=None))
    return cases


def documents():
    cases = make_cases()
    corpus = [{'name': name, 'contents': compact(source), 'benchmark': timed}
              for name, source, _, timed in cases]
    expected = {'modules': len(cases), 'timed_cases': sum(case[3] for case in cases),
                'names': [case[0] for case in cases],
                'fingerprints': [fingerprint(case[2]) for case in cases],
                'json_fingerprints': [fingerprint(case[1]) for case in cases]}
    return {'corpus.json': corpus, 'expected.json': expected}


class FixtureTests(unittest.TestCase):
    def test_frozen_documents(self):
        for name, value in documents().items():
            with self.subTest(name=name):
                self.assertEqual(json.loads((FIXTURES / name).read_text()), value)

    def test_case_coverage(self):
        cases = make_cases()
        self.assertEqual(len(cases), 17)
        self.assertEqual(sum(case[3] for case in cases), 5)
        self.assertEqual(sum('error' in case[2] for case in cases), 10)
        self.assertEqual(len({case[0] for case in cases}), len(cases))
        self.assertTrue(all('value' in outcome for _, _, outcome, timed in cases if timed))
        size = sum(len(compact(case[1]).encode()) for case in cases)
        self.assertGreater(size, 500000)
        self.assertLess(size, 1000000)

    def test_full_result_and_error_mutations_change_hash(self):
        cases = make_cases()
        nested = copy.deepcopy(cases[2][2])
        nested['value']['events'][321]['items'][2]['price'] += 0.125
        self.assertNotEqual(fingerprint(nested), fingerprint(cases[2][2]))
        failure = next(case[2] for case in cases if case[0] == 'nested-array-element-type')
        changed = {'error': failure['error'].replace('index 1:', 'index 0:')}
        self.assertNotEqual(fingerprint(changed), fingerprint(failure))

    def test_canonical_cross_runtime_conventions(self):
        self.assertEqual(fingerprint({'b': 1.0, 'a': 0.0}),
                         fingerprint({'a': 0, 'b': 1}))
        escaped = b'{"x":"\\u003c\\u003e\\u0026\\u2028\\u2029"}'
        self.assertEqual(fingerprint({'x': '<>&\u2028\u2029'}),
                         hashlib.sha256(escaped).hexdigest())


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--write', action='store_true', help='regenerate frozen fixtures')
    args = parser.parse_args()
    if args.write:
        FIXTURES.mkdir(parents=True, exist_ok=True)
        for name, value in documents().items():
            (FIXTURES / name).write_text(json.dumps(value, ensure_ascii=False, indent=2) + '\n')
    unittest.main(argv=[__file__])
