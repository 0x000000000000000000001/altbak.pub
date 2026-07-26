const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.Semigroup.Traversable/corefn.json'));

let foundPaths = [];
function findArray302(obj, path) {
  if (Array.isArray(obj)) {
    obj.forEach((val, i) => findArray302(val, path + '[' + i + ']'));
  } else if (obj !== null && typeof obj === 'object') {
    for (const k in obj) {
      findArray302(obj[k], path + '.' + k);
    }
  } else {
     foundPaths.push(path);
  }
}
findArray302(data.decls, 'decls');
console.log(foundPaths.filter(p => p.includes('[3]') && p.includes('[0]') && p.includes('[2]')).slice(0, 10).join('\n'));
