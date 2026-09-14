"""A bounded, typed ownership IR experiment, not a Purust modification.

Every output stage serializes its actual instruction graph before Rust emission.
All functions use the same lowering/passes; no name-specific Rust rewrite.
"""
from pathlib import Path
import copy
import json
import subprocess

HERE = Path(__file__).resolve().parent
BUILD = HERE / 'build'
STAGES = ['naive', 'pushdown', 'precise', 'fusion', 'specialized_drop', 'reuse', 'retained_fields']
LAYOUTS = {
    'List': {'fields': ['i64', 'ptr'], 'scalar': 0, 'children': [1]},
    'Tree': {'fields': ['ptr', 'i64', 'ptr'], 'scalar': 1, 'children': [0, 2]},
    'TreeAlt': {'fields': ['i64', 'ptr', 'ptr'], 'scalar': 0, 'children': [1, 2]},
}

def field(i): return ['field', i]
def scalar(op, *args): return [op, *args]
def call(value): return ['call', value]
def build(fields): return ['build', *fields]


def programs():
    result = []
    for name, layout in LAYOUTS.items():
        key = field(layout['scalar'])
        mapped = [call(field(i)) if ty == 'ptr' else scalar('add', field(i), ['int', 1])
                  for i, ty in enumerate(layout['fields'])]
        result.append({'name': name.lower() + '_map', 'layout': name, 'body': build(mapped)})
        if name == 'List':
            keep = [key, call(field(1))]
            body = ['if', scalar('even', key), build(keep), call(field(1))]
            result.append({'name': 'list_filter', 'layout': name, 'body': body})
        else:
            changed = [scalar('add', field(i), ['int', 1]) if ty == 'i64' else field(i)
                       for i, ty in enumerate(layout['fields'])]
            result.append({'name': name.lower() + '_update', 'layout': name, 'body': build(changed)})
            # A branch reads only one recursive field: eager dup can be sunk.
            result.append({'name': name.lower() + '_choose', 'layout': name,
                           'body': ['if', scalar('even', key), field(layout['children'][0]), field(layout['children'][1])]})
    return result


def lower(program):
    layout = LAYOUTS[program['layout']]
    types = {'n': 'ptr', **{f'f{i}': ty for i, ty in enumerate(layout['fields'])}}
    counter = [0]
    def temp(ty):
        counter[0] += 1
        name = f'v{counter[0]}'
        types[name] = ty
        return name
    def expr(node, ops):
        tag, *args = node
        if tag == 'field': return f'f{args[0]}'
        if tag == 'int':
            out = temp('i64'); ops.append({'op': 'const', 'out': out, 'value': args[0]}); return out
        if tag in ['add', 'even']:
            values = [expr(arg, ops) for arg in args]
            out = temp('i64' if tag == 'add' else 'bool')
            ops.append({'op': tag, 'out': out, 'args': values}); return out
        assert tag in ['call', 'build']
        values = [expr(arg, ops) for arg in args]
        passed = []
        for value in values:
            if types[value] == 'ptr':
                duplicate = temp('ptr')
                ops.append({'op': 'dup', 'out': duplicate, 'args': [value]})
                passed.append(duplicate)
            else: passed.append(value)
        out = temp('ptr')
        ops.append({'op': tag, 'out': out, 'args': passed})
        return out
    def block(node):
        ops = []
        if node[0] == 'if':
            test = expr(node[1], ops)
            term = {'op': 'if', 'test': test, 'yes': block(node[2]), 'no': block(node[3])}
        else:
            value = expr(node, ops)
            out = temp('ptr')
            ops.append({'op': 'dup', 'out': out, 'args': [value]})
            term = {'op': 'return', 'value': out}
        return {'ops': ops, 'term': term}
    body = block(program['body'])
    body['ops'] = [{'op': 'dup_field' if ty == 'ptr' else 'copy_field', 'out': f'f{i}', 'args': ['n'], 'index': i}
                   for i, ty in enumerate(layout['fields'])] + body['ops']
    return {'name': program['name'], 'layout': program['layout'], 'types': types,
            'entry': {'op': 'borrow_match', 'input': 'n'}, 'body': body,
            'source_body': program['body'], 'empty': {'op': 'new_empty'}}


