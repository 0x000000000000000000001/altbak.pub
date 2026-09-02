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
          console.log('Found identifier:', bind.identifier, bind.annotation.type);
          found = true;
        } else if (bind.bindType === 'NonRec' || bind.bindType === 'Rec') {
           let actualBind = bind.identifier ? bind : (bind.expression ? bind : null);
           // Actually in purescript corefn, binds are just objects with `identifier`, `expression`, `annotation`.
           // Wait, maybe the property is `identifier` ?
        }
      }
    }
    // pure walk
    for (let key in node) {
        if (key === 'identifier' && typeof node[key] === 'string' && node[key].includes('go')) {
            console.log("Found identifier key:", node[key]);
            if (node.annotation) console.log("Annotation type:", node.annotation.type);
        }
    }
    for (let key in node) {
      if (key !== 'annotation') walk(node[key]);
    }
  }
}
for (let decl of ast.decls) walk(decl);
