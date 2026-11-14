import { readFileSync } from 'node:fs';
import { createHash } from 'node:crypto';
import { performance } from 'node:perf_hooks';
const canonical = value => Array.isArray(value) ? value.map(canonical) : value !== null && typeof value === 'object' ? Object.fromEntries(Object.keys(value).sort().map(k => [k, canonical(value[k])])) : value;
const hash = value => createHash('sha256').update(JSON.stringify(canonical(value)).replace(/[<>&\u2028\u2029]/g, c => '\\u' + c.charCodeAt(0).toString(16).padStart(4, '0'))).digest('hex');
let sink;
export const drive = parse => decode => decodeText => encode => fingerprint => () => {
  const files = JSON.parse(readFileSync(process.env.DIAG_CORPUS, 'utf8'));
  const texts = files.map(x => x.contents);
  const parsed = texts.map(parse);
  const expectedJSON = texts.map(x => hash(JSON.parse(x)));
  const expectedAST = parsed.map(x => hash(JSON.parse(fingerprint(decode(x)))));
  // Invalid inputs are checked against the fixed oracle outside timed work.
  const indices = files.flatMap((file, i) => file.benchmark ? [i] : []);
  if (indices.length === 0) throw Error('No timed cases');
  const phases = {};
  for (const phase of ['parse', 'decode', 'combined']) {
    const samples = [];
    for (let pass = 0; pass < 7; pass++) {
      const results = new Array(indices.length);
      const start = performance.now();
      for (let slot = 0; slot < indices.length; slot++) {
        const i = indices[slot];
        const result = phase === 'parse' ? parse(texts[i]) : phase === 'decode' ? decode(parsed[i]) : decodeText(texts[i]);
        results[slot] = result;
        sink = result;
      }
      const elapsed = (performance.now() - start) * 1000;
      const fingerprints = results.map(x => hash(JSON.parse(phase === 'parse' ? encode(x) : fingerprint(x))));
      const expected = indices.map(i => phase === 'parse' ? expectedJSON[i] : expectedAST[i]);
      if (JSON.stringify(fingerprints) !== JSON.stringify(expected)) throw Error('Unstable ' + phase + ' output');
      if (pass >= 2) samples.push({time_us: elapsed});
    }
    phases[phase] = {samples, time_us: Math.min(...samples.map(x => x.time_us))};
  }
  console.log(JSON.stringify({backend:'js', modules:files.length, timed_cases:indices.length, names:files.map(f => f.name), fingerprints:expectedAST, json_fingerprints:expectedJSON, phases, node:process.version}));
};