def used(block):
    result = set()
    for op in block['ops']:
        result.update(op.get('args', []))
    term = block['term']
    if term['op'] == 'return': result.add(term['value'])
    else: result |= {term['test']} | used(term['yes']) | used(term['no'])
    return result


def clean_drops(block):
    block['ops'] = [op for op in block['ops'] if op['op'] != 'drop']
    if block['term']['op'] == 'if':
        clean_drops(block['term']['yes']); clean_drops(block['term']['no'])


def consuming(op, types):
    return {v for v in op.get('args', []) if types[v] == 'ptr'} if op['op'] in ['call', 'build', 'reuse', 'move', 'drop'] else set()


def drops(block, owners, types, precise):
    """Place explicit releases lexically or after the last use on each path."""
    ops = block['ops']; output = []; active = list(owners)
    term = block['term']
    for index, op in enumerate(ops):
        output.append(op)
        active = [v for v in active if v not in consuming(op, types)]
        if 'out' in op and types[op['out']] == 'ptr': active.append(op['out'])
        if precise:
            future = used({'ops': ops[index+1:], 'term': term})
            dead = [v for v in active if v not in future]
            output.extend({'op': 'drop', 'args': [v]} for v in reversed(dead))
            active = [v for v in active if v not in dead]
    if term['op'] == 'if':
        for arm in ['yes', 'no']: drops(term[arm], active, types, precise)
    else:
        output.extend({'op': 'drop', 'args': [v]} for v in reversed(active) if v != term['value'])
    block['ops'] = output


def pushdown(block):
    """Sink projected dup operations into only the branches that read them."""
    term = block['term']
    if term['op'] != 'if': return
    sunk = []
    for op in block['ops']:
        if op['op'] == 'dup_field' and not any(op['out'] in other.get('args', []) for other in block['ops']) and term['test'] != op['out']:
            for arm in ['yes', 'no']:
                if op['out'] in used(term[arm]): term[arm]['ops'].insert(0, copy.deepcopy(op))
            sunk.append(op)
    block['ops'] = [op for op in block['ops'] if op not in sunk]
    pushdown(term['yes']); pushdown(term['no'])


def fuse(block):
    ops = block['ops']; result = []; index = 0
    while index < len(ops):
        op = ops[index]
        if op['op'] == 'dup' and index+1 < len(ops) and ops[index+1] == {'op': 'drop', 'args': op['args']}:
            result.append({**op, 'op': 'move'}); index += 2
        else: result.append(op); index += 1
    block['ops'] = result
    if block['term']['op'] == 'if':
        fuse(block['term']['yes']); fuse(block['term']['no'])


def remove_projection(block):
    block['ops'] = [op for op in block['ops'] if op['op'] not in ['dup_field', 'copy_field'] and not (op['op'] == 'drop' and op['args'] == ['n'])]
    if block['term']['op'] == 'if':
        remove_projection(block['term']['yes']); remove_projection(block['term']['no'])


def reuse(block):
    # DSL branches contain at most one output constructor, so each token has
    # exactly one consumer or a release on every return path.
    for op in block['ops']:
        if op['op'] == 'build': op['op'] = 'reuse'; op['token'] = 'cell'
    if block['term']['op'] == 'if':
        reuse(block['term']['yes']); reuse(block['term']['no'])
    else:
        block['ops'].append({'op': 'drop_token', 'token': 'cell'})


def retained(program):
    body = program['source_body']; layout = LAYOUTS[program['layout']]
    if body[0] != 'build': return
    fields = body[1:]
    unchanged = [i for i, ty in enumerate(layout['fields']) if ty == 'ptr' and fields[i] == ['field', i]]
    if len(unchanged) != len(layout['children']): return
    scalar_fields = [i for i, ty in enumerate(layout['fields']) if ty == 'i64']
    def scalar_only(expr): return expr[0] in ['field', 'int', 'add'] and all(scalar_only(v) for v in expr[1:] if isinstance(v, list))
    if not all(scalar_only(fields[i]) for i in scalar_fields): return
    program['entry'] = {'op': 'retain_fields', 'input': 'n', 'unchanged_fields': unchanged,
                        'staged_copy_writes': [{'index': i, 'expr': fields[i]} for i in scalar_fields],
                        'shared_or_weak': copy.deepcopy(program['entry'])}


