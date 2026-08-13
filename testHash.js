const crypto = require('crypto');
function hashString(s) {
  const hash = crypto.createHash('sha256');
  hash.update(s);
  return parseInt(hash.digest('hex').substring(0, 8), 16).toString();
}
console.log(hashString('forall a. HeytingAlgebra a => a -> a -> a'));
console.log(hashString('forall a. HeytingAlgebra Any => Any -> Any -> Any'));
console.log(hashString('forall a. HeytingAlgebra Boolean => Boolean -> Boolean -> Boolean'));
