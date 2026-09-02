const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('output/Test.Primes/corefn.json', 'utf8'));
const filterDecl = ast.decls.find(d => d.identifier === 'filter' || (d.binds && d.binds.some(b => b.identifier === 'filter')));
console.log(JSON.stringify(filterDecl, null, 2));
