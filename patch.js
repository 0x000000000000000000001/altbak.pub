const fs = require('fs');
let code = fs.readFileSync('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  "processBindingGroup binds isRec =\n                    let\n                      _ = if isRec && mutRecBinds == Nothing then Debug.trace (\"Failed mutRec: \" <> String.joinWith \", \" (map (\\(Tuple (Ident name) val) -> name <> \": \" <> printTcoExprShape val) binds)) \\_ -> unit else unit",
  "processBindingGroup binds isRec ="
);
code = code.replace(
  "case mutRecBinds of",
  "let\n                        _ = if isRec && mutRecBinds == Nothing then Debug.trace (\"Failed mutRec: \" <> String.joinWith \", \" (map (\\(Tuple (Ident name) val) -> name <> \": \" <> printTcoExprShape val) binds)) \\_ -> unit else unit\n                      in\n                      case mutRecBinds of"
);
fs.writeFileSync('../gopurs/gopurs/src/Gopurs/CodeGen.purs', code);
