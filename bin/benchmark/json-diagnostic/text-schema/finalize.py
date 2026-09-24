#!/usr/bin/env python3
"""Archive and recheck the complete text-schema decoding campaign."""
import argparse, hashlib, importlib.util, json, shutil, statistics, subprocess
from pathlib import Path

root=Path(__file__).resolve().parents[4];workspace=root.parent
parser=argparse.ArgumentParser(description=__doc__)
parser.add_argument('--cycle',type=Path,required=True)
args=parser.parse_args();cycle=args.cycle.resolve()
sha=lambda p:hashlib.sha256(p.read_bytes()).hexdigest()
load=lambda p:json.loads(p.read_text())
def module(name,path):
    spec=importlib.util.spec_from_file_location(name,path);value=importlib.util.module_from_spec(spec);spec.loader.exec_module(value);return value
diagnostic=module('diagnostic',root/'bin/benchmark/json-diagnostic.py')
paired=module('paired',root/'bin/benchmark/json-diagnostic/paired.py')
before=load(cycle/'before.json');extra=load(cycle/'before-extra.json')
for section in [before['sources'],extra['sources'],load(cycle/'indexed-sources.json'),load(cycle/'object-slot-compiler.json')]:
    for relative,expected in section.items():assert sha(cycle/relative)==expected,relative
baseline=Path(before['baseline'])
assert sha(baseline/'benchmark')==before['baseline_binary_sha256']
assert sha(baseline/'manifest.json')==before['baseline_manifest_sha256']
current=cycle/'published';manifest=load(current/'manifest.json')
assert manifest['sources']==diagnostic.fingerprint('JsonDecoding')
assert manifest['binary_sha256']==sha(current/'benchmark')
assert manifest['js_sha256']==sha(current/'benchmark.mjs')

verified={}
oracle_path=root/'test/fixtures/json-decoding/expected.json';oracle=load(oracle_path)
for directory in sorted(cycle.glob('paired-*')):
    if not directory.is_dir():continue
    report=load(directory/'results.json')
    assert report['corpus_sha256']==sha(directory/'corpus.json')
    assert report['oracle_sha256']==sha(oracle_path)
    assert report['runner_sha256']==sha(root/'bin/benchmark/json-diagnostic/paired.py')
    outputs=0
    for mode,build in report['builds'].items():
        assert paired.provenance(Path(build['workspace']),'JsonDecoding')==build,(directory,mode)
        for index,result in enumerate(report['runs'][mode]):
            assert result==load(directory/f'{index:02}-{mode}.json')
            diagnostic.validate_result(result,oracle,'JsonDecoding')
            outputs+=oracle['timed_cases']*sum(len(p['samples']) for p in result['phases'].values())
        for phase in ['parse','decode','combined']:
            assert report['medians_us'][mode][phase]==statistics.median(r['phases'][phase]['time_us'] for r in report['runs'][mode])
            assert report['min_allocated_bytes_per_corpus'][mode][phase]==min(s['allocated_bytes'] for r in report['runs'][mode] for s in r['phases'][phase]['samples'])
    wins={phase:sum(c['phases'][phase]['time_us']<b['phases'][phase]['time_us'] for b,c in zip(report['runs']['baseline'],report['runs']['current'])) for phase in ['parse','decode','combined']}
    verified[directory.name]={'sampled_outputs':outputs,'winning_pairs':wins,'result_sha256':sha(directory/'results.json')}

for runtime in ['go','js']:
    directory=cycle/('official-final' if runtime=='go' else 'official-final-js');report=load(directory/'results.json')
    assert report['build']==manifest
    assert report['results'][runtime]==[load(p) for p in sorted(directory.glob(f'*-{runtime}.log'))]
    for result in report['results'][runtime]:diagnostic.validate_result(result,oracle,'JsonDecoding')
    for phase in ['parse','decode','combined']:
        assert report['medians_us'][runtime][phase]==statistics.median(r['phases'][phase]['time_us'] for r in report['results'][runtime])
        if runtime=='go':
            assert report['go_allocated_bytes_per_corpus'][phase]==min(s['allocated_bytes'] for r in report['results'][runtime] for s in r['phases'][phase]['samples'])
    verified[directory.name]={'sampled_outputs':sum(oracle['timed_cases']*sum(len(p['samples']) for p in r['phases'].values()) for r in report['results'][runtime])}

