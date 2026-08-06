const loop = loop$a0$copy => {
  let loop$a0 = loop$a0$copy, loop$c = true, loop$r;
  while (loop$c) {
    const v = loop$a0;
    if (v === 0) {
      loop$c = false;
      loop$r = 0;
      continue;
    }
    loop$a0 = v - 1 | 0;
  }
  return loop$r;
};
export {loop};
