#!/usr/bin/env python3
"""Count actual typed-plan paths in a disposable, untimed diagnostic build."""
import argparse
import json
import shutil
from pathlib import Path

from run import SPECIALIZED, FIXTURES, call, environment, generated_sources, sha


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--source', type=Path, required=True)
    parser.add_argument('--workspace', type=Path, required=True)
    args = parser.parse_args()
    source, work = args.source.resolve(), args.workspace.resolve()
    work.mkdir(exist_ok=False)
    (work / 'logs').mkdir()
    shutil.copytree(source / 'output', work / 'output')
    ffi = work / 'output/purescript/Data_Argonaut_Decode_Internal_Record_ffi.go'
    text = ffi.read_text()
    hooks = {
        'func (kind *fieldKind) decodeValue(': 'zzCoverage["kind_" + kind.tag]++',
        'func (plan *recordDecodePlan) decodeCustom(': 'zzCoverage["custom_element"]++',
        'func (plan *recordDecodePlan) decodeAny(': 'zzCoverage["direct_record"]++',
        'func (plan *recordDecodePlan) decode(': 'zzCoverage["generic_record"]++',
        'func (plan *recordDecodePlan) record(': 'zzCountSize("generic_record_size", len(plan.fields))',
    }
    if 'func (plan *recordDecodePlan) recordSmall(' in text:
        hooks['func (plan *recordDecodePlan) recordSmall('] = 'zzCountSize("compact_record_size", len(plan.fields))'
    if 'func decodeJsonObjectDirect(' in text:
        hooks['func decodeJsonObjectDirect('] = 'zzCoverage["identity_object"]++'
    if 'func Data_Argonaut_Decode_Internal_Record_TypedObject(' in text:
        hooks['func Data_Argonaut_Decode_Internal_Record_TypedObject('] = 'zzCoverage["object_decoder_construction"]++'
    if 'func Data_Argonaut_Decode_Internal_Record_BorrowObject(' in text:
        hooks['func Data_Argonaut_Decode_Internal_Record_BorrowObject('] = 'zzCoverage["borrow_object_attempt"]++'
        marker = 'return gopurs_runtime.Apply(identity.right, json)'
        if text.count(marker) != 1:
            raise SystemExit('Expected exactly one borrowed object return')
        text = text.replace(marker, 'zzCoverage["borrowed_identity_object"]++\n\t\t\t' + marker)
    if 'func getFieldDirect(' in text:
        hooks['func getFieldDirect('] = 'zzCoverage["field_accessor_attempt"]++'
        marker = '\tplan := tag.errors\n'
        if text.count(marker) != 1:
            raise SystemExit('Expected exactly one fused accessor body')
        text = text.replace(marker, marker + '\tzzCoverage["fused_field_accessor"]++\n')
    for marker, statement in hooks.items():
        if text.count(marker) != 1:
            raise SystemExit(f'Expected exactly one instrumentation point: {marker}')
        position = text.index(' {', text.index(marker)) + 2
        text = text[:position] + '\n\t' + statement + text[position:]
    marker = 'result := gopurs_runtime.Apply2(field.step, '
    if text.count(marker) != 2:
        raise SystemExit('Expected the direct fallback and generic field calls')
    text = text.replace(marker, 'zzCoverage["untagged_field_step"]++\n\t\t' + marker)
    ffi.write_text(text)
    shutil.copyfile(SPECIALIZED / 'coverage_main.go', work / 'output/purescript/zz_coverage.go')
    (work / 'output/main/main.go').write_text('package main\nimport "gopurs/output/purescript"\nfunc main() { purescript.ZzCoverageMain() }\n')
    env = environment({'GOMAXPROCS': '14'})
    call(['go', 'build', '-pgo=off', '-o', work / 'coverage', './main'],
         work / 'output', work / 'logs/build.log', env)
    env = environment({'DIAG_CORPUS': str(FIXTURES / 'corpus.json'), 'DIAG_ORACLE': str(FIXTURES / 'expected.json')})
    call([work / 'coverage'], work, work / 'counts.json', env)
    provenance = {'source_workspace': str(source), 'source_manifest': json.loads((source / 'manifest.json').read_text()),
                  'generated_sources': generated_sources(work), 'binary_sha256': sha(work / 'coverage'),
                  'instrumenter_sha256': sha(Path(__file__).resolve()), 'timing': False}
    (work / 'coverage-build.json').write_text(json.dumps(provenance, indent=2) + '\n')
    print((work / 'counts.json').read_text())


if __name__ == '__main__':
    main()
