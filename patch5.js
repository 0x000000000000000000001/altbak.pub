const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `                  if typeSig == Nothing then [ Tuple (sanitizeName name) { fullName, fArgs: Array.replicate (Array.length args) TypeValue, fRet: TypeValue, arity: Array.length args } ] else [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]`,
  `                  let _ = if typeSig == Nothing then Debug.trace ("No type signature for function: " <> name) (\\_ -> unit) else unit
                  in if typeSig == Nothing then [ Tuple (sanitizeName name) { fullName, fArgs: Array.replicate (Array.length args) TypeValue, fRet: TypeValue, arity: Array.length args } ] else [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
