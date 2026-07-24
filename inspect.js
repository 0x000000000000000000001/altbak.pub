const fs = require('fs');
const corefn = JSON.parse(fs.readFileSync('run/bak/go/output/Test.RBTree/corefn.json', 'utf8'));

function findBinding(id) {
  for (const d of corefn.decls) {
    if (d.identifier === id) return d;
    if (d.binds) {
      for (const b of d.binds) {
        if (b.identifier === id) return b;
      }
    }
  }
}

const balance = findBinding('balance');
const insert = findBinding('insert');

console.log("balance type:", balance.expression.type);
console.log("insert type:", insert.expression.type);

function inspect(expr, depth=0) {
  if (depth > 5) return "...";
  if (expr.type === 'Abs') return `Abs(${expr.argument}, ${inspect(expr.body, depth+1)})`;
  if (expr.type === 'App') return `App(${inspect(expr.abstraction, depth+1)}, ${inspect(expr.argument, depth+1)})`;
  if (expr.type === 'Var') return `Var(${expr.value.identifier})`;
  if (expr.type === 'Let') return `Let(...)`;
  if (expr.type === 'Case') return `Case(...)`;
  if (expr.type === 'Constructor') return `Constructor(...)`;
  if (expr.type === 'Literal') return `Literal(...)`;
  return expr.type;
}

console.log("balance:", inspect(balance.expression));
console.log("insert:", inspect(insert.expression));
