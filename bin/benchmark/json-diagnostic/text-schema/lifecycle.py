#!/usr/bin/env python3
"""Audit public decoder construction, two retained decodes and final collection."""
import argparse, hashlib, importlib.util, json, re, shutil, statistics, subprocess
from pathlib import Path

here=Path(__file__).resolve().parent;root=here.parents[3]
sha=lambda p:hashlib.sha256(p.read_bytes()).hexdigest()
load=lambda p:json.loads(p.read_text())
parser=argparse.ArgumentParser(description=__doc__)
parser.add_argument('action',choices=['build','measure'])
parser.add_argument('--baseline',type=Path)
parser.add_argument('--current',type=Path)
parser.add_argument('--output',type=Path,required=True)
args=parser.parse_args();out=args.output.resolve()
spec=importlib.util.spec_from_file_location('diagnostic',root/'bin/benchmark/json-diagnostic.py')
diagnostic=importlib.util.module_from_spec(spec);spec.loader.exec_module(diagnostic)
env=diagnostic.environment();env['DIAG_CORPUS']=str(out/'corpus.json')
oracle=load(root/'test/fixtures/json-decoding/expected.json')
expected=dict(zip(oracle['names'],oracle['fingerprints']))
if args.action=='build':
    assert args.baseline and args.current
    out.mkdir(parents=True)
    shutil.copy2(root/'test/fixtures/json-decoding/corpus.json',out/'corpus.json')
    shutil.copy2(here/'lifecycle.go',out/'main.go')
    builds={}
    for mode,base in [('baseline',args.baseline.resolve()),('current',args.current.resolve())]:
        work=out/mode
        shutil.copytree(base/'output',work/'output',ignore=shutil.ignore_patterns('*_test.go'))
        construction='parse,decode:=p.Get_Test_JsonDecoding_parse(),p.Get_Test_JsonDecoding_decode();combined=func(text string) rt.Value{return rt.Apply(decode,rt.Apply(parse,rt.Str(text)))}' if mode=='baseline' else 'decode:=p.Get_Test_JsonDecoding_decodeText();combined=func(text string) rt.Value{return rt.Apply(decode,rt.Str(text))}'
        (work/'output/main/main.go').write_text((out/'main.go').read_text().replace('// BUILD_DECODER',construction))
        command=['go','build','-pgo=off','-o',str(work/'audit'),'./main']
        with (work/'build.log').open('w') as log:subprocess.run(command,cwd=work/'output',env=dict(env,GOMAXPROCS='14'),stdout=log,stderr=subprocess.STDOUT,check=True)
        builds[mode]={'workspace':str(base),'source_manifest_sha256':sha(base/'manifest.json'),'binary_sha256':sha(work/'audit'),'command':command,'generated_go':{str(p.relative_to(work)):sha(p) for p in sorted((work/'output').rglob('*.go'))}}
    (out/'builds.json').write_text(json.dumps(builds,indent=2)+'\n')
else:
    builds=load(out/'builds.json');runs={'baseline':[],'current':[]}
    for index in range(6):
        for mode in (['baseline','current'] if index%2==0 else ['current','baseline']):
            assert sha(out/mode/'audit')==builds[mode]['binary_sha256']
            path=out/f'{index:02}-{mode}.json'
            with path.open('x') as log:subprocess.run([out/mode/'audit'],env=env,stdout=log,check=True)
            report=load(path)
            assert len(report['names'])==oracle['timed_cases']
            assert report['first_fingerprints']==report['second_fingerprints']==[expected[n] for n in report['names']]
            runs[mode].append(report)
    # inittrace is an allocation inventory, separate from the paired timings.
    initialization={}
    for mode in builds:
        with (out/f'init-{mode}.json').open('x') as stdout,(out/f'init-{mode}.log').open('x') as stderr:
            subprocess.run([out/mode/'audit'],env=dict(env,GODEBUG='inittrace=1'),stdout=stdout,stderr=stderr,check=True)
        report=load(out/f'init-{mode}.json')
        assert report['first_fingerprints']==report['second_fingerprints']==[expected[n] for n in report['names']]
        line=next(line for line in (out/f'init-{mode}.log').read_text().splitlines() if line.startswith('init gopurs/output/purescript '))
        match=re.search(r'(\d+) bytes, (\d+) allocs',line)
        initialization[mode]={'line':line,'allocated_bytes':int(match[1]),'allocations':int(match[2])}
    phases=['construction','first_combined','second_combined','release_collection']
    sums={mode:[sum(run[p]['time_us'] for p in phases) for run in reports] for mode,reports in runs.items()}
    result={'protocol':'six paired fresh processes per build; public combined getters; two full retained corpus passes; separately measured final collection after fingerprint scratch cleanup; GOMAXPROCS=1 GOGC=100',
        'builds':builds,'runs':runs,'source_sha256':{str(p):sha(p) for p in [Path(__file__).resolve(),out/'main.go']},
        'medians':{mode:{phase:{key:statistics.median(run[phase][key] for run in reports) for key in ['time_us','allocated_bytes']} for phase in phases} for mode,reports in runs.items()},
        'measured_interval_subtotals_us':sums,'subtotal_medians_us':{mode:statistics.median(values) for mode,values in sums.items()},
        'subtotal_winning_pairs':sum(c<b for b,c in zip(sums['baseline'],sums['current'])),
        'initialization':initialization,'validated_outputs':sum(len(run['names'])*2 for reports in runs.values() for run in reports)+2*2*oracle['timed_cases']}
    (out/'results.json').write_text(json.dumps(result,indent=2)+'\n')
    print(json.dumps({k:result[k] for k in ['medians','subtotal_medians_us','subtotal_winning_pairs','initialization','validated_outputs']},indent=2))
