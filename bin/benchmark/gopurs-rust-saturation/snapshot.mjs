// Queue a source-only hypothesis while another frozen candidate is measured.
// sources.json inventories the starting point; queued proposals may be edited.
// candidate.mjs freezes and hashes the actual build inputs before execution.
import assert from 'node:assert/strict';
import { existsSync, mkdirSync, readFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { copy, hash, manifest } from './common.mjs';
import { writeJson } from '../gopurs-aff/common.mjs';
const [archiveArg, label, candidate] = process.argv.slice(2);
assert(archiveArg && /^[a-z0-9-]+$/.test(label));
const archive = resolve(archiveArg), htdocs = resolve(dirname(fileURLToPath(import.meta.url)), '../../../..');
const base = join(archive, 'queued', label); assert(!existsSync(base)); mkdirSync(base, { recursive: true });
const roots = candidate ? [0, 1].map(i => join(archive, 'candidates', candidate, 'sources', String(i)))
  : [join(htdocs, 'gopurs/gopurs'), join(htdocs, 'purescript-backend-optimizer-gopurs')];
const sources = roots.map((root, index) => {
  const files = [...manifest(join(root, 'src')).map(file => ({ ...file, path: 'src/' + file.path })),
    ...['spago.yaml', 'spago.lock', 'package.json'].filter(file => existsSync(join(root, file)))
      .map(path => ({ path, sha256: hash(readFileSync(join(root, path))) }))];
  for (const file of files) copy(join(root, file.path), join(base, String(index), file.path));
  return { root, files };
});
writeJson(join(base, 'sources.json'), { candidate, sources });
console.log(`Queued ${label}: ${base}`);
