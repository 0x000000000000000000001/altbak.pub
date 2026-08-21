import re

with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'r') as f:
    content = f.read()

replacement = """  Accessor base (GetCtorField (Qualified mbMod _) _ (ProperName tyNameStr) (Ident ctorName) fieldIdx) ->
    let baseStr = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound alive false base
        baseTy = inferTypeExpr currentMod aritiesMap bound base
        modPrefix = getTyPrefix currentMod (Qualified mbMod (Ident tyNameStr))
        adtName = "crate::Adt_" <> modPrefix <> sanitizeIdent tyNameStr <> "_enum"
        underscores = String.joinWith ", " (Array.replicate fieldIdx "_")
        patternPrefix = if fieldIdx == 0 then "" else underscores <> ", "
        boxedBaseStr = boxUnbox (ADT tyNameStr [] []) baseTy baseStr -- boxUnbox to the ADT type if necessary
    in "(if let " <> adtName <> "::" <> sanitizeIdent ctorName <> "(" <> patternPrefix <> "val, ..) = &**(" <> boxedBaseStr <> ") { val.clone() } else { panic!(\\"Expected " <> ctorName <> "\\") })" """

content = content.replace('  Accessor base (GetCtorField _ _ _ _ _ fieldIdx) ->\n    let baseStr = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound alive false base\n    in baseStr <> ".unwrap_record().vals.as_ref().unwrap()[" <> show fieldIdx <> "].clone()"', replacement)

with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'w') as f:
    f.write(content)
print("Patched Accessor successfully")
