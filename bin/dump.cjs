const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Church/corefn.json', 'utf8'));
const c100k = corefn.decls.find(d => d.identifier === 'c100k');
console.log(JSON.stringify(c100k.annotation.type, null, 2));
