const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Polymorphism/corefn.json', 'utf8'));
const polyLoop = corefn.decls.find(d => d.identifier === 'polyLoop');
console.log(JSON.stringify(polyLoop, null, 2));
