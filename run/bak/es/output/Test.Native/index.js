const loopNative = loopNative$a0$copy => loopNative$a1$copy => {
  let loopNative$a0 = loopNative$a0$copy, loopNative$a1 = loopNative$a1$copy, loopNative$c = true, loopNative$r;
  while (loopNative$c) {
    const v = loopNative$a0, v1 = loopNative$a1;
    if (v === 0) {
      loopNative$c = false;
      loopNative$r = v1;
      continue;
    }
    loopNative$a0 = v - 1 | 0;
    loopNative$a1 = v1 + 1 | 0;
  }
  return loopNative$r;
};
export {loopNative};
