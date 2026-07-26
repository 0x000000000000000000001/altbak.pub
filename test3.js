const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.Semigroup.Traversable/corefn.json'));

function findArray302(obj, path) {
  if (Array.isArray(obj)) {
    if (path.endsWith('[3][0][2]')) {
       console.log("Found [3][0][2] at", path, obj);
    }
    obj.forEach((val, i) => findArray302(val, path + '[' + i + ']'));
  } else if (obj !== null && typeof obj === 'object') {
    for (const k in obj) {
      findArray302(obj[k], path + '.' + k);
    }
  }
}
findArray302(data, 'root');
