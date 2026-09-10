"""Validate exact emitted kernels with native Rc; no algorithm patch or timing."""
import sys
sys.dont_write_bytecode = True
import json
import re
import subprocess
import count


def main():
    sources = count.inputs()
    base = (count.PREVIOUS / 'count_harness.rs').read_text()
    base = base[base.index('fn validate('):].replace('tracked::Rc', 'std::rc::Rc')
    base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n', '', base)
    base = base.replace('fn main() {', 'fn original_checks() {')
    extra = (count.HERE.parent / 'rust-child-field-20260909/validation-extra.rs').read_text()
    black = extra[extra.index('fn b13_red_parent_cases()'):extra.index('pub fn extra_checks()')]
    black = black.replace('b13_red_parent_cases', 'post_black_parent_cases')
    black = black.replace('Tree::T(Color::R, left, 10, right)', 'Tree::T(Color::B, left, 10, right)')
    black = black.replace('Tree::T(Color::R, _, 10, _)', 'Tree::T(Color::B, _, 10, _)')
    black = black.replace('B13 red-parent cases', 'Post-call black-parent cases')
    result = {'method': 'Exact emitted kernel declarations extracted verbatim and compiled with native std::rc::Rc, rustc O1. Existing validation harnesses reused without algorithm patches.',
        'checks': 'Four rotations; 100000-key tree/depth22; 128 red-parent plus 128 black-parent sharing/weak combinations; 200 retained versions with i64 bounds; 512 mixed-sharing insertions with duplicates; ordering/red-black invariants; final weak expiry.',
        'variants': {}}
    for name, source in sources.items():
        path = count.BUILD / f'validate-{name}.rs'
        emitted = count.kernel(source['code'])
        path.write_text('#![allow(warnings)]\n' + emitted + base + extra + black
            + '\nfn main() { original_checks(); extra_checks(); post_black_parent_cases(); }\n')
        binary = count.BUILD / f'validate-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        result['variants'][name] = {key: value for key, value in source.items() if key != 'code'} | {
            'extracted_kernel_sha256': count.sha256(emitted.encode()), 'output': output.strip(), 'passed': True}
        print(name, output.strip(), flush=True)
    for source in sources.values():
        assert count.sha256(count.Path(source['path']).read_bytes()) == source['sha256'], 'Source changed during checks'
    (count.HERE / 'validation.json').write_text(json.dumps(result, indent=2) + '\n')


if __name__ == '__main__': main()
