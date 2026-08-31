const fs = require('fs');
const tcorefn = JSON.parse(fs.readFileSync('output/Data.Functor/corefn.json', 'utf8'));
const mapDecl = tcorefn.decls.find(d => d.identifier === 'map' || (d.binds && d.binds.some(b => b.identifier === 'map')));
console.log(JSON.stringify(mapDecl, null, 2));
