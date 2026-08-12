const mempty_ = dict => dict.mempty_;
const mappend_ = dict => dict.mappend_;
const polyLoop = dictMonoidish => {
  const mempty_1 = dictMonoidish.mempty_;
  return v => v1 => {
    if (v === 0) { return v1; }
    return polyLoop(dictMonoidish)(v - 1 | 0)(dictMonoidish.mappend_(v1)(mempty_1));
  };
};
const intMonoidish = {mempty_: 0, mappend_: a => b => a + b | 0};
const test = /* #__PURE__ */ polyLoop(intMonoidish)(10000000)(0);
export {intMonoidish, mappend_, mempty_, polyLoop, test};
