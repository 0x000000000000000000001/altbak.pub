const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Test.Primes/corefn.json', 'utf8'));

function findGoGo(node) {
    if (!node) return;
    if (typeof node !== 'object') return;
    
    if (node.type === 'Var' && node.value && node.value.identifier === 'go__go_2_1_8') {
        console.log("Found Var go__go_2_1_8!");
        console.log(JSON.stringify(node.ann.type, null, 2));
    }
    
    if (node.type === 'LetRec') {
        node.binds.forEach(b => {
            if (b.identifier === 'go__go_2_1_8') {
                console.log("Found LetRec go__go_2_1_8 binding!");
                console.log(JSON.stringify(b.expression.ann.type, null, 2));
            }
        });
    }
    
    for (const key in node) {
        if (Array.isArray(node[key])) {
            node[key].forEach(findGoGo);
        } else if (typeof node[key] === 'object') {
            findGoGo(node[key]);
        }
    }
}

data.decls.forEach(findGoGo);
