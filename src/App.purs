module App where

import Prelude
import Effect (Effect)
import Test.Fib as Fib
import Test.AstTree as AstTree

main :: Effect Unit
main = do
  AstTree.describe
  AstTree.act
  
  Fib.describe
  Fib.act