def pipeline(source):
    base = lower(source); stages = {}
    naive = copy.deepcopy(base); drops(naive['body'], ['n'], naive['types'], False); stages['naive'] = naive
    p = copy.deepcopy(base); pushdown(p['body']); drops(p['body'], ['n'], p['types'], False); stages['pushdown'] = copy.deepcopy(p)
    clean_drops(p['body']); drops(p['body'], ['n'], p['types'], True); stages['precise'] = copy.deepcopy(p)
    fuse(p['body']); stages['fusion'] = copy.deepcopy(p)
    p['entry'] = {'op': 'unpack_consuming', 'input': 'n',
                  'unique': ['detach_constructor_fields', 'release_parent_payload'],
                  'shared': ['dup_constructor_fields', 'drop_parent_owner']}
    remove_projection(p['body']); clean_drops(p['body'])
    field_owners = [f'f{i}' for i, ty in enumerate(LAYOUTS[p['layout']]['fields']) if ty == 'ptr']
    drops(p['body'], field_owners, p['types'], True); fuse(p['body']); stages['specialized_drop'] = copy.deepcopy(p)
    p['entry'] = {'op': 'reset', 'input': 'n', 'token': 'cell', 'requires': 'Rc::get_mut excludes strong aliases and Weak',
                  'unique': ['extract_fields_with_empty_sentinel', 'retain_cell_token'],
                  'shared_or_weak': ['unpack_consuming', 'empty_token']}
    p['empty'] = {'op': 'reuse_empty', 'token': 'cell'}; reuse(p['body']); stages['reuse'] = copy.deepcopy(p)
    retained(p); stages['retained_fields'] = copy.deepcopy(p)
    return stages


def emit_ops(block, program):
    lines = []; name = program['layout']; fn = program['name']
    for op in block['ops']:
        kind = op['op']; out = op.get('out'); args = op.get('args', [])
        if kind == 'dup_field': value = f'r{op["index"]}.clone()'
        elif kind == 'copy_field': value = f'*r{op["index"]}'
        elif kind == 'const': value = f'{op["value"]}_i64'
        elif kind == 'add': value = f'{args[0]}.wrapping_add({args[1]})'
        elif kind == 'even': value = f'({args[0]} % 2 == 0)'
        elif kind == 'dup': value = args[0] + '.clone()'
        elif kind == 'move': value = args[0]
        elif kind == 'call': value = f'{fn}({args[0]})'
        elif kind == 'build': value = f'Rc::new({name}::Node({", ".join(args)}))'
        elif kind == 'reuse': value = f'rebuild_{name}({name}::Node({", ".join(args)}), cell.take())'
        elif kind == 'drop': lines.append(f'drop({args[0]});'); continue
        elif kind == 'drop_token': lines.append('drop(cell.take());'); continue
        else: raise ValueError(kind)
        lines.append(f'let {out} = {value};')
    term = block['term']
    if term['op'] == 'return': lines.append(term['value'])
    else: lines.append(f'if {term["test"]} {{\n{emit_ops(term["yes"], program)}\n}} else {{\n{emit_ops(term["no"], program)}\n}}')
    return '\n'.join(lines)


