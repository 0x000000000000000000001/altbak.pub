const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Test.Polymorphism/corefn.json', 'utf8'));

function findGo(obj) {
  if (!obj) return;
  if (Array.isArray(obj)) return obj.forEach(findGo);
  if (typeof obj === 'object') {
    if (obj.identifier === 'go') {
      console.log(JSON.stringify(obj, null, 2));
    }
    for (let k in obj) findGo(obj[k]);
  }
}

findGo(data);
