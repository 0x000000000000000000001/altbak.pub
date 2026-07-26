const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.Semigroup.Traversable/corefn.json'));

let foundPaths = [];
function findArray302(obj, path) {
  if (Array.isArray(obj)) {
    if (obj.length > 2 && (obj[2] === null || typeof obj[2] !== 'object')) {
      console.log("Found array with non-object at index 2:", path, typeof obj[2], obj[2]);
    }
    obj.forEach((val, i) => findArray302(val, path + '[' + i + ']'));
  } else if (obj !== null && typeof obj === 'object') {
    for (const k in obj) {
      findArray302(obj[k], path + '.' + k);
    }
  }
}
findArray302(data, 'root');
