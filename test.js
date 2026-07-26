const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.Semigroup.Traversable/corefn.json'));

function findArray(obj, path) {
  if (Array.isArray(obj)) {
    if (obj.length > 0 && typeof obj[0] === 'string') {
      console.log("Found array of strings at:", path, obj);
    }
    obj.forEach((val, i) => findArray(val, path + '[' + i + ']'));
  } else if (obj !== null && typeof obj === 'object') {
    for (const k in obj) {
      findArray(obj[k], path + '.' + k);
    }
  }
}
findArray(data, 'root');
