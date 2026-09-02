const fs = require('fs');
const expr = JSON.parse(fs.readFileSync('sieve_ast.json'));

function extract(e) {
  if (e.type === 'Let') {
    for (const b of e.binds) {
      if (b.bindType === 'Rec') {
        for (const r of b.binds) {
          console.log('Found Rec binding:', r.identifier);
        }
      } else {
        console.log('Found NonRec binding:', b.identifier);
      }
    }
    extract(e.expression);
  } else if (e.type === 'Abs') {
    extract(e.body);
  } else if (e.type === 'Case') {
    for (const a of e.caseAlternatives) {
      if (!a.isGuarded) extract(a.expression);
      else {
        for (const g of a.caseGuards) extract(g.expression);
      }
    }
  }
}
extract(expr);
