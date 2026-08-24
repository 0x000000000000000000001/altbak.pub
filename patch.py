import re

with open('../phpurs/phpurs/src/Phpurs/CodeGen.purs', 'r') as f:
    content = f.read()

# Add isEffectNode and executeIfOpaque
content = content.replace(
    'freeVars :: TcoExpr -> Set String\nfreeVars (TcoExpr (TcoAnalysis { freeVars: fvs }) _) = fvs',
    '''freeVars :: TcoExpr -> Set String
freeVars (TcoExpr (TcoAnalysis { freeVars: fvs }) _) = fvs

isEffectNode :: TcoExpr -> Boolean
isEffectNode (TcoExpr _ syntax) = case syntax of
  EffectBind _ _ _ _ -> true
  EffectPure _ -> true
  EffectDefer _ -> false
  PrimEffect _ -> true
  UncurriedEffectApp _ _ -> true
  Let _ _ _ body -> isEffectNode body
  LetRec _ _ body -> isEffectNode body
  _ -> false

executeIfOpaque :: TcoExpr -> PhpExpr -> PhpExpr
executeIfOpaque expr phpExpr =
  if isEffectNode expr then phpExpr
  else PhpCall (PhpRaw "phpurs_execute_effect") [ phpExpr ]'''
)

# Replace function signature and body start
content = content.replace(
    '''translateExprImpl :: String -> Array String -> Map String String -> Map String String -> Maybe String -> Array String -> Boolean -> Int -> TcoExpr -> { stmts :: Array PhpStmt, expr :: PhpExpr, nextId :: Int }
translateExprImpl modNameStr recVars namedBound bound mbNamedVar loopCtx isTail nextId tcoExpr@(TcoExpr _ syntax) =
  let
    doTrace f =''',
    '''translateExprImpl_ :: String -> Array String -> Map String String -> Map String String -> Maybe String -> Array String -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: Array PhpStmt, expr :: PhpExpr, nextId :: Int }
translateExprImpl_ modNameStr recVars namedBound bound mbNamedVar loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr _ syntax) =
  let
    isEff = isEffectNode tcoExpr
  in
    if isEff && not inEffectBlock then
      let
        res = translateExprImpl_ modNameStr recVars namedBound bound mbNamedVar loopCtx false true nextId tcoExpr
        fvs = freeVars tcoExpr
        mappedFvs = map (\\v -> fromMaybe v (Map.lookup v bound)) (Array.fromFoldable fvs)
        useVars = Array.nub (map (\\mapped -> if Array.elem mapped recVars then "&" <> mapped else mapped) mappedFvs)
      in
        { stmts: [], expr: PhpFunction useVars [] "" (res.stmts <> [ PhpReturn res.expr ]), nextId: res.nextId }
    else
  let
    doTrace f ='''
)

def replacer(m):
    return f'translateExprImpl_ {m.group(1)} false {m.group(2)}'

content = re.sub(r'translateExprImpl (modNameStr.*?(?:\[\]|loopCtx|Nothing).*?(?:false|true|isTail)) ([a-zA-Z0-9_\.\+ \(\)]+) (expr|a|e|body|val|innerBody|def|condExpr|bodyExpr|fn|arg|argExpr|flatFn|e1|e2)', replacer, content)

content = content.replace(
    '''EffectBind (Just (Ident i)) (Level l) val body ->
    let
      oldVarName = localId (Just (Ident i)) (Level l)
      varName = oldVarName <> "_" <> show nextId
      resVal = translateExprImpl_ modNameStr recVars namedBound bound (Just varName) [] false false nextId val
      newBound = Map.insert oldVarName varName bound
      resBody = translateExprImpl_ modNameStr recVars namedBound newBound Nothing loopCtx isTail false (resVal.nextId + 1) body
    in
      { stmts: resVal.stmts <> [ PhpAssign varName resVal.expr ] <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }''',
    '''EffectBind (Just (Ident i)) (Level l) val body ->
    let
      oldVarName = localId (Just (Ident i)) (Level l)
      varName = oldVarName <> "_" <> show nextId
      resVal = translateExprImpl_ modNameStr recVars namedBound bound (Just varName) [] false true nextId val
      newBound = Map.insert oldVarName varName bound
      resBody = translateExprImpl_ modNameStr recVars namedBound newBound Nothing loopCtx isTail true (resVal.nextId + 1) body
      valExpr = executeIfOpaque val resVal.expr
    in
      { stmts: resVal.stmts <> [ PhpAssign varName valExpr ] <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }'''
)

content = content.replace(
    '''EffectBind Nothing (Level l) val body ->
    let
      oldVarName = localId Nothing (Level l)
      varName = oldVarName <> "_" <> show nextId
      resVal = translateExprImpl_ modNameStr recVars namedBound bound (Just varName) [] false false nextId val
      newBound = Map.insert oldVarName varName bound
      resBody = translateExprImpl_ modNameStr recVars namedBound newBound Nothing loopCtx isTail false (resVal.nextId + 1) body
    in
      { stmts: resVal.stmts <> [ PhpAssign varName resVal.expr ] <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }''',
    '''EffectBind Nothing (Level l) val body ->
    let
      oldVarName = localId Nothing (Level l)
      varName = oldVarName <> "_" <> show nextId
      resVal = translateExprImpl_ modNameStr recVars namedBound bound (Just varName) [] false true nextId val
      newBound = Map.insert oldVarName varName bound
      resBody = translateExprImpl_ modNameStr recVars namedBound newBound Nothing loopCtx isTail true (resVal.nextId + 1) body
      valExpr = executeIfOpaque val resVal.expr
    in
      { stmts: resVal.stmts <> [ PhpAssign varName valExpr ] <> resBody.stmts, expr: resBody.expr, nextId: resBody.nextId }'''
)

content = content.replace(
    '''EffectDefer e ->
    let
      res = translateExprImpl_ modNameStr recVars namedBound bound Nothing [] false false nextId e''',
    '''EffectDefer e ->
    let
      res = translateExprImpl_ modNameStr recVars namedBound bound Nothing [] false true nextId e'''
)

with open('../phpurs/phpurs/src/Phpurs/CodeGen.purs', 'w') as f:
    f.write(content)
