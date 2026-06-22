module App where

import Prelude
import Effect (Effect)
import Test.Fib as Fib
import Test.AstTree as AstTree
import Test.ListOps as ListOps

main :: Effect Unit
main = do
  AstTree.describe
  AstTree.act
  
  Fib.describe
  Fib.act
  
  ListOps.describe
  ListOps.act
