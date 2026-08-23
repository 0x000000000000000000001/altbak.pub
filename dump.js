const fs = require('fs');
const ast = JSON.parse(fs.readFileSync('run/bak/go/output/Test.Polymorphism/corefn.json'));

function printTree(node, depth = 0) {
    if (!node || typeof node !== 'object') return;
    
    let str = "  ".repeat(depth) + node.type;
    
    if (node.annotation && node.annotation.type) {
        let tyStr = node.annotation.type.type;
        if (tyStr === "ConstrainedType") tyStr += " (has constraints)";
        str += " :: " + tyStr;
    } else if (node.annotation) {
        str += " :: null";
    }
    
    if (node.type === "Var") {
        str += " (" + (node.value.identifier || node.value) + ")";
    } else if (node.type === "Abs") {
        str += " \\" + node.argument;
    } else if (node.type === "Let") {
        str += " [binds: " + node.binds.map(b => b.identifier).join(", ") + "]";
    }
    
    console.log(str);
    
    if (node.type === "Abs") {
        printTree(node.body, depth + 1);
    } else if (node.type === "App") {
        printTree(node.abstraction, depth + 1);
        printTree(node.argument, depth + 1);
    } else if (node.type === "Let") {
        node.binds.forEach(b => printTree(b.expression, depth + 1));
        printTree(node.expression, depth + 1);
    }
}

for (const decl of ast.decls) {
    if (decl.bindType === "NonRec" && decl.identifier === "polyLoop") {
        printTree(decl.expression);
    }
}
