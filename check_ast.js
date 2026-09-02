const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Primes/corefn.json'));

let found = false;
function walk(node) {
  if (Array.isArray(node)) {
    for (let child of node) walk(child);
  } else if (node && typeof node === 'object') {
    if (node.type === 'Let' || node.type === 'LetRec') {
      for (let bind of node.binds) {
        if (bind.identifier && bind.identifier.includes('go')) {
          console.log('Found binding:', bind.identifier, bind.annotation.type);
          found = true;
        }
      }
    }
    for (let key in node) {
      if (key !== 'annotation') walk(node[key]);
    }
  }
}
for (let decl of ast.decls) walk(decl);
if (!found) console.log('No go bindings found');
