import fs from 'fs';
const data = JSON.parse(fs.readFileSync('run/bak/rust/output/Data.Symbol/tcorefn.json', 'utf8'));
const reifySymbol = data.decls.find(b => b.bindType === 'NonRec' && b.ident === 'reifySymbol');
console.log(JSON.stringify(reifySymbol, null, 2));
