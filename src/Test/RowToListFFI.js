export const runRowToListFFI = function(limit) {
  let dummy = Math.floor(limit);
  let rec = { a: 1, b: "two", c: true, d: 4.0, e: "five" };
  
  let dict = dictCons(dictCons(dictCons(dictCons(dictCons(dictNil)))));
  return keys(dict, rec);
};

const dictNil = {
  keysImpl: function(_proxy) {
    return 0;
  }
};

function dictCons(dictTail) {
  return {
    keysImpl: function(_proxy) {
      return 1 + dictTail.keysImpl({});
    }
  };
}

function keysImpl(dict) {
  return dict.keysImpl;
}

// The record value is irrelevant, as in the PureScript RowToList dictionary.
function keys(dict, _record) {
  return keysImpl(dict)({});
}
