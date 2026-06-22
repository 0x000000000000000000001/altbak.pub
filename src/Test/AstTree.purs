module Test.AstTree where

import Prelude
import Effect (Effect)
import Effect.Console (logShow, log)

data Expr
  = Val Int
  | Add Expr Expr
  | Mul Expr Expr
  | Sub Expr Expr

eval :: Expr -> Int
eval (Val n) = n
eval (Add a b) = eval a + eval b
eval (Mul a b) = eval a * eval b
eval (Sub a b) = eval a - eval b

buildTree :: Int -> Expr
buildTree 0 = Val 1
buildTree n = Add (Mul (Val n) (buildTree (n - 1))) (Sub (buildTree (n - 1)) (Val 1))

describe :: Effect Unit
describe = log "AST Evaluation:"

act :: Effect Unit
act = logShow $ eval (buildTree 12)
