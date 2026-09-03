const fs = require('fs');
const data = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Primes/corefn.json', 'utf8'));
const decls = data.decls;
const filterDecl = decls.find(d => d.identifier === 'filter');
console.log(JSON.stringify(filterDecl, null, 2));
