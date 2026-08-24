const fs = require('fs');
let codeGen = fs.readFileSync('../phpurs/phpurs/src/Phpurs/CodeGen.purs', 'utf8');

codeGen = codeGen.replace(/freeVars :: TcoExpr -> Set String\nfreeVars \(TcoExpr \(TcoAnalysis \{ freeVars: fvs \}\) _\) = fvs/,
  `freeVars :: TcoExpr -> Set String\nfreeVars (TcoExpr (TcoAnalysis { freeVars: fvs }) _) = fvs\n\nisEffectNode :: TcoExpr -> Boolean\nisEffectNode (TcoExpr _ syntax) = case syntax of\n  EffectBind _ _ _ _ -> true\n  EffectPure _ -> true\n  EffectDefer _ -> false\n  PrimEffect _ -> true\n  UncurriedEffectApp _ _ -> true\n  Let _ _ _ body -> isEffectNode body\n  LetRec _ _ body -> isEffectNode body\n  _ -> false\n\nexecuteIfOpaque :: TcoExpr -> PhpExpr -> PhpExpr\nexecuteIfOpaque expr phpExpr =\n  if isEffectNode expr then phpExpr\n  else PhpCall (PhpRaw "phpurs_execute_effect") [ phpExpr ]`
);

codeGen = codeGen.replace(/translateExprImpl :: String -> Array String -> Map String String -> Map String String -> Maybe String -> Array String -> Boolean -> Int -> TcoExpr -> \{ stmts :: Array PhpStmt, expr :: PhpExpr, nextId :: Int \}\ntranslateExprImpl modNameStr recVars namedBound bound mbNamedVar loopCtx isTail nextId tcoExpr@\(TcoExpr _ syntax\) =/,
  `translateExprImpl_ :: String -> Array String -> Map String String -> Map String String -> Maybe String -> Array String -> Boolean -> Boolean -> Int -> TcoExpr -> { stmts :: Array PhpStmt, expr :: PhpExpr, nextId :: Int }\ntranslateExprImpl_ modNameStr recVars namedBound bound mbNamedVar loopCtx isTail inEffectBlock nextId tcoExpr@(TcoExpr _ syntax) =\n  let\n    isEff = isEffectNode tcoExpr\n  in\n    if isEff && not inEffectBlock then\n      let\n        res = translateExprImpl_ modNameStr recVars namedBound bound mbNamedVar loopCtx false true nextId tcoExpr\n        fvs = freeVars tcoExpr\n        mappedFvs = map (\\v -> fromMaybe v (Map.lookup v bound)) (Array.fromFoldable fvs)\n        useVars = Array.nub (map (\\mapped -> if Array.elem mapped recVars then "&" <> mapped else mapped) mappedFvs)\n      in\n        { stmts: [], expr: PhpFunction useVars [] "" (res.stmts <> [ PhpReturn res.expr ]), nextId: res.nextId }\n    else`
);

codeGen = codeGen.replace(/translateExprImpl /g, "translateExprImpl_ ");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound bound Nothing \[\] false ([a-zA-Z0-9_\.]+) expr/g, "translateExprImpl_ modNameStr recVars namedBound bound Nothing [] false false $1 expr");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound bound \(Just v\) loopCtx isTail ([a-zA-Z0-9_\.]+) body/g, "translateExprImpl_ modNameStr recVars namedBound bound (Just v) loopCtx isTail false $1 body");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound bound \(Just varName\) \[\] false ([a-zA-Z0-9_\.]+) val/g, "translateExprImpl_ modNameStr recVars namedBound bound (Just varName) [] false false $1 val");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound newBound Nothing loopCtx isTail ([a-zA-Z0-9_\.]+) body/g, "translateExprImpl_ modNameStr recVars namedBound newBound Nothing loopCtx isTail false $1 body");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound newBound \(Just varName\) \[\] true ([a-zA-Z0-9_\.]+) innerBody/g, "translateExprImpl_ modNameStr recVars namedBound newBound (Just varName) [] true false $1 innerBody");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound acc\.bound Nothing loopCtx isTail ([a-zA-Z0-9_\.]+) body/g, "translateExprImpl_ modNameStr recVars namedBound acc.bound Nothing loopCtx isTail false $1 body");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound bound Nothing loopCtx isTail ([a-zA-Z0-9_\.]+) def/g, "translateExprImpl_ modNameStr recVars namedBound bound Nothing loopCtx isTail false $1 def");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound bound Nothing \[\] false ([a-zA-Z0-9_\.]+) condExpr/g, "translateExprImpl_ modNameStr recVars namedBound bound Nothing [] false false $1 condExpr");

codeGen = codeGen.replace(/translateExprImpl_ modNameStr recVars namedBound bound Nothing loopCtx isTail ([a-zA-Z0-9_\.]+) bodyExpr/g, "translateExprImpl_ modNameStr recVars namedBound bound Nothing loopCtx isTail false $1 bodyExpr");


fs.writeFileSync('../phpurs/phpurs/src/Phpurs/CodeGen.purs', codeGen);