def emit(program):
    name = program['layout']; fields = LAYOUTS[name]['fields']; entry = program['entry']; pre = []
    if entry['op'] == 'retain_fields':
        def scalar_code(expr):
            if expr[0] == 'field': return f'*r{expr[1]}'
            if expr[0] == 'int': return f'{expr[1]}_i64'
            return f'({scalar_code(expr[1])}).wrapping_add({scalar_code(expr[2])})'
        refs = ', '.join(f'r{i}' for i in range(len(fields)))
        pre.append(f'if let Some({name}::Node({refs})) = Rc::get_mut(&mut n) {{')
        for write in entry['staged_copy_writes']: pre.append(f'let s{write["index"]} = {scalar_code(write["expr"])};')
        for write in entry['staged_copy_writes']: pre.append(f'*r{write["index"]} = s{write["index"]};')
        pre.append('event(7); return n; }'); entry = entry['shared_or_weak']
    if entry['op'] == 'borrow_match':
        subject = 'n.as_ref()'; pattern = ', '.join(f'r{i}' for i in range(len(fields)))
        empty = f'{{ drop(n); Rc::new({name}::Empty) }}'
    else:
        pattern = ', '.join(f'f{i}' for i in range(len(fields)))
        unpack = 'match Rc::try_unwrap(n) { Ok(value) => value, Err(owner) => { let value = owner.as_ref().clone(); drop(owner); value } }'
        if entry['op'] == 'reset':
            pre.append(f'let (payload, mut cell) = if let Some(slot) = Rc::get_mut(&mut n) {{ event(8); (std::mem::replace(slot, {name}::Empty), Some(n)) }} else {{ ({unpack}, None) }};')
            empty = f'rebuild_{name}({name}::Empty, cell.take())'
        else:
            pre.append(f'let payload = {unpack};'); empty = f'Rc::new({name}::Empty)'
        subject = 'payload'
    return f'fn {program["name"]}(mut n: Rc<{name}>) -> Rc<{name}> {{\n' + '\n'.join(pre) + f'\nmatch {subject} {{ {name}::Empty => {empty}, {name}::Node({pattern}) => {{\n{emit_ops(program["body"], program)}\n}} }}\n}}\n'


def check_program(program):
    """Reject ill-typed operands, use-after-move and leaked owned references."""
    types = program['types']; layout = LAYOUTS[program['layout']]
    entry = program['entry']
    if entry['op'] == 'retain_fields': entry = entry['shared_or_weak']
    borrowed = entry['op'] == 'borrow_match'
    values = {'n'} if borrowed else {f'f{i}' for i in range(len(layout['fields']))}
    owners = {v for v in values if types[v] == 'ptr'}
    def check(block, values, owners, token):
        values = set(values); owners = set(owners)
        for op in block['ops']:
            kind = op['op']; args = op.get('args', [])
            assert all(v in values for v in args), (program['name'], 'use after move', op)
            if kind in ['dup', 'move', 'drop', 'call']: assert all(types[v] == 'ptr' for v in args)
            if kind in ['dup_field', 'copy_field']:
                assert args == ['n'] and borrowed
                assert types[op['out']] == layout['fields'][op['index']]
            if kind in ['add', 'even']: assert all(types[v] == 'i64' for v in args)
            if kind in ['build', 'reuse']:
                assert [types[v] for v in args] == layout['fields']
                if kind == 'reuse': assert token; token = False
            if kind == 'drop_token': token = False
            removed = consuming(op, types)
            assert removed <= owners, ('non-owned consumption', op)
            owners -= removed; values -= removed
            if 'out' in op:
                assert op['out'] not in values, ('SSA redefinition', op)
                values.add(op['out'])
                if types[op['out']] == 'ptr': owners.add(op['out'])
        term = block['term']
        if term['op'] == 'if':
            assert term['test'] in values and types[term['test']] == 'bool'
            check(term['yes'], values, owners, token); check(term['no'], values, owners, token)
        else:
            assert term['value'] in owners and owners == {term['value']}, ('unreleased owner', program['name'], owners)
            assert not token, 'reset token escapes return'
    check(program['body'], values, owners, entry['op'] == 'reset')


def layout_code(name, layout):
    types = ['Rc<' + name + '>' if ty == 'ptr' else ty for ty in layout['fields']]
    return f'''#[derive(Clone)]
enum {name} {{ Empty, Node({', '.join(types)}) }}
fn rebuild_{name}(payload: {name}, cell: Option<Rc<{name}>>) -> Rc<{name}> {{
    if let Some(mut cell) = cell {{ *Rc::get_mut(&mut cell).unwrap() = payload; event(6); cell }} else {{ Rc::new(payload) }}
}}
'''


def input_value(name, variant=0):
    layout = LAYOUTS[name]
    def node(key, children):
        return tuple(key if ty == 'i64' else children[layout['children'].index(i)] for i, ty in enumerate(layout['fields']))
    if name == 'List':
        value = None
        for key in range(64+variant, variant, -1): value = node(key, [value])
        return value
    def tree(depth, key):
        return None if depth == 0 else node(key, [tree(depth-1, key*2), tree(depth-1, key*2+1)])
    return tree(4, 2+variant)


