module App where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench (runBench, formatNumber)
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
import Test.ArrayOps as ArrayOps
import TestLet as TestLet

main :: Effect Unit
main = do
  t1 <- runBench AstTree.describe AstTree.act
  t2 <- runBench Fib.describe Fib.act
  t3 <- runBench ListOps.describe ListOps.act
  t4 <- runBench TCO.describe TCO.act
  t5 <- runBench Records.describe Records.act
  t6 <- runBench Ackermann.describe Ackermann.act
  t7 <- runBench Church.describe Church.act
  t8 <- runBench Primes.describe Primes.act
  t9 <- runBench RBTree.describe RBTree.act
  t10 <- runBench Polymorphism.describe Polymorphism.act
  t11 <- runBench StateMonad.describe StateMonad.act
  t12 <- runBench LazyEvaluation.describe LazyEvaluation.act
  t13 <- runBench ArrayOps.describe ArrayOps.act

  let totalMs = (t1 / 1000.0) + (t2 / 1000.0) + (t3 / 1000.0) + (t4 / 1000.0) + (t5 / 1000.0) + (t6 / 1000.0) + (t7 / 1000.0) + (t8 / 1000.0) + (t9 / 1000.0) + (t10 / 1000.0) + (t11 / 1000.0) + (t12 / 1000.0) + (t13 / 1000.0)
  log $ "\n==================================================\n\nTotal exec time: " <> formatNumber totalMs <> " ms\n"
