import re

with open('/Users/0x1/Documents/htdocs/purescript-backend-optimizer/src/PureScript/Backend/Optimizer/Monomorphize.purs', 'r') as f:
    code = f.read()

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