def oracle(program, value):
    if value is None: return None
    def evaluate(expr):
        kind, *args = expr
        if kind == 'field': return value[args[0]]
        if kind == 'int': return args[0]
        if kind == 'add': return (evaluate(args[0]) + evaluate(args[1]) + 2**63) % 2**64 - 2**63
        if kind == 'even': return evaluate(args[0]) % 2 == 0
        if kind == 'call': return oracle(program, evaluate(args[0]))
        if kind == 'if': return evaluate(args[1] if evaluate(args[0]) else args[2])
        if kind == 'build': return tuple(evaluate(arg) for arg in args)
        raise ValueError(kind)
    return evaluate(program['body'])


def value_text(name, value):
    if value is None: return 'E'
    return 'N(' + ','.join(str(v) if ty == 'i64' else value_text(name, v) for ty, v in zip(LAYOUTS[name]['fields'], value)) + ')'


def value_rust(name, value):
    if value is None: return f'Rc::new({name}::Empty)'
    fields = [str(v) if ty == 'i64' else value_rust(name, v) for ty, v in zip(LAYOUTS[name]['fields'], value)]
    return f'Rc::new({name}::Node({", ".join(fields)}))'


def harness_code():
    code = []
    for name, layout in LAYOUTS.items():
        names = [f'a{i}' for i in range(len(layout['fields']))]
        values = [f'{v}.to_string()' if ty == 'i64' else f'view_{name}({v}.as_ref())' for ty, v in zip(layout['fields'], names)]
        pattern = ', '.join(names)
        code.append(f'''fn view_{name}(v:&{name})->String {{ match v {{ {name}::Empty=>"E".into(), {name}::Node({pattern})=>format!("N({','.join('{}' for _ in names)})",{','.join(values)}) }} }}
fn make_{name}()->Rc<{name}> {{ {value_rust(name, input_value(name))} }}
fn make_alt_{name}()->Rc<{name}> {{ {value_rust(name, input_value(name, 1))} }}
fn first_{name}(v:&Rc<{name}>)->Rc<{name}> {{ match v.as_ref() {{ {name}::Node({pattern})=>{names[layout['children'][0]]}.clone(), _=>panic!("empty fixture") }} }}
''')
    code.append('fn main() { assert_eq!(std::mem::size_of::<Rc<List>>(), std::mem::size_of::<std::rc::Rc<List>>());')
    for program in programs():
        name = program['layout']; fn = program['name']
        expected = value_text(name, oracle(program, input_value(name)))
        expected_alt = value_text(name, oracle(program, input_value(name, 1)))
        code.append(f'''for variant in 0..2 {{ for mode in 0..6 {{
clear_counts();
let original = if variant==0 {{ make_{name}() }} else {{ make_alt_{name}() }};
let snapshot = view_{name}(original.as_ref());
let saved = (mode==1 || mode==4).then(|| original.clone());
let child = (mode==2 || mode==4).then(|| first_{name}(&original));
let child_snapshot = child.as_ref().map(|c|view_{name}(c.as_ref()));
let weak = (mode==3 || mode==4).then(||Rc::downgrade(&original));
let weak_child = (mode==5).then(||{{let c=first_{name}(&original); (Rc::downgrade(&c),view_{name}(c.as_ref()))}});
let begin = counts();
let result = {fn}(original);
let end = counts();
assert_eq!(view_{name}(result.as_ref()), if variant==0 {{{json.dumps(expected)}}} else {{{json.dumps(expected_alt)}}}, "oracle {fn}");
if let Some(ref old)=saved {{ assert_eq!(view_{name}(old.as_ref()), snapshot); }}
if let Some(ref child)=child {{ assert_eq!(view_{name}(child.as_ref()), *child_snapshot.as_ref().unwrap()); }}
if let Some(old)=weak.as_ref().and_then(|w|w.upgrade()) {{ assert_eq!(view_{name}(&old),snapshot,"Weak observed mutation"); }}
if let Some((ref weak,ref old_shape))=weak_child {{ if let Some(old)=weak.upgrade() {{assert_eq!(view_{name}(&old),*old_shape,"child Weak observed mutation");}} }}
drop(result);drop(saved);drop(child);
if let Some(ref weak)=weak {{ assert!(weak.upgrade().is_none(),"owner escaped"); }}
if let Some((ref weak,_))=weak_child {{assert!(weak.upgrade().is_none(),"child owner escaped");}}
drop(weak);drop(weak_child);balance();
let delta:Vec<_>=end.iter().zip(begin).map(|(a,b)|a-b).collect();
println!("CASE {fn}-{{variant}} {{mode}} {{delta:?}}");
println!("PEAK {fn}-{{variant}} {{mode}} {{}}", peak());
}} }}
''')
    code.append('println!("PASS 96 oracle/ownership scenarios; all owner and payload balances"); }')
    return '\n'.join(code)


