const fs = require('fs');
const code = fs.readFileSync('run/bak/go/output/purescript/Test_Polymorphism.go', 'utf8');
console.log(code.split('\n').filter(l => l.includes('Call_Test_Polymorphism_polyLoopGo')).join('\n'));
