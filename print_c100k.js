const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Test.Church/corefn.json', 'utf8'));
const c100k = data.decls.find(d => d.identifier === 'c100k');
console.log(JSON.stringify(c100k, null, 2));
