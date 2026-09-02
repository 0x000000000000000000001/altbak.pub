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
import Test.RBTree as RBTree
import Test.Polymorphism as Polymorphism
import Test.StateMonad as StateMonad
import Test.LazyEvaluation as LazyEvaluation
import Test.ArrayOps as ArrayOps
import Test.RowToList as RowToList
import Test.Primes as Primes

warmup :: Effect Unit
warmup = do
  void AstTree.act -- 1
  void Fib.act -- 2
  void ListOps.act -- 3
  void TCO.act -- 4
  void Records.act -- 5
  void Ackermann.act -- 6
  void Church.act -- 7
  void Primes.act -- 8
  void RBTree.act -- 9
  void Polymorphism.act -- 10
  void StateMonad.act -- 11
  void LazyEvaluation.act -- 12
  void ArrayOps.act -- 13
  void RowToList.act -- 14

main :: Effect Unit
main = do
  log "Global warm-up in progress (this may take a moment)...\n"
  warmup
  warmup
  warmup
  log "Global warm-up complete. Starting benchmarks...\n"

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
  t14 <- runBench RowToList.describe RowToList.act

  let totalMs = (t1 / 1000.0) + (t2 / 1000.0) + (t3 / 1000.0) + (t4 / 1000.0) + (t5 / 1000.0) + (t6 / 1000.0) + (t7 / 1000.0) + (t8 / 1000.0) + (t9 / 1000.0) + (t10 / 1000.0) + (t11 / 1000.0) + (t12 / 1000.0) + (t13 / 1000.0) + (t14 / 1000.0)
  log $ "\n==================================================\n\nTotal exec time: " <> formatNumber totalMs <> " ms\n"
