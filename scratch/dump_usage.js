const fs = require('fs');
const data = JSON.parse(fs.readFileSync('./run/bak/rust/modes/test-RBTree/output/Test.RBTree/corefn.json', 'utf8'));

function walk(node, parent) {
  if (!node) return;
  
  if (node.annotation && node.annotation.usageCount === 1) {
    console.log(`USAGE_COUNT_1 on ${node.type}`);
  }
  
  if (node.type === 'Abs') {
    walk(node.argument, node);
  } else if (node.type === 'App') {
    walk(node.abstraction, node);
    walk(node.argument, node);
  } else if (node.type === 'Let') {
    node.binds.forEach(b => walk(b.expression, node));
    walk(node.expression, node);
  } else if (node.type === 'Case') {
    node.caseExpressions.forEach(e => walk(e, node));
    node.caseAlternatives.forEach(a => walk(a.expression, node));
  }
}
data.decls.forEach(d => walk(d.expression || d, null));
