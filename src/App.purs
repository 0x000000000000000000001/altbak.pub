module App where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Test.Fib as Fib
import Test.AstTree as AstTree

main :: Effect Unit
main = do
  log $ AstTree.description <> ":"
  AstTree.act
  
  log $ Fib.description <> ":"
  Fib.act
