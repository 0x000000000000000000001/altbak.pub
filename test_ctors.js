const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('output/Test.RBTree/corefn.json'));
const ctors = corefn.dataDecls.find(d => d.typeName === 'Tree').constructors.find(c => c.constructorName === 'T');
console.log(JSON.stringify(ctors, null, 2));