for name,build in [('validation-indexed','indexed'),('validation-slots','object-slots'),('validation-final','published'),('ownership-final',None),('ownership-published',None)]:
    directory=cycle/name;report=load(directory/'verification.json')
    assert report['status']=='passed'
    if build:assert report['source_manifest_sha256']==sha(cycle/build/'manifest.json')
    assert report['log_sha256']==sha(directory/'tests.log')
    assert report['tests_sha256']==sha(directory/'output/purescript'/('text_schema_test.go' if build else 'text_owned_test.go'))
    for relative,expected in report['generated_go'].items():assert sha(directory/relative)==expected
    if build:
        source_go={str(p.relative_to(cycle/build/'output')):sha(p) for p in (cycle/build/'output').rglob('*.go') if not p.name.endswith('_test.go')}
        tested_go={str(p.relative_to(directory/'output')):sha(p) for p in (directory/'output').rglob('*.go') if not p.name.endswith('_test.go')}
        assert source_go==tested_go
    if name in ['validation-slots','validation-final']:
        assert report['extra_corpus_sha256']==sha(directory/'extra-corpus.json')
        assert report['extra_expected_sha256']==sha(directory/'extra-expected.json')
    assert '\nPASS\n' in (directory/'tests.log').read_text()
    verified[name]={'status':'passed','go_files':len(report['generated_go'])}

compiler=workspace/'gopurs/gopurs'
for name,source in [('bootstrap-parity',cycle/'object-slot-compiler'),('bootstrap-final-parity',compiler)]:
    bootstrap=load(cycle/name/'result.json')
    assert sha(source/'bin/gopurs-native')==bootstrap['native_sha256']
    assert sha(source/'bin/gopurs.js')==bootstrap['node_sha256']
    for relative,expected in bootstrap['files'].items():
        assert sha(cycle/name/bootstrap.get('node_archive','node')/relative)==expected
        assert sha(Path(bootstrap['workspace'])/'output'/relative)==expected
    verified[name]={'go_files':bootstrap['matched']}

directory=cycle/'lifecycle-published';report=load(directory/'results.json')
expected=dict(zip(oracle['names'],oracle['fingerprints']))
for path,digest in report['source_sha256'].items():assert sha(Path(path))==digest
for mode,build in report['builds'].items():
    assert sha(directory/mode/'audit')==build['binary_sha256']
    assert sha(Path(build['workspace'])/'manifest.json')==build['source_manifest_sha256']
    for relative,digest in build['generated_go'].items():assert sha(directory/mode/relative)==digest
    for index,run in enumerate(report['runs'][mode]):
        assert run==load(directory/f'{index:02}-{mode}.json')
        assert run['first_fingerprints']==run['second_fingerprints']==[expected[n] for n in run['names']]
    initial=load(directory/f'init-{mode}.json')
    assert initial['first_fingerprints']==initial['second_fingerprints']==[expected[n] for n in initial['names']]
    assert report['initialization'][mode]['line'] in (directory/f'init-{mode}.log').read_text()
    for phase in ['construction','first_combined','second_combined','release_collection']:
        for key in ['time_us','allocated_bytes']:
            assert report['medians'][mode][phase][key]==statistics.median(r[phase][key] for r in report['runs'][mode])
verified['lifecycle']={'sampled_outputs':report['validated_outputs'],'subtotal_medians_us':report['subtotal_medians_us'],'subtotal_winning_pairs':report['subtotal_winning_pairs']}
phases=['construction','first_combined','second_combined','release_collection']
sums={mode:[sum(run[p]['time_us'] for p in phases) for run in runs] for mode,runs in report['runs'].items()}
assert sums==report['measured_interval_subtotals_us']
assert report['subtotal_medians_us']=={mode:statistics.median(values) for mode,values in sums.items()}
assert report['subtotal_winning_pairs']==sum(c<b for b,c in zip(sums['baseline'],sums['current']))

