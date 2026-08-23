// We'll write a node script to parse the JSON for Purs_Control_Extend and see what the AST is!
const fs = require('fs');
const json = JSON.parse(fs.readFileSync('output/purust_output/Purs_Control_Extend/module.json', 'utf8'));
const ast = json.decls.find(d => d.ident === 'extendFn');
console.log(JSON.stringify(ast, null, 2));
