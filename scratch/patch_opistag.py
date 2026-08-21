import re

with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'r') as f:
    content = f.read()

replacement = """      OpIsTag (Qualified mbMod (Ident ctorName)) -> 
        let
          aStrBoxed = boxUnbox Any aTy aStrRaw
          defaultFallback = "(" <> aStrBoxed <> ".unwrap_record().tag == \\"" <> ctorName <> "\\")"
        in case unwrapType aTy of
          ADT _ fqn _ -> 
            let enumName = "crate::Adt_" <> String.replaceAll (Pattern ".") (Replacement "_") (String.joinWith "_" fqn) <> "_enum"
                adtStr = boxUnbox aTy aTy aStrRaw
            in "(if let " <> enumName <> "::" <> sanitizeIdent ctorName <> "{ .. } = &**(" <> adtStr <> ") { true } else { false })"
          _ -> defaultFallback"""

content = content.replace('      OpIsTag (Qualified _ (Ident ctorName)) -> "(" <> boxUnbox Any aTy aStrRaw <> ".unwrap_record().tag == \\"" <> ctorName <> "\\")"', replacement)

with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'w') as f:
    f.write(content)
print("Patched OpIsTag successfully")