distribution=load(cycle/'distribution-final.json')
for name,path in [('baseline',baseline/'benchmark'),('candidate',current/'benchmark'),('compiler_before',cycle/'sources-before/gopurs/gopurs/bin/gopurs-native'),('compiler_after',compiler/'bin/gopurs-native')]:
    assert distribution['binaries'][name]=={'bytes':path.stat().st_size,'sha256':sha(path)}
assert distribution['go_mod_identical']==((baseline/'output/go.mod').read_bytes()==(current/'output/go.mod').read_bytes())
for name,path in [('baseline',baseline),('candidate',current)]:
    files=[p for p in (path/'output').rglob('*.go') if not p.name.endswith('_test.go')]
    assert distribution['generated_go_bytes'][name]==sum(p.stat().st_size for p in files)
    assert distribution['generated_go_files'][name]==len(files)

assert 'pass 9' in (cycle/'compiler-tests-final.log').read_text()
assert 'fail 0' in (cycle/'compiler-tests-final.log').read_text()
assert 'ok ' in (cycle/'record-plans.log').read_text() or 'ok\t' in (cycle/'record-plans.log').read_text()
for phase in ['javascript','go-race']:
    assert phase+': passed' in (cycle/'typed-plans-final.log').read_text()
for repo in [root,compiler,workspace/'altbak.pub',workspace/'gopurs/gopurs-argonaut-core',workspace/'purescript-backend-optimizer-gopurs']:
    subprocess.run(['git','diff','--check'],cwd=repo,check=True)

if not (cycle/'after.json').exists():
    paths={Path(p) for p in manifest['sources']}
    paths.update(workspace/Path(p).relative_to('sources-before') for p in before['sources'])
    paths.update(workspace/Path(p).relative_to('sources-before') for p in extra['sources'])
    paths.update([compiler/'src/Gopurs/DecoderSchemas.purs',compiler/'bin/gopurs.js',compiler/'tools/decoder-schemas.test.mjs',root/'README.md',workspace/'altbak.pub/README.md',root/'docs/benchmark-results/2026-09-25-json-text-schema.md'])
    paths.update(p for p in Path(__file__).resolve().parent.glob('*') if p.is_file())
    paths.update(p for p in (workspace/'gopurs/gopurs-argonaut-codecs/test').glob('*') if p.is_file())
    sources={}
    for path in sorted(paths):
        target=cycle/'sources-after'/path.relative_to(workspace);target.parent.mkdir(parents=True,exist_ok=True);shutil.copy2(path,target);sources[str(target.relative_to(cycle))]=sha(target)
    repos=[root,compiler,workspace/'altbak.pub']
    (cycle/'after.json').write_text(json.dumps({'sources':sources,'heads':{str(p):subprocess.check_output(['git','rev-parse','HEAD'],cwd=p,text=True).strip() for p in repos},'status':{str(p):subprocess.check_output(['git','status','--short'],cwd=p,text=True) for p in repos}},indent=2)+'\n')
after=load(cycle/'after.json')
for relative,expected in after['sources'].items():
    assert sha(cycle/relative)==expected
    assert sha(workspace/Path(relative).relative_to('sources-after'))==expected
evidence=[p for p in cycle.glob('*.json') if p.name!='final-verification.json']+[p for p in cycle.glob('*.log') if not p.name.startswith('finalize')]
evidence += list(cycle.glob('*/results.json'))+list(cycle.glob('*/verification.json'))
result={'status':'passed','starting_archive_files':len(before['sources'])+len(extra['sources']),'final_archive_files':len(after['sources']),'verification':verified,
    'sampled_outputs':sum(v.get('sampled_outputs',0) for v in verified.values()),'evidence_sha256':{str(p.relative_to(cycle)):sha(p) for p in sorted(set(evidence))}}
(cycle/'final-verification.json').write_text(json.dumps(result,indent=2)+'\n')
print(json.dumps({k:v for k,v in result.items() if k!='evidence_sha256'},indent=2))
