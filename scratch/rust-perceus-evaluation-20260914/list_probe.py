"""A scratch-only consumed-pattern substitution in the actual ListOps output."""
from pathlib import Path
import hashlib
import json
import re

HERE = Path(__file__).resolve().parent
OUT = HERE / 'list'
FROZEN = HERE / 'build/frozen'
SOURCE = FROZEN / 'Purs_Test_ListOps/src/lib.rs'


def digest(text):
    return hashlib.sha256(text.encode()).hexdigest()


def transform(original, consume, instrument=False):
    # Two copies: public filterEvens and the PBO-inlined filter inside sumEvens.
    pattern = re.compile(r'(    } else if /\* OpIsTag Debug: Test_ListOps_Cons -> Func \*/ matches!\(\(/\*[^\n]*?\*/(purs_local_[23])\)\.as_ref\(\), crate::List::Cons\(\.\.\)\) \{\n)(        if /\* Typed bool <- bool : PrimOp)')
    matches = list(pattern.finditer(original))
    assert len(matches) == 2, len(matches)
    result = original
    for match in reversed(matches):
        owner = match[2]
        begin = match.start()
        end = original.index('    };\n        }\n    }', match.end())
        block = original[begin:end]
        prefix = ''
        if instrument:
            prefix += f'        probe_filter_visit(std::rc::Rc::strong_count(&{owner}) == 1);\n'
        if consume:
            prefix += f'        let crate::List::Cons(probe_head, probe_tail) = std::rc::Rc::unwrap_or_clone({owner}) else {{ unreachable!() }};\n'
            head = '{ match (' + owner + ').as_ref() { crate::List::Cons(ref f, ..) => f.clone(), _ => unreachable!() } }'
            tail = '{ match (' + owner + ').as_ref() { crate::List::Cons(_, ref f, ..) => f.clone(), _ => unreachable!() } }'
            assert block.count(head) == 2 and block.count(tail) == 2
            block = block.replace(head, 'probe_head.clone()', 1)
            block = block.replace(head, 'probe_head')
            block = block.replace(tail, 'probe_tail')
        block = block[:len(match[1])] + prefix + block[len(match[1]):]
        result = result[:begin] + block + result[end:]
    if instrument:
        result += '''
std::thread_local! { static PROBE_FILTER: std::cell::Cell<[u64;2]> = const { std::cell::Cell::new([0,0]) }; }
fn probe_filter_visit(unique: bool) { PROBE_FILTER.with(|v| { let mut n=v.get(); n[if unique {0} else {1}]+=1; v.set(n); }); }
pub fn probe_filter_counts() -> [u64;2] { PROBE_FILTER.with(|v| v.replace([0,0])) }
'''
    return result


def prepare():
    OUT.mkdir(exist_ok=True)
    original = SOURCE.read_text()
    variants = {'before': original, 'consumed': transform(original, True)}
    for name, source in variants.items():
        (OUT / f'ListOps-{name}.rs').write_text(source)
        (OUT / f'ListOps-{name}-counted.rs').write_text(transform(original, name == 'consumed', True))
    (OUT / 'metadata.json').write_text(json.dumps({
        'source': str(SOURCE), 'sha256': {name: digest(value) for name, value in variants.items()},
        'scope': 'Only the two generated filter loop bodies change. Range, fold, boxing, closures, recursion thunks, accumulator clones and allocator are unchanged.',
        'hypothesis': 'Consuming the matched node transfers head/tail if its strong count is one; shared nodes clone their payload. This tests one missing consumed-pattern case, not a complete Perceus pass.'
    }, indent=2) + '\n')
    print('Two actual filter loops transformed; all surrounding generated code retained')


if __name__ == '__main__':
    prepare()
