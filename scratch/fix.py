import re

with open('scratch/Monomorphize.purs', 'r') as f:
    code = f.read()

bad1 = """                            let
                              rewrittenExpr = rewriteExpr globalAstMap Map.empty globalSubst astSubstFn resolvedExpr
                              localInstMap = collectAllLocalInsts [] Map.empty rewrittenExpr
                            specializedExpr = monomorphizeExpr modNameStr instMap localInstMap Map.empty rewrittenExpr"""

good1 = """                            rewrittenExpr = rewriteExpr globalAstMap Map.empty globalSubst astSubstFn resolvedExpr
                            localInstMap = collectAllLocalInsts [] Map.empty rewrittenExpr
                            specializedExpr = monomorphizeExpr modNameStr instMap localInstMap Map.empty rewrittenExpr"""
code = code.replace(bad1, good1)

bad2 = """                                    let localInstMap = collectAllLocalInsts [] Map.empty substitutedExpr
                                    specializedExpr = monomorphizeExpr caller currentMap localInstMap Map.empty substitutedExpr"""

good2 = """                                    localInstMap = collectAllLocalInsts [] Map.empty substitutedExpr
                                    specializedExpr = monomorphizeExpr caller currentMap localInstMap Map.empty substitutedExpr"""
code = code.replace(bad2, good2)

with open('scratch/Monomorphize.purs', 'w') as f:
    f.write(code)
print("fixed")
