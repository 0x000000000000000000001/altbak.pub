import re

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

pattern = r'''                        actualFieldType = case resObj\.exprType of
                          TypeStructPointer _ _ _ tArgs ->
                            case Map\.lookup key helpers\.ctorTypes of
                              Just ctorInfo ->
                                let
                                  env = Map\.fromFoldable \(Array\.zip ctorInfo\.vars tArgs\)
                                  genericTy = structFieldGoType \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.pointerAdtPaths \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.enumAdts \(unsafePerformEffect \(Ref\.read helpersRef\)\)\.elidedCtors ctorInfo\.vars modNameStr \(fromMaybe \(TypeVar ""\) \(Array\.index fields idx\)\)
                                in
                                  instantiateGenericGoType env genericTy
                              Nothing -> expectedType
                          _ -> expectedType

                        exprAccess =
                          if isNative then
                            GoConstructorAccess resObj\.expr monoStructName typeArgs idx true
                          else
                            GoConstructorAccess \(boxGoExpr resObj\.expr resObj\.exprType\) monoStructName typeArgs idx false
                      in
                        if isNative then
                          \{ stmts: resObj\.stmts, expr: coerceGoExpr exprAccess actualFieldType expectedType, exprType: expectedType, nextId: resObj\.nextId \}
                        else
                          \{ stmts: resObj\.stmts, expr: coerceGoExpr exprAccess expectedType expectedType, exprType: expectedType, nextId: resObj\.nextId \}'''

replacement = r'''                        exprAccessActualType = case Map.lookup key helpers.ctorTypes of
                          Just ctorInfo ->
                            let
                              env = Map.fromFoldable (Array.zip ctorInfo.vars typeArgs)
                              genericTy = structFieldGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors ctorInfo.vars modNameStr (fromMaybe (TypeVar "") (Array.index fields idx))
                            in
                              instantiateGenericGoType env genericTy
                          Nothing -> expectedType

                        exprAccess =
                          if isNative then
                            GoConstructorAccess resObj.expr monoStructName typeArgs idx true
                          else
                            GoConstructorAccess (boxGoExpr resObj.expr resObj.exprType) monoStructName typeArgs idx false
                      in
                        { stmts: resObj.stmts, expr: coerceGoExpr exprAccess exprAccessActualType expectedType, exprType: expectedType, nextId: resObj.nextId }'''

content = content.replace(pattern, replacement)

# Because python strings and regex can be tricky with indentation and newlines, let's just use string replace since it's an exact match!
with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(content)
