module Test.AstTreeFFICheatcode

type Expr = Val of int | Add of Expr * Expr | Mul of Expr * Expr | Sub of Expr * Expr
let rec build n =
    if n = 0 then Val 1
    else Add (Mul (Val n, build (n - 1)), Sub (build (n - 1), Val 1))
let rec eval = function
    | Val n -> n
    | Add (a, b) -> eval a + eval b
    | Mul (a, b) -> eval a * eval b
    | Sub (a, b) -> eval a - eval b
let runAstTreeFFICheatcode (n: obj) = eval (build (unbox<int> n)) :> obj
