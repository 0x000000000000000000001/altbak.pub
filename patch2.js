const fs = require('fs');
let code = fs.readFileSync('../gopurs/src/Gopurs/CodeGen.purs', 'utf8');
code = code.replace(
  `          paramsWithTypes = case getExprType tcoExpr of
            Func fArgs _ -> Array.zipWith (\\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)
            _ -> map (\\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args`,
  `          mbTopLevelInfo = case tcoIdent of
            Just topName -> Map.lookup topName moduleArities
            Nothing -> Nothing

          paramsWithTypes = case mbTopLevelInfo of
            Just info -> Array.zipWith (\\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (info.fArgs <> Array.replicate (Array.length args - Array.length info.fArgs) TypeValue)
            Nothing -> case getExprType tcoExpr of
              Func fArgs _ -> Array.zipWith (\\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)
              _ -> map (\\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args`
);
fs.writeFileSync('../gopurs/src/Gopurs/CodeGen.purs', code);
