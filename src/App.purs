module App where

import Prelude
import Effect (Effect)
import Test.Fib as Fib
import Test.AstTree as AstTree
import Test.ListOps as ListOps
import Test.TCO as TCO
import Test.Records as Records

main :: Effect Unit
main = do
  AstTree.describe
  AstTree.act
  
  Fib.describe
  Fib.act
  
  ListOps.describe
  ListOps.act
  
  TCO.describe
  TCO.act
  
  Records.describe
  Records.act
