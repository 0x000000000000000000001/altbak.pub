module App where

import Prelude
import Effect (Effect)
import Test.Fib as Fib
import Test.AstTree as AstTree
import Test.ListOps as ListOps
import Test.TCO as TCO
import Test.Records as Records
import Test.Ackermann as Ackermann
import Test.Church as Church
import Test.Primes as Primes
import Test.RBTree as RBTree
import Test.Polymorphism as Polymorphism
import Test.StateMonad as StateMonad
import Test.LazyEvaluation as LazyEvaluation
import Test.FileOps as FileOps
import Test.Async as Async

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
  
  Ackermann.describe
  Ackermann.act
  
  Church.describe
  Church.act
  
  Primes.describe
  Primes.act
  
  RBTree.describe
  RBTree.act
  
  Polymorphism.describe
  Polymorphism.act
  
  StateMonad.describe
  StateMonad.act
  
  LazyEvaluation.describe
  LazyEvaluation.act
  
  FileOps.describe
  FileOps.act
  
  Async.describe
  Async.act
