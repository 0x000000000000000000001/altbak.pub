// Build an isolated control with exactly this new proof disabled. All other
// compiled modules stay identical; no compiler source or live bundle changes.
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import * as esbuild from '../../../purust/purust/node_modules/esbuild/lib/main.js';

const here = dirname(fileURLToPath(import.meta.url));
const root = resolve(here, '../../../purust/purust');
await esbuild.build({
  stdin: { contents: `import { main } from ${JSON.stringify(resolve(root, 'output/Main/index.js'))}; main();`, resolveDir: root },
  bundle: true, platform: 'node', format: 'esm', outfile: resolve(here, 'build/purust-rule-disabled.js'),
  loader: { '.rs': 'base64', '.toml': 'base64' },
  plugins: [{ name: 'disable-new-proof-only', setup(build) {
    build.onLoad({ filter: /\/Purust\.FieldPermutations\/index\.js$/ }, args => {
      const contents = readFileSync(args.path, 'utf8');
      const original = 'var fieldPermutation = function (representation) {';
      assert.equal(contents.split(original).length, 2);
      return { contents: contents.replace(original,
        'var fieldPermutation = a => b => c => d => e => f => g => h => Data_Maybe.Nothing.value;\n' +
        'var disabledOriginalFieldPermutation = function (representation) {'), resolveDir: dirname(args.path), loader: 'js' };
    });
  } }],
});
