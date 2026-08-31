const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('./ast.json'));
function find(node) {
  if (!node) return;
  if (node.type === "Var") {
    console.log("VAR:", node.value.identifier, node.value.moduleName);
  } else if (node.type === "App") {
    find(node.abstraction);
    find(node.argument);
  } else if (node.type === "Abs") {
    find(node.argument);
  } else if (node.type === "Case") {
    node.caseExpressions.forEach(find);
    node.caseAlternatives.forEach(a => {
      let cg = a.caseGuarded;
      if (cg.expression) { // Unconditional
        find(cg.expression);
      } else { // Guarded
        cg.forEach(g => {
          find(g.guard);
          find(g.expression);
        });
      }
    });
  } else if (node.type === "Let") {
    node.binds.forEach(b => {
      if (b.bindType === "NonRec") find(b.expression);
      else b.binds.forEach(bb => find(bb.expression));
    });
    find(node.expression);
  } else if (node.type === "Constructor") {
    // nothing
  } else if (node.type === "Literal") {
    if (Array.isArray(node.value)) node.value.forEach(find);
    else if (node.value && typeof node.value === 'object' && node.value.type) find(node.value);
  }
}
find(ast);
