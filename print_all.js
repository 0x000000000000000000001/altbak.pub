const fs = require('fs');
const data = JSON.parse(fs.readFileSync('run/bak/go/output/purescript/Data_List_Types.gopurs-cache.json', 'utf8'));
function find(node) {
  if (!node) return;
  if (node.tag === 'EvalExtern') {
    console.log(node._1._2);
  }
  for (let key in node) {
    if (typeof node[key] === 'object') find(node[key]);
  }
}
find(data.directives);
