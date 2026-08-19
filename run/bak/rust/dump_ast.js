import fs from 'fs';
const data = JSON.parse(fs.readFileSync('./output/purust_output/Purs_Data_Symbol.json', 'utf8'));
const reifySymbol = data.bindings.find(b => b[0] === 'reifySymbol');
console.log(JSON.stringify(reifySymbol, null, 2));