def prepare():
    BUILD.mkdir(parents=True, exist_ok=True)
    all_programs = programs(); graphs = [pipeline(p) for p in all_programs]
    (HERE / 'source-ir.json').write_text(json.dumps({'layouts': LAYOUTS, 'programs': all_programs}, indent=2) + '\n')
    for stage in STAGES:
        ir = [graph[stage] for graph in graphs]
        for program in ir: check_program(program)
        (HERE / f'{stage}.ir.json').write_text(json.dumps(ir, indent=2) + '\n')
        code = '#![allow(warnings)]\n' + (HERE / 'runtime.rs').read_text()
        code += ''.join(layout_code(name, layout) for name, layout in LAYOUTS.items())
        code += ''.join(emit(p) for p in ir)
        code += harness_code()
        (BUILD / f'{stage}.rs').write_text(code)
    rejected = []
    def rejects(name, program):
        try: check_program(program)
        except AssertionError: rejected.append(name)
        else: raise AssertionError('invalid IR accepted: ' + name)
    invalid = copy.deepcopy(graphs[0]['fusion'])
    next(op for op in invalid['body']['ops'] if op['op'] == 'move')['args'] = ['n']
    rejects('use_after_drop', invalid)
    invalid = copy.deepcopy(graphs[0]['fusion'])
    constructor = next(op for op in invalid['body']['ops'] if op['op'] == 'build')
    constructor['args'][0] = constructor['args'][1]
    rejects('pointer_in_scalar_field', invalid)
    invalid = copy.deepcopy(graphs[0]['naive'])
    invalid['body']['ops'] = [op for op in invalid['body']['ops'] if op != {'op':'drop','args':['n']}]
    rejects('leaked_owner', invalid)
    invalid = copy.deepcopy(next(g['reuse'] for g in graphs if g['reuse']['name'] == 'tree_choose'))
    for arm in ['yes', 'no']:
        invalid['body']['term'][arm]['ops'] = [op for op in invalid['body']['term'][arm]['ops'] if op['op'] != 'drop_token']
    rejects('escaped_reset_token', invalid)
    (HERE / 'ir-checks.json').write_text(json.dumps({'typed_linear_programs_checked': len(graphs)*len(STAGES), 'invalid_ir_rejected': rejected}, indent=2)+'\n')
    print('Emitted 7 explicit IR stages and Rust programs for 8 functions / 3 layouts; no execution.')


def validate():
    results = {}
    for stage in STAGES:
        binary = BUILD / stage
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(BUILD / f'{stage}.rs'), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        (BUILD / f'{stage}.log').write_text(output)
        results[stage] = output.splitlines()
        print(stage, output.splitlines()[-1], flush=True)
        native = BUILD / f'{stage}-native'
        subprocess.run(['rustc', '--edition=2021', '--cfg', 'untracked', '-C', 'opt-level=1', str(BUILD / f'{stage}.rs'), '-o', str(native)], check=True)
        native_output = subprocess.check_output([str(native)], text=True)
        (BUILD / f'{stage}-native.log').write_text(native_output)
        assert native_output.splitlines()[-1] == output.splitlines()[-1]
        results[stage + '_native'] = {'passed': True, 'representation': 'std::rc::Rc', 'cases': 96}
    (HERE / 'validation-counts.json').write_text(json.dumps(results, indent=2) + '\n')


if __name__ == '__main__':
    import argparse
    parser = argparse.ArgumentParser(); parser.add_argument('mode', choices=['prepare', 'validate']); args = parser.parse_args()
    {'prepare': prepare, 'validate': validate}[args.mode]()
