#!/usr/bin/env python3
"""Describe the private Wasm batch callback using its supported function ABI."""
from pathlib import Path
import sys


def adapt(purs, js):
    declaration = 'foreign import measureBatch :: Int -> Int -> Effect Int -> Effect Number'
    if purs.count(declaration) != 1 or js.count('result = act();') != 1:
        raise RuntimeError('Unexpected shared batch harness; cannot verify the Wasm ABI adapter')
    # Effect Int is represented by a suspended unit->Int closure. The current
    # Wasm FFI lowers an Effect parameter to i32, whereas a function parameter
    # uses the correct closure reference. Its unused unit argument is boxed 0.
    purs = purs.replace('import Prelude\n', 'import Prelude\nimport Unsafe.Coerce (unsafeCoerce)\n', 1)
    purs = purs.replace(declaration,
                        'foreign import measureBatch :: Int -> Int -> (Int -> Int) -> Effect Number')
    for expected, count in [('expected', 1), ('out', 10)]:
        call = f'measureBatch iterations {expected} act'
        if purs.count(call) != count:
            raise RuntimeError(f'Expected {count} calls to {call}')
        purs = purs.replace(call, f'measureBatch iterations {expected} (unsafeCoerce act)')
    return purs, js.replace('result = act();', 'result = act(0);', 1)


if __name__ == '__main__':
    purs_path, js_path = map(Path, sys.argv[1:])
    purs, js = adapt(purs_path.read_text(), js_path.read_text())
    purs_path.write_text(purs)
    js_path.write_text(js)
    print('[wasm] Adapted the private batch callback to the Wasm function ABI.')
