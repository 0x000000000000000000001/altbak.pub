const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Polymorphism/corefn.json', 'utf8'));
const group = corefn.decls.find(d => d.bindType === 'Rec');
const polyLoopGo = group.binds.find(b => b.identifier === 'polyLoopGo');
console.log(JSON.stringify(polyLoopGo, null, 2));
