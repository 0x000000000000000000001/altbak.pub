const fs = require('fs');
const data = JSON.parse(fs.readFileSync('../gopurs/gopurs/output/Data.List.Types/corefn.json', 'utf8'));
data.decls.forEach(decl => {
  if (decl.binds) {
    decl.binds.forEach(b => {
      if (b.identifier === 'bindList') {
         console.log(JSON.stringify(b.expression.annotation.type, null, 2));
      }
    });
  }
});
