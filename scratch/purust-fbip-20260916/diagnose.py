from pathlib import Path
import os
import signal
import subprocess
import time

root = Path('/Users/0x1/Documents/htdocs')
scratch = Path(__file__).resolve().parent
source = root / 'purust/purust/bin/purust.js'
original = (scratch / 'before.mjs').read_text()
code = original.replace('var toBackendModule = function(v) {',
    'var toBackendModule = function(v) { console.error("PHASE optimize", v.name);')
code = code.replace('var modNameStr = replaceAll(".")("_")(unwrap9(backendMod.name));',
    'console.error("PHASE codegen", backendMod.name);\n            var modNameStr = replaceAll(".")("_")(unwrap9(backendMod.name));')
code = code.replace('var codegenExpr_ = function(valueEnums) {',
    'var __diagnosticCalls = 0;\nvar codegenExpr_ = function(valueEnums) {')
code = code.replace('return function(exprRaw) {\n                      var extractUsage',
    'return function(exprRaw) {\n'
    '                      if (++__diagnosticCalls % 10000 === 0) console.error("CODEGEN calls", __diagnosticCalls, currentMod, exprRaw.constructor.name);\n'
    '                      if (__diagnosticCalls > 200000) throw new Error("Codegen diagnostic budget exhausted: " + currentMod + " " + exprRaw.constructor.name);\n'
    '                      var extractUsage')
# Log only the root tag; printing an AST recursively obscures the diagnosis.
code = code.replace('var v1 = trace2("codegenExpr_ called with: " + take3(50)(printAST(exprRaw)))(function(v22) {\n                        return unit;\n                      });', 'var v1 = unit;')
(scratch / 'diagnostic.mjs').write_text(code)
command = ['node', '--expose-gc', '--stack-size=8192', '--max-old-space-size=4096',
    str(scratch / 'diagnostic.mjs'), '--main', 'AppX', '--source',
    str(root / 'altbak.pub/run/bak/rust/modes/test-RBTree/output'),
    '--out', str(scratch / 'before-output'), '--ffi-dir', str(root / 'purust')]
start = time.monotonic()
with (scratch / 'before-sampled.log').open('w') as log:
    process = subprocess.Popen(command, cwd=scratch, stdout=log, stderr=subprocess.STDOUT, start_new_session=True)
    try:
        result = process.wait(timeout=90)
    except subprocess.TimeoutExpired:
        os.killpg(process.pid, signal.SIGTERM)
        process.wait(timeout=5)
        result = 'timeout (90 seconds)'
print('result:', result, 'seconds:', round(time.monotonic() - start, 2))
lines = (scratch / 'before-sampled.log').read_text().splitlines()
print('log lines:', len(lines))
print('\n'.join(lines[-35:]))
