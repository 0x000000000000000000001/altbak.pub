"""Deterministic reference-contract cases, independent of the C++ decoders."""
import copy
import json
import random


def application(frozen):
    cases = list(frozen)
    rng = random.Random(2026092501)
    sample = {'version': 1, 'next': None,
              'users': [{'id': 7, 'name': 'owned-é🙂', 'active': True,
                         'profile': {'city': 'Paris', 'note': 'kept', 'scores': [1.5, 2]},
                         'tags': ['a', 'b']}],
              'events': [{'tag': 'view', 'path': '/home', 'duration': 17},
                         {'tag': 'purchase', 'orderId': 8,
                          'items': [{'sku': 'sku', 'quantity': 3, 'price': 2.5}]}]}
    replacements = [None, False, True, 0, 1, -1, 2.25, 'wrong', [], {}, 'é🙂', [None]]

    def mutate(value):
        if isinstance(value, dict) and value:
            key = rng.choice(list(value))
            if isinstance(value[key], (dict, list)) and rng.randrange(2):
                mutate(value[key])
            elif rng.randrange(4) == 0:
                del value[key]
            else:
                value[key] = copy.deepcopy(rng.choice(replacements))
        elif isinstance(value, list) and value:
            index = rng.randrange(len(value))
            if isinstance(value[index], (dict, list)) and rng.randrange(2):
                mutate(value[index])
            else:
                value[index] = copy.deepcopy(rng.choice(replacements))

    for index in range(800):
        value = copy.deepcopy(sample)
        for _ in range(1 + rng.randrange(3)):
            mutate(value)
        cases.append({'name': 'application-' + str(index), 'contents': json.dumps(value)})
    return cases


