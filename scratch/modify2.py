import re

with open('/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Monomorphize.purs', 'r') as f:
    code = f.read()

new_collect = """collectAllLocalInsts :: Array SpineArg -> LocalInstMap -> Expr Ann -> LocalInstMap
collectAllLocalInsts spine acc expr = case expr of
  ExprApp ann f arg ->
    let
      acc1 = collectAllLocalInsts (Array.cons (SpineApp arg) spine) acc f
      acc2 = collectAllLocalInsts [] acc1 arg
    in
      acc2
  ExprTypeApp ann f t ->
    collectAllLocalInsts (Array.cons (SpineTypeApp t) spine) acc f
  ExprVar _ (Qualified Nothing id) ->
    let
      typeArgs = getSpineTypeArgs spine
      args = getSpineArgs spine
      
      genericType = case getExprAnn expr of Ann a -> fromMaybe Any a.type
      substFromTypeArgs = buildSubst genericType typeArgs
      
      unifySpine :: ExprType -> Array (Expr Ann) -> Map String ExprType -> Map String ExprType
      unifySpine _ [] s = s
      unifySpine (ForAll _ t) args' s = unifySpine t args' s
      unifySpine (ConstrainedType constraints t) args' s =
        let numConstraints = Array.length constraints
        in unifySpine t (Array.drop numConstraints args') s
      unifySpine (Func paramTypes ret) args' s =
        let
          numParams = Array.length paramTypes
          appliedArgs = Array.take numParams args'
          remainingArgs = Array.drop numParams args'
          s1 = foldl (\\sAcc (Tuple paramType arg) ->
                 let actualType = case getExprAnn arg of Ann a -> fromMaybe Any a.type
                 in unify paramType actualType sAcc
               ) s (Array.zip paramTypes appliedArgs)
        in
          unifySpine ret remainingArgs s1
      unifySpine _ _ s = s
      
      finalSubst = unifySpine genericType args substFromTypeArgs
    in
      if not (Map.isEmpty finalSubst) then
        let
          stripForAlls = case _ t -> case t of
            ForAll _ b -> stripForAlls b
            x -> x
          instType = stripTypeVariables (substituteExprType finalSubst (stripForAlls genericType))
        in
          if (hasTypeVariables genericType) && (hasTypeVariables instType && instType == stripForAlls genericType) then
            acc
          else
            Map.insertWith (\\old new -> { genericType: new.genericType, insts: old.insts <> new.insts }) id { genericType, insts: [finalSubst] } acc
      else
        Map.insertWith (\\old new -> { genericType: new.genericType, insts: old.insts <> new.insts }) id { genericType, insts: [Map.empty] } acc

  ExprVar _ _ -> acc
  ExprLit _ lit -> foldl (\\a e -> collectAllLocalInsts [] a e) acc lit
  ExprConstructor _ _ _ _ -> acc
  ExprAccessor _ e _ -> collectAllLocalInsts [] acc e
  ExprUpdate _ e props -> foldl (\\a (Prop _ v) -> collectAllLocalInsts [] a v) (collectAllLocalInsts [] acc e) props
  ExprAbs _ _ e -> collectAllLocalInsts [] acc e
  ExprCase _ exprs alts ->
    foldl (\\a (CaseAlternative _ cg) -> case cg of
      Unconditional e' -> collectAllLocalInsts [] a e'
      Guarded guards -> foldl (\\a2 (Guard e1 e2) -> collectAllLocalInsts [] (collectAllLocalInsts [] a2 e1) e2) a guards
    ) (foldl (\\a e -> collectAllLocalInsts [] a e) acc exprs) alts
  ExprLet _ binds e -> 
    foldl (\\a b -> collectAllLocalBind a b) (collectAllLocalInsts spine acc e) binds

collectAllLocalBind :: LocalInstMap -> Bind Ann -> LocalInstMap
collectAllLocalBind acc = case _ of
  NonRec (Binding _ _ e) -> collectAllLocalInsts [] acc e
  Rec binds -> foldl (\\a (Binding _ _ e) -> collectAllLocalInsts [] a e) acc binds
"""

lines = code.split('\\n')
start_idx = -1
end_idx = -1
for i, line in enumerate(lines):
    if line.startswith("collectLocalExpr ::") and start_idx == -1:
        start_idx = i
    if line.startswith("collectTypesFromExpr ::") and end_idx == -1:
        end_idx = i

