import re

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'r') as f:
    content = f.read()

replacement = """                  Nothing ->
                    let
                      mbDirectCall = case getVar (unwrapTcoExpr fn) of
                        Just { mbMod, name } ->
                          let
                            isLocal = map (String.replaceAll (Pattern ".") (Replacement "_") <<< unwrap) mbMod == Just modNameStr || mbMod == Nothing
                            modPrefix = case mbMod of
                              Just mn -> String.replaceAll (Pattern ".") (Replacement "_") (unwrap mn)
                              Nothing -> modNameStr
                            fromModuleArities = if isLocal then Map.lookup name moduleArities else Nothing
                            fromTypeSig = case extractFuncType fn of
                              Just { fArgs, fRet } ->
                                Just { fullName: "Call_" <> modPrefix <> "_" <> sanitizeName name, fArgs: map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs, fRet: exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr fRet, arity: Array.length fArgs }
                              Nothing ->
                                Nothing
                            
                            entry = case fromTypeSig of
                              Just e | not isLocal -> Just e
                              _ -> fromModuleArities
                          in
                            case entry of
                              Just e ->
                                if Array.length args == e.arity && e.arity >= 1 then Just e else Nothing
                              Nothing -> Nothing
                        Nothing -> Nothing
                    in
                      case mbDirectCall of
                        Just callTarget ->
                          let
                            accArgs = foldl
                              ( \\acc arg ->
                                  let
                                    argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                                  in
                                    { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                              )
                              { stmts: StmtEmpty, exprs: [], exprTypes: [], nextId }
                              args
                              
                            callArgs = Array.mapWithIndex (\\i expected ->
                                let arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                    actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                in unboxGoExpr arg actual expected
                              ) callTarget.fArgs
                          in
                            { stmts: accArgs.stmts, expr: boxGoExpr (GoCall (GoVar callTarget.fullName) callArgs) callTarget.fRet, exprType: TypeValue, nextId: accArgs.nextId }
                        Nothing ->
                          let
                            resFn = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false nextId fn
                            accArgs = foldl
                              ( \\acc arg ->
                                  let
                                    argRes = translateExprImpl_ helpersRef (depth + 1) modNameStr recVars moduleArities bound Nothing [] false false acc.nextId arg
                                  in
                                    { stmts: acc.stmts <> argRes.stmts, exprs: Array.snoc acc.exprs argRes.expr, exprTypes: Array.snoc acc.exprTypes argRes.exprType, nextId: argRes.nextId }
                              )
                              { stmts: resFn.stmts, exprs: [], exprTypes: [], nextId: resFn.nextId }
                              args
                            len = Array.length args
                            goFuncName = if len >= 2 && len <= 10 then "UncurriedApp" <> show len else "UncurriedApp"
                          in
                            case resFn.exprType of
                              TypeFunc fArgs fRet | Array.length fArgs == len ->
                                let
                                  callArgs = Array.mapWithIndex (\\i expected ->
                                      let arg = fromMaybe (GoRaw "nil") (Array.index accArgs.exprs i)
                                          actual = fromMaybe TypeValue (Array.index accArgs.exprTypes i)
                                      in unboxGoExpr arg actual expected
                                    ) fArgs
                                in
                                  { stmts: accArgs.stmts, expr: boxGoExpr (GoCall resFn.expr callArgs) fRet, exprType: TypeValue, nextId: accArgs.nextId }
                              _ ->
                                let
                                  boxedArgs = Array.zipWith (\\arg actual -> boxGoExpr arg actual) accArgs.exprs accArgs.exprTypes
                                in
                                  { stmts: accArgs.stmts, expr: GoCall (GoSelector (GoVar "gopurs_runtime") goFuncName) (Array.cons (boxGoExpr resFn.expr resFn.exprType) boxedArgs), exprType: TypeValue, nextId: accArgs.nextId }"""

# We need to replace the `Nothing ->` block in `UncurriedApp fn args` branch.
import sys

lines = content.split('\n')

start_idx = -1
for i, line in enumerate(lines):
    if line.strip() == "Nothing ->" and "resFn = translateExprImpl_" in lines[i+2] and "UncurriedApp" in content[max(0, content.rfind("UncurriedApp fn args ->", 0, sum(len(x)+1 for x in lines[:i]))):]:
        start_idx = i
        break

if start_idx == -1:
    print("Could not find insertion point")
    sys.exit(1)

# Find end of the Nothing block
end_idx = start_idx
while end_idx < len(lines):
    if "UncurriedAbs args body -> liftIfNeeded \\_ ->" in lines[end_idx]:
        break
    end_idx += 1

end_idx -= 2 # keep empty lines

new_content = '\n'.join(lines[:start_idx]) + '\n' + replacement + '\n' + '\n'.join(lines[end_idx:])

with open('../gopurs/gopurs/src/Gopurs/CodeGen.purs', 'w') as f:
    f.write(new_content)

print("Patched CodeGen.purs successfully")
