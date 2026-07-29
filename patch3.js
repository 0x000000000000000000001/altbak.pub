const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `                  fullName = "Call_" <> sanitizeName name
                in
                  [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]`,
  `                  fullName = "Call_" <> sanitizeName name
                in
                  if typeSig == Nothing then unsafeCrashWith ("No type signature for function: " <> name) else [ Tuple (sanitizeName name) { fullName, fArgs: fArgsGo, fRet: fRetGo, arity: Array.length args } ]`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
