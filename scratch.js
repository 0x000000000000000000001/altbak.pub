const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.List.Lazy/corefn.json'));
const dropWhile = data.decls.find(d => d.identifier === 'dropWhile');
const goBind = dropWhile.expression.binds[0].binds[0];
console.log(JSON.stringify(goBind.expression.argument, null, 2));
