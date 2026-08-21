const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Polymorphism/corefn.json', 'utf8'));
const group = corefn.decls.find(d => Array.isArray(d) && d.some(b => b.identifier === 'polyLoopGo'));
console.log(JSON.stringify(group, null, 2));
