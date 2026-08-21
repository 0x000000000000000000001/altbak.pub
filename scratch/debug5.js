const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Polymorphism/corefn.json', 'utf8'));
const polyLoopGo = corefn.decls.find(d => d.identifier === 'polyLoopGo');
console.log(JSON.stringify(polyLoopGo, null, 2));
