import re

file_path = "/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs"

with open(file_path, "r") as f:
    content = f.read()

# Fix LetRec
letrec_body_pattern = r'bodyCode = codegenExpr_ currentMod allZeroArity allMacroBindings mbLoop aritiesMap globalClassFields bound alive inEffectBlock body'
letrec_body_replacement = r'''newMbLoop = case mbLoop of
        Just l | Array.any (\\(Tuple (Ident n) _) -> sanitizeIdent n == l.name) bindsArray -> Nothing
        _ -> mbLoop
      bodyCode = codegenExpr_ currentMod allZeroArity allMacroBindings newMbLoop aritiesMap globalClassFields bound alive inEffectBlock body'''

new_content = content.replace(letrec_body_pattern, letrec_body_replacement)

# Fix Let
let_pattern = r'Let _ \(Ident n\) val body ->\s+let\s+aliveForVal = Set\.union alive \(freeVariables body\)\s+valCode = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound aliveForVal false val\s+bodyCode = codegenExpr_ currentMod allZeroArity allMacroBindings mbLoop aritiesMap globalClassFields bound alive inEffectBlock body'

let_replacement = r'''Let _ (Ident n) val body ->
    let
      aliveForVal = Set.union alive (freeVariables body)
      valCode = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound aliveForVal false val
      newMbLoop = case mbLoop of
        Just l | sanitizeIdent n == l.name -> Nothing
        _ -> mbLoop
      bodyCode = codegenExpr_ currentMod allZeroArity allMacroBindings newMbLoop aritiesMap globalClassFields bound alive inEffectBlock body'''

new_content = re.sub(let_pattern, let_replacement, new_content, flags=re.MULTILINE | re.DOTALL)

with open(file_path, "w") as f:
    f.write(new_content)

print("Shadowing patched.")
