// Diagnostic only: preserve each test unchanged and redirect its Purust CLI
// child to the current bundle with fieldPermutation returning Nothing.
import childProcess from 'node:child_process';
import { syncBuiltinESMExports } from 'node:module';
import { appendFileSync, copyFileSync, existsSync, mkdirSync, readFileSync } from 'node:fs';
import { createHash } from 'node:crypto';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';

const scratch = fileURLToPath(new URL('./', import.meta.url));
const bundle = join(scratch, 'build/purust-rule-disabled.js');
const originalBundle = '/Users/0x1/Documents/htdocs/purust/purust/bin/purust.js';
const label = process.env.PURUST_BASELINE_TEST;
if (!label || !/^[a-z-]+$/.test(label)) throw new Error('PURUST_BASELINE_TEST is required');
const directory = join(scratch, 'build/baseline-test-evidence', label);
mkdirSync(directory, { recursive: true });
const hash = path => createHash('sha256').update(readFileSync(path)).digest('hex');
const originalSpawnSync = childProcess.spawnSync;
childProcess.spawnSync = function (command, args, options) {
  const redirected = Array.isArray(args) && args.includes(originalBundle);
  const effectiveArgs = redirected ? args.map(arg => arg === originalBundle ? bundle : arg) : args;
  const result = originalSpawnSync(command, effectiveArgs, options);
  if (redirected) {
    const outIndex = effectiveArgs.indexOf('--out');
    const output = outIndex < 0 ? null : effectiveArgs[outIndex + 1];
    const sources = [];
    if (output && result.status === 0) {
      for (const name of ['Purs_SharedNullaries', 'Purs_OtherNullaries', 'Purs_RecordRootMove', 'Purs_Foreign_Object']) {
        const path = join(output, name, 'src/lib.rs');
        if (!existsSync(path)) continue;
        const saved = join(directory, `${name}.rs`);
        copyFileSync(path, saved);
        sources.push({ name, original: path, saved, sha256: hash(path) });
      }
    }
    appendFileSync(join(directory, 'redirects.jsonl'), JSON.stringify({ command, args: effectiveArgs,
      cwd: options?.cwd, status: result.status, signal: result.signal, sources,
      bundle, bundleSha256: hash(bundle), originalBundle, originalBundleSha256: hash(originalBundle) }) + '\n');
  }
  return result;
};
syncBuiltinESMExports();
