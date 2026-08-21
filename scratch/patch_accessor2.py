import re

with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'r') as f:
    content = f.read()

replacement = """  Accessor base (GetCtorField (Qualified mbMod _) _ (ProperName tyNameStr) (Ident ctorName) fieldIdx) ->
    let baseStr = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound alive false base
        baseTy = inferTypeExpr currentMod aritiesMap bound base
        baseTyStr = codegenExprType true baseTy
        modPrefix = getTyPrefix currentMod (Qualified mbMod (Ident tyNameStr))
        adtName = "crate::Adt_" <> modPrefix <> sanitizeIdent tyNameStr <> "_enum"
        underscores = String.joinWith ", " (Array.replicate fieldIdx "_")
        patternPrefix = if fieldIdx == 0 then "" else underscores <> ", "
        boxedBaseStr = if baseTyStr == "crate::UnknownType" || baseTyStr == "crate::Value" then "(" <> baseStr <> ").unwrap_Adt_" <> modPrefix <> sanitizeIdent tyNameStr <> "()" else baseStr
    in "(if let " <> adtName <> "::" <> sanitizeIdent ctorName <> "(" <> patternPrefix <> "val, ..) = &**(" <> boxedBaseStr <> ") { val.clone() } else { panic!(\\"Expected " <> ctorName <> "\\") })" """

start_str = "  Accessor base (GetCtorField (Qualified mbMod _) _ (ProperName tyNameStr) (Ident ctorName) fieldIdx) ->"
start_idx = content.find(start_str)
end_idx = content.find("  Var (Qualified mbMod (Ident name)) ->")

if start_idx != -1 and end_idx != -1:
    new_content = content[:start_idx] + replacement + "\n" + content[end_idx:]
    with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'w') as f:
        f.write(new_content)
    print("Patched Accessor 2 successfully")
else:
    print("Could not find boundaries")
