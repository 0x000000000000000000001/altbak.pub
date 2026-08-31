const fs = require('fs');
const json = JSON.parse(fs.readFileSync('run/bak/go/output/Test.ArrayOps/corefn.json', 'utf8'));
const sumEvens = json.decls.find(d => d.identifier === 'sumEvens');
console.log(JSON.stringify(sumEvens, null, 2));
