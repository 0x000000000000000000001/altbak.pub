const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Polymorphism/corefn.json', 'utf8'));
const act = corefn.decls.find(d => d.identifier === 'act');
console.log(JSON.stringify(act, null, 2));
