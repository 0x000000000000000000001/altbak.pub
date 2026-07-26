const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.Semigroup.Traversable/corefn.json'));

function findArray(obj, path) {
  if (Array.isArray(obj)) {
    if (obj.length > 0) {
      console.log(path);
    }
    obj.forEach((val, i) => findArray(val, path + '[' + i + ']'));
  } else if (obj !== null && typeof obj === 'object') {
    for (const k in obj) {
      findArray(obj[k], path + '.' + k);
    }
  }
}
findArray(data.decls[3], 'decls[3]');
