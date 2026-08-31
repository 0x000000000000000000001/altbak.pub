const fs = require('fs');
const tcorefn = JSON.parse(fs.readFileSync('output/Data.Functor/corefn.json', 'utf8'));
const voidDecl = tcorefn.decls.find(d => d.identifier === 'void' || (d.binds && d.binds.some(b => b.identifier === 'void')));
console.log(JSON.stringify(voidDecl.expression.body.abstraction.abstraction.expression, null, 2));
