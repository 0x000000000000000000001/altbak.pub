const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Effect/corefn.json', 'utf8'));
const functorEffect = corefn.decls.find(d => d.binds && d.binds.some(b => b.identifier === 'functorEffect')) || corefn.decls.find(d => d.identifier === 'functorEffect');
console.log(JSON.stringify(functorEffect, null, 2));
