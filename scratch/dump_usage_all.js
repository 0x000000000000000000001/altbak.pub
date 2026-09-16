const fs = require('fs');
const data = JSON.parse(fs.readFileSync('./run/bak/rust/modes/test-RBTree/output/Test.RBTree/corefn.json', 'utf8'));
function walk(node) {
  if (!node) return;
  if (Array.isArray(node)) { node.forEach(walk); return; }
  if (typeof node !== 'object') return;
  if (node.annotation && node.annotation.usageCount === 1) {
    console.log(`USAGE_COUNT_1 on ${node.type}`);
  }
  Object.values(node).forEach(walk);
}
data.decls.forEach(walk);
