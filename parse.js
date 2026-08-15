const fs = require('fs');
const data = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Church/corefn.json'));
for (const decl of data.decls) {
  if (decl.bindType === 'Rec') {
    for (const bind of decl.binds) {
      if (bind.identifier === 'fromInt') {
        console.log(JSON.stringify(bind, null, 2));
      }
    }
  } else if (decl.identifier === 'fromInt') {
    console.log(JSON.stringify(decl, null, 2));
  }
}
