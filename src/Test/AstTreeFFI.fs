module Test.AstTreeFFI
type Expr = Value of int | Add of Expr * Expr
let rec eval = function Value n -> n | Add (l, r) -> eval l + eval r
let rec makeTree depth = if depth = 0 then Value 1 else Add (makeTree (depth - 1), makeTree (depth - 1))
let runAstTreeFFI (d: obj) = eval (makeTree (unbox<int> d)) :> obj
