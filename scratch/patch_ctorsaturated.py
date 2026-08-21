import re

with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'r') as f:
    content = f.read()

replacement = """  CtorSaturated (Qualified mbMod _) _ (ProperName tyNameStr) (Ident ctorName) fields ->
    let
      modPrefix = getTyPrefix currentMod (Qualified mbMod (Ident tyNameStr))
      structKey = modPrefix <> sanitizeIdent tyNameStr
      adtName = "crate::Adt_" <> modPrefix <> sanitizeIdent tyNameStr
    in case Map.lookup structKey globalClassFields of
      Just classFields -> 
        let
          structFieldsCode = String.joinWith ", " (Array.mapWithIndex (\\i (Tuple _ val) -> 
            let (Tuple fieldName expectedTy) = fromMaybe (Tuple ("field" <> show i) Any) (Array.index classFields i)
                subsequent = Array.drop (i + 1) fields
                aliveForV = Set.union alive (Array.foldl Set.union Set.empty (map (\\(Tuple _ sv) -> freeVariables sv) subsequent))
                valCode = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound aliveForV false val
                valTy = inferTypeExpr currentMod aritiesMap bound val
            in sanitizeIdent fieldName <> ": " <> boxUnbox expectedTy valTy valCode
          ) fields)
        in "crate::Value::" <> "Adt_" <> modPrefix <> sanitizeIdent tyNameStr <> "(perceus_ptr::PerceusPtr::new(" <> adtName <> "Dict { " <> structFieldsCode <> " }))"
      Nothing ->
        let
            boundVars = Map.toUnfoldable bound :: Array (Tuple String ExprType)
            fieldsAlive = Array.foldl Set.union Set.empty (map (\\(Tuple _ sv) -> freeVariables sv) fields)
            isSameADT targetName t = case unwrapType t of
              ADT n _ _ -> n == targetName || String.contains (Pattern ("." <> targetName)) n
              TypeApp a _ -> isSameADT targetName a
              _ -> false
            deadAdtVars = unsafePerformEffect do
              currentScope <- Ref.read globalCurrentScopeVars
              pure $ Array.filter (\\(Tuple name ty) -> 
                Set.member name currentScope &&
                not (Set.member name alive) &&
                true &&
                isSameADT tyNameStr ty
              ) boundVars
            
            reuseCode = unsafePerformEffect do
              consumed <- Ref.read globalConsumed
              currentScope <- Ref.read globalCurrentScopeVars
              let dbg = Array.foldl (\\acc (Tuple n t) -> acc <> " " <> n <> ":" <> printType t) "" boundVars
              let dbgDead = Array.foldl (\\acc (Tuple n t) -> acc <> " " <> n) "" deadAdtVars
              let dbgAlive = Array.foldl (\\acc n -> acc <> " " <> n) "" (Set.toUnfoldable alive :: Array String)
              let available = Array.filter (\\(Tuple name _) -> not (Set.member name consumed)) deadAdtVars
              let dbgTraceMsg = "CtorSaturated ctorName=" <> ctorName <> " tyNameStr=" <> tyNameStr <> "\\n  bound: " <> dbg <> "\\n  alive: " <> dbgAlive <> "\\n  dead: " <> dbgDead <> "\\n  scope: " <> show (Array.fromFoldable currentScope :: Array String)
              case Array.head available of
                Just (Tuple reuseName _) -> do
                  Ref.write (Set.insert reuseName consumed) globalConsumed
                  let innerAlive = Set.insert reuseName alive
                  let fieldsCode = if Array.null fields then "" else 
                        String.joinWith ", " (Array.mapWithIndex (\\i (Tuple _ val) -> 
                          let subsequent = Array.drop (i + 1) fields
                              aliveForV = Set.union innerAlive (Array.foldl Set.union Set.empty (map (\\(Tuple _ sv) -> freeVariables sv) subsequent))
                              valCode = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound aliveForV false val
                              valTy = inferTypeExpr currentMod aritiesMap bound val
                          in boxUnbox Any valTy valCode
                        ) fields)
                  let fieldsStr = if Array.null fields then "" else "(" <> fieldsCode <> ")"
                  let res = "{\\n" <>
                         "    let mut _reuse = " <> reuseName <> ";\\n" <>
                         "    {\\n" <>
                         "        let _mut = perceus_ptr::PerceusPtr::make_mut(_reuse.as_" <> "Adt_" <> modPrefix <> sanitizeIdent tyNameStr <> "_mut());\\n" <>
                         "        *_mut = " <> adtName <> "_enum::" <> sanitizeIdent ctorName <> fieldsStr <> ";\\n" <>
                         "    }\\n" <>
                         "    _reuse\\n" <>
                         "}"
                  pure $ Debug.trace ("KnotTying YES " <> dbgTraceMsg) \\_ -> res
                Nothing -> do
                  let fieldsCode = if Array.null fields then "" else 
                        String.joinWith ", " (Array.mapWithIndex (\\i (Tuple _ val) -> 
                          let subsequent = Array.drop (i + 1) fields
                              aliveForV = Set.union alive (Array.foldl Set.union Set.empty (map (\\(Tuple _ sv) -> freeVariables sv) subsequent))
                              valCode = codegenExpr_ currentMod allZeroArity allMacroBindings Nothing aritiesMap globalClassFields bound aliveForV false val
                              valTy = inferTypeExpr currentMod aritiesMap bound val
                          in boxUnbox Any valTy valCode
                        ) fields)
                  let fieldsStr = if Array.null fields then "" else "(" <> fieldsCode <> ")"
                  let res = "crate::Value::" <> "Adt_" <> modPrefix <> sanitizeIdent tyNameStr <> "(perceus_ptr::PerceusPtr::new(" <> adtName <> "_enum::" <> sanitizeIdent ctorName <> fieldsStr <> "))"
                  pure $ Debug.trace ("KnotTying NO " <> dbgTraceMsg) \\_ -> res
        in reuseCode
  CtorDef _ _ (Ident ctorName) _ -> "{ let _t: crate::UnknownType = unimplemented!(); _t } /* CtorDef not supported */\""
"""

start_str = "  CtorSaturated (Qualified mbMod _) _ (ProperName tyNameStr) (Ident ctorName) fields ->"
end_str = "  CtorDef _ _ (Ident ctorName) _ ->"

start_idx = content.find(start_str)
end_idx = content.find("  LetRec _ binds body ->")

if start_idx != -1 and end_idx != -1:
    new_content = content[:start_idx] + replacement + "\n" + content[end_idx:]
    with open('/Users/0x1/Documents/htdocs/purust/purust/src/Purust/CodeGen.purs', 'w') as f:
        f.write(new_content)
    print("Patched CtorSaturated successfully")
else:
    print("Could not find boundaries")
