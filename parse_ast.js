const fs = require('fs');
const json = JSON.parse(fs.readFileSync('output/Test.RBTree/corefn.json', 'utf8'));

function findVar(expr, name) {
  if (!expr) return;
  if (expr.type === 'Var' && expr.value.identifier === name) {
    console.log(JSON.stringify(expr.annotation.type, null, 2));
  }
  if (expr.type === 'App') {
    findVar(expr.abstraction, name);
    findVar(expr.argument, name);
  }
  if (expr.type === 'Abs') findVar(expr.body, name);
  if (expr.type === 'Let') {
    expr.binds.forEach(b => findVar(b.expression, name));
    findVar(expr.expression, name);
  }
  if (expr.type === 'Case') {
    expr.caseExpressions.forEach(e => findVar(e, name));
    expr.caseAlternatives.forEach(a => findVar(a.expression, name));
  }
}
json.decls.forEach(d => {
    if (d.bindType === 'NonRec') findVar(d.expression, 'lessThan');
    else d.binds.forEach(b => findVar(b.expression, 'lessThan'));
});
