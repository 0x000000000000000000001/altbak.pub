import { existsSync, readFileSync } from 'node:fs';
import { join, resolve } from 'node:path';
const [archiveArg, ...labels] = process.argv.slice(2);
const archive = resolve(archiveArg);
for (const label of labels) {
  const path = join(archive, 'runs', label, 'results.json');
  if (!existsSync(path)) {
    const proposal = join(archive, 'proposals', label, 'results.json');
    console.log(`${label}: no measured results; proposal ${existsSync(proposal) ? JSON.parse(readFileSync(proposal)).status : 'absent'}`);
    continue;
  }
  const result = JSON.parse(readFileSync(path));
  console.log(`${label}: ${result.status}; sum of medians ${JSON.stringify(result.total_ms)}`);
  for (const item of result.results) {
    if (!item.summary) continue;
    const medians = Object.entries(item.summary).map(([name, value]) => `${name}=${value.median_ms}`).join(' ');
    const pairs = Object.entries(item.paired).map(([name, value]) => `${name} ${(100 * (value.ratio_of_medians - 1)).toFixed(2)}%; ` +
      `paired95 [${value.bootstrap_95.map(ratio => (100 * (ratio - 1)).toFixed(2)).join(', ')}]% (${value.favorable}/${value.pairs})`).join('; ');
    console.log(`${item.name}: ${medians}; ${pairs}`);
  }
}
