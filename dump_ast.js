const fs = require('fs');
const tcorefn = JSON.parse(fs.readFileSync('output/Test.ListOps/corefn.json', 'utf8'));
const sumDecl = tcorefn.decls.filter(d => {
  if (d.bindType === 'NonRec') return d.identifier.includes('sum');
  if (d.bindType === 'Rec') return d.binds.some(b => b.identifier.includes('sum'));
  return false;
});
console.dir(sumDecl, {depth: null});