if start_idx != -1 and end_idx != -1:
    lines = lines[:start_idx] + [new_collect] + lines[end_idx:]
code = '\\n'.join(lines)

# 2. Add localInstMap to monomorphizeExpr and related functions
# First, monomorphizeBinding
old_mono_bind = """monomorphizeBinding :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Binding Ann -> Binding Ann
monomorphizeBinding modName instMap localDicts (Binding ann (Ident name) expr) =
  Binding ann (Ident name) (monomorphizeExpr modName instMap localDicts expr)"""

new_mono_bind = """monomorphizeBinding :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Binding Ann -> Binding Ann
monomorphizeBinding modName instMap localDicts (Binding ann (Ident name) expr) =
  let
    localInstMap = collectAllLocalInsts [] Map.empty expr
  in
    Binding ann (Ident name) (monomorphizeExpr modName instMap localInstMap localDicts expr)"""
code = code.replace(old_mono_bind, new_mono_bind)

# getInjectedBindsFor
old_inj = """                            specializedExpr = monomorphizeExpr modNameStr instMap Map.empty (rewriteExpr globalAstMap Map.empty globalSubst astSubstFn resolvedExpr)"""
new_inj = """                            rewrittenExpr = rewriteExpr globalAstMap Map.empty globalSubst astSubstFn resolvedExpr
                            localInstMap = collectAllLocalInsts [] Map.empty rewrittenExpr
                            specializedExpr = monomorphizeExpr modNameStr instMap localInstMap Map.empty rewrittenExpr"""
code = code.replace(old_inj, new_inj)

old_inj2 = """                                    specializedExpr = monomorphizeExpr caller currentMap Map.empty substitutedExpr"""
new_inj2 = """                                    localInstMap = collectAllLocalInsts [] Map.empty substitutedExpr
                                    specializedExpr = monomorphizeExpr caller currentMap localInstMap Map.empty substitutedExpr"""
code = code.replace(old_inj2, new_inj2)


# Replace signatures and calls
code = code.replace("monomorphizeExpr :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Expr Ann -> Expr Ann",
                    "monomorphizeExpr :: String -> InstantiationMap -> LocalInstMap -> Map Ident (Expr Ann) -> Expr Ann -> Expr Ann")

code = code.replace("monomorphizeExpr modName instMap localDicts expr =",
                    "monomorphizeExpr modName instMap localInstMap localDicts expr =")

# Calls
code = code.replace("monomorphizeExpr modName instMap localDicts", "monomorphizeExpr modName instMap localInstMap localDicts")
code = code.replace("monomorphizeExpr modName instMap newLocalDicts", "monomorphizeExpr modName instMap localInstMap newLocalDicts")

code = code.replace("monomorphizeBindLocal :: String -> InstantiationMap -> Map Ident (Expr Ann) -> Bind Ann -> Bind Ann",
                    "monomorphizeBindLocal :: String -> InstantiationMap -> LocalInstMap -> Map Ident (Expr Ann) -> Bind Ann -> Bind Ann")
code = code.replace("monomorphizeBindLocal modName instMap localDicts =",
                    "monomorphizeBindLocal modName instMap localInstMap localDicts =")
code = code.replace("monomorphizeBindLocal modName instMap localDicts", "monomorphizeBindLocal modName instMap localInstMap localDicts")
code = code.replace("monomorphizeBindLocal modName instMap newLocalDicts", "monomorphizeBindLocal modName instMap localInstMap newLocalDicts")

code = code.replace("monomorphizeAlt modName instMap localDicts (CaseAlternative ann guard) =",
                    "monomorphizeAlt modName instMap localInstMap localDicts (CaseAlternative ann guard) =")
code = code.replace("monomorphizeAlt modName instMap localDicts", "monomorphizeAlt modName instMap localInstMap localDicts")

code = code.replace("monomorphizeCaseGuard modName instMap localDicts", "monomorphizeCaseGuard modName instMap localInstMap localDicts")
code = code.replace("monomorphizeGuard modName instMap localDicts", "monomorphizeGuard modName instMap localInstMap localDicts")
code = code.replace("monomorphizeProp modName instMap localDicts", "monomorphizeProp modName instMap localInstMap localDicts")

# Remove the localInstMap computation inside monomorphizeExpr
old_local_compute = "      localInstMap = collectLocalExpr boundIds Map.empty (ExprLet ann binds e)\\n"
code = code.replace(old_local_compute, "")

with open('/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Monomorphize.purs', 'w') as f:
    f.write(code)

print("done")
