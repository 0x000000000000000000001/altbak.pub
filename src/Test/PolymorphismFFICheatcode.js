
class StrShowCheat { showLen() { return 5; } }
class ArrShowCheat { showLen() { return 3; } }

export const runPolymorphismFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  let s1 = new StrShowCheat();
  let s2 = new ArrShowCheat();
  for (let i = 0; i < n; i++) {
    sum += s1.showLen() + s2.showLen();
  }
  return sum;
};
