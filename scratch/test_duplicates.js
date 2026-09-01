const fs = require('fs');
const content = fs.readFileSync('run/bak/go/output/purescript/Control_Monad_Gen.go', 'utf-8');
const lines = content.split('\n');
const rebox = lines.filter(l => l.startsWith('func Rebox_Control_Monad_Gen_3094389156_3240988860'));
console.log("Found matches:", rebox.length);
