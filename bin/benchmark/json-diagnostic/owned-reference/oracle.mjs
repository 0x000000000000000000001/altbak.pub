import { readFileSync, writeFileSync } from 'node:fs';
import { resolve } from 'node:path';
import { pathToFileURL } from 'node:url';
import { createHash } from 'node:crypto';

const [suite, output, corpus, destination] = process.argv.slice(2);
const load = name => import(pathToFileURL(resolve(output, name, 'index.js')));
const test = await load('Test.' + suite);
const either = await load('Data.Either');
const decoder = suite === 'JsonTypedAst'
  ? (await load('PureScript.Backend.Optimizer.CoreFn.Json')).decodeModule
  : test.decode;
const canonical = value => Array.isArray(value) ? value.map(canonical)
  : value !== null && typeof value === 'object'
    ? Object.fromEntries(Object.keys(value).sort().map(key => [key, canonical(value[key])])) : value;
const hash = text => createHash('sha256').update(JSON.stringify(canonical(JSON.parse(text)))
  .replace(/[<>&\u2028\u2029]/g, c => '\\u' + c.charCodeAt(0).toString(16).padStart(4, '0'))).digest('hex');
const reports = JSON.parse(readFileSync(corpus)).map(({ name, contents }) => {
  try {
    const value = decoder(JSON.parse(contents));
    const accepted = value instanceof either.Right;
    return { name, accepted, fingerprint: accepted
      ? hash(test.fingerprint(suite === 'JsonTypedAst' ? value.value0 : value)) : null };
  } catch { return { name, accepted: false, fingerprint: null }; }
});
writeFileSync(destination, JSON.stringify(reports, null, 2) + '\n');
