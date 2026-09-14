module AppFFICheatcode where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench (runBench, formatNumber)
import Test.AstTreeFFICheatcode as AstTreeFFICheatcode
import Test.FibFFICheatcode as FibFFICheatcode
import Test.ListOpsFFICheatcode as ListOpsFFICheatcode
import Test.TCOFFICheatcode as TCOFFICheatcode
import Test.RecordsFFICheatcode as RecordsFFICheatcode
import Test.AckermannFFICheatcode as AckermannFFICheatcode
import Test.ChurchFFICheatcode as ChurchFFICheatcode
import Test.PrimesFFICheatcode as PrimesFFICheatcode
import Test.RBTreeFFICheatcode as RBTreeFFICheatcode
import Test.PolymorphismFFICheatcode as PolymorphismFFICheatcode
import Test.StateMonadFFICheatcode as StateMonadFFICheatcode
import Test.LazyEvaluationFFICheatcode as LazyEvaluationFFICheatcode
import Test.ArrayOpsFFICheatcode as ArrayOpsFFICheatcode
import Test.RowToListFFICheatcode as RowToListFFICheatcode

warmup :: Effect Unit
warmup = do
  void AstTreeFFICheatcode.act
  void FibFFICheatcode.act
  void ListOpsFFICheatcode.act
  void TCOFFICheatcode.act
  void RecordsFFICheatcode.act
  void AckermannFFICheatcode.act
  void ChurchFFICheatcode.act
  void PrimesFFICheatcode.act
  void RBTreeFFICheatcode.act
  void PolymorphismFFICheatcode.act
  void StateMonadFFICheatcode.act
  void LazyEvaluationFFICheatcode.act
  void ArrayOpsFFICheatcode.act
  void RowToListFFICheatcode.act

main :: Effect Unit
main = do
  log "Global warm-up in progress (this may take a moment)...\n"
  warmup
  warmup
  warmup
  log "Global warm-up complete. Starting benchmarks...\n"

  t1 <- runBench AstTreeFFICheatcode.describe AstTreeFFICheatcode.act
  t2 <- runBench FibFFICheatcode.describe FibFFICheatcode.act
  t3 <- runBench ListOpsFFICheatcode.describe ListOpsFFICheatcode.act
  t4 <- runBench TCOFFICheatcode.describe TCOFFICheatcode.act
  t5 <- runBench RecordsFFICheatcode.describe RecordsFFICheatcode.act
  t6 <- runBench AckermannFFICheatcode.describe AckermannFFICheatcode.act
  t7 <- runBench ChurchFFICheatcode.describe ChurchFFICheatcode.act
  t8 <- runBench PrimesFFICheatcode.describe PrimesFFICheatcode.act
  t9 <- runBench RBTreeFFICheatcode.describe RBTreeFFICheatcode.act
  t10 <- runBench PolymorphismFFICheatcode.describe PolymorphismFFICheatcode.act
  t11 <- runBench StateMonadFFICheatcode.describe StateMonadFFICheatcode.act
  t12 <- runBench LazyEvaluationFFICheatcode.describe LazyEvaluationFFICheatcode.act
  t13 <- runBench ArrayOpsFFICheatcode.describe ArrayOpsFFICheatcode.act
  t14 <- runBench RowToListFFICheatcode.describe RowToListFFICheatcode.act

  let totalMs = (t1 / 1000.0) + (t2 / 1000.0) + (t3 / 1000.0) + (t4 / 1000.0) + (t5 / 1000.0) + (t6 / 1000.0) + (t7 / 1000.0) + (t8 / 1000.0) + (t9 / 1000.0) + (t10 / 1000.0) + (t11 / 1000.0) + (t12 / 1000.0) + (t13 / 1000.0) + (t14 / 1000.0)
  log $ "\n==================================================\n\nTotal exec time: " <> formatNumber totalMs <> " ms\n"
