const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `                  fvs = fromMaybe Set.empty (Array.index laterFvs acc.fieldIdx)
                  filteredBound = foldl (\\b k -> Map.delete k b) bound (Array.fromFoldable fvs)
                  resVal = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities filteredBound Nothing [] false false acc.nextId val
                  expectedType = case Array.index fieldTypes acc.fieldIdx of`,
  `                  fvs = fromMaybe Set.empty (Array.index laterFvs acc.fieldIdx)
                  resVal = unsafePerformEffect do
                    originalReused <- Ref.read globalReusedVars
                    Ref.modify_ (Set.union fvs) globalReusedVars
                    let r = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId val
                    Ref.write originalReused globalReusedVars
                    pure r
                  expectedType = case Array.index fieldTypes acc.fieldIdx of`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