def typed(frozen):
    cases = list(frozen)
    rng = random.Random(2026092502)
    span = {'start': [1, 1], 'end': [1, 2]}

    def ann(**extra):
        return {'sourceSpan': span, 'meta': None, 'type': None, **extra}

    def binding(identity):
        return {'bindingId': identity, 'maxUses': 1, 'hasEscapingUseContext': False}

    def var(name, identity=None):
        facts = {} if identity is None else {'variableUse': {'bindingId': identity, 'lastLocalUse': True}}
        return {'type': 'Var', 'annotation': ann(**facts), 'value': {'identifier': name}}

    def lam(name, identity, body):
        facts = {} if identity is None else {'bindingUsage': binding(identity)}
        return {'type': 'Abs', 'annotation': ann(**facts), 'argument': name, 'body': body}

    def local(name, identity, expr):
        return {'bindType': 'NonRec', 'identifier': name,
                'annotation': ann(**({} if identity is None else {'bindingUsage': binding(identity)})),
                'expression': expr}

    def module(expr=None, table=None):
        return {'moduleName': ['Fixture'], 'modulePath': 'Fixture.purs', 'sourceSpan': span,
                'imports': [], 'exports': ['identity'], 'reExports': {}, 'foreign': [], 'comments': [],
                'decls': [local('identity', None, expr or lam('x', 0, var('x', 0)))],
                'typeTable': table or []}

    def add(name, value):
        cases.append({'name': name, 'contents': json.dumps(value)})

    for index in range(600):
        # Nested scope identity, shadowing without facts, and globally repeated
        # IDs are varied independently of the expression's JSON field order.
        depth = rng.randrange(1, 8)
        frames = [(rng.choice(['x', 'y', 'z']), rng.choice([i, i, None, 0])) for i in range(depth)]
        name = rng.choice(['x', 'y', 'z', 'free'])
        identity = next((identity for ident, identity in reversed(frames) if ident == name), None)
        leaf = var(name, identity if index % 2 else rng.randrange(depth + 2))
        expr = leaf
        for name, identity in reversed(frames):
            expr = lam(name, identity, expr)
        add('scope-' + str(index), module(expr))

    for key, values in [('bindingId', [-1, 0.5, 2147483648, '0', None]),
                        ('maxUses', [-1, 0.5, '1', None, 0, 2147483648, 9007199254740992, 1e30]),
                        ('hasEscapingUseContext', [0, 'false', None, False, True])]:
        for value in values:
            m = module()
            m['decls'][0]['expression']['annotation']['bindingUsage'][key] = value
            add('usage-number-' + key + '-' + str(value), m)
    for value in [False, None, True, 0, 'true']:
        m = module()
        m['decls'][0]['expression']['body']['annotation']['variableUse']['lastLocalUse'] = value
        add('last-local-' + str(value), m)
    m = module(); m['decls'][0]['annotation']['bindingUsage'] = binding(9); add('global-facts', m)
    m = module(); m['imports'] = [{'moduleName': ['Other'], 'annotation': ann(bindingUsage=binding(9))}]; add('import-facts', m)
    m = module(); m['decls'][0]['expression']['body']['value']['moduleName'] = ['Fixture']; add('qualified-use', m)
    m = module(); m['decls'][0]['expression']['body']['annotation']['bindingUsage'] = binding(9); add('binding-on-variable', m)
    m = module(); m['decls'][0]['expression']['annotation']['variableUse'] = {'bindingId': 0}; add('variable-on-binding', m)
    m = module(); m['decls'].append(local('other', None, lam('y', 0, var('y', 0)))); add('global-duplicate-id', m)

    for recursive in [False, True]:
        for forward in [False, True]:
            bindings = [local('a', 1, var('b', 2) if forward else var('x', 0)),
                        local('b', 2, var('a', 1))]
            group = [{'bindType': 'Rec', 'binds': bindings}] if recursive else bindings
            expr = {'type': 'Let', 'annotation': ann(), 'binds': group, 'expression': var('b', 2)}
            add(f'let-{recursive}-{forward}', module(lam('x', 0, expr)))

    simple = {'binderType': 'VarBinder', 'annotation': ann(bindingUsage=binding(1)), 'identifier': 'y'}
    patterns = [simple,
                {'binderType': 'NamedBinder', 'annotation': ann(bindingUsage=binding(2)), 'identifier': 'alias', 'binder': simple},
                {'binderType': 'LiteralBinder', 'annotation': ann(), 'literal': {'literalType': 'ArrayLiteral', 'value': [simple]}},
                {'binderType': 'LiteralBinder', 'annotation': ann(), 'literal': {'literalType': 'ObjectLiteral', 'value': [['field', simple]]}},
                {'binderType': 'ConstructorBinder', 'annotation': ann(), 'typeName': {'identifier': 'T'}, 'constructorName': {'identifier': 'C'}, 'binders': [simple]}]
    for i, pattern in enumerate(patterns):
        for guarded in [False, True]:
            for valid in [False, True]:
                leaf = var('y', 1 if valid else 0)
                alt = {'binders': [pattern], 'isGuarded': guarded}
                alt.update({'expressions': [{'guard': var('x', 0), 'expression': leaf}]} if guarded else {'expression': leaf})
                expr = {'type': 'Case', 'annotation': ann(), 'caseExpressions': [var('x', 0)], 'caseAlternatives': [alt]}
                add(f'pattern-{i}-{guarded}-{valid}', module(lam('x', 0, expr)))

    # Forward/missing references and cycles must settle in the same ascending
    # order as PureScript, even for table entries no executable node mentions.
    for index in range(700):
        count = rng.randrange(1, 10)
        table = []
        ref = lambda: rng.randrange(-1, count + 2)
        for _ in range(count):
            kind = rng.randrange(12)
            entry = ['Int', 'Any', {'type': 'TypeVar', 'name': 'a'}, {'TypeVar': 'legacy'},
                     {'type': 'Array', 'element': ref()}, {'type': 'Record', 'row': ref()},
                     {'type': 'ForAll', 'vars': ['a'], 'body': ref()},
                     {'type': 'TypeApp', 'constructor': ref(), 'args': [ref()]},
                     {'type': 'Adt', 'fqn': ['Data', 'T'], 'args': [ref()]},
                     {'type': 'Func', 'args': [ref()], 'ret': ref()},
                     {'type': 'Row', 'fields': [{'label': 'field', 'type': ref()}], 'tail': ref()},
                     {'type': 'ConstrainedType', 'constraints': [{'fqn': ['Data', 'C'], 'args': [ref()]}], 'body': ref()}][kind]
            table.append(entry)
        m = module(table=table)
        if index % 9 == 0:
            m['typeTable'].append({'type': 'Unknown'})
        if index % 3:
            m['foreign'] = ['v' + str(i) for i in range(count)]
            m['foreignAnnotations'] = {name: ann(type=i) for i, name in enumerate(m['foreign'])}
        add('type-graph-' + str(index), m)
    return cases
