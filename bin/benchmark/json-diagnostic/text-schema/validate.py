#!/usr/bin/env python3
"""Validate public text decoding in a separate copy of a preserved diagnostic."""
import argparse, hashlib, importlib.util, json, shutil, subprocess
from pathlib import Path

root=Path(__file__).resolve().parents[4]
parser=argparse.ArgumentParser(description=__doc__)
parser.add_argument('--workspace',type=Path,required=True)
parser.add_argument('--output',type=Path,required=True)
args=parser.parse_args();source=args.workspace.resolve();out=args.output.resolve()
sha=lambda p:hashlib.sha256(p.read_bytes()).hexdigest()
manifest=json.loads((source/'manifest.json').read_text());assert sha(source/'benchmark')==manifest['binary_sha256']
out.mkdir(parents=True)
shutil.copytree(source/'output',out/'output',ignore=shutil.ignore_patterns('*_test.go'))
test=Path(__file__).with_name('validate_test.go')
shutil.copy2(test,out/'output/purescript/text_schema_test.go')
spec=importlib.util.spec_from_file_location('diagnostic',root/'bin/benchmark/json-diagnostic.py')
diagnostic=importlib.util.module_from_spec(spec);spec.loader.exec_module(diagnostic)
env=diagnostic.environment();env['DIAG_CORPUS']=str(root/'test/fixtures/json-decoding/corpus.json')
corpus=root/'var/benchmark/json-direct-20260924/application-differential/corpus.json'
expected=root/'var/benchmark/json-schema-20260924/application-differential/current.json'
summary=json.loads((expected.parent/'summary.json').read_text())
assert sha(corpus)==summary['corpus_sha256']
shutil.copy2(corpus,out/'extra-corpus.json');shutil.copy2(expected,out/'extra-expected.json')
env.update(DIAG_EXTRA_CORPUS=str(out/'extra-corpus.json'),DIAG_EXTRA_EXPECTED=str(out/'extra-expected.json'))
with (out/'tests.log').open('w') as log:
    subprocess.run(['go','test','-race','-count=1','-run','^TestTextSchema','-v','./purescript'],cwd=out/'output',env=env,stdout=log,stderr=subprocess.STDOUT,check=True)
(out/'verification.json').write_text(json.dumps({'status':'passed','source_manifest_sha256':sha(source/'manifest.json'),'tests_sha256':sha(test),'log_sha256':sha(out/'tests.log'),'extra_corpus_sha256':sha(corpus),'extra_expected_sha256':sha(expected),'generated_go':{str(p.relative_to(out)):sha(p) for p in sorted((out/'output').rglob('*.go'))}},indent=2)+'\n')
print(out)
