function hashString(s) {
  var hash = 5381, i = s.length;
  while(i) {
    hash = (hash * 33) ^ s.charCodeAt(--i);
  }
  return hash >>> 0;
}
console.log(hashString("Data_Test_RBTree_T"));
