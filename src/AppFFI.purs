module AppFFI where

import Prelude
import Effect (Effect)
import Effect.Console (log)
import Bench (runBench, formatNumber)
import Test.AstTreeFFI as AstTreeFFI
import Test.FibFFI as FibFFI
import Test.ListOpsFFI as ListOpsFFI
import Test.TCOFFI as TCOFFI
import Test.RecordsFFI as RecordsFFI
import Test.AckermannFFI as AckermannFFI
import Test.ChurchFFI as ChurchFFI
import Test.PrimesFFI as PrimesFFI
import Test.RBTreeFFI as RBTreeFFI
import Test.PolymorphismFFI as PolymorphismFFI
import Test.StateMonadFFI as StateMonadFFI
import Test.LazyEvaluationFFI as LazyEvaluationFFI
import Test.ArrayOpsFFI as ArrayOpsFFI
import Test.RowToListFFI as RowToListFFI

main :: Effect Unit
main = do
  t1 <- runBench AstTreeFFI.describe AstTreeFFI.act
  t2 <- runBench FibFFI.describe FibFFI.act
  t3 <- runBench ListOpsFFI.describe ListOpsFFI.act
  t4 <- runBench TCOFFI.describe TCOFFI.act
  t5 <- runBench RecordsFFI.describe RecordsFFI.act
  t6 <- runBench AckermannFFI.describe AckermannFFI.act
  t7 <- runBench ChurchFFI.describe ChurchFFI.act
  t8 <- runBench PrimesFFI.describe PrimesFFI.act
  t9 <- runBench RBTreeFFI.describe RBTreeFFI.act
  t10 <- runBench PolymorphismFFI.describe PolymorphismFFI.act
  t11 <- runBench StateMonadFFI.describe StateMonadFFI.act
  t12 <- runBench LazyEvaluationFFI.describe LazyEvaluationFFI.act
  t13 <- runBench ArrayOpsFFI.describe ArrayOpsFFI.act
  t14 <- runBench RowToListFFI.describe RowToListFFI.act

  let totalMs = (t1 / 1000.0) + (t2 / 1000.0) + (t3 / 1000.0) + (t4 / 1000.0) + (t5 / 1000.0) + (t6 / 1000.0) + (t7 / 1000.0) + (t8 / 1000.0) + (t9 / 1000.0) + (t10 / 1000.0) + (t11 / 1000.0) + (t12 / 1000.0) + (t13 / 1000.0) + (t14 / 1000.0)
  log $ "\n==================================================\n\nTotal exec time: " <> formatNumber totalMs <> " ms\n"
